/** The schema version implemented by this binding. */
export const ARTIFACT_PROTOCOL_VERSION = "1.0" as const;

export type ArtifactProtocolVersion = typeof ARTIFACT_PROTOCOL_VERSION;

/** Base fields carried by every top-level protocol record. */
export interface ArtifactProtocolRecord {
  schemaVersion: ArtifactProtocolVersion;
}

/** Opaque, non-empty serialized identifiers. Their generation is host-defined. */
export type ArtifactId = string;
export type ArtifactVersionId = string;
export type ArtifactLinkId = string;
export type ArtifactSpecId = string;
export type ArtifactRequirementId = string;
export type ArtifactSubmissionId = string;
export type ArtifactVerificationId = string;

/** Host extension data. Values must be JSON-serializable when encoded. */
export type ArtifactMetadata = Record<string, unknown>;

export const ARTIFACT_VALUE_TYPES = [
  "text",
  "number",
  "boolean",
  "currency",
  "date",
  "datetime",
  "time",
  "location",
  "file",
  "image",
  "video",
  "audio",
  "link",
  "structured",
  "reference",
  "signature",
  "collection",
] as const;

export type ArtifactValueType = (typeof ARTIFACT_VALUE_TYPES)[number];

export type ArtifactKind = string;
export type ArtifactRole = string;

export interface ActorReference<TType extends string = string> {
  type: TType;
  id?: string;
  displayName?: string;
}

export interface ArtifactScopeReference<TType extends string = string> {
  type: TType;
  id: string;
}

export interface ArtifactSubjectReference<TType extends string = string> {
  type: TType;
  id: string;
  scope?: Record<string, string>;
}

export interface Artifact<
  TKind extends string = ArtifactKind,
  TValueType extends ArtifactValueType = ArtifactValueType,
> extends ArtifactProtocolRecord {
  id: ArtifactId;
  scope?: ArtifactScopeReference;
  kind: TKind;
  valueType: TValueType;
  title?: string;
  description?: string;
  currentVersionId?: ArtifactVersionId;
  createdBy: ActorReference;
  createdAt: string;
  updatedAt: string;
  archivedAt?: string;
  metadata?: ArtifactMetadata;
}

export interface InlineArtifactSource<TValue = unknown> {
  type: "inline";
  value: TValue;
  mediaType?: string;
}

export const LOCAL_ARTIFACT_SYNC_STATES = [
  "local_only",
  "pending_upload",
  "uploading",
  "uploaded",
  "failed",
] as const;

export type LocalArtifactSyncState =
  (typeof LOCAL_ARTIFACT_SYNC_STATES)[number];

export interface LocalArtifactSource {
  type: "local";
  localId: string;
  filename?: string;
  mediaType?: string;
  size?: number;
  syncState: LocalArtifactSyncState;
  remoteVersionId?: ArtifactVersionId;
}

export interface ObjectArtifactSource {
  type: "object";
  objectId: string;
  filename?: string;
  mediaType?: string;
  size?: number;
  storageProvider?: string;
}

export interface UrlArtifactSource {
  type: "url";
  url: string;
  provider?: string;
  mediaType?: string;
}

export interface HostedArtifactSource {
  type: "hosted";
  recordType: string;
  recordId: string;
}

export interface ProviderArtifactSource<
  TProvider extends string = string,
  TReference = unknown,
> {
  type: "provider";
  provider: TProvider;
  reference: TReference;
}

export const ARTIFACT_SOURCE_TYPES = [
  "inline",
  "local",
  "object",
  "url",
  "hosted",
  "provider",
] as const;

export type ArtifactSource =
  | InlineArtifactSource
  | LocalArtifactSource
  | ObjectArtifactSource
  | UrlArtifactSource
  | HostedArtifactSource
  | ProviderArtifactSource;

export const ARTIFACT_INTEGRITY_ALGORITHMS = [
  "sha256",
  "sha384",
  "sha512",
] as const;

export type ArtifactIntegrityAlgorithm =
  (typeof ARTIFACT_INTEGRITY_ALGORITHMS)[number];

export interface ArtifactIntegrity {
  algorithm: ArtifactIntegrityAlgorithm;
  hash: string;
  size?: number;
  verifiedAt?: string;
}

export interface ArtifactVersion<
  TSource extends ArtifactSource = ArtifactSource,
> extends ArtifactProtocolRecord {
  id: ArtifactVersionId;
  artifactId: ArtifactId;
  version: number;
  source: TSource;
  integrity?: ArtifactIntegrity;
  createdBy: ActorReference;
  createdAt: string;
  note?: string;
  metadata?: ArtifactMetadata;
}

export interface ArtifactLink<
  TRole extends string = ArtifactRole,
  TSubjectType extends string = string,
> extends ArtifactProtocolRecord {
  id: ArtifactLinkId;
  artifactId: ArtifactId;
  /** Omit to follow the logical artifact; provide to pin immutable content. */
  artifactVersionId?: ArtifactVersionId;
  subject: ArtifactSubjectReference<TSubjectType>;
  role: TRole;
  note?: string;
  createdBy: ActorReference;
  createdAt: string;
  metadata?: ArtifactMetadata;
}

