// Code generated from canonical JSON Schemas; DO NOT EDIT.
package artifact

import "encoding/json"

type WireArtifactIntegrity struct {
	Algorithm string `json:"algorithm"`
	Hash string `json:"hash"`
	Size *int64 `json:"size,omitempty"`
	VerifiedAt WireArtifactProtocolCommonContractsTimestamp `json:"verifiedAt,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactLink struct {
	SchemaVersion WireArtifactProtocolCommonContractsSchemaVersion `json:"schemaVersion"`
	ID WireArtifactProtocolCommonContractsArtifactLinkID `json:"id"`
	ArtifactID WireArtifactProtocolCommonContractsArtifactID `json:"artifactId"`
	ArtifactVersionID WireArtifactProtocolCommonContractsArtifactVersionID `json:"artifactVersionId,omitempty"`
	Subject WireArtifactProtocolCommonContractsArtifactSubjectReference `json:"subject"`
	Role string `json:"role"`
	Note *string `json:"note,omitempty"`
	CreatedBy WireArtifactProtocolCommonContractsActorReference `json:"createdBy"`
	CreatedAt WireArtifactProtocolCommonContractsTimestamp `json:"createdAt"`
	Metadata *WireArtifactProtocolCommonContractsArtifactMetadata `json:"metadata,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactSource = json.RawMessage

type WireInlineArtifactSource struct {
	Type string `json:"type"`
	Value json.RawMessage `json:"value"`
	MediaType *string `json:"mediaType,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireLocalArtifactSource struct {
	Type string `json:"type"`
	LocalID string `json:"localId"`
	Filename *string `json:"filename,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	Size *int64 `json:"size,omitempty"`
	SyncState string `json:"syncState"`
	RemoteVersionID WireArtifactProtocolCommonContractsArtifactVersionID `json:"remoteVersionId,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireObjectArtifactSource struct {
	Type string `json:"type"`
	ObjectID string `json:"objectId"`
	Filename *string `json:"filename,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	Size *int64 `json:"size,omitempty"`
	StorageProvider *string `json:"storageProvider,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireURLArtifactSource struct {
	Type string `json:"type"`
	URL string `json:"url"`
	Provider *string `json:"provider,omitempty"`
	MediaType *string `json:"mediaType,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireHostedArtifactSource struct {
	Type string `json:"type"`
	RecordType string `json:"recordType"`
	RecordID string `json:"recordId"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireProviderArtifactSource struct {
	Type string `json:"type"`
	Provider string `json:"provider"`
	Reference json.RawMessage `json:"reference"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactSpecification struct {
	Schema string `json:"schema"`
	Version int64 `json:"version"`
	Value json.RawMessage `json:"value"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactVersion struct {
	SchemaVersion WireArtifactProtocolCommonContractsSchemaVersion `json:"schemaVersion"`
	ID WireArtifactProtocolCommonContractsArtifactVersionID `json:"id"`
	ArtifactID WireArtifactProtocolCommonContractsArtifactID `json:"artifactId"`
	Version int64 `json:"version"`
	Source WireArtifactSource `json:"source"`
	Integrity *WireArtifactIntegrity `json:"integrity,omitempty"`
	Specification *WireArtifactSpecification `json:"specification,omitempty"`
	CreatedBy WireArtifactProtocolCommonContractsActorReference `json:"createdBy"`
	CreatedAt WireArtifactProtocolCommonContractsTimestamp `json:"createdAt"`
	Note *string `json:"note,omitempty"`
	Metadata *WireArtifactProtocolCommonContractsArtifactMetadata `json:"metadata,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifact struct {
	SchemaVersion WireArtifactProtocolCommonContractsSchemaVersion `json:"schemaVersion"`
	ID WireArtifactProtocolCommonContractsArtifactID `json:"id"`
	SpecID WireArtifactProtocolCommonContractsArtifactSpecID `json:"specId"`
	Scope *WireArtifactProtocolCommonContractsArtifactScopeReference `json:"scope,omitempty"`
	Kind string `json:"kind"`
	ValueType WireArtifactProtocolCommonContractsArtifactValueType `json:"valueType"`
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Specification *WireArtifactSpecification `json:"specification,omitempty"`
	CurrentVersionID WireArtifactProtocolCommonContractsArtifactVersionID `json:"currentVersionId,omitempty"`
	CreatedBy WireArtifactProtocolCommonContractsActorReference `json:"createdBy"`
	CreatedAt WireArtifactProtocolCommonContractsTimestamp `json:"createdAt"`
	UpdatedAt WireArtifactProtocolCommonContractsTimestamp `json:"updatedAt"`
	ArchivedAt WireArtifactProtocolCommonContractsTimestamp `json:"archivedAt,omitempty"`
	Metadata *WireArtifactProtocolCommonContractsArtifactMetadata `json:"metadata,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactAccessPolicy struct {
	Read WireArtifactAccessPolicyRules `json:"read,omitempty"`
	Write WireArtifactAccessPolicyRules `json:"write,omitempty"`
	Submit WireArtifactAccessPolicyRules `json:"submit,omitempty"`
	Verify WireArtifactAccessPolicyRules `json:"verify,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactAccessPolicyRules = []WireArtifactAccessPolicyRule

type WireArtifactAccessPolicyRule struct {
	Actors []string `json:"actors"`
	Condition WireArtifactCondition `json:"condition,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactCondition = WireArtifactConditionCondition

type WireArtifactConditionCondition = json.RawMessage

type WireArtifactConditionState struct {
	Kind string `json:"kind"`
	Namespace string `json:"namespace"`
	In []string `json:"in"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactConditionActor struct {
	Kind string `json:"kind"`
	In []string `json:"in"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactConditionArtifactExists struct {
	Kind string `json:"kind"`
	Artifact string `json:"artifact"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactConditionArtifactValue struct {
	Kind string `json:"kind"`
	Artifact string `json:"artifact"`
	Operator string `json:"operator"`
	Value json.RawMessage `json:"value"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactConditionAnd struct {
	Kind string `json:"kind"`
	Conditions []WireArtifactConditionCondition `json:"conditions"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactConditionOr struct {
	Kind string `json:"kind"`
	Conditions []WireArtifactConditionCondition `json:"conditions"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactConditionNot struct {
	Kind string `json:"kind"`
	Condition WireArtifactConditionCondition `json:"condition"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactLifecyclePolicy struct {
	CreateAt *string `json:"createAt,omitempty"`
	EditableDuring []string `json:"editableDuring,omitempty"`
	SubmitDuring []string `json:"submitDuring,omitempty"`
	LockAt *string `json:"lockAt,omitempty"`
	InvalidateOn []string `json:"invalidateOn,omitempty"`
	Condition WireArtifactCondition `json:"condition,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactPresentationPolicy struct {
	Label *string `json:"label,omitempty"`
	HelpText *string `json:"helpText,omitempty"`
	Order *float64 `json:"order,omitempty"`
	Display *string `json:"display,omitempty"`
	Config map[string]any `json:"config,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactPrivacyPolicy struct {
	Classification string `json:"classification"`
	Reveal []WireArtifactRevealRule `json:"reveal,omitempty"`
	Masking *WireArtifactMaskingPolicy `json:"masking,omitempty"`
	Encryption *WireArtifactEncryptionPolicy `json:"encryption,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactRevealRule struct {
	Actors []string `json:"actors"`
	When WireArtifactCondition `json:"when,omitempty"`
	Representation string `json:"representation"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactMaskingPolicy struct {
	Strategy string `json:"strategy"`
	Config map[string]any `json:"config,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactEncryptionPolicy struct {
	Required bool `json:"required"`
	Level *string `json:"level,omitempty"`
	KeyScope *string `json:"keyScope,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactProviderPolicy struct {
	Actors []string `json:"actors"`
	Mode *string `json:"mode,omitempty"`
	Delegation *string `json:"delegation,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactRequirementPolicy struct {
	Mode string `json:"mode"`
	Condition WireArtifactCondition `json:"condition,omitempty"`
	Blocks []string `json:"blocks,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactRetentionPolicy = json.RawMessage

type WireArtifactRetentionPolicyForever struct {
	Policy string `json:"policy"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactRetentionPolicyDuration struct {
	Policy string `json:"policy"`
	Days int64 `json:"days"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactRetentionPolicyUntil struct {
	Policy string `json:"policy"`
	Date WireArtifactProtocolCommonContractsDate `json:"date"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactRetentionPolicyHostDefined struct {
	Policy string `json:"policy"`
	Key string `json:"key"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValidationPolicy struct {
	Mode *string `json:"mode,omitempty"`
	Schema map[string]any `json:"schema,omitempty"`
	Rules []WireArtifactValidationRule `json:"rules,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValidationRule struct {
	Type string `json:"type"`
	Config map[string]any `json:"config,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactVerificationPolicy struct {
	Required bool `json:"required"`
	Methods []string `json:"methods,omitempty"`
	Actors []string `json:"actors,omitempty"`
	Condition WireArtifactCondition `json:"condition,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactProtocolCommonContracts = json.RawMessage

type WireArtifactProtocolCommonContractsSchemaVersion = string

type WireArtifactProtocolCommonContractsOpaqueIDentifier = string

type WireArtifactProtocolCommonContractsArtifactID = WireArtifactProtocolCommonContractsOpaqueIDentifier

type WireArtifactProtocolCommonContractsArtifactVersionID = WireArtifactProtocolCommonContractsOpaqueIDentifier

type WireArtifactProtocolCommonContractsArtifactLinkID = WireArtifactProtocolCommonContractsOpaqueIDentifier

type WireArtifactProtocolCommonContractsArtifactSpecID = WireArtifactProtocolCommonContractsOpaqueIDentifier

type WireArtifactProtocolCommonContractsArtifactSubmissionID = WireArtifactProtocolCommonContractsOpaqueIDentifier

type WireArtifactProtocolCommonContractsArtifactVerificationID = WireArtifactProtocolCommonContractsOpaqueIDentifier

type WireArtifactProtocolCommonContractsArtifactRequirementID = WireArtifactProtocolCommonContractsOpaqueIDentifier

type WireArtifactProtocolCommonContractsDate = string

type WireArtifactProtocolCommonContractsTimestamp = string

type WireArtifactProtocolCommonContractsArtifactValueType = string

type WireArtifactProtocolCommonContractsActorReference struct {
	Type string `json:"type"`
	ID WireArtifactProtocolCommonContractsOpaqueIDentifier `json:"id,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactProtocolCommonContractsArtifactScopeReference struct {
	Type string `json:"type"`
	ID WireArtifactProtocolCommonContractsOpaqueIDentifier `json:"id"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactProtocolCommonContractsArtifactSubjectReference struct {
	Type string `json:"type"`
	ID WireArtifactProtocolCommonContractsOpaqueIDentifier `json:"id"`
	Scope map[string]any `json:"scope,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactProtocolCommonContractsArtifactMetadata map[string]any

type WireArtifactSubmission struct {
	SchemaVersion WireArtifactProtocolCommonContractsSchemaVersion `json:"schemaVersion"`
	ID WireArtifactProtocolCommonContractsArtifactSubmissionID `json:"id"`
	ArtifactID WireArtifactProtocolCommonContractsArtifactID `json:"artifactId"`
	ArtifactVersionID WireArtifactProtocolCommonContractsArtifactVersionID `json:"artifactVersionId,omitempty"`
	SubmittedBy WireArtifactProtocolCommonContractsActorReference `json:"submittedBy"`
	Value json.RawMessage `json:"value,omitempty"`
	SubmittedAt WireArtifactProtocolCommonContractsTimestamp `json:"submittedAt"`
	Context *WireArtifactSubmissionContext `json:"context,omitempty"`
	Metadata *WireArtifactProtocolCommonContractsArtifactMetadata `json:"metadata,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactSubmissionContext struct {
	Latitude *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	DeviceID *string `json:"deviceId,omitempty"`
	IPAddress json.RawMessage `json:"ipAddress,omitempty"`
	UserAgent *string `json:"userAgent,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactVerification struct {
	SchemaVersion WireArtifactProtocolCommonContractsSchemaVersion `json:"schemaVersion"`
	ID WireArtifactProtocolCommonContractsArtifactVerificationID `json:"id"`
	ArtifactID WireArtifactProtocolCommonContractsArtifactID `json:"artifactId"`
	ArtifactVersionID WireArtifactProtocolCommonContractsArtifactVersionID `json:"artifactVersionId,omitempty"`
	SubmissionID WireArtifactProtocolCommonContractsArtifactSubmissionID `json:"submissionId,omitempty"`
	Status string `json:"status"`
	Method *string `json:"method,omitempty"`
	VerifiedBy *WireArtifactProtocolCommonContractsActorReference `json:"verifiedBy,omitempty"`
	Reason *string `json:"reason,omitempty"`
	CreatedAt WireArtifactProtocolCommonContractsTimestamp `json:"createdAt"`
	VerifiedAt WireArtifactProtocolCommonContractsTimestamp `json:"verifiedAt,omitempty"`
	Metadata *WireArtifactProtocolCommonContractsArtifactMetadata `json:"metadata,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactRequirement struct {
	SchemaVersion WireArtifactProtocolCommonContractsSchemaVersion `json:"schemaVersion"`
	ID WireArtifactProtocolCommonContractsArtifactRequirementID `json:"id"`
	Key *string `json:"key,omitempty"`
	AllowedKinds []string `json:"allowedKinds,omitempty"`
	AllowedValueTypes []WireArtifactProtocolCommonContractsArtifactValueType `json:"allowedValueTypes,omitempty"`
	MinimumCount *int64 `json:"minimumCount,omitempty"`
	MaximumCount *int64 `json:"maximumCount,omitempty"`
	Required bool `json:"required"`
	SpecID WireArtifactProtocolCommonContractsArtifactSpecID `json:"specId,omitempty"`
	Metadata *WireArtifactProtocolCommonContractsArtifactMetadata `json:"metadata,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactSpecSnapshot struct {
	SchemaVersion WireArtifactProtocolCommonContractsSchemaVersion `json:"schemaVersion"`
	SourceSpecID WireArtifactProtocolCommonContractsArtifactSpecID `json:"sourceSpecId"`
	SourceVersion int64 `json:"sourceVersion"`
	Key string `json:"key"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Kind string `json:"kind"`
	ValueType WireArtifactProtocolCommonContractsArtifactValueType `json:"valueType"`
	Config map[string]any `json:"config,omitempty"`
	Provider *WireArtifactProviderPolicy `json:"provider,omitempty"`
	Requirement *WireArtifactRequirementPolicy `json:"requirement,omitempty"`
	Lifecycle *WireArtifactLifecyclePolicy `json:"lifecycle,omitempty"`
	Access *WireArtifactAccessPolicy `json:"access,omitempty"`
	Privacy *WireArtifactPrivacyPolicy `json:"privacy,omitempty"`
	Validation *WireArtifactValidationPolicy `json:"validation,omitempty"`
	Verification *WireArtifactVerificationPolicy `json:"verification,omitempty"`
	Retention WireArtifactRetentionPolicy `json:"retention,omitempty"`
	Presentation *WireArtifactPresentationPolicy `json:"presentation,omitempty"`
	Metadata *WireArtifactProtocolCommonContractsArtifactMetadata `json:"metadata,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactSpec struct {
	SchemaVersion WireArtifactProtocolCommonContractsSchemaVersion `json:"schemaVersion"`
	ID WireArtifactProtocolCommonContractsArtifactSpecID `json:"id"`
	Key string `json:"key"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Version int64 `json:"version"`
	Kind string `json:"kind"`
	ValueType WireArtifactProtocolCommonContractsArtifactValueType `json:"valueType"`
	Config map[string]any `json:"config,omitempty"`
	Provider *WireArtifactProviderPolicy `json:"provider,omitempty"`
	Requirement *WireArtifactRequirementPolicy `json:"requirement,omitempty"`
	Lifecycle *WireArtifactLifecyclePolicy `json:"lifecycle,omitempty"`
	Access *WireArtifactAccessPolicy `json:"access,omitempty"`
	Privacy *WireArtifactPrivacyPolicy `json:"privacy,omitempty"`
	Validation *WireArtifactValidationPolicy `json:"validation,omitempty"`
	Verification *WireArtifactVerificationPolicy `json:"verification,omitempty"`
	Retention WireArtifactRetentionPolicy `json:"retention,omitempty"`
	Presentation *WireArtifactPresentationPolicy `json:"presentation,omitempty"`
	Metadata *WireArtifactProtocolCommonContractsArtifactMetadata `json:"metadata,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchema = json.RawMessage

type WireArtifactValueSchemaText struct {
	ValueType string `json:"valueType"`
	MinLength *int64 `json:"minLength,omitempty"`
	MaxLength *int64 `json:"maxLength,omitempty"`
	Multiline *bool `json:"multiline,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaNumber struct {
	ValueType string `json:"valueType"`
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	Integer *bool `json:"integer,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaBoolean struct {
	ValueType string `json:"valueType"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaCurrency struct {
	ValueType string `json:"valueType"`
	Currencies []string `json:"currencies,omitempty"`
	MinimumMinorUnits *int64 `json:"minimumMinorUnits,omitempty"`
	MaximumMinorUnits *int64 `json:"maximumMinorUnits,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaDate struct {
	ValueType string `json:"valueType"`
	Minimum WireArtifactProtocolCommonContractsDate `json:"minimum,omitempty"`
	Maximum WireArtifactProtocolCommonContractsDate `json:"maximum,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaDatetime struct {
	ValueType string `json:"valueType"`
	Minimum WireArtifactProtocolCommonContractsTimestamp `json:"minimum,omitempty"`
	Maximum WireArtifactProtocolCommonContractsTimestamp `json:"maximum,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaTime struct {
	ValueType string `json:"valueType"`
	Minimum *string `json:"minimum,omitempty"`
	Maximum *string `json:"maximum,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaLocation struct {
	ValueType string `json:"valueType"`
	Mode string `json:"mode"`
	RequireCoordinates *bool `json:"requireCoordinates,omitempty"`
	AllowManualEntry *bool `json:"allowManualEntry,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaFile struct {
	MinFiles *int64 `json:"minFiles,omitempty"`
	MaxFiles *int64 `json:"maxFiles,omitempty"`
	AcceptedMimeTypes []string `json:"acceptedMimeTypes,omitempty"`
	MaxSizeBytes *int64 `json:"maxSizeBytes,omitempty"`
	ValueType string `json:"valueType"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaImage struct {
	MinFiles *int64 `json:"minFiles,omitempty"`
	MaxFiles *int64 `json:"maxFiles,omitempty"`
	AcceptedMimeTypes []string `json:"acceptedMimeTypes,omitempty"`
	MaxSizeBytes *int64 `json:"maxSizeBytes,omitempty"`
	ValueType string `json:"valueType"`
	RequireTimestamp *bool `json:"requireTimestamp,omitempty"`
	RequireLocation *bool `json:"requireLocation,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaVideo struct {
	MinFiles *int64 `json:"minFiles,omitempty"`
	MaxFiles *int64 `json:"maxFiles,omitempty"`
	AcceptedMimeTypes []string `json:"acceptedMimeTypes,omitempty"`
	MaxSizeBytes *int64 `json:"maxSizeBytes,omitempty"`
	ValueType string `json:"valueType"`
	MaxDurationSeconds *float64 `json:"maxDurationSeconds,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaAudio struct {
	MinFiles *int64 `json:"minFiles,omitempty"`
	MaxFiles *int64 `json:"maxFiles,omitempty"`
	AcceptedMimeTypes []string `json:"acceptedMimeTypes,omitempty"`
	MaxSizeBytes *int64 `json:"maxSizeBytes,omitempty"`
	ValueType string `json:"valueType"`
	MaxDurationSeconds *float64 `json:"maxDurationSeconds,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaFileBase struct {
	MinFiles *int64 `json:"minFiles,omitempty"`
	MaxFiles *int64 `json:"maxFiles,omitempty"`
	AcceptedMimeTypes []string `json:"acceptedMimeTypes,omitempty"`
	MaxSizeBytes *int64 `json:"maxSizeBytes,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaLink struct {
	ValueType string `json:"valueType"`
	AllowedSchemes []string `json:"allowedSchemes,omitempty"`
	AllowedHosts []string `json:"allowedHosts,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaStructured struct {
	ValueType string `json:"valueType"`
	JSONSchema map[string]any `json:"jsonSchema,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaReference struct {
	ValueType string `json:"valueType"`
	Providers []string `json:"providers,omitempty"`
	ResourceTypes []string `json:"resourceTypes,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaSignature struct {
	ValueType string `json:"valueType"`
	Methods []string `json:"methods,omitempty"`
	RequireTimestamp *bool `json:"requireTimestamp,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireArtifactValueSchemaCollection struct {
	ValueType string `json:"valueType"`
	ItemSchema WireArtifactValueSchema `json:"itemSchema"`
	MinItems *int64 `json:"minItems,omitempty"`
	MaxItems *int64 `json:"maxItems,omitempty"`
	UniqueItems *bool `json:"uniqueItems,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

func (value *WireArtifactIntegrity) UnmarshalJSON(data []byte) error {
	type plain WireArtifactIntegrity
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactIntegrity(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"algorithm": {}, "hash": {}, "size": {}, "verifiedAt": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactIntegrity) MarshalJSON() ([]byte, error) {
	type plain WireArtifactIntegrity
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"algorithm": {}, "hash": {}, "size": {}, "verifiedAt": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactLink) UnmarshalJSON(data []byte) error {
	type plain WireArtifactLink
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactLink(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schemaVersion": {}, "id": {}, "artifactId": {}, "artifactVersionId": {}, "subject": {}, "role": {}, "note": {}, "createdBy": {}, "createdAt": {}, "metadata": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactLink) MarshalJSON() ([]byte, error) {
	type plain WireArtifactLink
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schemaVersion": {}, "id": {}, "artifactId": {}, "artifactVersionId": {}, "subject": {}, "role": {}, "note": {}, "createdBy": {}, "createdAt": {}, "metadata": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireInlineArtifactSource) UnmarshalJSON(data []byte) error {
	type plain WireInlineArtifactSource
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireInlineArtifactSource(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "value": {}, "mediaType": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireInlineArtifactSource) MarshalJSON() ([]byte, error) {
	type plain WireInlineArtifactSource
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "value": {}, "mediaType": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireLocalArtifactSource) UnmarshalJSON(data []byte) error {
	type plain WireLocalArtifactSource
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireLocalArtifactSource(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "localId": {}, "filename": {}, "mediaType": {}, "size": {}, "syncState": {}, "remoteVersionId": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireLocalArtifactSource) MarshalJSON() ([]byte, error) {
	type plain WireLocalArtifactSource
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "localId": {}, "filename": {}, "mediaType": {}, "size": {}, "syncState": {}, "remoteVersionId": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireObjectArtifactSource) UnmarshalJSON(data []byte) error {
	type plain WireObjectArtifactSource
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireObjectArtifactSource(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "objectId": {}, "filename": {}, "mediaType": {}, "size": {}, "storageProvider": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireObjectArtifactSource) MarshalJSON() ([]byte, error) {
	type plain WireObjectArtifactSource
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "objectId": {}, "filename": {}, "mediaType": {}, "size": {}, "storageProvider": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireURLArtifactSource) UnmarshalJSON(data []byte) error {
	type plain WireURLArtifactSource
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireURLArtifactSource(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "url": {}, "provider": {}, "mediaType": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireURLArtifactSource) MarshalJSON() ([]byte, error) {
	type plain WireURLArtifactSource
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "url": {}, "provider": {}, "mediaType": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireHostedArtifactSource) UnmarshalJSON(data []byte) error {
	type plain WireHostedArtifactSource
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireHostedArtifactSource(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "recordType": {}, "recordId": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireHostedArtifactSource) MarshalJSON() ([]byte, error) {
	type plain WireHostedArtifactSource
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "recordType": {}, "recordId": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireProviderArtifactSource) UnmarshalJSON(data []byte) error {
	type plain WireProviderArtifactSource
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireProviderArtifactSource(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "provider": {}, "reference": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireProviderArtifactSource) MarshalJSON() ([]byte, error) {
	type plain WireProviderArtifactSource
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "provider": {}, "reference": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactSpecification) UnmarshalJSON(data []byte) error {
	type plain WireArtifactSpecification
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactSpecification(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schema": {}, "version": {}, "value": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactSpecification) MarshalJSON() ([]byte, error) {
	type plain WireArtifactSpecification
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schema": {}, "version": {}, "value": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactVersion) UnmarshalJSON(data []byte) error {
	type plain WireArtifactVersion
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactVersion(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schemaVersion": {}, "id": {}, "artifactId": {}, "version": {}, "source": {}, "integrity": {}, "specification": {}, "createdBy": {}, "createdAt": {}, "note": {}, "metadata": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactVersion) MarshalJSON() ([]byte, error) {
	type plain WireArtifactVersion
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schemaVersion": {}, "id": {}, "artifactId": {}, "version": {}, "source": {}, "integrity": {}, "specification": {}, "createdBy": {}, "createdAt": {}, "note": {}, "metadata": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifact) UnmarshalJSON(data []byte) error {
	type plain WireArtifact
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifact(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schemaVersion": {}, "id": {}, "specId": {}, "scope": {}, "kind": {}, "valueType": {}, "title": {}, "description": {}, "specification": {}, "currentVersionId": {}, "createdBy": {}, "createdAt": {}, "updatedAt": {}, "archivedAt": {}, "metadata": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifact) MarshalJSON() ([]byte, error) {
	type plain WireArtifact
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schemaVersion": {}, "id": {}, "specId": {}, "scope": {}, "kind": {}, "valueType": {}, "title": {}, "description": {}, "specification": {}, "currentVersionId": {}, "createdBy": {}, "createdAt": {}, "updatedAt": {}, "archivedAt": {}, "metadata": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactAccessPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactAccessPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactAccessPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"read": {}, "write": {}, "submit": {}, "verify": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactAccessPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactAccessPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"read": {}, "write": {}, "submit": {}, "verify": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactAccessPolicyRule) UnmarshalJSON(data []byte) error {
	type plain WireArtifactAccessPolicyRule
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactAccessPolicyRule(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"actors": {}, "condition": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactAccessPolicyRule) MarshalJSON() ([]byte, error) {
	type plain WireArtifactAccessPolicyRule
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"actors": {}, "condition": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactConditionState) UnmarshalJSON(data []byte) error {
	type plain WireArtifactConditionState
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactConditionState(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"kind": {}, "namespace": {}, "in": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactConditionState) MarshalJSON() ([]byte, error) {
	type plain WireArtifactConditionState
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"kind": {}, "namespace": {}, "in": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactConditionActor) UnmarshalJSON(data []byte) error {
	type plain WireArtifactConditionActor
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactConditionActor(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"kind": {}, "in": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactConditionActor) MarshalJSON() ([]byte, error) {
	type plain WireArtifactConditionActor
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"kind": {}, "in": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactConditionArtifactExists) UnmarshalJSON(data []byte) error {
	type plain WireArtifactConditionArtifactExists
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactConditionArtifactExists(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"kind": {}, "artifact": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactConditionArtifactExists) MarshalJSON() ([]byte, error) {
	type plain WireArtifactConditionArtifactExists
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"kind": {}, "artifact": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactConditionArtifactValue) UnmarshalJSON(data []byte) error {
	type plain WireArtifactConditionArtifactValue
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactConditionArtifactValue(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"kind": {}, "artifact": {}, "operator": {}, "value": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactConditionArtifactValue) MarshalJSON() ([]byte, error) {
	type plain WireArtifactConditionArtifactValue
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"kind": {}, "artifact": {}, "operator": {}, "value": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactConditionAnd) UnmarshalJSON(data []byte) error {
	type plain WireArtifactConditionAnd
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactConditionAnd(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"kind": {}, "conditions": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactConditionAnd) MarshalJSON() ([]byte, error) {
	type plain WireArtifactConditionAnd
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"kind": {}, "conditions": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactConditionOr) UnmarshalJSON(data []byte) error {
	type plain WireArtifactConditionOr
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactConditionOr(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"kind": {}, "conditions": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactConditionOr) MarshalJSON() ([]byte, error) {
	type plain WireArtifactConditionOr
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"kind": {}, "conditions": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactConditionNot) UnmarshalJSON(data []byte) error {
	type plain WireArtifactConditionNot
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactConditionNot(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"kind": {}, "condition": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactConditionNot) MarshalJSON() ([]byte, error) {
	type plain WireArtifactConditionNot
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"kind": {}, "condition": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactLifecyclePolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactLifecyclePolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactLifecyclePolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"createAt": {}, "editableDuring": {}, "submitDuring": {}, "lockAt": {}, "invalidateOn": {}, "condition": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactLifecyclePolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactLifecyclePolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"createAt": {}, "editableDuring": {}, "submitDuring": {}, "lockAt": {}, "invalidateOn": {}, "condition": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactPresentationPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactPresentationPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactPresentationPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"label": {}, "helpText": {}, "order": {}, "display": {}, "config": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactPresentationPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactPresentationPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"label": {}, "helpText": {}, "order": {}, "display": {}, "config": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactPrivacyPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactPrivacyPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactPrivacyPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"classification": {}, "reveal": {}, "masking": {}, "encryption": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactPrivacyPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactPrivacyPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"classification": {}, "reveal": {}, "masking": {}, "encryption": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactRevealRule) UnmarshalJSON(data []byte) error {
	type plain WireArtifactRevealRule
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactRevealRule(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"actors": {}, "when": {}, "representation": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactRevealRule) MarshalJSON() ([]byte, error) {
	type plain WireArtifactRevealRule
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"actors": {}, "when": {}, "representation": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactMaskingPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactMaskingPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactMaskingPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"strategy": {}, "config": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactMaskingPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactMaskingPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"strategy": {}, "config": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactEncryptionPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactEncryptionPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactEncryptionPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"required": {}, "level": {}, "keyScope": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactEncryptionPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactEncryptionPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"required": {}, "level": {}, "keyScope": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactProviderPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactProviderPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactProviderPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"actors": {}, "mode": {}, "delegation": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactProviderPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactProviderPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"actors": {}, "mode": {}, "delegation": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactRequirementPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactRequirementPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactRequirementPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"mode": {}, "condition": {}, "blocks": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactRequirementPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactRequirementPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"mode": {}, "condition": {}, "blocks": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactRetentionPolicyForever) UnmarshalJSON(data []byte) error {
	type plain WireArtifactRetentionPolicyForever
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactRetentionPolicyForever(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"policy": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactRetentionPolicyForever) MarshalJSON() ([]byte, error) {
	type plain WireArtifactRetentionPolicyForever
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"policy": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactRetentionPolicyDuration) UnmarshalJSON(data []byte) error {
	type plain WireArtifactRetentionPolicyDuration
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactRetentionPolicyDuration(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"policy": {}, "days": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactRetentionPolicyDuration) MarshalJSON() ([]byte, error) {
	type plain WireArtifactRetentionPolicyDuration
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"policy": {}, "days": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactRetentionPolicyUntil) UnmarshalJSON(data []byte) error {
	type plain WireArtifactRetentionPolicyUntil
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactRetentionPolicyUntil(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"policy": {}, "date": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactRetentionPolicyUntil) MarshalJSON() ([]byte, error) {
	type plain WireArtifactRetentionPolicyUntil
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"policy": {}, "date": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactRetentionPolicyHostDefined) UnmarshalJSON(data []byte) error {
	type plain WireArtifactRetentionPolicyHostDefined
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactRetentionPolicyHostDefined(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"policy": {}, "key": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactRetentionPolicyHostDefined) MarshalJSON() ([]byte, error) {
	type plain WireArtifactRetentionPolicyHostDefined
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"policy": {}, "key": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValidationPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValidationPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValidationPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"mode": {}, "schema": {}, "rules": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValidationPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValidationPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"mode": {}, "schema": {}, "rules": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValidationRule) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValidationRule
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValidationRule(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "config": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValidationRule) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValidationRule
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "config": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactVerificationPolicy) UnmarshalJSON(data []byte) error {
	type plain WireArtifactVerificationPolicy
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactVerificationPolicy(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"required": {}, "methods": {}, "actors": {}, "condition": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactVerificationPolicy) MarshalJSON() ([]byte, error) {
	type plain WireArtifactVerificationPolicy
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"required": {}, "methods": {}, "actors": {}, "condition": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactProtocolCommonContractsActorReference) UnmarshalJSON(data []byte) error {
	type plain WireArtifactProtocolCommonContractsActorReference
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactProtocolCommonContractsActorReference(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "id": {}, "displayName": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactProtocolCommonContractsActorReference) MarshalJSON() ([]byte, error) {
	type plain WireArtifactProtocolCommonContractsActorReference
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "id": {}, "displayName": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactProtocolCommonContractsArtifactScopeReference) UnmarshalJSON(data []byte) error {
	type plain WireArtifactProtocolCommonContractsArtifactScopeReference
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactProtocolCommonContractsArtifactScopeReference(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "id": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactProtocolCommonContractsArtifactScopeReference) MarshalJSON() ([]byte, error) {
	type plain WireArtifactProtocolCommonContractsArtifactScopeReference
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "id": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactProtocolCommonContractsArtifactSubjectReference) UnmarshalJSON(data []byte) error {
	type plain WireArtifactProtocolCommonContractsArtifactSubjectReference
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactProtocolCommonContractsArtifactSubjectReference(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "id": {}, "scope": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactProtocolCommonContractsArtifactSubjectReference) MarshalJSON() ([]byte, error) {
	type plain WireArtifactProtocolCommonContractsArtifactSubjectReference
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "id": {}, "scope": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactSubmission) UnmarshalJSON(data []byte) error {
	type plain WireArtifactSubmission
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactSubmission(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schemaVersion": {}, "id": {}, "artifactId": {}, "artifactVersionId": {}, "submittedBy": {}, "value": {}, "submittedAt": {}, "context": {}, "metadata": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactSubmission) MarshalJSON() ([]byte, error) {
	type plain WireArtifactSubmission
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schemaVersion": {}, "id": {}, "artifactId": {}, "artifactVersionId": {}, "submittedBy": {}, "value": {}, "submittedAt": {}, "context": {}, "metadata": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactSubmissionContext) UnmarshalJSON(data []byte) error {
	type plain WireArtifactSubmissionContext
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactSubmissionContext(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"latitude": {}, "longitude": {}, "deviceId": {}, "ipAddress": {}, "userAgent": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactSubmissionContext) MarshalJSON() ([]byte, error) {
	type plain WireArtifactSubmissionContext
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"latitude": {}, "longitude": {}, "deviceId": {}, "ipAddress": {}, "userAgent": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactVerification) UnmarshalJSON(data []byte) error {
	type plain WireArtifactVerification
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactVerification(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schemaVersion": {}, "id": {}, "artifactId": {}, "artifactVersionId": {}, "submissionId": {}, "status": {}, "method": {}, "verifiedBy": {}, "reason": {}, "createdAt": {}, "verifiedAt": {}, "metadata": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactVerification) MarshalJSON() ([]byte, error) {
	type plain WireArtifactVerification
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schemaVersion": {}, "id": {}, "artifactId": {}, "artifactVersionId": {}, "submissionId": {}, "status": {}, "method": {}, "verifiedBy": {}, "reason": {}, "createdAt": {}, "verifiedAt": {}, "metadata": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactRequirement) UnmarshalJSON(data []byte) error {
	type plain WireArtifactRequirement
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactRequirement(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schemaVersion": {}, "id": {}, "key": {}, "allowedKinds": {}, "allowedValueTypes": {}, "minimumCount": {}, "maximumCount": {}, "required": {}, "specId": {}, "metadata": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactRequirement) MarshalJSON() ([]byte, error) {
	type plain WireArtifactRequirement
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schemaVersion": {}, "id": {}, "key": {}, "allowedKinds": {}, "allowedValueTypes": {}, "minimumCount": {}, "maximumCount": {}, "required": {}, "specId": {}, "metadata": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactSpecSnapshot) UnmarshalJSON(data []byte) error {
	type plain WireArtifactSpecSnapshot
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactSpecSnapshot(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schemaVersion": {}, "sourceSpecId": {}, "sourceVersion": {}, "key": {}, "name": {}, "description": {}, "kind": {}, "valueType": {}, "config": {}, "provider": {}, "requirement": {}, "lifecycle": {}, "access": {}, "privacy": {}, "validation": {}, "verification": {}, "retention": {}, "presentation": {}, "metadata": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactSpecSnapshot) MarshalJSON() ([]byte, error) {
	type plain WireArtifactSpecSnapshot
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schemaVersion": {}, "sourceSpecId": {}, "sourceVersion": {}, "key": {}, "name": {}, "description": {}, "kind": {}, "valueType": {}, "config": {}, "provider": {}, "requirement": {}, "lifecycle": {}, "access": {}, "privacy": {}, "validation": {}, "verification": {}, "retention": {}, "presentation": {}, "metadata": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactSpec) UnmarshalJSON(data []byte) error {
	type plain WireArtifactSpec
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactSpec(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"schemaVersion": {}, "id": {}, "key": {}, "name": {}, "description": {}, "version": {}, "kind": {}, "valueType": {}, "config": {}, "provider": {}, "requirement": {}, "lifecycle": {}, "access": {}, "privacy": {}, "validation": {}, "verification": {}, "retention": {}, "presentation": {}, "metadata": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactSpec) MarshalJSON() ([]byte, error) {
	type plain WireArtifactSpec
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"schemaVersion": {}, "id": {}, "key": {}, "name": {}, "description": {}, "version": {}, "kind": {}, "valueType": {}, "config": {}, "provider": {}, "requirement": {}, "lifecycle": {}, "access": {}, "privacy": {}, "validation": {}, "verification": {}, "retention": {}, "presentation": {}, "metadata": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaText) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaText
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaText(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "minLength": {}, "maxLength": {}, "multiline": {}, "pattern": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaText) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaText
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "minLength": {}, "maxLength": {}, "multiline": {}, "pattern": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaNumber) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaNumber
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaNumber(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "minimum": {}, "maximum": {}, "integer": {}, "multipleOf": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaNumber) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaNumber
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "minimum": {}, "maximum": {}, "integer": {}, "multipleOf": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaBoolean) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaBoolean
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaBoolean(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaBoolean) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaBoolean
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaCurrency) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaCurrency
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaCurrency(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "currencies": {}, "minimumMinorUnits": {}, "maximumMinorUnits": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaCurrency) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaCurrency
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "currencies": {}, "minimumMinorUnits": {}, "maximumMinorUnits": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaDate) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaDate
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaDate(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "minimum": {}, "maximum": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaDate) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaDate
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "minimum": {}, "maximum": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaDatetime) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaDatetime
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaDatetime(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "minimum": {}, "maximum": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaDatetime) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaDatetime
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "minimum": {}, "maximum": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaTime) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaTime
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaTime(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "minimum": {}, "maximum": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaTime) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaTime
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "minimum": {}, "maximum": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaLocation) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaLocation
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaLocation(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "mode": {}, "requireCoordinates": {}, "allowManualEntry": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaLocation) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaLocation
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "mode": {}, "requireCoordinates": {}, "allowManualEntry": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaFile) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaFile
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaFile(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}, "valueType": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaFile) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaFile
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}, "valueType": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaImage) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaImage
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaImage(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}, "valueType": {}, "requireTimestamp": {}, "requireLocation": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaImage) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaImage
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}, "valueType": {}, "requireTimestamp": {}, "requireLocation": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaVideo) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaVideo
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaVideo(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}, "valueType": {}, "maxDurationSeconds": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaVideo) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaVideo
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}, "valueType": {}, "maxDurationSeconds": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaAudio) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaAudio
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaAudio(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}, "valueType": {}, "maxDurationSeconds": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaAudio) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaAudio
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}, "valueType": {}, "maxDurationSeconds": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaFileBase) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaFileBase
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaFileBase(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaFileBase) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaFileBase
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"minFiles": {}, "maxFiles": {}, "acceptedMimeTypes": {}, "maxSizeBytes": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaLink) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaLink
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaLink(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "allowedSchemes": {}, "allowedHosts": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaLink) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaLink
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "allowedSchemes": {}, "allowedHosts": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaStructured) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaStructured
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaStructured(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "jsonSchema": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaStructured) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaStructured
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "jsonSchema": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaReference) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaReference
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaReference(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "providers": {}, "resourceTypes": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaReference) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaReference
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "providers": {}, "resourceTypes": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaSignature) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaSignature
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaSignature(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "methods": {}, "requireTimestamp": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaSignature) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaSignature
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "methods": {}, "requireTimestamp": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireArtifactValueSchemaCollection) UnmarshalJSON(data []byte) error {
	type plain WireArtifactValueSchemaCollection
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireArtifactValueSchemaCollection(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"valueType": {}, "itemSchema": {}, "minItems": {}, "maxItems": {}, "uniqueItems": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireArtifactValueSchemaCollection) MarshalJSON() ([]byte, error) {
	type plain WireArtifactValueSchemaCollection
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"valueType": {}, "itemSchema": {}, "minItems": {}, "maxItems": {}, "uniqueItems": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}
