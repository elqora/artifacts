import assert from "node:assert/strict";
import { cp, mkdtemp, readFile, writeFile } from "node:fs/promises";
import os from "node:os";
import path from "node:path";
import test from "node:test";
import { fileURLToPath } from "node:url";
import { generatedOutputs } from "../../generators/generate-bindings.mjs";
import { loadSchemaModel } from "../../generators/lib/schema-model.mjs";
import { renderTypeScript } from "../../generators/lib/render-typescript.mjs";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "../..");
const artifactId = "https://artifact-sdk.dev/schema/1.1/artifact/artifact.schema.json";
const conditionId = "https://artifact-sdk.dev/schema/1.1/policy/condition.schema.json";
const githubId = "https://artifact-sdk.dev/schema/1.1/extensions/github/github-reference.schema.json";

test("schema model loads every core and GitHub schema and resolves references", async () => {
  const model = await loadSchemaModel(root);
  assert.equal(model.documents.length, 25);
  const artifact = model.byId.get(artifactId);
  const valueType = model.resolve(artifact.schema.properties.valueType.$ref, artifact.id);
  assert.deepEqual(valueType.node.enum.slice(0, 3), ["text", "number", "boolean"]);
});

test("model preserves recursion, unions, composition, and open properties", async () => {
  const model = await loadSchemaModel(root);
  const condition = model.byId.get(conditionId).schema;
  assert.equal(condition.$defs.condition.oneOf.length, 7);
  assert.equal(condition.$defs.not.properties.condition.$ref, "#/$defs/condition");
  assert.equal(model.byId.get(githubId).schema.$defs.pullRequest.allOf.length, 2);
  assert.equal(model.byId.get(artifactId).schema.additionalProperties, true);
});

test("rendering is deterministic and exposes stable generated names", async () => {
  const model = await loadSchemaModel(root);
  const first = renderTypeScript(model);
  const second = renderTypeScript(model);
  assert.equal(first, second);
  assert.match(first, /export type WireArtifact =/);
  assert.match(first, /WireArtifactConditionAnd/);
});

async function copiedRepository() {
  const temporary = await mkdtemp(path.join(os.tmpdir(), "artifact-codegen-"));
  await cp(path.join(root, "spec"), path.join(temporary, "spec"), { recursive: true });
  await cp(path.join(root, "extensions", "github", "spec"), path.join(temporary, "extensions", "github", "spec"), { recursive: true });
  return temporary;
}

test("a canonical schema field addition changes generated language output", async () => {
  const temporary = await copiedRepository();
  const before = await generatedOutputs(temporary);
  const artifactPath = path.join(temporary, "spec", "artifact", "artifact.schema.json");
  const schema = JSON.parse(await readFile(artifactPath, "utf8"));
  schema.properties.futureCanonicalField = { type: "string" };
  await writeFile(artifactPath, JSON.stringify(schema));
  const after = await generatedOutputs(temporary);
  assert.notEqual(before.get("packages/typescript/src/generated/wire.ts"), after.get("packages/typescript/src/generated/wire.ts"));
  assert.notEqual(before.get("packages/go/schema_wire_gen.go"), after.get("packages/go/schema_wire_gen.go"));
  assert.notEqual(before.get("packages/php/src/Dto/Artifact.php"), after.get("packages/php/src/Dto/Artifact.php"));
});

test("generation fails loudly for unsupported structural schema keywords", async () => {
  const temporary = await copiedRepository();
  const artifactPath = path.join(temporary, "spec", "artifact", "artifact.schema.json");
  const schema = JSON.parse(await readFile(artifactPath, "utf8"));
  schema.unevaluatedProperties = false;
  await writeFile(artifactPath, JSON.stringify(schema));
  await assert.rejects(loadSchemaModel(temporary), /Unsupported schema keyword unevaluatedProperties/);
});
