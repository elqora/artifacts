import { mkdir, readFile, writeFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import { loadSchemaModel, schemaVocabularies } from "./lib/schema-model.mjs";
import { renderTypeScript } from "./lib/render-typescript.mjs";
import { renderGo } from "./lib/render-go.mjs";
import { renderPhpDtos, renderPhpEnums } from "./lib/render-php.mjs";

export const repositoryRoot = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");

export async function generatedOutputs(root = repositoryRoot) {
  const model = await loadSchemaModel(root);
  const vocab = schemaVocabularies(model);
  const quoteGo = (values) => values.map((value) => `\t"${value}",`).join("\n");
  const go = `// Code generated from canonical JSON Schemas; DO NOT EDIT.
package artifact

var ArtifactValueTypes = []string{
${quoteGo(vocab.valueTypes)}
}

var ArtifactSourceTypes = []string{
${quoteGo(vocab.sourceTypes)}
}

var LocalArtifactSyncStates = []string{
${quoteGo(vocab.syncStates)}
}

var ArtifactIntegrityAlgorithms = []string{
${quoteGo(vocab.integrityAlgorithms)}
}

var ArtifactConditionKinds = []string{
${quoteGo(vocab.conditionKinds)}
}

var ArtifactValueOperators = []string{
${quoteGo(vocab.valueOperators)}
}

var ArtifactPrivacyClassifications = []string{
${quoteGo(vocab.privacyClassifications)}
}

var ArtifactPrivacyRepresentations = []string{
${quoteGo(vocab.privacyRepresentations)}
}

var ArtifactVerificationStatuses = []string{
${quoteGo(vocab.verificationStatuses)}
}
`;

  const phpArray = (values) => values.map((value) => `        '${value}',`).join("\n");
  const php = `<?php

declare(strict_types=1);

// Code generated from canonical JSON Schemas; DO NOT EDIT.
namespace Elqora\\Artifact;

final class Vocabulary
{
    public const PROTOCOL_VERSION = '1.0';
    public const ARTIFACT_VALUE_TYPES = [
${phpArray(vocab.valueTypes)}
    ];
    public const ARTIFACT_SOURCE_TYPES = [
${phpArray(vocab.sourceTypes)}
    ];
    public const LOCAL_SYNC_STATES = [
${phpArray(vocab.syncStates)}
    ];
    public const INTEGRITY_ALGORITHMS = [
${phpArray(vocab.integrityAlgorithms)}
    ];
    public const CONDITION_KINDS = [
${phpArray(vocab.conditionKinds)}
    ];
    public const VALUE_OPERATORS = [
${phpArray(vocab.valueOperators)}
    ];
    public const PRIVACY_CLASSIFICATIONS = [
${phpArray(vocab.privacyClassifications)}
    ];
    public const PRIVACY_REPRESENTATIONS = [
${phpArray(vocab.privacyRepresentations)}
    ];
    public const VERIFICATION_STATUSES = [
${phpArray(vocab.verificationStatuses)}
    ];
}
`;

  const tsVocabulary = `// Code generated from canonical JSON Schemas; DO NOT EDIT.
export const ARTIFACT_VALUE_TYPES = ${JSON.stringify(vocab.valueTypes, null, 2)} as const;
export const ARTIFACT_SOURCE_TYPES = ${JSON.stringify(vocab.sourceTypes, null, 2)} as const;
export const LOCAL_ARTIFACT_SYNC_STATES = ${JSON.stringify(vocab.syncStates, null, 2)} as const;
export const ARTIFACT_INTEGRITY_ALGORITHMS = ${JSON.stringify(vocab.integrityAlgorithms, null, 2)} as const;
export const ARTIFACT_CONDITION_KINDS = ${JSON.stringify(vocab.conditionKinds, null, 2)} as const;
export const ARTIFACT_VALUE_OPERATORS = ${JSON.stringify(vocab.valueOperators, null, 2)} as const;
export const ARTIFACT_PRIVACY_CLASSIFICATIONS = ${JSON.stringify(vocab.privacyClassifications, null, 2)} as const;
export const ARTIFACT_PRIVACY_REPRESENTATIONS = ${JSON.stringify(vocab.privacyRepresentations, null, 2)} as const;
export const ARTIFACT_VERIFICATION_STATUSES = ${JSON.stringify(vocab.verificationStatuses, null, 2)} as const;
`;

  return new Map([
    ["packages/typescript/src/generated/wire.ts", renderTypeScript(model)],
    ["packages/typescript/src/generated/vocabulary.ts", tsVocabulary],
    ["extensions/github/typescript/src/generated/wire.ts", renderTypeScript(model, { extension: true })],
    ["packages/go/schema_wire_gen.go", renderGo(model)],
    ["extensions/github/go/schema_wire_gen.go", renderGo(model, { extension: true, packageName: "githubartifact" })],
    ["packages/go/vocabulary_gen.go", go],
    ["packages/php/src/Vocabulary.php", php],
    ...renderPhpEnums(vocab),
    ...renderPhpDtos(model),
  ]);
}

export async function runGeneration({ root = repositoryRoot, check = false } = {}) {
  const outputs = await generatedOutputs(root);
  const stale = [];
  for (const [relative, content] of outputs) {
    const target = path.join(root, relative);
    if (check) {
      const current = await readFile(target, "utf8").catch(() => "");
      if (current.replaceAll("\r\n", "\n") !== content) stale.push(relative);
    } else {
      await mkdir(path.dirname(target), { recursive: true });
      await writeFile(target, content, "utf8");
    }
  }
  if (stale.length > 0) throw new Error(`Generated bindings are stale: ${stale.join(", ")}`);
  return outputs;
}

if (process.argv[1] && path.resolve(process.argv[1]) === fileURLToPath(import.meta.url)) {
  await runGeneration({ check: process.argv.includes("--check") });
}
