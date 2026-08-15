import { pascal } from "./schema-model.mjs";

const preferred = {
  state: "ArtifactStateCondition", actor: "ArtifactActorCondition", artifactExists: "ArtifactExistsCondition",
  artifactValue: "ArtifactValueCondition", and: "ArtifactAndCondition", or: "ArtifactOrCondition", not: "ArtifactNotCondition",
  rule: "ArtifactAccessRule", context: "ArtifactSubmissionContext",
  text: "TextArtifactValueSchema", number: "NumberArtifactValueSchema", boolean: "BooleanArtifactValueSchema",
  currency: "CurrencyArtifactValueSchema", date: "DateArtifactValueSchema", datetime: "DatetimeArtifactValueSchema",
  time: "TimeArtifactValueSchema", location: "LocationArtifactValueSchema", file: "FileArtifactValueSchema",
  image: "ImageArtifactValueSchema", video: "VideoArtifactValueSchema", audio: "AudioArtifactValueSchema",
  link: "LinkArtifactValueSchema", structured: "StructuredArtifactValueSchema", reference: "ReferenceArtifactValueSchema",
  signature: "SignatureArtifactValueSchema", collection: "CollectionArtifactValueSchema",
};
const methodName = (value) => value.replace(/^[^A-Za-z_]+/, "").replace(/[^A-Za-z0-9_]/g, "_");

function isObject(node) { return Boolean(node.type === "object" || node.properties || (node.allOf ?? []).length); }

function flatten(model, node, document, seen = new Set()) {
  const properties = { ...(node.properties ?? {}) }; const required = new Set(node.required ?? []);
  for (const part of (node.allOf ?? []).filter((item) => !item.if && !item.then)) {
    let target = part; let targetDocument = document;
    if (part.$ref) { const resolved = model.resolve(part.$ref, document.id); target = resolved.node; targetDocument = resolved.document; }
    const marker = `${targetDocument.id}:${JSON.stringify(target)}`; if (seen.has(marker)) continue;
    const nested = flatten(model, target, targetDocument, new Set([...seen, marker]));
    Object.assign(properties, nested.properties); nested.required.forEach((key) => required.add(key));
  }
  return { properties, required };
}

function phpType(model, node, document) {
  if (node.$ref) {
    const resolved = model.resolve(node.$ref, document.id);
    return phpType(model, resolved.node, resolved.document);
  }
  if (node.const !== undefined || node.enum || node.type === "string" || node.format) return "string";
  if (node.type === "integer") return "int";
  if (node.type === "number") return "float|int";
  if (node.type === "boolean") return "bool";
  if (node.type === "array" || node.type === "object" || node.properties || node.oneOf || node.anyOf || node.allOf) return "array";
  return "mixed";
}

function renderDto(model, name, node, document, namespace) {
  const { properties, required } = flatten(model, node, document);
  const known = Object.keys(properties);
  const accessors = Object.entries(properties).map(([key, schema]) => {
    const type = phpType(model, schema, document); const optional = !required.has(key); const returnType = type === "mixed" ? "mixed" : optional ? (type.includes("|") ? `${type}|null` : `?${type}`) : type;
    const expression = optional ? `return $this->data['${key}'] ?? null;` : `return $this->data['${key}'];`;
    const doc = type === "array" ? `    /** @return ${optional ? "array<mixed>|null" : "array<mixed>"} */\n` : "";
    return `${doc}    public function ${methodName(key)}(): ${returnType}\n    {\n        ${expression}\n    }\n\n    public function has${pascal(key)}(): bool\n    {\n        return array_key_exists('${key}', $this->data);\n    }`;
  }).join("\n\n");
  const requiredChecks = [...required].map((key) => `        if (!array_key_exists('${key}', $data)) { throw new InvalidArgumentException('Missing required field ${key} for ${name}.'); }`).join("\n");
  const knownArray = known.map((key) => `'${key}' => true`).join(", ");
  return `<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace ${namespace};

use InvalidArgumentException;

final readonly class ${name}
{
    /** @param array<string, mixed> $data */
    private function __construct(private array $data) {}

    /** @param array<string, mixed> $data */
    public static function fromArray(array $data): self
    {
${requiredChecks || "        // This schema has no required fields."}
        return new self($data);
    }

${accessors}

    /** @return array<string, mixed> */
    public function unknownFields(): array
    {
        return array_diff_key($this->data, [${knownArray}]);
    }

    /** @return array<string, mixed> */
    public function toArray(): array
    {
        return $this->data;
    }
}
`;
}

export function renderPhpDtos(model) {
  const outputs = new Map(); const used = new Set();
  for (const document of model.documents) {
    const namespaceRoot = document.extension ? "extensions/github/php/Dto" : "packages/php/src/Dto";
    const phpNamespace = document.extension ? "Elqora\\Artifact\\GitHub\\Dto" : "Elqora\\Artifact\\Dto";
    if (document.schema.title && isObject(document.schema) && !document.schema.oneOf && !document.schema.$ref) {
      const name = pascal(document.schema.title); used.add(`${namespaceRoot}/${name}`);
      outputs.set(`${namespaceRoot}/${name}.php`, renderDto(model, name, document.schema, document, phpNamespace));
    }
    for (const [key, definition] of Object.entries(document.schema.$defs ?? {})) {
      if (!isObject(definition) || definition.oneOf || definition.$ref || key === "fileBase" || key === "rules" || key === "condition") continue;
      let name = definition.title ? pascal(definition.title) : preferred[key] ?? `${pascal(document.schema.title ?? "Artifact")} ${pascal(key)}`.replaceAll(" ", "");
      const marker = `${namespaceRoot}/${name}`; if (used.has(marker)) name = `${pascal(document.schema.title ?? "Artifact")}${pascal(key)}`;
      used.add(`${namespaceRoot}/${name}`);
      outputs.set(`${namespaceRoot}/${name}.php`, renderDto(model, name, definition, document, phpNamespace));
    }
  }
  return outputs;
}

export function renderPhpEnums(vocab) {
  const definitions = [
    ["ArtifactValueType", vocab.valueTypes], ["ArtifactSourceType", vocab.sourceTypes],
    ["LocalArtifactSyncState", vocab.syncStates], ["ArtifactIntegrityAlgorithm", vocab.integrityAlgorithms],
    ["ArtifactConditionKind", vocab.conditionKinds], ["ArtifactValueOperator", vocab.valueOperators],
    ["ArtifactPrivacyClassification", vocab.privacyClassifications], ["ArtifactPrivacyRepresentation", vocab.privacyRepresentations],
    ["ArtifactVerificationStatus", vocab.verificationStatuses],
  ];
  return new Map(definitions.map(([name, values]) => [`packages/php/src/Enum/${name}.php`, `<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\\Artifact\\Enum;

enum ${name}: string
{
${values.map((value) => `    case V_${value.toUpperCase().replace(/[^A-Z0-9]+/g, "_")} = '${value}';`).join("\n")}
}
`]));
}