export const ARTIFACT_CONDITION_KINDS = [
  "state", "actor", "artifact_exists", "artifact_value", "and", "or", "not",
] as const;

export const ARTIFACT_VALUE_OPERATORS = [
  "eq", "neq", "gt", "gte", "lt", "lte", "contains", "in",
] as const;

export type ArtifactValueOperator = (typeof ARTIFACT_VALUE_OPERATORS)[number];

export interface ArtifactStateCondition {
  kind: "state";
  namespace: string;
  in: string[];
}

export interface ArtifactActorCondition {
  kind: "actor";
  in: string[];
}

export interface ArtifactExistsCondition {
  kind: "artifact_exists";
  artifact: string;
}

export interface ArtifactValueCondition<TValue = unknown> {
  kind: "artifact_value";
  artifact: string;
  operator: ArtifactValueOperator;
  value: TValue;
}

export interface ArtifactAndCondition { kind: "and"; conditions: ArtifactCondition[]; }
export interface ArtifactOrCondition { kind: "or"; conditions: ArtifactCondition[]; }
export interface ArtifactNotCondition { kind: "not"; condition: ArtifactCondition; }

export type ArtifactCondition =
  | ArtifactStateCondition | ArtifactActorCondition | ArtifactExistsCondition
  | ArtifactValueCondition | ArtifactAndCondition | ArtifactOrCondition
  | ArtifactNotCondition;

export interface ArtifactProviderPolicy {
  actors: string[];
  mode?: "single" | "any" | "all";
  delegation?: "forbidden" | "allowed";
}

export interface ArtifactRequirementPolicy {
  mode: "required" | "optional" | "conditional";
  condition?: ArtifactCondition;
  blocks?: string[];
}

export interface ArtifactLifecyclePolicy {
  createAt?: string;
  editableDuring?: string[];
  submitDuring?: string[];
  lockAt?: string;
  invalidateOn?: string[];
  condition?: ArtifactCondition;
}

export interface ArtifactAccessRule {
  actors: string[];
  condition?: ArtifactCondition;
}

export interface ArtifactAccessPolicy {
  read?: ArtifactAccessRule[];
  write?: ArtifactAccessRule[];
  submit?: ArtifactAccessRule[];
  verify?: ArtifactAccessRule[];
}

export const ARTIFACT_PRIVACY_CLASSIFICATIONS = [
  "public", "internal", "private", "sensitive", "restricted",
] as const;
export type ArtifactPrivacyClassification =
  (typeof ARTIFACT_PRIVACY_CLASSIFICATIONS)[number];

export const ARTIFACT_PRIVACY_REPRESENTATIONS = [
  "hidden", "masked", "approximate", "full",
] as const;
export type ArtifactPrivacyRepresentation =
  (typeof ARTIFACT_PRIVACY_REPRESENTATIONS)[number];

export interface ArtifactRevealRule {
  actors: string[];
  when?: ArtifactCondition;
  representation: ArtifactPrivacyRepresentation;
}

export interface ArtifactMaskingPolicy {
  strategy: string;
  config?: Record<string, unknown>;
}

export interface ArtifactEncryptionPolicy {
  required: boolean;
  level?: string;
  keyScope?: string;
}

export interface ArtifactPrivacyPolicy {
  classification: ArtifactPrivacyClassification;
  reveal?: ArtifactRevealRule[];
  masking?: ArtifactMaskingPolicy;
  encryption?: ArtifactEncryptionPolicy;
}

export interface ArtifactValidationRule {
  type: string;
  config?: Record<string, unknown>;
}

export interface ArtifactValidationPolicy {
  mode?: "strict" | "lenient";
  rules?: ArtifactValidationRule[];
}

export interface ArtifactVerificationPolicy {
  required: boolean;
  methods?: string[];
  actors?: string[];
  condition?: ArtifactCondition;
}

export type ArtifactRetentionPolicy =
  | { policy: "forever" }
  | { policy: "duration"; days: number }
  | { policy: "until"; date: string }
  | { policy: "host_defined"; key: string };

export interface ArtifactPresentationPolicy {
  label?: string;
  helpText?: string;
  order?: number;
  display?: string;
  config?: Record<string, unknown>;
}

