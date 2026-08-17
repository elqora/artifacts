export type * from "./generated/wire.js";
import type * as Generated from "./generated/wire.js";
import {
  ARTIFACT_CONDITION_KINDS,
  ARTIFACT_INTEGRITY_ALGORITHMS,
  ARTIFACT_PRIVACY_CLASSIFICATIONS,
  ARTIFACT_PRIVACY_REPRESENTATIONS,
  ARTIFACT_SOURCE_TYPES,
  ARTIFACT_VALUE_OPERATORS,
  ARTIFACT_VALUE_TYPES,
  ARTIFACT_VERIFICATION_STATUSES,
  LOCAL_ARTIFACT_SYNC_STATES,
} from "./generated/vocabulary.js";
export {
  ARTIFACT_CONDITION_KINDS,
  ARTIFACT_INTEGRITY_ALGORITHMS,
  ARTIFACT_PRIVACY_CLASSIFICATIONS,
  ARTIFACT_PRIVACY_REPRESENTATIONS,
  ARTIFACT_SOURCE_TYPES,
  ARTIFACT_VALUE_OPERATORS,
  ARTIFACT_VALUE_TYPES,
  ARTIFACT_VERIFICATION_STATUSES,
  LOCAL_ARTIFACT_SYNC_STATES,
};

/** The schema version implemented by this binding. */
export const ARTIFACT_PROTOCOL_VERSION = "1.1" as const;

export type ArtifactProtocolVersion = typeof ARTIFACT_PROTOCOL_VERSION;

