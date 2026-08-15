import { pascal } from "./schema-model.mjs";

function pathStem(id) { return id.split("/").at(-1).replace(".schema.json", ""); }
function goName(value) {
  return pascal(value).replace(/Id/g, "ID").replace(/Url/g, "URL").replace(/Ip/g, "IP").replace(/Json/g, "JSON").replace(/Sha/g, "SHA");
}

function registry(model, extension) {
  const names = new Map();
  for (const document of model.documents.filter((item) => item.extension === extension)) {
    if (document.schema.title) names.set(document.id, `Wire${goName(document.schema.title)}`);
    for (const [key, definition] of Object.entries(document.schema.$defs ?? {})) {
      names.set(`${document.id}#/$defs/${key}`, `Wire${goName(definition.title ?? `${document.schema.title ?? pathStem(document.id)} ${key}`)}`);
    }
  }
  return names;
}

export function renderGo(model, { extension = false, packageName = "artifact" } = {}) {
  const names = registry(model, extension);
  const declarations = [];
  const methods = [];

  const typeOf = (node, document, optional = false) => {
    if (node.$ref) {
      const resolved = model.resolve(node.$ref, document.id);
      const name = names.get(resolved.ref);
      if (name) return optional && isObject(resolved.node) ? `*${name}` : name;
      return typeOf(resolved.node, resolved.document, optional);
    }
    if (node.oneOf || node.anyOf) return "json.RawMessage";
    if (node.allOf && !isObject(node)) return "json.RawMessage";
    if (Object.hasOwn(node, "const") || node.enum || node.type === "string") return optional ? "*string" : "string";
    if (node.type === "integer") return optional ? "*int64" : "int64";
    if (node.type === "number") return optional ? "*float64" : "float64";
    if (node.type === "boolean") return optional ? "*bool" : "bool";
    if (node.type === "array") return `[]${typeOf(node.items ?? {}, document)}`;
    if (node.type === "object" || node.properties) return "map[string]any";
    return "json.RawMessage";
  };

  const isObject = (node) => Boolean(node.type === "object" || node.properties || objectParts(node).length);
  const objectParts = (node) => (node.allOf ?? []).filter((part) => !part.if && !part.then);
  const flattened = (node, document, seen = new Set()) => {
    const properties = { ...(node.properties ?? {}) };
    const required = new Set(node.required ?? []);
    for (const part of objectParts(node)) {
      let target = part; let targetDocument = document;
      if (part.$ref) { const resolved = model.resolve(part.$ref, document.id); target = resolved.node; targetDocument = resolved.document; }
      const key = `${targetDocument.id}:${JSON.stringify(target)}`;
      if (seen.has(key)) continue;
      const nested = flattened(target, targetDocument, new Set([...seen, key]));
      Object.assign(properties, nested.properties);
      nested.required.forEach((item) => required.add(item));
    }
    return { properties, required };
  };

  const declare = (name, node, document) => {
    if (node.oneOf || node.anyOf || (node.$ref && !node.properties)) {
      declarations.push(`type ${name} = ${typeOf(node, document)}`);
      return;
    }
    const { properties, required } = flattened(node, document);
    if ((node.type === "object" || properties) && Object.keys(properties).length === 0 && node.additionalProperties === true) {
      declarations.push(`type ${name} map[string]any`);
      return;
    }
    if (isObject(node)) {
      const fields = Object.entries(properties).map(([jsonName, schema]) => {
        const optional = !required.has(jsonName);
        return `\t${goName(jsonName)} ${typeOf(schema, document, optional)} \`json:"${jsonName}${optional ? ",omitempty" : ""}"\``;
      });
      fields.push('\tUnknownFields map[string]json.RawMessage `json:"-"`');
      declarations.push(`type ${name} struct {\n${fields.join("\n")}\n}`);
      const known = Object.keys(properties).map((key) => `"${key}": {}`).join(", ");
      methods.push(`func (value *${name}) UnmarshalJSON(data []byte) error {
\ttype plain ${name}
\tvar decoded plain
\tif err := json.Unmarshal(data, &decoded); err != nil { return err }
\t*value = ${name}(decoded)
\tvar fields map[string]json.RawMessage
\tif err := json.Unmarshal(data, &fields); err != nil { return err }
\tknown := map[string]struct{}{${known}}
\tfor key := range known { delete(fields, key) }
\tif len(fields) > 0 { value.UnknownFields = fields }
\treturn nil
}

func (value ${name}) MarshalJSON() ([]byte, error) {
\ttype plain ${name}
\tknownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
\tvar fields map[string]json.RawMessage
\tif err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
\treserved := map[string]struct{}{${known}}
\tfor key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
\treturn json.Marshal(fields)
}`);
      return;
    }
    declarations.push(`type ${name} = ${typeOf(node, document)}`);
  };

  for (const document of model.documents.filter((item) => item.extension === extension)) {
    if (document.schema.title) declare(names.get(document.id), document.schema, document);
    for (const [key, definition] of Object.entries(document.schema.$defs ?? {})) declare(names.get(`${document.id}#/$defs/${key}`), definition, document);
  }
  return `// Code generated from canonical JSON Schemas; DO NOT EDIT.\npackage ${packageName}\n\nimport "encoding/json"\n\n${declarations.join("\n\n")}\n\n${methods.join("\n\n")}\n`;
}
