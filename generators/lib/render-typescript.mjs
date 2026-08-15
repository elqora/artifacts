import { pascal } from "./schema-model.mjs";

function literal(value) {
  return JSON.stringify(value);
}

function registry(model, extension) {
  const names = new Map();
  for (const document of model.documents.filter((item) => item.extension === extension)) {
    if (document.schema.title && document.schema.type !== undefined || document.schema.oneOf || document.schema.$ref) {
      names.set(document.id, `Wire${pascal(document.schema.title ?? pathStem(document.id))}`);
    }
    for (const [key, definition] of Object.entries(document.schema.$defs ?? {})) {
      const base = definition.title ?? `${document.schema.title ?? pathStem(document.id)} ${key}`;
      names.set(`${document.id}#/$defs/${key}`, `Wire${pascal(base)}`);
    }
  }
  return names;
}

function pathStem(id) {
  return id.split("/").at(-1).replace(".schema.json", "");
}

export function renderTypeScript(model, { extension = false } = {}) {
  const names = registry(model, extension);
  const declarations = [];

  const typeOf = (node, document, stack = new Set()) => {
    if (node.$ref) {
      const resolved = model.resolve(node.$ref, document.id);
      const name = names.get(resolved.ref);
      if (name) return name;
      return typeOf(resolved.node, resolved.document, stack);
    }
    if (Object.hasOwn(node, "const")) return literal(node.const);
    if (node.enum) return node.enum.map(literal).join(" | ");
    if (node.oneOf) return node.oneOf.map((item) => typeOf(item, document, stack)).join(" | ");
    if (node.anyOf) return node.anyOf.map((item) => typeOf(item, document, stack)).join(" | ");
    if (node.type === "array") return `Array<${typeOf(node.items ?? {}, document, stack)}>`;
    if (node.type === "string") return "string";
    if (node.type === "integer" || node.type === "number") return "number";
    if (node.type === "boolean") return "boolean";
    if (node.type === "object" || node.properties || node.additionalProperties) {
      const required = new Set(node.required ?? []);
      const fields = Object.entries(node.properties ?? {}).map(([key, value]) => `  ${JSON.stringify(key)}${required.has(key) ? "" : "?"}: ${typeOf(value, document, stack)};`);
      if (node.additionalProperties === true) fields.push("  [key: string]: unknown;");
      else if (node.additionalProperties && typeof node.additionalProperties === "object") fields.push(`  [key: string]: ${typeOf(node.additionalProperties, document, stack)};`);
      return `{\n${fields.join("\n")}\n}`;
    }
    if (node.allOf) return node.allOf.filter((item) => !item.if && !item.then).map((item) => typeOf(item, document, stack)).filter((value) => value !== "unknown").join(" & ") || "unknown";
    return "unknown";
  };

  const declare = (name, node, document) => {
    const description = node.description ? `/** ${node.description.replaceAll("*/", "* /")} */\n` : "";
    declarations.push(`${description}export type ${name} = ${typeOf(node, document)};`);
  };

  for (const document of model.documents.filter((item) => item.extension === extension)) {
    const rootName = names.get(document.id);
    if (rootName) declare(rootName, document.schema, document);
    for (const [key, definition] of Object.entries(document.schema.$defs ?? {})) {
      declare(names.get(`${document.id}#/$defs/${key}`), definition, document);
    }
  }

  return `// Code generated from canonical JSON Schemas; DO NOT EDIT.\n\n${declarations.join("\n\n")}\n`;
}