/** Base fields carried by every top-level protocol record. */
export interface ArtifactProtocolRecord extends Pick<Generated.WireArtifact, "schemaVersion"> {
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

/** Schema-governed semantic interpretation of an artifact or immutable artifact version. */
export interface ArtifactSpecification<TValue = unknown> extends Omit<Generated.WireArtifactSpecification, "value"> {
  schema: string;
  version: number;
  value: TValue;
}

export interface Artifact<
  TKind extends string = ArtifactKind,
  TValueType extends ArtifactValueType = ArtifactValueType,
  TSpecification = unknown,
> extends ArtifactProtocolRecord, Omit<Generated.WireArtifact, "scope" | "kind" | "valueType" | "specification" | "createdBy" | "metadata"> {
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

export interface InlineArtifactSource<TValue = unknown> extends Omit<Generated.WireInlineArtifactSource, "value"> {
  type: "inline";
  value: TValue;
  mediaType?: string;
}

export type LocalArtifactSyncState =
  (typeof LOCAL_ARTIFACT_SYNC_STATES)[number];

export interface LocalArtifactSource extends Generated.WireLocalArtifactSource {
  type: "local";
  localId: string;
  filename?: string;
  mediaType?: string;
  size?: number;
  syncState: LocalArtifactSyncState;
  remoteVersionId?: ArtifactVersionId;
}

export interface ObjectArtifactSource extends Generated.WireObjectArtifactSource {
  type: "object";
  objectId: string;
  filename?: string;
  mediaType?: string;
  size?: number;
  storageProvider?: string;
}

export interface UrlArtifactSource extends Generated.WireUrlArtifactSource {
  type: "url";
  url: string;
  provider?: string;
  mediaType?: string;
}

export interface HostedArtifactSource extends Generated.WireHostedArtifactSource {
  type: "hosted";
  recordType: string;
  recordId: string;
}

export interface ProviderArtifactSource<
  TProvider extends string = string,
  TReference = unknown,
> extends Omit<Generated.WireProviderArtifactSource, "provider" | "reference"> {
  type: "provider";
  provider: TProvider;
  reference: TReference;
}

export type ArtifactSource =
  | InlineArtifactSource
  | LocalArtifactSource
  | ObjectArtifactSource
  | UrlArtifactSource
  | HostedArtifactSource
  | ProviderArtifactSource;

export type ArtifactIntegrityAlgorithm =
  (typeof ARTIFACT_INTEGRITY_ALGORITHMS)[number];

export interface ArtifactIntegrity extends Generated.WireArtifactIntegrity {
  algorithm: ArtifactIntegrityAlgorithm;
  hash: string;
  size?: number;
  verifiedAt?: string;
}

export interface ArtifactVersion<
  TSource extends ArtifactSource = ArtifactSource,
  TSpecification = unknown,
> extends ArtifactProtocolRecord, Omit<Generated.WireArtifactVersion, "source" | "integrity" | "specification" | "createdBy" | "metadata"> {
  id: ArtifactVersionId;
  artifactId: ArtifactId;
  version: number;
  source: TSource;
  integrity?: ArtifactIntegrity;
  specification?: ArtifactSpecification<TSpecification>;
  createdBy: ActorReference;
  createdAt: string;
  note?: string;
  metadata?: ArtifactMetadata;
}

export interface ArtifactLink<
  TRole extends string = ArtifactRole,
  TSubjectType extends string = string,
> extends ArtifactProtocolRecord, Omit<Generated.WireArtifactLink, "subject" | "role" | "createdBy" | "metadata"> {
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

export type ArtifactValueOperator = (typeof ARTIFACT_VALUE_OPERATORS)[number];

export interface ArtifactStateCondition extends Generated.WireArtifactConditionState {
  kind: "state";
  namespace: string;
  in: string[];
}

export interface ArtifactActorCondition extends Generated.WireArtifactConditionActor {
  kind: "actor";
  in: string[];
}

export interface ArtifactExistsCondition extends Generated.WireArtifactConditionArtifactExists {
  kind: "artifact_exists";
  artifact: string;
}

export interface ArtifactValueCondition<TValue = unknown> extends Omit<Generated.WireArtifactConditionArtifactValue, "value"> {
  kind: "artifact_value";
  artifact: string;
  operator: ArtifactValueOperator;
  value: TValue;
}

export interface ArtifactAndCondition extends Omit<Generated.WireArtifactConditionAnd, "conditions"> { kind: "and"; conditions: ArtifactCondition[]; }
export interface ArtifactOrCondition extends Omit<Generated.WireArtifactConditionOr, "conditions"> { kind: "or"; conditions: ArtifactCondition[]; }
export interface ArtifactNotCondition extends Omit<Generated.WireArtifactConditionNot, "condition"> { kind: "not"; condition: ArtifactCondition; }

export type ArtifactCondition =
  | ArtifactStateCondition | ArtifactActorCondition | ArtifactExistsCondition
  | ArtifactValueCondition | ArtifactAndCondition | ArtifactOrCondition
  | ArtifactNotCondition;

export interface ArtifactProviderPolicy extends Generated.WireArtifactProviderPolicy {
  actors: string[];
  mode?: "single" | "any" | "all";
  delegation?: "forbidden" | "allowed";
}

export interface ArtifactRequirementPolicy extends Omit<Generated.WireArtifactRequirementPolicy, "condition"> {
  mode: "required" | "optional" | "conditional";
  condition?: ArtifactCondition;
  blocks?: string[];
}

export interface ArtifactLifecyclePolicy extends Omit<Generated.WireArtifactLifecyclePolicy, "condition"> {
  createAt?: string;
  editableDuring?: string[];
  submitDuring?: string[];
  lockAt?: string;
  invalidateOn?: string[];
  condition?: ArtifactCondition;
}

export interface ArtifactAccessRule extends Omit<Generated.WireArtifactAccessPolicyRule, "condition"> {
  actors: string[];
  condition?: ArtifactCondition;
}

export interface ArtifactAccessPolicy extends Omit<Generated.WireArtifactAccessPolicy, "read" | "write" | "submit" | "verify"> {
  read?: ArtifactAccessRule[];
  write?: ArtifactAccessRule[];
  submit?: ArtifactAccessRule[];
  verify?: ArtifactAccessRule[];
}

export type ArtifactPrivacyClassification =
  (typeof ARTIFACT_PRIVACY_CLASSIFICATIONS)[number];

export type ArtifactPrivacyRepresentation =
  (typeof ARTIFACT_PRIVACY_REPRESENTATIONS)[number];

export interface ArtifactRevealRule extends Omit<Generated.WireArtifactRevealRule, "when"> {
  actors: string[];
  when?: ArtifactCondition;
  representation: ArtifactPrivacyRepresentation;
}

export interface ArtifactMaskingPolicy extends Generated.WireArtifactMaskingPolicy {
  strategy: string;
  config?: Record<string, unknown>;
}

export interface ArtifactEncryptionPolicy extends Generated.WireArtifactEncryptionPolicy {
  required: boolean;
  level?: string;
  keyScope?: string;
}

export interface ArtifactPrivacyPolicy extends Omit<Generated.WireArtifactPrivacyPolicy, "reveal" | "masking" | "encryption"> {
  classification: ArtifactPrivacyClassification;
  reveal?: ArtifactRevealRule[];
  masking?: ArtifactMaskingPolicy;
  encryption?: ArtifactEncryptionPolicy;
}

export interface ArtifactValidationRule extends Generated.WireArtifactValidationRule {
  type: string;
  config?: Record<string, unknown>;
}

export interface ArtifactValidationPolicy extends Omit<Generated.WireArtifactValidationPolicy, "rules"> {
  mode?: "strict" | "lenient";
  rules?: ArtifactValidationRule[];
}

export interface ArtifactVerificationPolicy extends Omit<Generated.WireArtifactVerificationPolicy, "condition"> {
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

export interface ArtifactPresentationPolicy extends Generated.WireArtifactPresentationPolicy {
  label?: string;
  helpText?: string;
  order?: number;
  display?: string;
  config?: Record<string, unknown>;
}

export interface TextArtifactValueSchema extends Generated.WireArtifactValueSchemaText { valueType: "text"; }
export interface NumberArtifactValueSchema extends Generated.WireArtifactValueSchemaNumber { valueType: "number"; }
export interface BooleanArtifactValueSchema extends Generated.WireArtifactValueSchemaBoolean { valueType: "boolean"; }
export interface CurrencyArtifactValueSchema extends Generated.WireArtifactValueSchemaCurrency { valueType: "currency"; }
export interface DateArtifactValueSchema extends Generated.WireArtifactValueSchemaDate { valueType: "date"; }
export interface DatetimeArtifactValueSchema extends Generated.WireArtifactValueSchemaDatetime { valueType: "datetime"; }
export interface TimeArtifactValueSchema extends Generated.WireArtifactValueSchemaTime { valueType: "time"; }
export interface LocationArtifactValueSchema extends Generated.WireArtifactValueSchemaLocation { valueType: "location"; }
export interface FileArtifactValueSchema extends Generated.WireArtifactValueSchemaFile { valueType: "file"; }
export interface ImageArtifactValueSchema extends Generated.WireArtifactValueSchemaImage { valueType: "image"; }
export interface VideoArtifactValueSchema extends Generated.WireArtifactValueSchemaVideo { valueType: "video"; }
export interface AudioArtifactValueSchema extends Generated.WireArtifactValueSchemaAudio { valueType: "audio"; }
export interface LinkArtifactValueSchema extends Generated.WireArtifactValueSchemaLink { valueType: "link"; }
export interface StructuredArtifactValueSchema extends Generated.WireArtifactValueSchemaStructured { valueType: "structured"; }
export interface ReferenceArtifactValueSchema extends Generated.WireArtifactValueSchemaReference { valueType: "reference"; }
export interface SignatureArtifactValueSchema extends Generated.WireArtifactValueSchemaSignature { valueType: "signature"; }
export interface CollectionArtifactValueSchema extends Omit<Generated.WireArtifactValueSchemaCollection, "itemSchema"> { valueType: "collection"; itemSchema: ArtifactValueSchema; }

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
> extends ArtifactProtocolRecord, Omit<Generated.WireArtifactSpec, "kind" | "valueType" | "config" | "provider" | "requirement" | "lifecycle" | "access" | "privacy" | "validation" | "verification" | "retention" | "presentation" | "metadata"> {
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

export interface ArtifactRequirement extends ArtifactProtocolRecord, Omit<Generated.WireArtifactRequirement, "metadata"> {
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

export interface ArtifactSubmissionContext extends Generated.WireArtifactSubmissionContext {
  latitude?: number;
  longitude?: number;
  deviceId?: string;
  ipAddress?: string;
  userAgent?: string;
}

export interface ArtifactSubmission<TValue = unknown> extends ArtifactProtocolRecord, Omit<Generated.WireArtifactSubmission, "submittedBy" | "value" | "context" | "metadata"> {
  id: ArtifactSubmissionId;
  artifactId: ArtifactId;
  artifactVersionId?: ArtifactVersionId;
  submittedBy: ActorReference;
  value?: TValue;
  submittedAt: string;
  context?: ArtifactSubmissionContext;
  metadata?: ArtifactMetadata;
}

export type ArtifactVerificationStatus =
  (typeof ARTIFACT_VERIFICATION_STATUSES)[number];

export interface ArtifactVerification extends ArtifactProtocolRecord, Omit<Generated.WireArtifactVerification, "verifiedBy" | "metadata"> {
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
