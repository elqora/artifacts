// Code generated from canonical JSON Schemas; DO NOT EDIT.

export type WireArtifactIntegrity = {
  "algorithm": "sha256" | "sha384" | "sha512";
  "hash": string;
  "size"?: number;
  "verifiedAt"?: WireArtifactProtocolCommonContractsTimestamp;
  [key: string]: unknown;
};

export type WireArtifactLink = {
  "schemaVersion": WireArtifactProtocolCommonContractsSchemaVersion;
  "id": WireArtifactProtocolCommonContractsArtifactLinkId;
  "artifactId": WireArtifactProtocolCommonContractsArtifactId;
  "artifactVersionId"?: WireArtifactProtocolCommonContractsArtifactVersionId;
  "subject": WireArtifactProtocolCommonContractsArtifactSubjectReference;
  "role": string;
  "note"?: string;
  "createdBy": WireArtifactProtocolCommonContractsActorReference;
  "createdAt": WireArtifactProtocolCommonContractsTimestamp;
  "metadata"?: WireArtifactProtocolCommonContractsArtifactMetadata;
  [key: string]: unknown;
};

export type WireArtifactSource = WireInlineArtifactSource | WireLocalArtifactSource | WireObjectArtifactSource | WireUrlArtifactSource | WireHostedArtifactSource | WireProviderArtifactSource;

export type WireInlineArtifactSource = {
  "type": "inline";
  "value": unknown;
  "mediaType"?: string;
  [key: string]: unknown;
};

export type WireLocalArtifactSource = {
  "type": "local";
  "localId": string;
  "filename"?: string;
  "mediaType"?: string;
  "size"?: number;
  "syncState": "local_only" | "pending_upload" | "uploading" | "uploaded" | "failed";
  "remoteVersionId"?: WireArtifactProtocolCommonContractsArtifactVersionId;
  [key: string]: unknown;
};

export type WireObjectArtifactSource = {
  "type": "object";
  "objectId": string;
  "filename"?: string;
  "mediaType"?: string;
  "size"?: number;
  "storageProvider"?: string;
  [key: string]: unknown;
};

export type WireUrlArtifactSource = {
  "type": "url";
  "url": string;
  "provider"?: string;
  "mediaType"?: string;
  [key: string]: unknown;
};

export type WireHostedArtifactSource = {
  "type": "hosted";
  "recordType": string;
  "recordId": string;
  [key: string]: unknown;
};

export type WireProviderArtifactSource = {
  "type": "provider";
  "provider": string;
  "reference": unknown;
  [key: string]: unknown;
};

export type WireArtifactVersion = {
  "schemaVersion": WireArtifactProtocolCommonContractsSchemaVersion;
  "id": WireArtifactProtocolCommonContractsArtifactVersionId;
  "artifactId": WireArtifactProtocolCommonContractsArtifactId;
  "version": number;
  "source": WireArtifactSource;
  "integrity"?: WireArtifactIntegrity;
  "createdBy": WireArtifactProtocolCommonContractsActorReference;
  "createdAt": WireArtifactProtocolCommonContractsTimestamp;
  "note"?: string;
  "metadata"?: WireArtifactProtocolCommonContractsArtifactMetadata;
  [key: string]: unknown;
};

export type WireArtifact = {
  "schemaVersion": WireArtifactProtocolCommonContractsSchemaVersion;
  "id": WireArtifactProtocolCommonContractsArtifactId;
  "scope"?: WireArtifactProtocolCommonContractsArtifactScopeReference;
  "kind": string;
  "valueType": WireArtifactProtocolCommonContractsArtifactValueType;
  "title"?: string;
  "description"?: string;
  "currentVersionId"?: WireArtifactProtocolCommonContractsArtifactVersionId;
  "createdBy": WireArtifactProtocolCommonContractsActorReference;
  "createdAt": WireArtifactProtocolCommonContractsTimestamp;
  "updatedAt": WireArtifactProtocolCommonContractsTimestamp;
  "archivedAt"?: WireArtifactProtocolCommonContractsTimestamp;
  "metadata"?: WireArtifactProtocolCommonContractsArtifactMetadata;
  [key: string]: unknown;
};

