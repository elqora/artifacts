Artifact SDK — Language-Neutral Specification and Implementation Design

1. Purpose

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

«An artifact is a typed, identifiable, versionable piece of information, evidence, output, or reference that can participate in relationships and workflows.»

The SDK defines what these records mean.

The host application defines what happens because of them.

---

2. Architectural Principle

The architecture should follow this rule:

Artifact SDK
    = shared language and contracts

Host Application
    = behavior and implementation

The SDK should not know what an Errand, Milestone, Review, Project, Delivery, Challenge, or Approval is.

Applications provide those concepts through extensible identifiers and subject references.

The protocol therefore standardizes structure, while allowing applications to define their own vocabulary.

---

3. Language-Neutral Design

TypeScript is used throughout this document as the reference notation because it makes the contracts easy to understand.

TypeScript must not be considered the authoritative implementation.

Eventually the authoritative representation should be a language-neutral schema such as JSON Schema or an equivalent schema format.

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

4. Core Model

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

5. Identifier Types

All entity identifiers should be opaque strings.

export type ArtifactId = string;
export type ArtifactVersionId = string;
export type ArtifactSpecId = string;
export type ArtifactLinkId = string;
export type ArtifactSubmissionId = string;
export type ArtifactVerificationId = string;

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

6. Protocol Versioning

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

A base protocol record may eventually expose:

export interface ArtifactProtocolRecord {
  schemaVersion: string;
}

Example:

{
  schemaVersion: "1.1"
}

Changing an artifact's content must not alter the protocol version.

Changing an artifact definition or instance semantic specification must not alter
the protocol version.

---

7. Actor References

The Artifact SDK must not depend on a host application's "User" model.

Artifact operations therefore use generic actor references.

export interface ActorReference<
  TType extends string = string
> {
  type: TType;

  id?: string;

  displayName?: string;
}

Examples:

const githubActor: ActorReference = {
  type: "github_user",
  id: "12345678",
  displayName: "Davy",
};

const runner: ActorReference = {
  type: "runner",
  id: "runner_123",
};

const system: ActorReference = {
  type: "system",
};

The SDK does not interpret these actor types.

---

8. Artifact Scope

Artifacts may optionally exist inside a broad application-level scope.

export interface ArtifactScopeReference {
  type: string;
  id: string;
}

Examples:

{
  type: "project",
  id: "prj_123"
}

{
  type: "errand",
  id: "err_456"
}

The scope is not the same thing as an "ArtifactLink".

Scope describes the broad owning boundary.

Links describe specific relationships and usages.

---

9. Artifact Kind vs Artifact Value Type

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

10. Artifact Kind

"kind" represents what the artifact means in the consuming domain.

It should remain open.

export type ArtifactKind = string;

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

11. Artifact Value Type

The value type represents the fundamental representation or capability.

This vocabulary should be more standardized.

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

Examples:

{
  kind: "delivery_evidence",
  valueType: "image"
}

{
  kind: "implementation",
  valueType: "reference"
}

{
  kind: "purchase_receipt",
  valueType: "structured"
}

{
  kind: "client_signature",
  valueType: "signature"
}

---

12. Artifact

"Artifact" represents the logical identity of the thing.

It should not directly contain provider-specific content.

export interface Artifact<
  TKind extends string = string,
  TValueType extends string = ArtifactValueType,
  TSpecification = unknown
