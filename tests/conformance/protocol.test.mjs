import assert from "node:assert/strict";
import { readdir, readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import test from "node:test";
import Ajv2020 from "ajv/dist/2020.js";
import addFormats from "ajv-formats";
import {
  ARTIFACT_INTEGRITY_ALGORITHMS,
  ARTIFACT_PROTOCOL_VERSION,
  ARTIFACT_SOURCE_TYPES,
  ARTIFACT_VALUE_TYPES,
  LOCAL_ARTIFACT_SYNC_STATES,
} from "../../packages/typescript/dist/index.js";

const repositoryRoot = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  "../..",
);
const schemaRoot = path.join(repositoryRoot, "spec");
const extensionSchemaRoot = path.join(repositoryRoot, "extensions");
const fixtureRoot = path.join(repositoryRoot, "tests", "fixtures");

const schemaIds = {
  protocol:
    "https://artifact-sdk.dev/schema/1.0/protocol/protocol.schema.json",
  artifact:
    "https://artifact-sdk.dev/schema/1.0/artifact/artifact.schema.json",
  version:
    "https://artifact-sdk.dev/schema/1.0/artifact/artifact-version.schema.json",
  source:
    "https://artifact-sdk.dev/schema/1.0/artifact/artifact-source.schema.json",
  integrity:
    "https://artifact-sdk.dev/schema/1.0/artifact/artifact-integrity.schema.json",
  link: "https://artifact-sdk.dev/schema/1.0/artifact/artifact-link.schema.json",
  valueSchema: "https://artifact-sdk.dev/schema/1.0/specification/artifact-value-schema.schema.json",
  spec: "https://artifact-sdk.dev/schema/1.0/specification/artifact-spec.schema.json",
  snapshot: "https://artifact-sdk.dev/schema/1.0/specification/artifact-spec-snapshot.schema.json",
  requirement: "https://artifact-sdk.dev/schema/1.0/specification/artifact-requirement.schema.json",
  condition: "https://artifact-sdk.dev/schema/1.0/policy/condition.schema.json",
  providerPolicy: "https://artifact-sdk.dev/schema/1.0/policy/provider-policy.schema.json",
  requirementPolicy: "https://artifact-sdk.dev/schema/1.0/policy/requirement-policy.schema.json",
  lifecyclePolicy: "https://artifact-sdk.dev/schema/1.0/policy/lifecycle-policy.schema.json",
  accessPolicy: "https://artifact-sdk.dev/schema/1.0/policy/access-policy.schema.json",
  privacyPolicy: "https://artifact-sdk.dev/schema/1.0/policy/privacy-policy.schema.json",
  validationPolicy: "https://artifact-sdk.dev/schema/1.0/policy/validation-policy.schema.json",
  verificationPolicy: "https://artifact-sdk.dev/schema/1.0/policy/verification-policy.schema.json",
  retentionPolicy: "https://artifact-sdk.dev/schema/1.0/policy/retention-policy.schema.json",
  presentationPolicy: "https://artifact-sdk.dev/schema/1.0/policy/presentation-policy.schema.json",
  submission: "https://artifact-sdk.dev/schema/1.0/runtime/submission.schema.json",
  verification: "https://artifact-sdk.dev/schema/1.0/runtime/verification.schema.json",
  githubSource: "https://artifact-sdk.dev/schema/1.0/extensions/github/github-source.schema.json",
};

async function findJsonFiles(directory) {
  const entries = await readdir(directory, { withFileTypes: true });
  const nested = await Promise.all(
    entries.map((entry) => {
      const target = path.join(directory, entry.name);
      return entry.isDirectory() ? findJsonFiles(target) : [target];
    }),
  );
  return nested.flat().filter((file) => file.endsWith(".json"));
}

async function readJson(file) {
  return JSON.parse(await readFile(file, "utf8"));
}

const ajv = new Ajv2020({ allErrors: true, strict: true });
addFormats(ajv);
for (const schemaFile of await findJsonFiles(schemaRoot)) {
  ajv.addSchema(await readJson(schemaFile));
}
for (const schemaFile of await findJsonFiles(extensionSchemaRoot)) {
  if (schemaFile.includes(`${path.sep}spec${path.sep}`)) {
    ajv.addSchema(await readJson(schemaFile));
  }
}