export type WireArtifactAccessPolicy = {
  "read"?: WireArtifactAccessPolicyRules;
  "write"?: WireArtifactAccessPolicyRules;
  "submit"?: WireArtifactAccessPolicyRules;
  "verify"?: WireArtifactAccessPolicyRules;
  [key: string]: unknown;
};

export type WireArtifactAccessPolicyRules = Array<WireArtifactAccessPolicyRule>;

export type WireArtifactAccessPolicyRule = {
  "actors": Array<string>;
  "condition"?: WireArtifactCondition;
  [key: string]: unknown;
};

/** A language-neutral recursive condition AST. Resolution of host identifiers is host-defined. */
export type WireArtifactCondition = WireArtifactConditionCondition;

export type WireArtifactConditionCondition = WireArtifactConditionState | WireArtifactConditionActor | WireArtifactConditionArtifactExists | WireArtifactConditionArtifactValue | WireArtifactConditionAnd | WireArtifactConditionOr | WireArtifactConditionNot;

export type WireArtifactConditionState = {
  "kind": "state";
  "namespace": string;
  "in": Array<string>;
  [key: string]: unknown;
};

export type WireArtifactConditionActor = {
  "kind": "actor";
  "in": Array<string>;
  [key: string]: unknown;
};

export type WireArtifactConditionArtifactExists = {
  "kind": "artifact_exists";
  "artifact": string;
  [key: string]: unknown;
};

export type WireArtifactConditionArtifactValue = {
  "kind": "artifact_value";
  "artifact": string;
  "operator": "eq" | "neq" | "gt" | "gte" | "lt" | "lte" | "contains" | "in";
  "value": unknown;
  [key: string]: unknown;
};

export type WireArtifactConditionAnd = {
  "kind": "and";
  "conditions": Array<WireArtifactConditionCondition>;
  [key: string]: unknown;
};

export type WireArtifactConditionOr = {
  "kind": "or";
  "conditions": Array<WireArtifactConditionCondition>;
  [key: string]: unknown;
};

export type WireArtifactConditionNot = {
  "kind": "not";
  "condition": WireArtifactConditionCondition;
  [key: string]: unknown;
};

export type WireArtifactLifecyclePolicy = {
  "createAt"?: string;
  "editableDuring"?: Array<string>;
  "submitDuring"?: Array<string>;
  "lockAt"?: string;
  "invalidateOn"?: Array<string>;
  "condition"?: WireArtifactCondition;
  [key: string]: unknown;
};

export type WireArtifactPresentationPolicy = {
  "label"?: string;
  "helpText"?: string;
  "order"?: number;
  "display"?: string;
  "config"?: {
  [key: string]: unknown;
};
  [key: string]: unknown;
};

export type WireArtifactPrivacyPolicy = {
  "classification": "public" | "internal" | "private" | "sensitive" | "restricted";
  "reveal"?: Array<WireArtifactRevealRule>;
  "masking"?: WireArtifactMaskingPolicy;
  "encryption"?: WireArtifactEncryptionPolicy;
  [key: string]: unknown;
};

export type WireArtifactRevealRule = {
  "actors": Array<string>;
  "when"?: WireArtifactCondition;
  "representation": "hidden" | "masked" | "approximate" | "full";
  [key: string]: unknown;
};

export type WireArtifactMaskingPolicy = {
  "strategy": string;
  "config"?: {
  [key: string]: unknown;
};
  [key: string]: unknown;
};

export type WireArtifactEncryptionPolicy = {
  "required": boolean;
  "level"?: string;
  "keyScope"?: string;
  [key: string]: unknown;
};

export type WireArtifactProviderPolicy = {
  "actors": Array<string>;
  "mode"?: "single" | "any" | "all";
  "delegation"?: "forbidden" | "allowed";
  [key: string]: unknown;
};

export type WireArtifactRequirementPolicy = {
  "mode": "required" | "optional" | "conditional";
  "condition"?: WireArtifactCondition;
  "blocks"?: Array<string>;
  [key: string]: unknown;
};

export type WireArtifactRetentionPolicy = WireArtifactRetentionPolicyForever | WireArtifactRetentionPolicyDuration | WireArtifactRetentionPolicyUntil | WireArtifactRetentionPolicyHostDefined;

