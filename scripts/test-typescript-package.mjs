import assert from "node:assert/strict";
import { mkdtempSync, mkdirSync, readFileSync, readdirSync, rmSync, writeFileSync } from "node:fs";
import { tmpdir } from "node:os";
import path from "node:path";
import { spawnSync } from "node:child_process";
import { fileURLToPath } from "node:url";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const workspace = path.join(root, "packages", "typescript");

function run(command, args, cwd = root) {
  const result = spawnSync(command, args, {
    cwd,
    encoding: "utf8",
    env: { ...process.env, npm_config_audit: "false", npm_config_fund: "false" },
    shell: process.platform === "win32" && command === "npm",
  });
  if (result.error) throw result.error;
  if (result.status !== 0) {
    process.stdout.write(result.stdout ?? "");
    process.stderr.write(result.stderr ?? "");
    throw new Error(`${command} ${args.join(" ")} exited with ${result.status}`);
  }
  return result.stdout ?? "";
}

function filesBelow(directory, relative = "") {
  return readdirSync(path.join(directory, relative), { withFileTypes: true }).flatMap((entry) => {
    const child = path.join(relative, entry.name);
    return entry.isDirectory() ? filesBelow(directory, child) : [child.replaceAll("\\", "/")];
  });
}

const scratch = mkdtempSync(path.join(tmpdir(), "elqora-artifacts-package-"));