const validateArtifact = ajv.getSchema(schemaIds.artifact);
const validateVersion = ajv.getSchema(schemaIds.version);
const validateSource = ajv.getSchema(schemaIds.source);
const validateIntegrity = ajv.getSchema(schemaIds.integrity);
const validateLink = ajv.getSchema(schemaIds.link);
const validators = Object.fromEntries(
  Object.entries(schemaIds).map(([name, id]) => [name, ajv.getSchema(id)]),
);

assert.ok(validateArtifact);
assert.ok(validateVersion);
assert.ok(validateSource);
assert.ok(validateIntegrity);
assert.ok(validateLink);
for (const [name, validate] of Object.entries(validators)) {
  assert.ok(validate, `missing compiled schema: ${name}`);
}

function assertValid(validate, value) {
  assert.equal(validate(value), true, JSON.stringify(validate.errors, null, 2));
}

function assertInvalid(validate, value) {
  assert.equal(validate(value), false, "expected schema validation to fail");
}

const fixtureValidators = new Map([
  ["artifact", validateArtifact],
  ["version", validateVersion],
  ["link", validateLink],
  ["specification", null],
  ["value-schema", validators.valueSchema],
  ["runtime", null],
  ["condition", validators.condition],
  ["provider", validators.githubSource],
]);

test("all canonical record fixtures conform to their schemas", async () => {
  for (const [directory, validate] of fixtureValidators) {
    for (const file of await findJsonFiles(path.join(fixtureRoot, directory))) {
      const value = await readJson(file);
      const filename = path.basename(file);
      const selected = validate ?? (
        directory === "runtime"
          ? (filename.includes("verification") ? validators.verification : validators.submission)
          : (filename.includes("snapshot") ? validators.snapshot : filename.includes("requirement") ? validators.requirement : validators.spec)
      );
      assertValid(selected, value);
    }
  }
});

test("all standardized ArtifactValueSchema discriminators and recursive collections validate", async () => {
  const values = [
    { valueType: "text", minLength: 1 }, { valueType: "number", integer: true },
    { valueType: "boolean" }, { valueType: "currency", currencies: ["NGN", "USD"] },
    { valueType: "date", minimum: "2026-01-01" },
    { valueType: "datetime", maximum: "2027-01-01T00:00:00Z" },
    { valueType: "time", minimum: "08:00:00Z" },
    { valueType: "location", mode: "point_and_address" },
    { valueType: "file", maxFiles: 2 }, { valueType: "image", requireLocation: true },
    { valueType: "video", maxDurationSeconds: 90 }, { valueType: "audio", maxDurationSeconds: 300 },
    { valueType: "link", allowedSchemes: ["https"] },
    { valueType: "structured", jsonSchema: { type: "object" } },
    { valueType: "reference", providers: ["github"] },
    { valueType: "signature", methods: ["drawn"] },
    await readJson(path.join(fixtureRoot, "value-schema", "nested-collection.json")),
  ];
  for (const value of values) assertValid(validators.valueSchema, value);
  assertInvalid(validators.valueSchema, { valueType: "custom" });
});

test("recursive conditions accept nested ASTs and reject empty or executable-like nodes", async () => {
  assertValid(validators.condition, await readJson(path.join(fixtureRoot, "condition", "nested.json")));
  assertInvalid(validators.condition, await readJson(path.join(fixtureRoot, "invalid", "condition", "empty-and.json")));
  assertInvalid(validators.condition, await readJson(path.join(fixtureRoot, "invalid", "condition", "callback-like.json")));
});

test("every declarative policy fixture conforms and closed policy vocabulary rejects drift", async () => {
  const mapping = {
    provider: validators.providerPolicy, requirement: validators.requirementPolicy,
    lifecycle: validators.lifecyclePolicy, access: validators.accessPolicy,
    privacy: validators.privacyPolicy, validation: validators.validationPolicy,
    verification: validators.verificationPolicy, retention: validators.retentionPolicy,
    presentation: validators.presentationPolicy,
  };
  for (const [name, validate] of Object.entries(mapping)) {
    assertValid(validate, await readJson(path.join(fixtureRoot, "policy", `${name}.json`)));
  }
  assertInvalid(validators.privacyPolicy, { classification: "secret" });
  assertInvalid(validators.retentionPolicy, { policy: "duration", days: -1 });
  assertInvalid(validators.requirementPolicy, { mode: "conditional" });
});