export type WireArtifactRetentionPolicyForever = {
  "policy": "forever";
  [key: string]: unknown;
};

export type WireArtifactRetentionPolicyDuration = {
  "policy": "duration";
  "days": number;
  [key: string]: unknown;
};

export type WireArtifactRetentionPolicyUntil = {
  "policy": "until";
  "date": WireArtifactProtocolCommonContractsDate;
  [key: string]: unknown;
};

export type WireArtifactRetentionPolicyHostDefined = {
  "policy": "host_defined";
  "key": string;
  [key: string]: unknown;
};

export type WireArtifactValidationPolicy = {
  "mode"?: "strict" | "lenient";
  "rules"?: Array<WireArtifactValidationRule>;
  [key: string]: unknown;
};

export type WireArtifactValidationRule = {
  "type": string;
  "config"?: {
  [key: string]: unknown;
};
  [key: string]: unknown;
};

export type WireArtifactVerificationPolicy = {
  "required": boolean;
  "methods"?: Array<string>;
  "actors"?: Array<string>;
  "condition"?: WireArtifactCondition;
  [key: string]: unknown;
};

/** The major.minor version of the serialized Artifact Protocol record. */
export type WireArtifactProtocolCommonContractsSchemaVersion = "1.0";

/** A non-empty, case-sensitive identifier whose generation and internal form are host-defined. */
export type WireArtifactProtocolCommonContractsOpaqueIdentifier = string;

export type WireArtifactProtocolCommonContractsArtifactId = WireArtifactProtocolCommonContractsOpaqueIdentifier;

export type WireArtifactProtocolCommonContractsArtifactVersionId = WireArtifactProtocolCommonContractsOpaqueIdentifier;

export type WireArtifactProtocolCommonContractsArtifactLinkId = WireArtifactProtocolCommonContractsOpaqueIdentifier;

export type WireArtifactProtocolCommonContractsArtifactSpecId = WireArtifactProtocolCommonContractsOpaqueIdentifier;

export type WireArtifactProtocolCommonContractsArtifactSubmissionId = WireArtifactProtocolCommonContractsOpaqueIdentifier;

export type WireArtifactProtocolCommonContractsArtifactVerificationId = WireArtifactProtocolCommonContractsOpaqueIdentifier;

export type WireArtifactProtocolCommonContractsArtifactRequirementId = WireArtifactProtocolCommonContractsOpaqueIdentifier;

/** An RFC 3339 full-date string. */
export type WireArtifactProtocolCommonContractsDate = string;

/** An RFC 3339 date-time string. */
export type WireArtifactProtocolCommonContractsTimestamp = string;

export type WireArtifactProtocolCommonContractsArtifactValueType = "text" | "number" | "boolean" | "currency" | "date" | "datetime" | "time" | "location" | "file" | "image" | "video" | "audio" | "link" | "structured" | "reference" | "signature" | "collection";

export type WireArtifactProtocolCommonContractsActorReference = {
  "type": string;
  "id"?: WireArtifactProtocolCommonContractsOpaqueIdentifier;
  "displayName"?: string;
  [key: string]: unknown;
};

export type WireArtifactProtocolCommonContractsArtifactScopeReference = {
  "type": string;
  "id": WireArtifactProtocolCommonContractsOpaqueIdentifier;
  [key: string]: unknown;
};

export type WireArtifactProtocolCommonContractsArtifactSubjectReference = {
  "type": string;
  "id": WireArtifactProtocolCommonContractsOpaqueIdentifier;
  "scope"?: {
  [key: string]: string;
};
  [key: string]: unknown;
};

/** Host extensions represented by arbitrary JSON values. */
export type WireArtifactProtocolCommonContractsArtifactMetadata = {
  [key: string]: unknown;
};

export type WireArtifactSubmission = {
  "schemaVersion": WireArtifactProtocolCommonContractsSchemaVersion;
  "id": WireArtifactProtocolCommonContractsArtifactSubmissionId;
  "artifactId": WireArtifactProtocolCommonContractsArtifactId;
  "artifactVersionId"?: WireArtifactProtocolCommonContractsArtifactVersionId;
  "submittedBy": WireArtifactProtocolCommonContractsActorReference;
  "value"?: unknown;
  "submittedAt": WireArtifactProtocolCommonContractsTimestamp;
  "context"?: WireArtifactSubmissionContext;
  "metadata"?: WireArtifactProtocolCommonContractsArtifactMetadata;
  [key: string]: unknown;
};

