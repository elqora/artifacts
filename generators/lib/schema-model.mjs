import { readdir, readFile } from "node:fs/promises";
import path from "node:path";

const structuralKeywords = new Set([
  "$schema", "$id", "$ref", "$defs", "title", "description", "type",
  "required", "properties", "additionalProperties", "items", "oneOf", "anyOf", "allOf",
  "const", "enum", "if", "then", "format", "pattern", "minimum", "maximum",
  "exclusiveMinimum", "exclusiveMaximum", "minLength", "maxLength", "minItems",
  "maxItems", "uniqueItems", "multipleOf", "x-artifact-codegen",
]);

async function jsonFiles(directory) {
  const entries = await readdir(directory, { withFileTypes: true });
  const files = await Promise.all(entries.map((entry) => {
    const target = path.join(directory, entry.name);
    return entry.isDirectory() ? jsonFiles(target) : [target];
  }));
  return files.flat().filter((file) => file.endsWith(".schema.json"));
}

function pointer(root, fragment) {
  if (!fragment || fragment === "#") return root;
  if (!fragment.startsWith("#/")) throw new Error(`Unsupported JSON pointer: ${fragment}`);
  return fragment.slice(2).split("/").reduce((value, part) => {
    const key = part.replaceAll("~1", "/").replaceAll("~0", "~");
    if (value?.[key] === undefined) throw new Error(`Unresolvable JSON pointer: ${fragment}`);
    return value[key];
  }, root);
}

function assertSupported(node, location = "schema") {
  if (!node || typeof node !== "object" || Array.isArray(node)) return;
  for (const key of Object.keys(node)) {
    if (!structuralKeywords.has(key)) throw new Error(`Unsupported schema keyword ${key} at ${location}`);
  }
  for (const key of ["properties", "$defs"]) {
    for (const [name, child] of Object.entries(node[key] ?? {})) assertSupported(child, `${location}/${key}/${name}`);
  }
  for (const key of ["items", "additionalProperties", "if", "then"]) {
    if (node[key] && typeof node[key] === "object") assertSupported(node[key], `${location}/${key}`);
  }
  for (const key of ["oneOf", "anyOf", "allOf"]) {
    for (const [index, child] of (node[key] ?? []).entries()) assertSupported(child, `${location}/${key}/${index}`);
  }
}

export async function loadSchemaModel(repositoryRoot) {
  const roots = [path.join(repositoryRoot, "spec"), path.join(repositoryRoot, "extensions", "github", "spec")];
  const files = (await Promise.all(roots.map(jsonFiles))).flat().sort();
  const documents = [];
  const byId = new Map();
  for (const file of files) {
    const schema = JSON.parse(await readFile(file, "utf8"));
    if (!schema.$id) throw new Error(`Schema has no $id: ${file}`);
    assertSupported(schema, schema.$id);
    const document = { file, schema, id: schema.$id, extension: file.includes(`${path.sep}extensions${path.sep}`) };
    documents.push(document);
    byId.set(schema.$id, document);
  }

  function resolve(ref, baseId) {
    const [targetPart, fragment = ""] = ref.split("#", 2);
    const targetId = targetPart || baseId;
    const document = byId.get(targetId);
    if (!document) throw new Error(`Unresolvable schema reference ${ref} from ${baseId}`);
    return { document, node: pointer(document.schema, fragment ? `#${fragment}` : "#"), ref: `${targetId}${fragment ? `#${fragment}` : ""}` };
  }

  for (const document of documents) {
    const visit = (node) => {
      if (!node || typeof node !== "object") return;
      if (node.$ref) resolve(node.$ref, document.id);
      for (const value of Object.values(node)) {
        if (Array.isArray(value)) value.forEach(visit);
        else if (value && typeof value === "object") visit(value);
      }
    };
    visit(document.schema);
  }

  return { documents, byId, resolve };
}

export function schemaVocabularies(model) {
  const get = (id) => model.byId.get(id)?.schema ?? (() => { throw new Error(`Missing schema ${id}`); })();
  const protocol = get("https://artifact-sdk.dev/schema/1.0/protocol/protocol.schema.json");
  const source = get("https://artifact-sdk.dev/schema/1.0/artifact/artifact-source.schema.json");
  const integrity = get("https://artifact-sdk.dev/schema/1.0/artifact/artifact-integrity.schema.json");
  const condition = get("https://artifact-sdk.dev/schema/1.0/policy/condition.schema.json");
  const privacy = get("https://artifact-sdk.dev/schema/1.0/policy/privacy-policy.schema.json");
  const verification = get("https://artifact-sdk.dev/schema/1.0/runtime/verification.schema.json");
  return {
    valueTypes: protocol.$defs.artifactValueType.enum,
    sourceTypes: source.oneOf.map(({ $ref }) => $ref.replace("#/$defs/", "")),
    syncStates: source.$defs.local.properties.syncState.enum,
    integrityAlgorithms: integrity.properties.algorithm.enum,
    conditionKinds: condition.$defs.condition.oneOf.map(({ $ref }) => condition.$defs[$ref.replace("#/$defs/", "")].properties.kind.const),
    valueOperators: condition.$defs.artifactValue.properties.operator.enum,
    privacyClassifications: privacy.properties.classification.enum,
    privacyRepresentations: privacy.$defs.revealRule.properties.representation.enum,
    verificationStatuses: verification.properties.status.enum,
  };
}

export function pascal(value) {
  return value.replace(/(^|[^A-Za-z0-9]+)([A-Za-z0-9])/g, (_, __, char) => char.toUpperCase()).replace(/[^A-Za-z0-9]/g, "");
}