test("spec config representation agrees with the containing valueType", async () => {
  const valid = await readJson(path.join(fixtureRoot, "specification", "delivery-evidence-v3.json"));
  assertValid(validators.spec, valid);
  assert.equal(valid.config.valueType, valid.valueType);
  const goalStyleConfig = structuredClone(valid);
  delete goalStyleConfig.config.valueType;
  assertValid(validators.spec, goalStyleConfig);
  const invalid = await readJson(path.join(fixtureRoot, "invalid", "specification", "mismatched-config.json"));
  assertInvalid(validators.spec, invalid);
});

test("a captured v3 specification snapshot is not changed by the live v4 definition", async () => {
  const snapshot = await readJson(path.join(fixtureRoot, "specification", "delivery-evidence-v3-snapshot.json"));
  const current = await readJson(path.join(fixtureRoot, "specification", "delivery-evidence-v4.json"));
  assertValid(validators.snapshot, snapshot);
  assertValid(validators.spec, current);
  assert.equal(snapshot.sourceSpecId, current.id);
  assert.equal(snapshot.sourceVersion, 3);
  assert.equal(current.version, 4);
  assert.equal(snapshot.config.minFiles, 1);
  assert.equal(current.config.minFiles, 3);
});

test("submissions and verification preserve exact-version provenance", async () => {
  const submission = await readJson(path.join(fixtureRoot, "runtime", "delivery-submission.json"));
  const verification = await readJson(path.join(fixtureRoot, "runtime", "client-verification.json"));
  assertValid(validators.submission, submission);
  assertValid(validators.verification, verification);
  assert.equal(verification.artifactId, submission.artifactId);
  assert.equal(verification.artifactVersionId, submission.artifactVersionId);
  assert.equal(verification.submissionId, submission.id);
  assertInvalid(validators.verification, await readJson(path.join(fixtureRoot, "invalid", "runtime", "unknown-verification-status.json")));
});

test("all ArtifactSource discriminators validate with their required fields", () => {
  const sources = [
    { type: "inline", value: null },
    { type: "local", localId: "local_1", syncState: "local_only" },
    { type: "object", objectId: "object_1" },
    { type: "url", url: "https://example.test/resource" },
    { type: "hosted", recordType: "report", recordId: "report_1" },
    { type: "provider", provider: "forge", reference: { id: "change_1" } },
  ];

  for (const source of sources) assertValid(validateSource, source);
  assertInvalid(validateSource, { type: "object" });
  assertInvalid(validateSource, { type: "unknown", id: "value" });
});

test("required fields, versions, timestamps, and null semantics are enforced", () => {
  assertInvalid(validateArtifact, {
    schemaVersion: "1.0",
    id: "art_1",
    kind: "example",
    valueType: "text",
    createdAt: "2026-08-15T09:00:00Z",
    updatedAt: "2026-08-15T09:00:00Z",
  });
  assertInvalid(validateArtifact, {
    schemaVersion: "1.0",
    id: "art_1",
    kind: "example",
    valueType: "text",
    title: null,
    createdBy: { type: "system" },
    createdAt: "not-a-timestamp",
    updatedAt: "2026-08-15T09:00:00Z",
  });
  assertInvalid(validateVersion, {
    schemaVersion: "1.0",
    id: "artver_1",
    artifactId: "art_1",
    version: 0,
    source: { type: "inline", value: true },
    createdBy: { type: "system" },
    createdAt: "2026-08-15T09:00:00Z",
  });
});