export type WireArtifactSubmissionContext = {
  "latitude"?: number;
  "longitude"?: number;
  "deviceId"?: string;
  "ipAddress"?: unknown | unknown;
  "userAgent"?: string;
  [key: string]: unknown;
};

export type WireArtifactVerification = {
  "schemaVersion": WireArtifactProtocolCommonContractsSchemaVersion;
  "id": WireArtifactProtocolCommonContractsArtifactVerificationId;
  "artifactId": WireArtifactProtocolCommonContractsArtifactId;
  "artifactVersionId"?: WireArtifactProtocolCommonContractsArtifactVersionId;
  "submissionId"?: WireArtifactProtocolCommonContractsArtifactSubmissionId;
  "status": "pending" | "verified" | "rejected" | "waived";
  "method"?: string;
  "verifiedBy"?: WireArtifactProtocolCommonContractsActorReference;
  "reason"?: string;
  "createdAt": WireArtifactProtocolCommonContractsTimestamp;
  "verifiedAt"?: WireArtifactProtocolCommonContractsTimestamp;
  "metadata"?: WireArtifactProtocolCommonContractsArtifactMetadata;
  [key: string]: unknown;
};

export type WireArtifactRequirement = {
  "schemaVersion": WireArtifactProtocolCommonContractsSchemaVersion;
  "id": WireArtifactProtocolCommonContractsArtifactRequirementId;
  "key"?: string;
  "allowedKinds"?: Array<string>;
  "allowedValueTypes"?: Array<WireArtifactProtocolCommonContractsArtifactValueType>;
  "minimumCount"?: number;
  "maximumCount"?: number;
  "required": boolean;
  "specId"?: WireArtifactProtocolCommonContractsArtifactSpecId;
  "metadata"?: WireArtifactProtocolCommonContractsArtifactMetadata;
  [key: string]: unknown;
};

/** An immutable copy of a specification revision captured by a host. */
export type WireArtifactSpecSnapshot = {
  "schemaVersion": WireArtifactProtocolCommonContractsSchemaVersion;
  "sourceSpecId": WireArtifactProtocolCommonContractsArtifactSpecId;
  "sourceVersion": number;
  "key": string;
  "name": string;
  "description"?: string;
  "kind": string;
  "valueType": WireArtifactProtocolCommonContractsArtifactValueType;
  "config"?: {
  "valueType"?: WireArtifactProtocolCommonContractsArtifactValueType;
  [key: string]: unknown;
};
  "provider"?: WireArtifactProviderPolicy;
  "requirement"?: WireArtifactRequirementPolicy;
  "lifecycle"?: WireArtifactLifecyclePolicy;
  "access"?: WireArtifactAccessPolicy;
  "privacy"?: WireArtifactPrivacyPolicy;
  "validation"?: WireArtifactValidationPolicy;
  "verification"?: WireArtifactVerificationPolicy;
  "retention"?: WireArtifactRetentionPolicy;
  "presentation"?: WireArtifactPresentationPolicy;
  "metadata"?: WireArtifactProtocolCommonContractsArtifactMetadata;
  [key: string]: unknown;
};

export type WireArtifactSpec = {
  "schemaVersion": WireArtifactProtocolCommonContractsSchemaVersion;
  "id": WireArtifactProtocolCommonContractsArtifactSpecId;
  "key": string;
  "name": string;
  "description"?: string;
  "version": number;
  "kind": string;
  "valueType": WireArtifactProtocolCommonContractsArtifactValueType;
  "config"?: {
  "valueType"?: WireArtifactProtocolCommonContractsArtifactValueType;
  [key: string]: unknown;
};
  "provider"?: WireArtifactProviderPolicy;
  "requirement"?: WireArtifactRequirementPolicy;
  "lifecycle"?: WireArtifactLifecyclePolicy;
  "access"?: WireArtifactAccessPolicy;
  "privacy"?: WireArtifactPrivacyPolicy;
  "validation"?: WireArtifactValidationPolicy;
  "verification"?: WireArtifactVerificationPolicy;
  "retention"?: WireArtifactRetentionPolicy;
  "presentation"?: WireArtifactPresentationPolicy;
  "metadata"?: WireArtifactProtocolCommonContractsArtifactMetadata;
  [key: string]: unknown;
};