> {
  id: ArtifactId;

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

An Artifact is the logical object.

Example:

Authentication demonstration

The physical MP4, YouTube link, or provider reference is represented by an "ArtifactVersion".

`ArtifactSpecification` is schema-identified semantic interpretation data for
an actual artifact. It is distinct from `ArtifactSpec`, which describes an
expected artifact and policy. Artifact-level specifications describe stable
logical meaning; version-level specifications describe one content revision.

export interface ArtifactSpecification<T = unknown> {
  schema: string;
  version: number;
  value: T;
}

---

13. Artifact Metadata

The protocol needs an extension mechanism.

export type ArtifactMetadata =
  Record<string, unknown>;

Metadata should not become a dumping ground.

If a property becomes important to interoperability or core behavior, it should become a formal protocol field.

Use metadata for host extensions and non-core information.

---

14. ArtifactVersion

Artifacts can change over time.

The logical identity should remain stable while physical content changes through versions.

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

15. ArtifactSource

An "ArtifactSource" describes where the content or reference represented by a particular version comes from.

The core SDK should support broad source categories.

export type ArtifactSource =
  | InlineArtifactSource
  | LocalArtifactSource
  | ObjectArtifactSource
  | UrlArtifactSource
  | HostedArtifactSource
  | ProviderArtifactSource;

---

16. InlineArtifactSource

Inline values are useful for text, structured data, numbers, coordinates, declarations, and similar values.

export interface InlineArtifactSource {
  type: "inline";

  value: unknown;

  mediaType?: string;
}

Example:

{
  type: "inline",
  value: "Package delivered successfully."
}

Structured example:

{
  type: "inline",
  value: {
    merchant: "Example Store",
    amount: 12500,
    currency: "NGN",
    transactionReference: "TX-1289"
  }
}

---

17. LocalArtifactSource

Local sources support offline-first systems.

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

A host can therefore create artifacts before network upload succeeds.

The Artifact ID remains stable.

---

18. ObjectArtifactSource

Uploaded objects should use a storage-neutral representation.

export interface ObjectArtifactSource {
  type: "object";

  objectId: string;

  filename?: string;

  mediaType?: string;

  size?: number;

  storageProvider?: string;
}

The SDK should not assume:

Amazon S3
Cloudflare R2
MinIO
Azure Blob Storage
Google Cloud Storage

The host decides how "objectId" maps to physical storage.

---

19. UrlArtifactSource

External web resources can be represented as URLs.

export interface UrlArtifactSource {
  type: "url";

  url: string;

  provider?: string;

  mediaType?: string;
}

Possible examples:

- YouTube;
- Vimeo;
- Google Drive;
- Dropbox;
- public file URLs;
- documentation pages;
- external evidence pages.

---

20. HostedArtifactSource

A host application may manage some artifacts internally.

export interface HostedArtifactSource {
  type: "hosted";

  recordType: string;

  recordId: string;
}

Example:

{
  type: "hosted",
  recordType: "generated_report",
  recordId: "report_983"
}

---

21. ProviderArtifactSource

External providers should not be hardcoded into the global SDK.

The protocol should expose a generic provider source.

export interface ProviderArtifactSource<
  TProvider extends string = string,
  TReference = unknown
> {
  type: "provider";

  provider: TProvider;

  reference: TReference;
}

Provider-specific packages can define strong reference types.

---

22. GitHub Provider Example

A GitHub extension could define:

export type GitHubArtifactReference =
  | GitHubPullRequestReference
  | GitHubCommitReference
  | GitHubIssueReference
  | GitHubDiscussionReference
  | GitHubDeploymentReference
  | GitHubWorkflowRunReference
  | GitHubCheckRunReference;

Pull request:

export interface GitHubPullRequestReference {
  resource: "pull_request";

  repositoryId: number;

  number: number;

  nodeId?: string;
}

Commit:

export interface GitHubCommitReference {
  resource: "commit";

  repositoryId: number;

  sha: string;
}

Issue:

export interface GitHubIssueReference {
  resource: "issue";

  repositoryId: number;

  number: number;

  nodeId?: string;
}

Usage:

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

GitLab support should not require modifying the Artifact core specification.

A GitLab extension should simply define another provider-reference contract.

---

23. ArtifactIntegrity

Artifacts used as evidence may require proof that their content has not changed.

export type ArtifactIntegrityAlgorithm =
  | "sha256"
  | "sha384"
  | "sha512";

export interface ArtifactIntegrity {
  algorithm: ArtifactIntegrityAlgorithm;

  hash: string;

  size?: number;

  verifiedAt?: string;
}

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

24. ArtifactLink

Being evidence, a deliverable, an attachment, or review material is not an intrinsic property of an artifact.

It is a relationship.

The SDK therefore uses "ArtifactLink".

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

An ArtifactLink may optionally pin a specific version.

If no version is specified, the relationship follows the logical artifact.

If a version is specified, the relationship refers to that exact content revision.

---

25. ArtifactSubjectReference

The Artifact SDK must not contain application-specific subject unions.

Instead:

export interface ArtifactSubjectReference<
  TType extends string = string
> {
  type: TType;

  id: string;

  scope?: Record<string, string>;
}

Project Manager may define:

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

Errands may define:

export type ErrandsArtifactSubjectType =
  | "errand"
  | "assignment"
  | "delivery"
  | "completion"
  | "dispute"
  | "payment";

The global Artifact SDK needs to understand neither list.

---

26. Artifact Role

Artifact roles should also remain extensible.

export type ArtifactRole = string;

Project Manager may define:

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

Errands might define:

export type ErrandsArtifactRole =
  | "delivery_evidence"
  | "completion_evidence"
  | "purchase_evidence"
  | "dispute_evidence"
  | "identity_evidence";

The principle is:

«Standardize the relationship structure, not every possible relationship vocabulary.»

---

27. Example of Artifact Reuse

Suppose a video exists once:

Authentication demonstration

It may have one Artifact:

const artifact: Artifact = {
  id: "art_auth_demo",

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

It can then have multiple links:

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

The artifact does not change because its context changes.

---

28. ArtifactSpec

Some applications need to declare what artifacts are expected before an actual artifact exists.

This concept comes from workflow-driven systems such as Errands.

The global protocol should model that separately as "ArtifactSpec".

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

"ArtifactSpec" is not an Artifact.

It describes what an expected artifact should look like and the policies governing it.

---

29. ArtifactSpecSnapshot

Definitions can evolve.

A running task or workflow should not silently change because an administrator edits an artifact specification later.

Therefore specifications need immutable snapshots.

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

30. ArtifactProviderPolicy

Provider policy answers:

«Who is expected to supply this artifact?»

It must not be confused with a storage/provider source.

export interface ArtifactProviderPolicy {
  actors: string[];

  mode?: "single" | "any" | "all";

  delegation?: "forbidden" | "allowed";
}

Example:

{
  actors: ["runner"],
  mode: "single"
}

Another application:

{
  actors: ["developer", "designer"],
  mode: "any"
}

Actor identifiers remain host-defined.

---

31. ArtifactRequirementPolicy

Requiredness should not be reduced to:

required: true

because workflow-oriented systems need to know what an artifact requirement blocks.

export interface ArtifactRequirementPolicy {
  mode:
    | "required"
    | "optional"
    | "conditional";

  condition?: ArtifactCondition;

  blocks?: string[];
}

Example:

{
  mode: "required",

  blocks: [
    "submit_completion"
  ]
}

Project Manager might instead define:

{
  mode: "required",

  blocks: [
    "request_milestone_approval"
  ]
}

The operation identifiers remain domain-specific strings.

---

32. ArtifactLifecyclePolicy

Lifecycle policy defines when an artifact may exist or change.

export interface ArtifactLifecyclePolicy {
  createAt?: string;

  editableDuring?: string[];

  submitDuring?: string[];

  lockAt?: string;

  invalidateOn?: string[];

  condition?: ArtifactCondition;
}

Errands example:

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

Project Manager example:

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

The SDK does not interpret these stage names globally.

---

33. ArtifactCondition

Conditions must be language-neutral and serializable.

Never encode policies using language-specific callbacks.

Do not do:

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

---

34. State Condition

export interface ArtifactStateCondition {
  kind: "state";

  namespace: string;

  in: string[];
}

Example:

{
  kind: "state",

  namespace: "errand",

  in: [
    "accepted",
    "in_progress"
  ]
}

---

35. Actor Condition

export interface ArtifactActorCondition {
  kind: "actor";

  in: string[];
}

Example:

{
  kind: "actor",
  in: ["runner", "admin"]
}

---

36. Artifact Exists Condition

export interface ArtifactExistsCondition {
  kind: "artifact_exists";

  artifact: string;
}

The artifact reference may correspond to a spec key, artifact key, or another host-defined identifier.

The exact resolution rule should be documented by the protocol or extension layer.

---

37. Artifact Value Condition

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

Example:

{
  kind: "artifact_value",

  artifact: "item_unavailable",

  operator: "eq",

  value: true
}

---

38. Logical Conditions

export interface ArtifactAndCondition {
  kind: "and";

  conditions: ArtifactCondition[];
}

export interface ArtifactOrCondition {
  kind: "or";

  conditions: ArtifactCondition[];
}

export interface ArtifactNotCondition {
  kind: "not";

  condition: ArtifactCondition;
}

Example:

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

This is safely serializable to JSON and interpretable in any implementation language.

---

39. ArtifactAccessPolicy

Access rules should remain distinct from privacy rules.

export interface ArtifactAccessPolicy {
  read?: ArtifactAccessRule[];

  write?: ArtifactAccessRule[];

  submit?: ArtifactAccessRule[];

  verify?: ArtifactAccessRule[];
}

export interface ArtifactAccessRule {
  actors: string[];

  condition?: ArtifactCondition;
}

Example:

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

---

40. ArtifactPrivacyPolicy

Access asks:

«Is this actor allowed to access the artifact?»

Privacy asks:

«What representation of the artifact is this actor allowed to receive?»

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

---

41. ArtifactRevealRule

export interface ArtifactRevealRule {
  actors: string[];

  when?: ArtifactCondition;

  representation:
    | "hidden"
    | "masked"
    | "approximate"
    | "full";
}

This supports progressive disclosure.

For example, a runner may see an approximate destination before accepting a job:

{
  actors: ["runner"],

  when: {
    kind: "state",
    namespace: "errand",
    in: ["published", "offered"]
  },

  representation: "approximate"
}

After acceptance:

{
  actors: ["runner"],

  when: {
    kind: "state",
    namespace: "errand",
    in: ["accepted", "in_progress"]
  },

  representation: "full"
}

---

42. ArtifactMaskingPolicy

A basic extensible form could be:

export interface ArtifactMaskingPolicy {
  strategy: string;

  config?: Record<string, unknown>;
}

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

43. ArtifactEncryptionPolicy

The protocol may describe encryption expectations without implementing encryption itself.

export interface ArtifactEncryptionPolicy {
  required: boolean;

  level?: string;

  keyScope?: string;
}

The host decides how actual encryption is performed.

---

44. ArtifactSubmission

Artifact values should not necessarily be silently overwritten.

Operational submission history should be first-class.

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

---

45. ArtifactSubmissionContext

export interface ArtifactSubmissionContext {
  latitude?: number;

  longitude?: number;

  deviceId?: string;

  ipAddress?: string;

  userAgent?: string;
}

This supports evidence-sensitive workflows where submission provenance matters.

The protocol defines the data shape.

The host decides which values are safe and legal to collect.

---

46. ArtifactVerification

Verification must not be represented as:

verified: true

Verification is an event/record with provenance.

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

47. ArtifactVerificationPolicy

Artifact specifications may define verification requirements.

export interface ArtifactVerificationPolicy {
  required: boolean;

  methods?: string[];

  actors?: string[];

  condition?: ArtifactCondition;
}

Example:

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

The policy describes what is expected.

"ArtifactVerification" records what actually happened.

---

48. ArtifactValidationPolicy

export interface ArtifactValidationPolicy {
  mode?: "strict" | "lenient";

  rules?: ArtifactValidationRule[];
}

export interface ArtifactValidationRule {
  type: string;

  config?: Record<string, unknown>;
}

Example:

{
  type: "file_count",

  config: {
    min: 1,
    max: 4
  }
}

Another:

{
  type: "mime_type",

  config: {
    allowed: [
      "image/jpeg",
      "image/png"
    ]
  }
}

The protocol may later standardize common validation-rule identifiers.

---

49. ArtifactRetentionPolicy

Retention should describe expectations, not storage implementation.

export type ArtifactRetentionPolicy =
  | {
      policy: "forever";
    }
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

{
  policy: "duration",
  days: 365
}

---

50. ArtifactPresentationPolicy

Presentation hints can exist without coupling the SDK to a UI framework.

export interface ArtifactPresentationPolicy {
  label?: string;

  helpText?: string;

  order?: number;

  display?: string;

  config?: Record<string, unknown>;
}

These should remain hints only.

The host may ignore them.

The protocol must never depend on React, Vue, Flutter, SwiftUI, or another presentation framework.

---

51. Artifact Value Schemas

Value-type-specific schemas can add stronger typing and interoperable constraints.

Example image schema:

export interface ImageArtifactValueSchema {
  valueType: "image";

  minFiles?: number;

  maxFiles?: number;

  acceptedMimeTypes?: string[];

  maxSizeBytes?: number;

  requireTimestamp?: boolean;

  requireLocation?: boolean;
}

Location:

export interface LocationArtifactValueSchema {
  valueType: "location";

  mode:
    | "point"
    | "address"
    | "point_and_address";

  requireCoordinates?: boolean;

  allowManualEntry?: boolean;
}

Text:

export interface TextArtifactValueSchema {
  valueType: "text";

  minLength?: number;

  maxLength?: number;

  multiline?: boolean;

  pattern?: string;
}

File:

export interface FileArtifactValueSchema {
  valueType: "file";

  minFiles?: number;

  maxFiles?: number;

  acceptedMimeTypes?: string[];

  maxSizeBytes?: number;
}

Collection:

export interface CollectionArtifactValueSchema {
  valueType: "collection";

  itemSchema: ArtifactValueSchema;

  minItems?: number;

  maxItems?: number;
}

Eventually:

export type ArtifactValueSchema =
  | ImageArtifactValueSchema
  | LocationArtifactValueSchema
  | TextArtifactValueSchema
  | FileArtifactValueSchema
  | CollectionArtifactValueSchema;

Additional value types can be introduced over time.

---

52. ArtifactRequirement

Some domain objects may require artifacts without owning a full ArtifactSpec.

A generic requirement model can represent this.

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

Example:

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

A project-management deliverable can use this without making the Artifact itself aware that it is a deliverable.

---

53. Project Manager Example

A developer submits a GitHub pull request.

Artifact:

const artifact: Artifact = {
  id: "art_auth_impl",

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

Version:

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

Deliverable relationship:

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

Review relationship:

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

The same artifact participates in multiple contexts without duplication.

---

54. Errands Example

An Errands workflow requires delivery photographic evidence.

Specification:

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

Actual artifact:

const artifact: Artifact = {
  id: "art_delivery_1",

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

Artifact version:

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

Submission:

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

Verification:

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

This proves that two fundamentally different applications can use the same artifact vocabulary.

---

55. What the SDK Must Not Implement

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

56. Optional SDK Utilities

Language bindings may eventually provide deterministic helper utilities.

Examples:

export function isArtifactArchived(
  artifact: Artifact
): boolean {
  return artifact.archivedAt !== undefined;
}

export function isVersionPinned(
  link: ArtifactLink
): boolean {
  return link.artifactVersionId !== undefined;
}

Potential helpers include:

- validation against generated schema;
- serialization;
- deserialization;
- protocol-version compatibility checks;
- artifact-reference helpers;
- integrity verification helpers;
- condition-AST parsing;
- safe discriminated-union helpers.

These utilities are optional conveniences.

They must not turn the SDK into a host-specific runtime.

---

57. Open vs Closed Vocabularies

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

58. Extension Packages

Provider or domain-specific contracts should eventually be separable.

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

59. Suggested Repository Structure

An initial repository could look like:

artifacts/
├── README.md
├── Go.md
├── LICENSE
│
├── spec/
│   ├── protocol/
│   │   └── protocol.schema.json
│   │
│   ├── artifact/
│   │   ├── artifact.schema.json
│   │   ├── artifact-version.schema.json
│   │   ├── artifact-source.schema.json
│   │   ├── artifact-integrity.schema.json
│   │   └── artifact-link.schema.json
│   │
│   ├── specification/
│   │   ├── artifact-spec.schema.json
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
│   │   └── retention-policy.schema.json
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
├── generators/
│
├── tests/
│   ├── fixtures/
│   ├── conformance/
│   └── compatibility/
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

The exact repository structure may change, but the conceptual boundaries should remain.

---

60. Conformance Fixtures

Because multiple languages will implement the protocol, the repository should eventually include canonical JSON fixtures.

Example:

tests/fixtures/artifact/basic.json
tests/fixtures/artifact/versioned-file.json
tests/fixtures/artifact/github-reference.json
tests/fixtures/spec/delivery-evidence.json
tests/fixtures/submission/delivery-photo.json
tests/fixtures/verification/client-confirmation.json

All generated or manually implemented language SDKs should be able to deserialize and validate the same fixtures.

This gives the project actual cross-language conformance rather than merely similar-looking interfaces.

---

61. Serialization Rules

Serialized forms should use stable discriminators.

For example:

{
  "type": "provider",
  "provider": "github",
  "reference": {
    "resource": "pull_request",
    "repositoryId": 123,
    "number": 55
  }
}

Avoid representations whose meaning depends on language-specific class names.

Dates should use interoperable formats such as ISO 8601 strings unless the protocol later defines another canonical representation.

Optional fields should have clear missing/null semantics.

The specification should eventually explicitly define:

- whether absent and "null" mean the same thing;
- number precision expectations;
- date/time formats;
- identifier case sensitivity;
- enum case sensitivity;
- unknown-field behavior;
- forward-compatible parsing behavior.

---

62. Forward Compatibility

The protocol should be designed so future SDKs can safely encounter records created by newer versions.

Where possible:

- additions should be backward-compatible;
- unknown metadata should be preserved;
- extension identifiers should not cause parser failure;
- provider-specific references should remain opaque to implementations that do not understand them.

Breaking structural changes should require a protocol-major-version change.

---

63. Security Boundary

The Artifact SDK describes security-related intent but must not create false guarantees.

For example:

privacy: {
  classification: "restricted"
}

does not itself make content restricted.

The host must enforce the policy.

Likewise:

encryption: {
  required: true
}

does not encrypt anything.

The documentation must explicitly state:

«Policy records are declarations. Security exists only when a host correctly enforces them.»

This distinction must remain clear throughout the SDK.

---

64. Initial Implementation Priorities

The initial implementation should focus on the stable shared core before attempting advanced runtime policy evaluation.

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

The first iteration does not require a full policy engine.

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

Start with providers that are actually needed.

For example:

GitHub

Do not prematurely add dozens of provider definitions.

Provider extensions should prove the extension mechanism rather than expand the Artifact core.

---

65. Primary Design Rules for Implementers

The implementation should preserve these rules.

Rule 1

Artifact identity must remain separate from artifact content.

Artifact != ArtifactVersion

Rule 2

Provider-specific physical references belong to ArtifactSource.

ArtifactVersion → ArtifactSource

Rule 3

Usage is a relationship.

Do not add fields such as:

isDeliverable
isEvidence
reviewId
milestoneId

to Artifact.

Use ArtifactLink.

Rule 4

Version-sensitive relationships must support version pinning.

Approval of v3 must not automatically apply to v4.

Rule 5

Artifact specifications are not actual artifacts.

ArtifactSpec != Artifact

Rule 6

Definitions must be serializable.

Do not store executable callbacks in ArtifactSpec or policy definitions.

Rule 7

Domain vocabulary remains extensible.

Do not hardcode application concepts such as:

milestone
errand
runner
review
submit_completion

into the Artifact core specification.

Rule 8

The SDK does not implement host workflows.

Rule 9

Security policies are declarations that hosts must enforce.

Rule 10

All language SDKs must represent the same protocol semantics.

---

66. Final Domain Vocabulary

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

67. Final Objective

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