test("closed vocabularies reject unknown values", () => {
  const baseArtifact = {
    schemaVersion: "1.0",
    id: "art_1",
    kind: "custom_kind",
    valueType: "custom_value_type",
    createdBy: { type: "custom_actor" },
    createdAt: "2026-08-15T09:00:00Z",
    updatedAt: "2026-08-15T09:00:00Z",
  };
  assertInvalid(validateArtifact, baseArtifact);
  assertInvalid(validateSource, {
    type: "local",
    localId: "local_1",
    syncState: "queued",
  });
  assertInvalid(validateIntegrity, {
    algorithm: "md5",
    hash: "0".repeat(32),
  });
  assertInvalid(validateIntegrity, {
    algorithm: "sha256",
    hash: "ABC",
  });
});

test("host-defined vocabularies and unknown additive properties remain open", () => {
  assertValid(validateArtifact, {
    schemaVersion: "1.0",
    id: "opaque:ART-1",
    kind: "host_specific_kind",
    valueType: "reference",
    createdBy: { type: "host_actor_type", hostField: 7 },
    createdAt: "2026-08-15T09:00:00+01:00",
    updatedAt: "2026-08-15T09:00:00+01:00",
    futureProtocolField: { enabled: true },
  });
  assertValid(validateLink, {
    schemaVersion: "1.0",
    id: "alink_1",
    artifactId: "art_1",
    subject: { type: "host_subject_type", id: "subject_1" },
    role: "host_role",
    createdBy: { type: "host_actor_type" },
    createdAt: "2026-08-15T09:00:00Z",
  });
  assertValid(validateSource, {
    type: "provider",
    provider: "host_provider",
    reference: { resource: "host_resource", nested: [1, true, null] },
  });
});

test("logical-following and exact-version-pinned links both conform", async () => {
  const follows = await readJson(
    path.join(fixtureRoot, "link", "follows-artifact.json"),
  );
  const pinned = await readJson(
    path.join(fixtureRoot, "link", "pinned-version.json"),
  );
  assertValid(validateLink, follows);
  assert.equal("artifactVersionId" in follows, false);
  assertValid(validateLink, pinned);
  assert.equal(pinned.artifactVersionId, "artver_design_1");
});

test("the multiple-version scenario preserves immutable historical pinning", async () => {
  const scenarioFile = path.join(
    fixtureRoot,
    "scenario",
    "multiple-versions.json",
  );
  const scenario = await readJson(scenarioFile);
  const resolveFixture = (relative) =>
    readJson(path.resolve(path.dirname(scenarioFile), relative));
  const artifact = await resolveFixture(scenario.artifactFixture);
  const versions = await Promise.all(scenario.versionFixtures.map(resolveFixture));
  const followingLink = await resolveFixture(scenario.followingLinkFixture);
  const pinnedLink = await resolveFixture(scenario.pinnedLinkFixture);

  assertValid(validateArtifact, artifact);
  for (const version of versions) {
    assertValid(validateVersion, version);
    assert.equal(version.artifactId, artifact.id);
  }
  assert.deepEqual(
    versions.map(({ version }) => version),
    [1, 2],
  );
  assert.equal(artifact.currentVersionId, versions[1].id);
  assert.equal("artifactVersionId" in followingLink, false);
  assert.equal(pinnedLink.artifactVersionId, versions[0].id);
});

test("TypeScript runtime vocabularies match canonical schema vocabularies", () => {
  const protocolSchema = ajv.getSchema(schemaIds.protocol)?.schema;
  const sourceSchema = validateSource.schema;
  const integritySchema = validateIntegrity.schema;

  assert.equal(ARTIFACT_PROTOCOL_VERSION, "1.0");
  assert.deepEqual(
    [...ARTIFACT_VALUE_TYPES],
    protocolSchema.$defs.artifactValueType.enum,
  );
  assert.deepEqual(
    [...LOCAL_ARTIFACT_SYNC_STATES],
    sourceSchema.$defs.local.properties.syncState.enum,
  );
  assert.deepEqual(
    [...ARTIFACT_SOURCE_TYPES],
    sourceSchema.oneOf.map(({ $ref }) => $ref.replace("#/$defs/", "")),
  );
  assert.deepEqual(
    [...ARTIFACT_INTEGRITY_ALGORITHMS],
    integritySchema.properties.algorithm.enum,
  );
});