/** Standardized representation constraints. Artifact kind remains a separate host-defined semantic identifier. */
export type WireArtifactValueSchema = WireArtifactValueSchemaText | WireArtifactValueSchemaNumber | WireArtifactValueSchemaBoolean | WireArtifactValueSchemaCurrency | WireArtifactValueSchemaDate | WireArtifactValueSchemaDatetime | WireArtifactValueSchemaTime | WireArtifactValueSchemaLocation | WireArtifactValueSchemaFile | WireArtifactValueSchemaImage | WireArtifactValueSchemaVideo | WireArtifactValueSchemaAudio | WireArtifactValueSchemaLink | WireArtifactValueSchemaStructured | WireArtifactValueSchemaReference | WireArtifactValueSchemaSignature | WireArtifactValueSchemaCollection;

export type WireArtifactValueSchemaText = {
  "valueType": "text";
  "minLength"?: number;
  "maxLength"?: number;
  "multiline"?: boolean;
  "pattern"?: string;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaNumber = {
  "valueType": "number";
  "minimum"?: number;
  "maximum"?: number;
  "integer"?: boolean;
  "multipleOf"?: number;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaBoolean = {
  "valueType": "boolean";
  [key: string]: unknown;
};

export type WireArtifactValueSchemaCurrency = {
  "valueType": "currency";
  "currencies"?: Array<string>;
  "minimumMinorUnits"?: number;
  "maximumMinorUnits"?: number;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaDate = {
  "valueType": "date";
  "minimum"?: WireArtifactProtocolCommonContractsDate;
  "maximum"?: WireArtifactProtocolCommonContractsDate;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaDatetime = {
  "valueType": "datetime";
  "minimum"?: WireArtifactProtocolCommonContractsTimestamp;
  "maximum"?: WireArtifactProtocolCommonContractsTimestamp;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaTime = {
  "valueType": "time";
  "minimum"?: string;
  "maximum"?: string;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaLocation = {
  "valueType": "location";
  "mode": "point" | "address" | "point_and_address";
  "requireCoordinates"?: boolean;
  "allowManualEntry"?: boolean;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaFile = WireArtifactValueSchemaFileBase & {
  "valueType": "file";
  [key: string]: unknown;
};

export type WireArtifactValueSchemaImage = WireArtifactValueSchemaFileBase & {
  "valueType": "image";
  "requireTimestamp"?: boolean;
  "requireLocation"?: boolean;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaVideo = WireArtifactValueSchemaFileBase & {
  "valueType": "video";
  "maxDurationSeconds"?: number;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaAudio = WireArtifactValueSchemaFileBase & {
  "valueType": "audio";
  "maxDurationSeconds"?: number;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaFileBase = {
  "minFiles"?: number;
  "maxFiles"?: number;
  "acceptedMimeTypes"?: Array<string>;
  "maxSizeBytes"?: number;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaLink = {
  "valueType": "link";
  "allowedSchemes"?: Array<string>;
  "allowedHosts"?: Array<string>;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaStructured = {
  "valueType": "structured";
  "jsonSchema"?: {
  [key: string]: unknown;
};
  [key: string]: unknown;
};

export type WireArtifactValueSchemaReference = {
  "valueType": "reference";
  "providers"?: Array<string>;
  "resourceTypes"?: Array<string>;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaSignature = {
  "valueType": "signature";
  "methods"?: Array<string>;
  "requireTimestamp"?: boolean;
  [key: string]: unknown;
};

export type WireArtifactValueSchemaCollection = {
  "valueType": "collection";
  "itemSchema": WireArtifactValueSchema;
  "minItems"?: number;
  "maxItems"?: number;
  "uniqueItems"?: boolean;
  [key: string]: unknown;
};