try {
  const packOutput = run("npm", [
    "pack",
    "--json",
    "--pack-destination",
    scratch,
    "--workspace",
    "@elqora/artifacts",
  ]);
  const packed = JSON.parse(packOutput)[0];
  assert.equal(packed.name, "@elqora/artifacts");
  assert.equal(packed.version, "0.2.0");

  const packedPaths = packed.files.map(({ path: filePath }) => filePath).sort();
  const allowed = /^(LICENSE|README\.md|package\.json|dist\/(?:generated\/)?[^/]+\.(?:js|js\.map|d\.ts|d\.ts\.map))$/;
  assert.deepEqual(
    packedPaths.filter((filePath) => !allowed.test(filePath)),
    [],
    "the npm tarball contains files outside the release allowlist",
  );
  for (const required of ["LICENSE", "README.md", "package.json", "dist/index.js", "dist/index.d.ts"]) {
    assert(packedPaths.includes(required), `the npm tarball is missing ${required}`);
  }

  const consumer = path.join(scratch, "consumer");
  mkdirSync(consumer);
  writeFileSync(path.join(consumer, "package.json"), `${JSON.stringify({
    name: "artifact-sdk-package-consumer",
    private: true,
    type: "module",
  }, null, 2)}\n`);
  writeFileSync(path.join(consumer, "tsconfig.json"), `${JSON.stringify({
    compilerOptions: {
      target: "ES2022",
      module: "NodeNext",
      moduleResolution: "NodeNext",
      strict: true,
      exactOptionalPropertyTypes: true,
      noUncheckedIndexedAccess: true,
      noEmit: true,
    },
    include: ["consumer.ts"],
  }, null, 2)}\n`);
  writeFileSync(path.join(consumer, "consumer.ts"), `import {
  ARTIFACT_PROTOCOL_VERSION,
  ArtifactEditor,
  ArtifactSpecEditor,
  type Artifact,
  type ArtifactSpecification,
  type ArtifactLink,
  type ArtifactSource,
  type ArtifactSpec,
  type ArtifactSubmission,
  type ArtifactVerification,
  type ArtifactVersion,
  type WireArtifact,
} from "@elqora/artifacts";

const artifact = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "artifact_1",
  specId: "spec_1",
  kind: "delivery_evidence",
  valueType: "image",
  createdBy: { type: "user", id: "user_1" },
  createdAt: "2026-08-15T09:00:00Z",
  updatedAt: "2026-08-15T09:00:00Z",
  specification: {
    schema: "image-subject",
    version: 1,
    value: { subject: "Delivery entrance" },
  },
} satisfies Artifact<"delivery_evidence", "image", { subject: string }>;

const artifactSpecification = artifact.specification satisfies ArtifactSpecification<{ subject: string }>;

const source = { type: "url", url: "https://example.test/evidence.png" } satisfies ArtifactSource;
const version = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "version_1",
  artifactId: artifact.id,
  version: 1,
  source,
  createdBy: { type: "user", id: "user_1" },
  createdAt: "2026-08-15T09:00:00Z",
  specification: {
    schema: "image-region",
    version: 1,
    value: { x: 120, y: 45, width: 310, height: 180 },
  },
} satisfies ArtifactVersion<typeof source, { x: number; y: number; width: number; height: number }>;
const link = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "link_1",
  artifactId: artifact.id,
  artifactVersionId: version.id,
  subject: { type: "task", id: "task_1" },
  role: "evidence",
  createdBy: { type: "user", id: "user_1" },
  createdAt: "2026-08-15T09:00:00Z",
} satisfies ArtifactLink<"evidence", "task">;
const spec = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "spec_1",
  key: "evidence",
  name: "Evidence",
  version: 1,
  kind: "delivery_evidence",
  valueType: "image",
} satisfies ArtifactSpec<"delivery_evidence", "image">;
const opened = ArtifactEditor.open(artifact, { specs: [spec] });
if (!opened.ok) throw opened.error;
const editResult = opened.value.transaction((tx) => tx.setTitle("Edited evidence"));
if (!editResult.ok) throw editResult.error;
const specEditor = ArtifactSpecEditor.open(spec);
const specEditResult = specEditor.transaction((tx) => tx.setName("Edited evidence spec"));
if (!specEditResult.ok) throw specEditResult.error;
const submission = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "submission_1",
  artifactId: artifact.id,
  artifactVersionId: version.id,
  submittedBy: { type: "user", id: "user_1" },
  submittedAt: "2026-08-15T09:01:00Z",
} satisfies ArtifactSubmission;
const verification = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "verification_1",
  artifactId: artifact.id,
  artifactVersionId: version.id,
  submissionId: submission.id,
  status: "verified",
  createdAt: "2026-08-15T09:02:00Z",
} satisfies ArtifactVerification;
const wire: WireArtifact = artifact;
void [artifactSpecification, link, spec, verification, wire, opened, specEditor];
`);
  writeFileSync(path.join(consumer, "runtime.mjs"), `import {
  ARTIFACT_PROTOCOL_VERSION,
  ARTIFACT_SOURCE_TYPES,
  ARTIFACT_VALUE_TYPES,
} from "@elqora/artifacts";
import assert from "node:assert/strict";

assert.equal(ARTIFACT_PROTOCOL_VERSION, "1.1");
assert(ARTIFACT_SOURCE_TYPES.includes("provider"));
assert(ARTIFACT_VALUE_TYPES.includes("collection"));
`);

  const tarball = path.join(scratch, packed.filename);
  run("npm", ["install", "--ignore-scripts", "--no-package-lock", tarball], consumer);
  run(process.execPath, [path.join(root, "node_modules", "typescript", "bin", "tsc"), "-p", "tsconfig.json"], consumer);
  run(process.execPath, ["runtime.mjs"], consumer);

  const installed = path.join(consumer, "node_modules", "@elqora", "artifacts");
  const mapFiles = filesBelow(path.join(installed, "dist")).filter((filePath) => filePath.endsWith(".map"));
  assert(mapFiles.length > 0, "the installed package has no source maps");
  for (const mapFile of mapFiles) {
    const sourceMap = JSON.parse(readFileSync(path.join(installed, "dist", mapFile), "utf8"));
    assert(
      Array.isArray(sourceMap.sourcesContent) && sourceMap.sourcesContent.some((source) => typeof source === "string"),
      `${mapFile} does not embed its source content`,
    );
  }

  process.stdout.write("@elqora/artifacts tarball, clean install, types, runtime exports, and source maps passed.\n");
} finally {
  rmSync(scratch, { recursive: true, force: true });
}
