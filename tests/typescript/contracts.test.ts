import {
  ARTIFACT_PROTOCOL_VERSION,
  type Artifact,
  type ArtifactLink,
  type ArtifactSource,
  type ArtifactValueType,
  type ArtifactVersion,
  type ProviderArtifactSource,
  type ArtifactCondition,
  type ArtifactSpec,
  type ArtifactSpecSnapshot,
  type ArtifactSubmission,
  type ArtifactVerification,
  type ArtifactValidationPolicy,
  type WireArtifact,
  type WireArtifactSpec,
  type WireArtifactSubmission,
  type WireArtifactVerification,
  type WireProviderArtifactSource,
} from "../../packages/typescript/src/index.js";

const actor = { type: "host_actor", id: "actor_1" };

const artifact = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "art_1",
  specId: "artspec_host_defined",
  kind: "host_defined_kind",
  valueType: "reference",
  createdBy: actor,
  createdAt: "2026-08-15T09:00:00Z",
  updatedAt: "2026-08-15T09:00:00Z",
} satisfies Artifact<"host_defined_kind", "reference">;

const providerSource = {
  type: "provider",
  provider: "host_provider",
  reference: { resource: "change", id: 42 },
} satisfies ProviderArtifactSource<
  "host_provider",
  { resource: "change"; id: number }
>;

const version = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "artver_1",
  artifactId: artifact.id,
  version: 1,
  source: providerSource,
  createdBy: actor,
  createdAt: "2026-08-15T09:00:00Z",
} satisfies ArtifactVersion<typeof providerSource>;

const pinnedLink = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "alink_1",
  artifactId: artifact.id,
  artifactVersionId: version.id,
  subject: { type: "host_subject", id: "subject_1" },
  role: "host_role",
  createdBy: actor,
  createdAt: "2026-08-15T09:00:00Z",
} satisfies ArtifactLink<"host_role", "host_subject">;

function sourceIdentity(source: ArtifactSource): string {
  switch (source.type) {
    case "inline":
      return String(source.value);
    case "local":
      return source.localId;
    case "object":
      return source.objectId;
    case "url":
      return source.url;
    case "hosted":
      return `${source.recordType}:${source.recordId}`;
    case "provider":
      return source.provider;
  }
}

type Assert<T extends true> = T;
type InvalidValueTypeIsRejected = Assert<
  "host_defined_value_type" extends ArtifactValueType ? false : true
>;

void pinnedLink;
void sourceIdentity;
const typeAssertion: InvalidValueTypeIsRejected = true;
void typeAssertion;

const nestedCondition = {
  kind: "and",
  conditions: [
    { kind: "state", namespace: "workflow", in: ["active"] },
    { kind: "not", condition: { kind: "artifact_exists", artifact: "waiver" } },
  ],
} satisfies ArtifactCondition;

const imageSpec = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "artspec_photo",
  key: "photo",
  name: "Photo",
  version: 3,
  kind: "evidence_photo",
  valueType: "image",
  config: { valueType: "image", minFiles: 1, requireLocation: true },
  requirement: { mode: "conditional", condition: nestedCondition },
  privacy: {
    classification: "restricted",
    reveal: [{ actors: ["reviewer"], representation: "masked" }],
    encryption: { required: true },
  },
  validation: {
    schema: {
      locations: "required|object",
      "locations.name": ["required", "string"],
    },
  } satisfies ArtifactValidationPolicy,
} satisfies ArtifactSpec<"evidence_photo", "image">;

const { id: sourceSpecId, version: sourceVersion, ...snapshotDefinition } = imageSpec;
const snapshot = {
  ...snapshotDefinition,
  sourceSpecId,
  sourceVersion,
} satisfies ArtifactSpecSnapshot<"evidence_photo", "image">;

const submission = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "sub_1",
  artifactId: artifact.id,
  artifactVersionId: version.id,
  submittedBy: actor,
  submittedAt: "2026-08-15T09:10:00Z",
  context: { latitude: 6.5244, longitude: 3.3792 },
} satisfies ArtifactSubmission;

const verification = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "ver_1",
  artifactId: submission.artifactId,
  artifactVersionId: submission.artifactVersionId,
  submissionId: submission.id,
  status: "verified",
  method: "host_method",
  createdAt: "2026-08-15T09:11:00Z",
} satisfies ArtifactVerification;

void snapshot;
void verification;

// The ergonomic generic facade must remain assignable to schema-generated wire contracts.
const generatedArtifact: WireArtifact = artifact;
const generatedProviderSource: WireProviderArtifactSource = providerSource;
const generatedSpec: WireArtifactSpec = imageSpec;
const generatedSubmission: WireArtifactSubmission = submission;
const generatedVerification: WireArtifactVerification = verification;
void generatedArtifact;
void generatedProviderSource;
void generatedSpec;
void generatedSubmission;
void generatedVerification;