export interface TextArtifactValueSchema { valueType: "text"; minLength?: number; maxLength?: number; multiline?: boolean; pattern?: string; }
export interface NumberArtifactValueSchema { valueType: "number"; minimum?: number; maximum?: number; integer?: boolean; multipleOf?: number; }
export interface BooleanArtifactValueSchema { valueType: "boolean"; }
export interface CurrencyArtifactValueSchema { valueType: "currency"; currencies?: string[]; minimumMinorUnits?: number; maximumMinorUnits?: number; }
export interface DateArtifactValueSchema { valueType: "date"; minimum?: string; maximum?: string; }
export interface DatetimeArtifactValueSchema { valueType: "datetime"; minimum?: string; maximum?: string; }
export interface TimeArtifactValueSchema { valueType: "time"; minimum?: string; maximum?: string; }
export interface LocationArtifactValueSchema { valueType: "location"; mode: "point" | "address" | "point_and_address"; requireCoordinates?: boolean; allowManualEntry?: boolean; }
export interface FileArtifactValueSchema { valueType: "file"; minFiles?: number; maxFiles?: number; acceptedMimeTypes?: string[]; maxSizeBytes?: number; }
export interface ImageArtifactValueSchema extends Omit<FileArtifactValueSchema, "valueType"> { valueType: "image"; requireTimestamp?: boolean; requireLocation?: boolean; }
export interface VideoArtifactValueSchema extends Omit<FileArtifactValueSchema, "valueType"> { valueType: "video"; maxDurationSeconds?: number; }
export interface AudioArtifactValueSchema extends Omit<FileArtifactValueSchema, "valueType"> { valueType: "audio"; maxDurationSeconds?: number; }
export interface LinkArtifactValueSchema { valueType: "link"; allowedSchemes?: string[]; allowedHosts?: string[]; }
export interface StructuredArtifactValueSchema { valueType: "structured"; jsonSchema?: Record<string, unknown>; }
export interface ReferenceArtifactValueSchema { valueType: "reference"; providers?: string[]; resourceTypes?: string[]; }
export interface SignatureArtifactValueSchema { valueType: "signature"; methods?: string[]; requireTimestamp?: boolean; }
export interface CollectionArtifactValueSchema { valueType: "collection"; itemSchema: ArtifactValueSchema; minItems?: number; maxItems?: number; uniqueItems?: boolean; }

export type ArtifactValueSchema =
  | TextArtifactValueSchema | NumberArtifactValueSchema | BooleanArtifactValueSchema
  | CurrencyArtifactValueSchema | DateArtifactValueSchema | DatetimeArtifactValueSchema
  | TimeArtifactValueSchema | LocationArtifactValueSchema | FileArtifactValueSchema
  | ImageArtifactValueSchema | VideoArtifactValueSchema | AudioArtifactValueSchema
  | LinkArtifactValueSchema | StructuredArtifactValueSchema | ReferenceArtifactValueSchema
  | SignatureArtifactValueSchema | CollectionArtifactValueSchema;

export type ArtifactValueConfig<TValueType extends ArtifactValueType> =
  Omit<Extract<ArtifactValueSchema, { valueType: TValueType }>, "valueType"> &
  { valueType?: TValueType };

export interface ArtifactSpec<
  TKind extends string = ArtifactKind,
  TValueType extends ArtifactValueType = ArtifactValueType,
  TConfig extends object = ArtifactValueConfig<TValueType>,
> extends ArtifactProtocolRecord {
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

export type ArtifactSpecSnapshot<
  TKind extends string = ArtifactKind,
  TValueType extends ArtifactValueType = ArtifactValueType,
  TConfig extends object = ArtifactValueConfig<TValueType>,
> = Omit<ArtifactSpec<TKind, TValueType, TConfig>, "id" | "version"> & {
  sourceSpecId: ArtifactSpecId;
  sourceVersion: number;
};

export interface ArtifactRequirement extends ArtifactProtocolRecord {
  id: ArtifactRequirementId;
  key?: string;
  allowedKinds?: string[];
  allowedValueTypes?: ArtifactValueType[];
  minimumCount?: number;
  maximumCount?: number;
  required: boolean;
  specId?: ArtifactSpecId;
  metadata?: ArtifactMetadata;
}

export interface ArtifactSubmissionContext {
  latitude?: number;
  longitude?: number;
  deviceId?: string;
  ipAddress?: string;
  userAgent?: string;
}

export interface ArtifactSubmission<TValue = unknown> extends ArtifactProtocolRecord {
  id: ArtifactSubmissionId;
  artifactId: ArtifactId;
  artifactVersionId?: ArtifactVersionId;
  submittedBy: ActorReference;
  value?: TValue;
  submittedAt: string;
  context?: ArtifactSubmissionContext;
  metadata?: ArtifactMetadata;
}

export const ARTIFACT_VERIFICATION_STATUSES = [
  "pending", "verified", "rejected", "waived",
] as const;
export type ArtifactVerificationStatus =
  (typeof ARTIFACT_VERIFICATION_STATUSES)[number];

export interface ArtifactVerification extends ArtifactProtocolRecord {
  id: ArtifactVerificationId;
  artifactId: ArtifactId;
  artifactVersionId?: ArtifactVersionId;
  submissionId?: ArtifactSubmissionId;
  status: ArtifactVerificationStatus;
  method?: string;
  verifiedBy?: ActorReference;
  reason?: string;
  createdAt: string;
  verifiedAt?: string;
  metadata?: ArtifactMetadata;
}
