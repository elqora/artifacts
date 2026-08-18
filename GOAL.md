# Artifact SDK — Language-Neutral Specification and Implementation Design

This document is the authoritative architecture and implementation brief for the
Artifact SDK. It defines the shared, language-neutral artifact protocol and the
boundaries between the SDK and host applications.

## Contents

- [Part I — Foundations](#part-i--foundations)
- [Part II — Core model](#part-ii--core-model)
- [Part III — Records, sources, and integrity](#part-iii--records-sources-and-integrity)
- [Part IV — Relationships](#part-iv--relationships)
- [Part V — Specifications and policies](#part-v--specifications-and-policies)
- [Part VI — Runtime records](#part-vi--runtime-records)
- [Part VII — Cross-host examples](#part-vii--cross-host-examples)
- [Part VIII — Protocol operation and evolution](#part-viii--protocol-operation-and-evolution)
- [Part IX — Implementer reference](#part-ix--implementer-reference)

## Part I — Foundations

### 1. Purpose

The Artifact SDK is intended to define a language-neutral artifact protocol that can be reused across unrelated applications.

It is not an artifact storage service, UI package, workflow engine, ORM layer, upload service, or application-specific implementation.

The SDK defines a canonical vocabulary and data model for describing:

- artifacts;
- artifact versions;
- artifact sources;
- artifact relationships;
- artifact specifications;
- artifact requirements;
- submissions;
- verification;
- integrity;
- access and privacy policies;
- lifecycle constraints;
- validation;
- retention;
- extensible host-specific semantics.

The same artifact specification must be usable by radically different systems.

Examples include:

- Errands and task marketplaces;
- project-management systems;
- freelancer platforms;
- approval workflows;
- compliance systems;
- audit systems;
- issue trackers;
- CRM systems;
- legal evidence systems;
- insurance claims;
- education systems;
- content-production systems;
- software-development tooling;
- delivery and fulfillment platforms.

The central abstraction is:

> An artifact is a typed, identifiable, versionable piece of information,
> evidence, output, or reference that can participate in relationships and
> workflows.

The SDK defines what these records mean.

The host application defines what happens because of them.

---

### 2. Architectural Principle

The architecture should follow this rule:

```text
Artifact SDK
    = shared language and contracts

Host Application
    = behavior and implementation
```

The SDK should not know what an Errand, Milestone, Review, Project, Delivery, Challenge, or Approval is.

Applications provide those concepts through extensible identifiers and subject references.

The protocol therefore standardizes structure, while allowing applications to define their own vocabulary.

---

### 3. Language-Neutral Design

TypeScript is used throughout this document as the reference notation because it makes the contracts easy to understand.

TypeScript must not be considered the authoritative implementation.

The authoritative representation is a language-neutral schema such as JSON Schema or an equivalent schema format. Language bindings are derived from that canonical representation.

Conceptually:

Canonical Artifact Specification
              │
              ▼
       Machine-readable schemas
              │
     ┌────────┼────────┬─────────┐
     ▼        ▼        ▼         ▼
 TypeScript   PHP      Go      Python
     SDK      SDK      SDK       SDK

Generated SDKs should expose equivalent domain contracts.

Generated code should not be manually edited.

Changes should originate from the canonical specification and regenerate language bindings.

---

## Part II — Core model

### 4. Core Model

The artifact model must not collapse every concern into one giant record.

The primary objects are:

ArtifactSpec
    ↓ describes or constrains

Artifact
    ↓ has

ArtifactVersion
    ↓ references

ArtifactSource

Relationships are separate:

Artifact
    ↓

ArtifactLink
    ↓

ArtifactSubjectReference

Operational activity is also separate:

Artifact
    ↓

ArtifactSubmission
    ↓

ArtifactVerification

The complete conceptual relationship is:

                           ArtifactSpec
                                │
                                │ describes / governs
                                ▼
                             Artifact
                                │
              ┌─────────────────┼─────────────────┐
              │                 │                 │
              ▼                 ▼                 ▼
      ArtifactVersion      ArtifactLink     ArtifactSubmission
              │                 │                 │
              ▼                 ▼                 ▼
      ArtifactSource     SubjectReference    Verification
              │
              ▼
      ArtifactIntegrity

---

### 5. Identifier Types

All entity identifiers should be opaque strings.

```typescript
export type ArtifactId = string;
export type ArtifactVersionId = string;
export type ArtifactSpecId = string;
export type ArtifactLinkId = string;
export type ArtifactSubmissionId = string;
export type ArtifactVerificationId = string;
```

The protocol should not prescribe whether a host uses:

- UUID;
- ULID;
- Snowflake;
- database IDs;
- provider IDs;
- another stable identifier system.

The host controls ID generation.

This is also important for offline-first environments, where objects may need IDs before they have reached a server.

---

### 6. Protocol Versioning

Protocol versioning must be distinguished from artifact-definition, artifact-content,
and instance-semantic-specification versions.

These represent four separate concepts.

Protocol Version
    = version of the Artifact SDK/schema itself

ArtifactSpec Version
    = revision of an artifact definition

ArtifactVersion
    = revision of the actual artifact content

ArtifactSpecification Version
    = revision of the schema that interprets instance semantic data

Every base protocol record exposes:

```typescript
export interface ArtifactProtocolRecord {
  schemaVersion: string;
}
```

Example:

```json
{
  schemaVersion: "1.1"
}
```

Changing an artifact's content must not alter the protocol version.

Changing an artifact definition or instance semantic specification must not alter
the protocol version.

---

### 7. Actor References

The Artifact SDK must not depend on a host application's "User" model.

Artifact operations therefore use generic actor references.

```typescript
export interface ActorReference<
  TType extends string = string
> {
  type: TType;

  id?: string;

  displayName?: string;
}
```

Examples:

```typescript
const githubActor: ActorReference = {
  type: "github_user",
  id: "12345678",
  displayName: "Davy",
};
```

```typescript
const runner: ActorReference = {
  type: "runner",
  id: "runner_123",
};
```

```typescript
const system: ActorReference = {
  type: "system",
};
```

The SDK does not interpret these actor types.

---

### 8. Artifact Scope

Artifacts may optionally exist inside a broad application-level scope.

```typescript
export interface ArtifactScopeReference {
  type: string;
  id: string;
}
```

Examples:

```json
{
  type: "project",
  id: "prj_123"
}
```

```json
{
  type: "errand",
  id: "err_456"
}
```

The scope is not the same thing as an "ArtifactLink".

Scope describes the broad owning boundary.

Links describe specific relationships and usages.

---

### 9. Artifact Kind vs Artifact Value Type

The protocol must distinguish semantic meaning from representation.

These are different concepts.

For example:

delivery_evidence

is a semantic meaning.

image

is a representation.

Similarly:

implementation

might be represented as:

reference

pointing at a GitHub pull request.

The SDK therefore separates:

Artifact.kind

from:

Artifact.valueType

---

### 10. Artifact Kind

"kind" represents what the artifact means in the consuming domain.

It should remain open.

```typescript
export type ArtifactKind = string;
```

Examples:

delivery_evidence
purchase_receipt
implementation
design_mockup
contract
identity_document
proof_of_delivery
handover_document
review_material
invoice
completion_evidence

The SDK should not define an exhaustive enum for these values.

Applications own their artifact vocabulary.

---

### 11. Artifact Value Type

The value type represents the fundamental representation or capability.

This vocabulary should be more standardized.

```typescript
export type ArtifactValueType =
  | "text"
  | "number"
  | "boolean"
  | "currency"
  | "date"
  | "datetime"
  | "time"
  | "location"
  | "file"
  | "image"
  | "video"
  | "audio"
  | "link"
  | "structured"
  | "reference"
  | "signature"
  | "collection";
```

Examples:

```json
{
  kind: "delivery_evidence",
  valueType: "image"
}
```

```json
{
  kind: "implementation",
  valueType: "reference"
}
```

```json
{
  kind: "purchase_receipt",
  valueType: "structured"
}
```

```json
{
  kind: "client_signature",
  valueType: "signature"
}
```

---

## Part III — Records, sources, and integrity

### 12. Artifact

"Artifact" represents the logical identity of the thing.

It should not directly contain provider-specific content.

```typescript
export interface Artifact<
  TKind extends string = string,
  TValueType extends string = ArtifactValueType,
  TSpecification = unknown
> {
  id: ArtifactId;

  specId: ArtifactSpecId;

  scope?: ArtifactScopeReference;

  kind: TKind;

  valueType: TValueType;

  title?: string;

  description?: string;

  specification?: ArtifactSpecification<TSpecification>;

  currentVersionId?: ArtifactVersionId;

  createdBy: ActorReference;

  createdAt: string;

  updatedAt: string;

  archivedAt?: string;

  metadata?: ArtifactMetadata;
}
```

`specId` is required and is the authoritative link to the `ArtifactSpec` that
defines the artifact. `kind` and `valueType` remain serialized cached
projections for efficient consumption, but they must match the selected
specification. Hosts and editors must change them by selecting a specification,
not by editing them independently.

An Artifact is the logical object.

Example:

Authentication demonstration

The physical MP4, YouTube link, or provider reference is represented by an "ArtifactVersion".

`ArtifactSpecification` is schema-identified semantic interpretation data for
an actual artifact. It is distinct from `ArtifactSpec`, which describes an
expected artifact and policy. Artifact-level specifications describe stable
logical meaning; version-level specifications describe one content revision.

```typescript
export interface ArtifactSpecification<T = unknown> {
  schema: string;
  version: number;
  value: T;
}
```

---

### 13. Artifact Metadata

The protocol needs an extension mechanism.

```typescript
export type ArtifactMetadata =
  Record<string, unknown>;
```

Metadata should not become a dumping ground.

If a property becomes important to interoperability or core behavior, it should become a formal protocol field.

Use metadata for host extensions and non-core information.

---

### 14. ArtifactVersion

Artifacts can change over time.

The logical identity should remain stable while physical content changes through versions.

```typescript
export interface ArtifactVersion {
  id: ArtifactVersionId;

  artifactId: ArtifactId;

  version: number;

  source: ArtifactSource;

  integrity?: ArtifactIntegrity;

  specification?: ArtifactSpecification;

  createdBy: ActorReference;

  createdAt: string;

  note?: string;

  metadata?: ArtifactMetadata;
}
```

Example:

Homepage Design
├── v1
├── v2
├── v3
└── v4

The Artifact remains:

Homepage Design

while each revision has a different ArtifactVersion.

This is critical for immutable approval and evidence history.

---

### 15. ArtifactSource

An "ArtifactSource" describes where the content or reference represented by a particular version comes from.

The core SDK should support broad source categories.

```typescript
export type ArtifactSource =
  | InlineArtifactSource
  | LocalArtifactSource
  | ObjectArtifactSource
  | UrlArtifactSource
  | HostedArtifactSource
  | ProviderArtifactSource;
```

---

### 16. InlineArtifactSource

Inline values are useful for text, structured data, numbers, coordinates, declarations, and similar values.

```typescript
export interface InlineArtifactSource {
  type: "inline";

  value: unknown;

  mediaType?: string;
}
```

Example:

```json
{
  type: "inline",
  value: "Package delivered successfully."
}
```

Structured example:

```json
{
  type: "inline",
  value: {
    merchant: "Example Store",
    amount: 12500,
    currency: "NGN",
    transactionReference: "TX-1289"
  }
}
```

---

### 17. LocalArtifactSource

Local sources support offline-first systems.

```typescript
export interface LocalArtifactSource {
  type: "local";

  localId: string;

  filename?: string;

  mediaType?: string;

  size?: number;

  syncState:
    | "local_only"
    | "pending_upload"
    | "uploading"
    | "uploaded"
    | "failed";

  remoteVersionId?: ArtifactVersionId;
}
```

A host can therefore create artifacts before network upload succeeds.

The Artifact ID remains stable.

---

### 18. ObjectArtifactSource

Uploaded objects should use a storage-neutral representation.

```typescript
export interface ObjectArtifactSource {
  type: "object";

  objectId: string;

  filename?: string;

  mediaType?: string;

  size?: number;

  storageProvider?: string;
}
```

The SDK should not assume:

Amazon S3
Cloudflare R2
MinIO
Azure Blob Storage
Google Cloud Storage

The host decides how "objectId" maps to physical storage.

---

### 19. UrlArtifactSource

External web resources can be represented as URLs.

```typescript
export interface UrlArtifactSource {
  type: "url";

  url: string;

  provider?: string;

  mediaType?: string;
}
```

Possible examples:

- YouTube;
- Vimeo;
- Google Drive;
- Dropbox;
- public file URLs;
- documentation pages;
- external evidence pages.

---

### 20. HostedArtifactSource

A host application may manage some artifacts internally.

```typescript
export interface HostedArtifactSource {
  type: "hosted";

  recordType: string;

  recordId: string;
}
```

Example:

```json
{
  type: "hosted",
  recordType: "generated_report",
  recordId: "report_983"
}
```

---

### 21. ProviderArtifactSource

External providers should not be hardcoded into the global SDK.

The protocol should expose a generic provider source.

```typescript
export interface ProviderArtifactSource<
  TProvider extends string = string,
  TReference = unknown
> {
  type: "provider";

  provider: TProvider;

  reference: TReference;
}
```

Provider-specific packages can define strong reference types.

---

### 22. GitHub Provider Example

A GitHub extension could define:

```typescript
export type GitHubArtifactReference =
  | GitHubPullRequestReference
  | GitHubCommitReference
  | GitHubIssueReference
  | GitHubDiscussionReference
  | GitHubDeploymentReference
  | GitHubWorkflowRunReference
  | GitHubCheckRunReference;
```

Pull request:

```typescript
export interface GitHubPullRequestReference {
  resource: "pull_request";

  repositoryId: number;

  number: number;

  nodeId?: string;
}
```

Commit:

```typescript
export interface GitHubCommitReference {
  resource: "commit";

  repositoryId: number;

  sha: string;
}
```

Issue:

```typescript
export interface GitHubIssueReference {
  resource: "issue";

  repositoryId: number;

  number: number;

  nodeId?: string;
}
```

Usage:

```typescript
const source: ProviderArtifactSource<
  "github",
  GitHubArtifactReference
> = {
  type: "provider",

  provider: "github",

  reference: {
    resource: "pull_request",
    repositoryId: 982134,
    number: 57,
  },
};
```

GitLab support should not require modifying the Artifact core specification.

A GitLab extension should simply define another provider-reference contract.

---

### 23. ArtifactIntegrity

Artifacts used as evidence may require proof that their content has not changed.

```typescript
export type ArtifactIntegrityAlgorithm =
  | "sha256"
  | "sha384"
  | "sha512";
```

```typescript
export interface ArtifactIntegrity {
  algorithm: ArtifactIntegrityAlgorithm;

  hash: string;

  size?: number;

  verifiedAt?: string;
}
```

Integrity is especially important for:

- accepted deliverables;
- contract versions;
- approvals;
- handover packages;
- compliance records;
- audit evidence;
- payment evidence;
- signed documents;
- dispute evidence.

Approval of an artifact version should refer to exact content, not merely a mutable logical artifact.

---

## Part IV — Relationships

### 24. ArtifactLink

Being evidence, a deliverable, an attachment, or review material is not an intrinsic property of an artifact.

It is a relationship.

The SDK therefore uses "ArtifactLink".

```typescript
export interface ArtifactLink<
  TRole extends string = string,
  TSubjectType extends string = string
> {
  id: ArtifactLinkId;

  artifactId: ArtifactId;

  artifactVersionId?: ArtifactVersionId;

  subject: ArtifactSubjectReference<TSubjectType>;

  role: TRole;

  note?: string;

  createdBy: ActorReference;

  createdAt: string;

  metadata?: ArtifactMetadata;
}
```

An ArtifactLink may optionally pin a specific version.

If no version is specified, the relationship follows the logical artifact.

If a version is specified, the relationship refers to that exact content revision.

---

### 25. ArtifactSubjectReference

The Artifact SDK must not contain application-specific subject unions.

Instead:

```typescript
export interface ArtifactSubjectReference<
  TType extends string = string
> {
  type: TType;

  id: string;

  scope?: Record<string, string>;
}
```

Project Manager may define:

```typescript
export type ProjectArtifactSubjectType =
  | "project"
  | "version"
  | "milestone"
  | "criterion"
  | "deliverable"
  | "task"
  | "work_package"
  | "challenge"
  | "review"
  | "approval"
  | "discussion_message"
  | "meeting"
  | "contract";
```

Errands may define:

```typescript
export type ErrandsArtifactSubjectType =
  | "errand"
  | "assignment"
  | "delivery"
  | "completion"
  | "dispute"
  | "payment";
```

The global Artifact SDK needs to understand neither list.

---

### 26. Artifact Role

Artifact roles should also remain extensible.

```typescript
export type ArtifactRole = string;
```

Project Manager may define:

```typescript
export type ProjectArtifactRole =
  | "implementation"
  | "deliverable"
  | "evidence"
  | "verification"
  | "challenge_evidence"
  | "response_evidence"
  | "review_evidence"
  | "documentation"
  | "approval"
  | "handover"
  | "attachment"
  | "reference";
```

Errands might define:

```typescript
export type ErrandsArtifactRole =
  | "delivery_evidence"
  | "completion_evidence"
  | "purchase_evidence"
  | "dispute_evidence"
  | "identity_evidence";
```

The principle is:

«Standardize the relationship structure, not every possible relationship vocabulary.»

---

### 27. Example of Artifact Reuse

Suppose a video exists once:

Authentication demonstration

It may have one Artifact:

```typescript
const artifact: Artifact = {
  id: "art_auth_demo",

  specId: "artspec_authentication_demo",

  scope: {
    type: "project",
    id: "prj_123",
  },

  kind: "authentication_demo",

  valueType: "video",

  title: "Authentication demonstration",

  currentVersionId: "artver_auth_demo_1",

  createdBy: currentUser,

  createdAt: now,

  updatedAt: now,
};
```

It can then have multiple links:

```typescript
const deliverableLink: ArtifactLink = {
  id: "alink_1",

  artifactId: artifact.id,

  artifactVersionId: artifact.currentVersionId,

  subject: {
    type: "deliverable",
    id: "del_auth_demo",
  },

  role: "deliverable",

  createdBy: currentUser,

  createdAt: now,
};
```

```typescript
const criterionLink: ArtifactLink = {
  id: "alink_2",

  artifactId: artifact.id,

  artifactVersionId: artifact.currentVersionId,

  subject: {
    type: "criterion",
    id: "criterion_login",
  },

  role: "evidence",

  createdBy: currentUser,

  createdAt: now,
};
```

```typescript
const reviewLink: ArtifactLink = {
  id: "alink_3",

  artifactId: artifact.id,

  artifactVersionId: artifact.currentVersionId,

  subject: {
    type: "review",
    id: "review_12",
  },

  role: "review_evidence",

  createdBy: currentUser,

  createdAt: now,
};
```

The artifact does not change because its context changes.

---

## Part V — Specifications and policies

### 28. ArtifactSpec

Some applications need to declare what artifacts are expected before an actual artifact exists.

This concept comes from workflow-driven systems such as Errands.

The global protocol should model that separately as "ArtifactSpec".

```typescript
export interface ArtifactSpec<
  TKind extends string = string,
  TValueType extends string = ArtifactValueType,
  TConfig extends object = Record<string, unknown>
> {
  id: ArtifactSpecId;

  key: string;

  name: string;

  description?: string;

  version: number;

  kind: TKind;

  valueType: TValueType;

  config?: TConfig;

  provider?: ArtifactProviderPolicy;

  requirement?: ArtifactRequirementPolicy;

  lifecycle?: ArtifactLifecyclePolicy;

  access?: ArtifactAccessPolicy;

  privacy?: ArtifactPrivacyPolicy;

  validation?: ArtifactValidationPolicy;

  verification?: ArtifactVerificationPolicy;

  retention?: ArtifactRetentionPolicy;

  presentation?: ArtifactPresentationPolicy;

  metadata?: ArtifactMetadata;
}
```

"ArtifactSpec" is not an Artifact.

It describes what an expected artifact should look like and the policies governing it.

#### Policy responsibility boundaries

The properties of an `ArtifactSpec` have deliberately separate responsibilities:

| Property | Responsibility |
| --- | --- |
| `kind`, `valueType` | The semantic meaning and fundamental representation expected by the host. |
| `config` | Representation-specific shape and constraints, such as image count, text length, or collection item schema. |
| `validation` | Declarative checks a host applies when accepting a value or version, including host-defined validation rules. |
| `provider`, `requirement`, `lifecycle` | Who supplies the artifact, whether it is required, and when the host permits it to change. |
| `access`, `privacy`, `verification`, `retention` | Governance of access, disclosure, trust, and preservation. |
| `presentation` | Optional UI hints that do not affect protocol validity. |

`config` must not become a second policy bag. A host-specific rule belongs in its
dedicated policy when one exists; `metadata` remains the extension point for
non-interoperable information.

---

### 29. ArtifactSpecSnapshot

Definitions can evolve.

A running task or workflow should not silently change because an administrator edits an artifact specification later.

Therefore specifications need immutable snapshots.

```typescript
export interface ArtifactSpecSnapshot<
  TKind extends string = string,
  TValueType extends string = ArtifactValueType,
  TConfig extends object = Record<string, unknown>
> {
  sourceSpecId: ArtifactSpecId;

  sourceVersion: number;

  key: string;

  name: string;

  description?: string;

  kind: TKind;

  valueType: TValueType;

  config?: TConfig;

  provider?: ArtifactProviderPolicy;

  requirement?: ArtifactRequirementPolicy;

  lifecycle?: ArtifactLifecyclePolicy;

  access?: ArtifactAccessPolicy;

  privacy?: ArtifactPrivacyPolicy;

  validation?: ArtifactValidationPolicy;

  verification?: ArtifactVerificationPolicy;

  retention?: ArtifactRetentionPolicy;

  presentation?: ArtifactPresentationPolicy;

  metadata?: ArtifactMetadata;
}
```

Example:

delivery_evidence v3

may require:

1 image

while later:

delivery_evidence v4

may require:

3 images + GPS

Existing workflows based on v3 must continue using the v3 snapshot.

---

### 30. ArtifactProviderPolicy

Provider policy answers:

«Who is expected to supply this artifact?»

It must not be confused with `ProviderArtifactSource`: this policy identifies
the host-defined actor category expected to supply an artifact, whereas a
provider source identifies the external system from which a version originates.

```typescript
export interface ArtifactProviderPolicy {
  actors: string[];

  mode?: "single" | "any" | "all";

  delegation?: "forbidden" | "allowed";
}
```

Example:

```json
{
  actors: ["runner"],
  mode: "single"
}
```

Another application:

```json
{
  actors: ["developer", "designer"],
  mode: "any"
}
```

Actor identifiers remain host-defined.

---

### 31. ArtifactRequirementPolicy

Requiredness should not be reduced to:

required: true

because workflow-oriented systems need to know what an artifact requirement blocks.

```typescript
export interface ArtifactRequirementPolicy {
  mode:
    | "required"
    | "optional"
    | "conditional";

  condition?: ArtifactCondition;

  blocks?: string[];
}
```

Example:

```json
{
  mode: "required",

  blocks: [
    "submit_completion"
  ]
}
```

Project Manager might instead define:

```json
{
  mode: "required",

  blocks: [
    "request_milestone_approval"
  ]
}
```

The operation identifiers remain domain-specific strings.

---

### 32. ArtifactLifecyclePolicy

Lifecycle policy defines when an artifact may exist or change.

```typescript
export interface ArtifactLifecyclePolicy {
  createAt?: string;

  editableDuring?: string[];

  submitDuring?: string[];

  lockAt?: string;

  invalidateOn?: string[];

  condition?: ArtifactCondition;
}
```

Errands example:

```json
{
  createAt: "accepted",

  editableDuring: [
    "accepted",
    "in_progress"
  ],

  submitDuring: [
    "in_progress"
  ],

  lockAt: "completion_submitted"
}
```

Project Manager example:

```json
{
  createAt: "active",

  editableDuring: [
    "active",
    "changes_requested"
  ],

  submitDuring: [
    "active"
  ],

  lockAt: "approved"
}
```

The SDK does not interpret these stage names globally.

Each host must document the state namespace and the meaning of its stage values.
Lifecycle policies constrain artifact operations; they do not themselves create
or transition host workflow state.

---

### 33. ArtifactCondition

Conditions must be language-neutral and serializable.

Never encode policies using language-specific callbacks.

Do not do:

```typescript
condition: ctx => ctx.task.status === "accepted"

Instead, conditions should form a declarative AST.

export type ArtifactCondition =
  | ArtifactStateCondition
  | ArtifactActorCondition
  | ArtifactExistsCondition
  | ArtifactValueCondition
  | ArtifactAndCondition
  | ArtifactOrCondition
  | ArtifactNotCondition;
```

---

### 34. State Condition

```typescript
export interface ArtifactStateCondition {
  kind: "state";

  namespace: string;

  in: string[];
}
```

Example:

```json
{
  kind: "state",

  namespace: "errand",

  in: [
    "accepted",
    "in_progress"
  ]
}
```

---

### 35. Actor Condition

```typescript
export interface ArtifactActorCondition {
  kind: "actor";

  in: string[];
}
```

Example:

```json
{
  kind: "actor",
  in: ["runner", "admin"]
}
```

---

### 36. Artifact Exists Condition

```typescript
export interface ArtifactExistsCondition {
  kind: "artifact_exists";

  artifact: string;
}
```

The artifact reference may correspond to a spec key, artifact key, or another host-defined identifier.

The exact resolution rule should be documented by the protocol or extension layer.

A core implementation must therefore treat this value as an opaque reference.
An extension or host policy must define whether it resolves, for example, by
spec key, artifact key, logical artifact ID, or another namespace.

---

### 37. Artifact Value Condition

```typescript
export interface ArtifactValueCondition {
  kind: "artifact_value";

  artifact: string;

  operator:
    | "eq"
    | "neq"
    | "gt"
    | "gte"
    | "lt"
    | "lte"
    | "contains"
    | "in";

  value: unknown;
}
```

Example:

```json
{
  kind: "artifact_value",

  artifact: "item_unavailable",

  operator: "eq",

  value: true
}
```

---

### 38. Logical Conditions

```typescript
export interface ArtifactAndCondition {
  kind: "and";

  conditions: ArtifactCondition[];
}
```

```typescript
export interface ArtifactOrCondition {
  kind: "or";

  conditions: ArtifactCondition[];
}
```

```typescript
export interface ArtifactNotCondition {
  kind: "not";

  condition: ArtifactCondition;
}
```

Example:

```typescript
const condition: ArtifactCondition = {
  kind: "and",

  conditions: [
    {
      kind: "state",
      namespace: "errand",
      in: ["in_progress"],
    },
    {
      kind: "artifact_value",
      artifact: "item_unavailable",
      operator: "eq",
      value: true,
    },
  ],
};
```

This is safely serializable to JSON and interpretable in any implementation language.

---

### 39. ArtifactAccessPolicy

Access rules should remain distinct from privacy rules.

```typescript
export interface ArtifactAccessPolicy {
  read?: ArtifactAccessRule[];

  write?: ArtifactAccessRule[];

  submit?: ArtifactAccessRule[];

  verify?: ArtifactAccessRule[];
}
```

```typescript
export interface ArtifactAccessRule {
  actors: string[];

  condition?: ArtifactCondition;
}
```

The protocol declares access intent but does not prescribe authorization
semantics. Hosts must define their default decision, how multiple matching rules
combine, and how access rules interact with privacy representation rules.

Example:

```json
{
  read: [
    {
      actors: ["runner", "admin"]
    },
    {
      actors: ["client"],
      condition: {
        kind: "state",
        namespace: "errand",
        in: ["completion_submitted", "completed"]
      }
    }
  ]
}
```

---

### 40. ArtifactPrivacyPolicy

Access asks:

«Is this actor allowed to access the artifact?»

Privacy asks:

«What representation of the artifact is this actor allowed to receive?»

```typescript
export interface ArtifactPrivacyPolicy {
  classification:
    | "public"
    | "internal"
    | "private"
    | "sensitive"
    | "restricted";

  reveal?: ArtifactRevealRule[];

  masking?: ArtifactMaskingPolicy;

  encryption?: ArtifactEncryptionPolicy;
}
```

---

### 41. ArtifactRevealRule

```typescript
export interface ArtifactRevealRule {
  actors: string[];

  when?: ArtifactCondition;

  representation:
    | "hidden"
    | "masked"
    | "approximate"
    | "full";
}
```

This supports progressive disclosure.

When several reveal rules match, the host must apply a documented precedence
rule. In the absence of an applicable rule, the host must use its documented
default representation; the core protocol does not select one.

For example, a runner may see an approximate destination before accepting a job:

```json
{
  actors: ["runner"],

  when: {
    kind: "state",
    namespace: "errand",
    in: ["published", "offered"]
  },

  representation: "approximate"
}
```

After acceptance:

```json
{
  actors: ["runner"],

  when: {
    kind: "state",
    namespace: "errand",
    in: ["accepted", "in_progress"]
  },

  representation: "full"
}
```

---

### 42. ArtifactMaskingPolicy

A basic extensible form could be:

```typescript
export interface ArtifactMaskingPolicy {
  strategy: string;

  config?: Record<string, unknown>;
}
```

Potential future strategies:

partial
redact
truncate
hash
approximate_location
last_four
custom

The exact standardized strategies can evolve later.

---

### 43. ArtifactEncryptionPolicy

The protocol may describe encryption expectations without implementing encryption itself.

```typescript
export interface ArtifactEncryptionPolicy {
  required: boolean;

  level?: string;

  keyScope?: string;
}
```

The host decides how actual encryption is performed.

---

## Part VI — Runtime records

### 44. ArtifactSubmission

Artifact values should not necessarily be silently overwritten.

Operational submission history should be first-class.

```typescript
export interface ArtifactSubmission<
  TValue = unknown
> {
  id: ArtifactSubmissionId;

  artifactId: ArtifactId;

  artifactVersionId?: ArtifactVersionId;

  submittedBy: ActorReference;

  value?: TValue;

  submittedAt: string;

  context?: ArtifactSubmissionContext;

  metadata?: ArtifactMetadata;
}
```

A submission is operational history, not automatically a new content revision.
It may point at an immutable `ArtifactVersion`, contain a proposed value for a
host to process, or both. When a host treats submitted content as evidence or
otherwise needs historical immutability, it must create and reference an
`ArtifactVersion`.

---

### 45. ArtifactSubmissionContext

```typescript
export interface ArtifactSubmissionContext {
  latitude?: number;

  longitude?: number;

  deviceId?: string;

  ipAddress?: string;

  userAgent?: string;
}
```

This supports evidence-sensitive workflows where submission provenance matters.

The protocol defines the data shape.

The host decides which values are safe and legal to collect.

---

### 46. ArtifactVerification

Verification must not be represented as:

verified: true

Verification is an event/record with provenance.

```typescript
export interface ArtifactVerification {
  id: ArtifactVerificationId;

  artifactId: ArtifactId;

  artifactVersionId?: ArtifactVersionId;

  submissionId?: ArtifactSubmissionId;

  status:
    | "pending"
    | "verified"
    | "rejected"
    | "waived";

  method?: string;

  verifiedBy?: ActorReference;

  reason?: string;

  createdAt: string;

  verifiedAt?: string;

  metadata?: ArtifactMetadata;
}
```

Methods remain extensible.

Examples:

manual
client_confirmation
platform_review
automatic
external_provider
github_status
signature_check
human_review

---

### 47. ArtifactVerificationPolicy

Artifact specifications may define verification requirements.

```typescript
export interface ArtifactVerificationPolicy {
  required: boolean;

  methods?: string[];

  actors?: string[];

  condition?: ArtifactCondition;
}
```

Example:

```json
{
  required: true,

  methods: [
    "client_confirmation",
    "platform_review"
  ],

  actors: [
    "client",
    "platform"
  ]
}
```

The policy describes what is expected.

"ArtifactVerification" records what actually happened.

When `condition` is present, the host evaluates it before applying the policy.
If it evaluates to false, this verification policy does not apply for that
context; the host still records any verification events it performs.

---

### 48. ArtifactValidationPolicy

```typescript
export interface ArtifactValidationPolicy {
  mode?: "strict" | "lenient";

  schema?: ArtifactValidationSchema;

  rules?: ArtifactValidationRule[];
}
```

`schema` is a serializable Laravel-style rule map applied to the semantic value
inside `Artifact.specification` or `ArtifactVersion.specification`. It does not
validate the specification envelope (`schema` and `version`) or the physical
`ArtifactSource`.

```typescript
export type ArtifactValidationSchema = Record<string, string | string[]>;
```

The rule-map keys are exact field paths, including nested paths such as
`locations.name`. A value may be a pipe-form rule string or Bunwire's canonical
array form:

```typescript
const policy = {
  schema: {
    locations: "required|object",
    "locations.name": ["required", "string"],
  },
};
```

The host selects the applicable specification context and invokes Bunwire (or a
semantically equivalent implementation). Wildcards, executable callbacks, and
language-specific rule objects are not serialized in the protocol.

```typescript
export interface ArtifactValidationRule {
  type: string;

  config?: Record<string, unknown>;
}
```

Example:

```json
{
  type: "file_count",

  config: {
    min: 1,
    max: 4
  }
}
```

Another:

```json
{
  type: "mime_type",

  config: {
    allowed: [
      "image/jpeg",
      "image/png"
    ]
  }
}
```

The protocol may later standardize common validation-rule identifiers.

---

### 49. ArtifactRetentionPolicy

Retention should describe expectations, not storage implementation.

```typescript
export type ArtifactRetentionPolicy =
  | {
      policy: "forever";
    }
```
  | {
      policy: "duration";
      days: number;
    }
  | {
      policy: "until";
      date: string;
    }
  | {
      policy: "host_defined";
      key: string;
    };

Example:

```json
{
  policy: "duration",
  days: 365
}
```

---

### 50. ArtifactPresentationPolicy

Presentation hints can exist without coupling the SDK to a UI framework.

```typescript
export interface ArtifactPresentationPolicy {
  label?: string;

  helpText?: string;

  order?: number;

  display?: string;

  config?: Record<string, unknown>;
}
```

These should remain hints only.

The host may ignore them.

The protocol must never depend on React, Vue, Flutter, SwiftUI, or another presentation framework.

---

### 51. Artifact Value Schemas

Value-type-specific schemas can add stronger typing and interoperable constraints.

The following are illustrative value schemas. The canonical schema and generated
bindings define representation-specific schemas for every standardized
`ArtifactValueType`.

Example image schema:

```typescript
export interface ImageArtifactValueSchema {
  valueType: "image";

  minFiles?: number;

  maxFiles?: number;

  acceptedMimeTypes?: string[];

  maxSizeBytes?: number;

  requireTimestamp?: boolean;

  requireLocation?: boolean;
}
```

Location:

```typescript
export interface LocationArtifactValueSchema {
  valueType: "location";

  mode:
    | "point"
    | "address"
    | "point_and_address";

  requireCoordinates?: boolean;

  allowManualEntry?: boolean;
}
```

Text:

```typescript
export interface TextArtifactValueSchema {
  valueType: "text";

  minLength?: number;

  maxLength?: number;

  multiline?: boolean;

  pattern?: string;
}
```

File:

```typescript
export interface FileArtifactValueSchema {
  valueType: "file";

  minFiles?: number;

  maxFiles?: number;

  acceptedMimeTypes?: string[];

  maxSizeBytes?: number;
}
```

Collection:

```typescript
export interface CollectionArtifactValueSchema {
  valueType: "collection";

  itemSchema: ArtifactValueSchema;

  minItems?: number;

  maxItems?: number;
}
```

The value-schema union is extended as standardized value types gain their own
constraints:

```typescript
export type ArtifactValueSchema =
  | ImageArtifactValueSchema
  | LocationArtifactValueSchema
  | TextArtifactValueSchema
  | FileArtifactValueSchema
  | CollectionArtifactValueSchema;
```

Additional value types can be introduced over time.

---

### 52. ArtifactRequirement

Some domain objects may require artifacts without owning a full ArtifactSpec.

A generic requirement model can represent this.

```typescript
export interface ArtifactRequirement {
  id: string;

  key?: string;

  allowedKinds?: string[];

  allowedValueTypes?: ArtifactValueType[];

  minimumCount?: number;

  maximumCount?: number;

  required: boolean;

  specId?: ArtifactSpecId;

  metadata?: ArtifactMetadata;
}
```

Example:

```typescript
const requirement: ArtifactRequirement = {
  id: "req_demo",

  allowedKinds: [
    "authentication_demo"
  ],

  allowedValueTypes: [
    "video",
    "link"
  ],

  minimumCount: 1,

  required: true,
};
```

A project-management deliverable can use this without making the Artifact itself aware that it is a deliverable.

---

## Part VII — Cross-host examples

### 53. Project Manager Example

A developer submits a GitHub pull request.

Artifact:

```typescript
const artifact: Artifact = {
  id: "art_auth_impl",

  specId: "artspec_authentication_implementation",

  scope: {
    type: "project",
    id: "prj_123",
  },

  kind: "implementation",

  valueType: "reference",

  title: "Authentication implementation",

  currentVersionId: "artver_auth_impl_1",

  createdBy: {
    type: "github_user",
    id: "123456",
  },

  createdAt: "2026-08-15T09:00:00Z",

  updatedAt: "2026-08-15T09:00:00Z",
};
```

Version:

```typescript
const version: ArtifactVersion = {
  id: "artver_auth_impl_1",

  artifactId: artifact.id,

  version: 1,

  source: {
    type: "provider",

    provider: "github",

    reference: {
      resource: "pull_request",
      repositoryId: 845121,
      number: 58,
    },
  },

  createdBy: {
    type: "github_user",
    id: "123456",
  },

  createdAt: "2026-08-15T09:00:00Z",
};
```

Deliverable relationship:

```typescript
const deliverableLink: ArtifactLink = {
  id: "alink_deliverable",

  artifactId: artifact.id,

  artifactVersionId: version.id,

  subject: {
    type: "deliverable",
    id: "del_auth",
  },

  role: "deliverable",

  createdBy: currentUser,

  createdAt: now,
};
```

Review relationship:

```typescript
const reviewLink: ArtifactLink = {
  id: "alink_review",

  artifactId: artifact.id,

  artifactVersionId: version.id,

  subject: {
    type: "review",
    id: "review_12",
  },

  role: "review_evidence",

  createdBy: currentUser,

  createdAt: now,
};
```

The same artifact participates in multiple contexts without duplication.

---

### 54. Errands Example

An Errands workflow requires delivery photographic evidence.

Specification:

```typescript
const deliveryEvidenceSpec: ArtifactSpec = {
  id: "artspec_delivery_evidence",

  key: "delivery_evidence",

  name: "Delivery evidence",

  description:
    "Photographic evidence that the delivery was completed.",

  version: 3,

  kind: "delivery_evidence",

  valueType: "image",

  config: {
    minFiles: 1,
    maxFiles: 4,
    requireTimestamp: true,
    requireLocation: true,
  },

  provider: {
    actors: ["runner"],
    mode: "single",
  },

  requirement: {
    mode: "required",
    blocks: ["submit_completion"],
  },

  lifecycle: {
    createAt: "accepted",

    editableDuring: [
      "accepted",
      "in_progress"
    ],

    submitDuring: [
      "in_progress"
    ],

    lockAt: "completion_submitted",
  },

  access: {
    read: [
      {
        actors: [
          "runner",
          "platform",
          "admin"
        ]
      },
      {
        actors: ["client"],
        condition: {
          kind: "state",
          namespace: "errand",
          in: [
            "completion_submitted",
            "under_review",
            "completed"
          ]
        }
      }
    ],

    write: [
      {
        actors: ["runner"],
        condition: {
          kind: "state",
          namespace: "errand",
          in: ["in_progress"]
        }
      }
    ]
  },

  privacy: {
    classification: "private",
  },

  verification: {
    required: true,

    methods: [
      "client_confirmation",
      "platform_review"
    ],

    actors: [
      "client",
      "platform"
    ]
  },

  retention: {
    policy: "duration",
    days: 365
  }
};
```

Actual artifact:

```typescript
const artifact: Artifact = {
  id: "art_delivery_1",

  specId: "artspec_delivery_evidence",

  scope: {
    type: "errand",
    id: "err_100",
  },

  kind: "delivery_evidence",

  valueType: "image",

  title: "Delivery photograph",

  currentVersionId: "artver_delivery_1",

  createdBy: {
    type: "runner",
    id: "runner_28",
  },

  createdAt: now,

  updatedAt: now,
};
```

Artifact version:

```typescript
const version: ArtifactVersion = {
  id: "artver_delivery_1",

  artifactId: artifact.id,

  version: 1,

  source: {
    type: "object",

    objectId: "obj_9384",

    filename: "delivery.jpg",

    mediaType: "image/jpeg",

    size: 2_450_933,
  },

  integrity: {
    algorithm: "sha256",

    hash: "example-hash",

    size: 2_450_933,
  },

  createdBy: {
    type: "runner",
    id: "runner_28",
  },

  createdAt: now,
};
```

Submission:

```typescript
const submission: ArtifactSubmission = {
  id: "sub_1",

  artifactId: artifact.id,

  artifactVersionId: version.id,

  submittedBy: {
    type: "runner",
    id: "runner_28",
  },

  submittedAt: now,

  context: {
    latitude: 6.5244,
    longitude: 3.3792,
  },
};
```

Verification:

```typescript
const verification: ArtifactVerification = {
  id: "ver_1",

  artifactId: artifact.id,

  artifactVersionId: version.id,

  submissionId: submission.id,

  status: "verified",

  method: "client_confirmation",

  verifiedBy: {
    type: "client",
    id: "client_18",
  },

  createdAt: now,

  verifiedAt: now,
};
```

This proves that two fundamentally different applications can use the same artifact vocabulary.

---

## Part VIII — Protocol operation and evolution

### 55. What the SDK Must Not Implement

The core SDK must not implement application behavior.

It should not determine:

- how files are uploaded;
- where data is persisted;
- which database is used;
- how authentication works;
- how authorization is enforced;
- how storage credentials work;
- how GitHub APIs are called;
- how URLs are signed;
- how encryption is performed;
- how workflow states transition;
- how UI is rendered;
- how artifacts are displayed;
- how payment is released;
- how a milestone is approved;
- how an Errand is completed;
- how notifications are delivered;
- how disputes are resolved.

The protocol may describe policies related to these behaviors.

The host executes them.

---

### 56. Optional SDK Utilities

Language bindings may eventually provide deterministic helper utilities.

Examples:

```typescript
export function isArtifactArchived(
  artifact: Artifact
): boolean {
  return artifact.archivedAt !== undefined;
}
```

```typescript
export function isVersionPinned(
  link: ArtifactLink
): boolean {
  return link.artifactVersionId !== undefined;
}
```

Potential helpers include:

- validation against generated schema;
- serialization;
- deserialization;
- protocol-version compatibility checks;
- artifact-reference helpers;
- integrity verification helpers;
- condition-AST parsing;
- safe discriminated-union helpers.

The TypeScript package may also provide small programmatic editors for generic
artifact and specification content:

- `ArtifactEditor`;
- `ArtifactSpecEditor`; and
- shared callback transactions with local undo/redo history.

These editors operate on in-memory protocol data only. They do not persist
changes, execute validation policies, advance specification versions, publish
definitions, or implement host workflows.

These utilities are optional conveniences.

They must not turn the SDK into a host-specific runtime.

---

### 57. Open vs Closed Vocabularies

The protocol must deliberately distinguish standardized concepts from extensible application concepts.

Good candidates for closed or centrally managed protocol vocabularies

ArtifactValueType
ArtifactSource.type
ArtifactIntegrityAlgorithm
Condition AST node kinds
verification status
synchronization status
privacy representations

Concepts that should remain extensible

Artifact.kind
ArtifactLink.role
ArtifactSubjectReference.type
ArtifactProviderPolicy.actors
ArtifactLifecyclePolicy stages
ArtifactRequirementPolicy.blocks
ArtifactVerification.method
host-specific metadata
provider identifiers
provider resource identifiers

This prevents the Artifact SDK from becoming a catalogue of every possible business domain.

---

### 58. Extension Packages

Provider and domain-specific contracts are separate from the core protocol. The
GitHub extension is the first implementation of this boundary.

Potential structure:

artifact-sdk/
├── spec/
├── typescript/
├── php/
├── go/
└── docs/

artifact-github/
artifact-gitlab/
artifact-figma/
artifact-jira/
artifact-drive/

A provider package can supply:

- typed provider references;
- optional provider-specific metadata schemas;
- optional validation rules;
- convenience adapters.

The Artifact core should remain provider-neutral.

---

### 59. Current repository layout

The repository is currently organised around the following protocol boundaries:

```text
artifacts/
├── README.md
├── GOAL.md
├── LICENSE
│
├── spec/
│   ├── protocol/
│   │   └── protocol.schema.json
│   │
│   ├── artifact/
│   │   ├── artifact.schema.json
│   │   ├── artifact-specification.schema.json
│   │   ├── artifact-version.schema.json
│   │   ├── artifact-source.schema.json
│   │   ├── artifact-integrity.schema.json
│   │   └── artifact-link.schema.json
│   │
│   ├── specification/
│   │   ├── artifact-spec.schema.json
│   │   ├── artifact-spec-snapshot.schema.json
│   │   ├── artifact-requirement.schema.json
│   │   └── artifact-value-schema.schema.json
│   │
│   ├── policy/
│   │   ├── condition.schema.json
│   │   ├── provider-policy.schema.json
│   │   ├── requirement-policy.schema.json
│   │   ├── lifecycle-policy.schema.json
│   │   ├── access-policy.schema.json
│   │   ├── privacy-policy.schema.json
│   │   ├── validation-policy.schema.json
│   │   ├── verification-policy.schema.json
│   │   ├── retention-policy.schema.json
│   │   └── presentation-policy.schema.json
│   │
│   └── runtime/
│       ├── submission.schema.json
│       └── verification.schema.json
│
├── packages/
│   ├── typescript/
│   ├── php/
│   └── go/
│
├── extensions/
│   └── github/
│
├── generators/
├── scripts/
│
├── tests/
│   ├── fixtures/
│   ├── conformance/
│   ├── generator/
│   ├── php/
│   └── typescript/
│
└── docs/
    ├── architecture.md
    ├── terminology.md
    ├── versioning.md
    ├── extensions.md
    ├── conditions.md
    └── examples/
        ├── errands.md
        └── project-manager.md
```

The exact repository structure may change, but the conceptual boundaries should remain.

---

### 60. Conformance Fixtures

The repository contains canonical JSON fixtures used by the conformance suite
and language-package tests.

Example:

```text
tests/fixtures/artifact/basic.json
tests/fixtures/artifact/versioned-file.json
tests/fixtures/artifact/github-reference.json
tests/fixtures/spec/delivery-evidence.json
tests/fixtures/submission/delivery-photo.json
tests/fixtures/verification/client-confirmation.json
```

All generated or manually implemented language SDKs should be able to deserialize and validate the same fixtures.

This gives the project actual cross-language conformance rather than merely similar-looking interfaces.

---

### 61. Serialization Rules

Serialized forms should use stable discriminators.

For example:

```json
{
  "type": "provider",
  "provider": "github",
  "reference": {
    "resource": "pull_request",
    "repositoryId": 123,
    "number": 55
  }
}
```

Avoid representations whose meaning depends on language-specific class names.

Dates use the canonical RFC 3339 date-time representation defined by the protocol schema. Producers should emit UTC where practical.

Optional fields should have clear missing/null semantics.

The canonical schemas and serialization guidance define:

- whether absent and "null" mean the same thing;
- number precision expectations;
- date/time formats;
- identifier case sensitivity;
- enum case sensitivity;
- unknown-field behavior;
- forward-compatible parsing behavior.

---

### 62. Forward Compatibility

The protocol should be designed so future SDKs can safely encounter records created by newer versions.

Where possible:

- additions should be backward-compatible;
- unknown metadata should be preserved;
- extension identifiers should not cause parser failure;
- provider-specific references should remain opaque to implementations that do not understand them.

Breaking structural changes should require a protocol-major-version change.

---

### 63. Security Boundary

The Artifact SDK describes security-related intent but must not create false guarantees.

For example:

```json
privacy: {
  classification: "restricted"
}
```

does not itself make content restricted.

The host must enforce the policy.

Likewise:

```json
encryption: {
  required: true
}
```

does not encrypt anything.

The documentation must explicitly state:

«Policy records are declarations. Security exists only when a host correctly enforces them.»

This distinction must remain clear throughout the SDK.

---

### 64. Implementation status and extension priorities

The repository now contains the stable shared core, canonical schemas,
generated TypeScript/PHP/Go bindings, conformance fixtures, and an initial
GitHub extension. The staged order below remains the required order for new
protocol work and for implementations in another language.

Recommended order:

Phase 1 — Core Specification

Implement:

Artifact
ArtifactVersion
ArtifactSource
ArtifactIntegrity
ArtifactLink
ArtifactSubjectReference
ActorReference
ArtifactMetadata

Define JSON schemas.

Provide TypeScript contracts.

Add canonical fixtures.

---

Phase 2 — Artifact Specifications

Implement:

ArtifactSpec
ArtifactSpecSnapshot
ArtifactRequirement
ArtifactValueType
ArtifactValueSchema

Add schema-version and definition-version rules.

---

Phase 3 — Runtime Records

Implement:

ArtifactSubmission
ArtifactSubmissionContext
ArtifactVerification

Ensure version-pinned evidence is supported.

---

Phase 4 — Declarative Policies

Implement:

ArtifactCondition
ArtifactProviderPolicy
ArtifactRequirementPolicy
ArtifactLifecyclePolicy
ArtifactAccessPolicy
ArtifactPrivacyPolicy
ArtifactValidationPolicy
ArtifactVerificationPolicy
ArtifactRetentionPolicy
ArtifactPresentationPolicy

The protocol does not require a full policy engine.

The schema can exist before every language provides evaluators.

---

Phase 5 — Additional SDK Bindings

Generate or implement bindings for:

TypeScript
PHP
Go

Additional languages may follow.

Each language must conform to the same fixtures.

---

Phase 6 — Provider Extensions

Add providers only when they are actually needed.

For example:

GitHub

Do not prematurely add dozens of provider definitions.

Provider extensions should prove the extension mechanism rather than expand the Artifact core.

---

## Part IX — Implementer reference

### 65. Primary Design Rules for Implementers

The implementation should preserve these rules.

#### Rule 1 — Separate identity from content

Artifact identity must remain separate from artifact content.

Artifact != ArtifactVersion

#### Rule 2 — Keep physical references in sources

Provider-specific physical references belong to ArtifactSource.

ArtifactVersion → ArtifactSource

#### Rule 3 — Model usage as a relationship

Usage is a relationship.

Do not add fields such as:

isDeliverable
isEvidence
reviewId
milestoneId

to Artifact.

Use ArtifactLink.

#### Rule 4 — Support version pinning

Version-sensitive relationships must support version pinning.

Approval of v3 must not automatically apply to v4.

#### Rule 5 — Keep definitions separate from artifacts

Artifact specifications are not actual artifacts.

ArtifactSpec != Artifact

#### Rule 6 — Keep definitions serializable

Definitions must be serializable.

Do not store executable callbacks in ArtifactSpec or policy definitions.

#### Rule 7 — Preserve extensible vocabulary

Domain vocabulary remains extensible.

Do not hardcode application concepts such as:

milestone
errand
runner
review
submit_completion

into the Artifact core specification.

#### Rule 8 — Do not implement host workflows

The SDK does not implement host workflows.

#### Rule 9 — Treat security policies as declarations

Security policies are declarations that hosts must enforce.

#### Rule 10 — Preserve cross-language semantics

All language SDKs must represent the same protocol semantics.

---

### 66. Final Domain Vocabulary

The initial Artifact SDK should revolve around these terms:

Artifact
ArtifactVersion
ArtifactSource
ArtifactIntegrity

ArtifactLink
ArtifactSubjectReference
ArtifactRole

ArtifactSpec
ArtifactSpecSnapshot
ArtifactRequirement
ArtifactValueType
ArtifactValueSchema

ArtifactSubmission
ArtifactSubmissionContext
ArtifactVerification

ArtifactProviderPolicy
ArtifactRequirementPolicy
ArtifactLifecyclePolicy
ArtifactAccessPolicy
ArtifactPrivacyPolicy
ArtifactValidationPolicy
ArtifactVerificationPolicy
ArtifactRetentionPolicy
ArtifactPresentationPolicy

ArtifactCondition

ActorReference
ArtifactScopeReference
ArtifactMetadata

These form the initial language of the subsystem.

---

### 67. Final Objective

The project should result in a reusable, language-neutral Artifact Specification that allows completely different applications to communicate using the same conceptual model.

Errands should be able to use it for:

delivery evidence
purchase evidence
location information
completion requirements
runner submissions
client verification
disputes
progressive disclosure

Project Manager should be able to use it for:

deliverables
GitHub pull requests
commits
demo videos
review evidence
approval materials
challenge evidence
handover documents
versioned project output

Neither application should appear in the core protocol.

They are examples and consumers only.

The Artifact SDK provides the grammar.

Applications provide the vocabulary.

Hosts provide the behavior.

The final architectural principle is:

Artifact Specification
    = defines the shared artifact language

Language SDKs
    = expose that language naturally in each ecosystem

Host Applications
    = implement persistence, workflows, policy enforcement,
      providers, storage, UI, and business behavior

The Artifact SDK should therefore become infrastructure rather than application code: a reusable protocol for representing produced work, evidence, references, revisions, relationships, submissions, and verification across software systems.
