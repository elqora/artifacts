// Package artifact exposes Artifact Protocol 1.0 records without implementing host behavior.
package artifact

import "encoding/json"

const ProtocolVersion = "1.0"

type Metadata map[string]any
type ActorReference struct {
	Type        string `json:"type"`
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
type ScopeReference struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}
type SubjectReference struct {
	Type  string            `json:"type"`
	ID    string            `json:"id"`
	Scope map[string]string `json:"scope,omitempty"`
}

type Artifact struct {
	SchemaVersion    string          `json:"schemaVersion"`
	ID               string          `json:"id"`
	Scope            *ScopeReference `json:"scope,omitempty"`
	Kind             string          `json:"kind"`
	ValueType        string          `json:"valueType"`
	Title            string          `json:"title,omitempty"`
	Description      string          `json:"description,omitempty"`
	CurrentVersionID string          `json:"currentVersionId,omitempty"`
	CreatedBy        ActorReference  `json:"createdBy"`
	CreatedAt        string          `json:"createdAt"`
	UpdatedAt        string          `json:"updatedAt"`
	ArchivedAt       string          `json:"archivedAt,omitempty"`
	Metadata         Metadata        `json:"metadata,omitempty"`
}

type InlineArtifactSource struct {
	Type      string          `json:"type"`
	Value     json.RawMessage `json:"value"`
	MediaType string          `json:"mediaType,omitempty"`
}
type LocalArtifactSource struct {
	Type            string `json:"type"`
	LocalID         string `json:"localId"`
	Filename        string `json:"filename,omitempty"`
	MediaType       string `json:"mediaType,omitempty"`
	Size            *int64 `json:"size,omitempty"`
	SyncState       string `json:"syncState"`
	RemoteVersionID string `json:"remoteVersionId,omitempty"`
}
type ObjectArtifactSource struct {
	Type            string `json:"type"`
	ObjectID        string `json:"objectId"`
	Filename        string `json:"filename,omitempty"`
	MediaType       string `json:"mediaType,omitempty"`
	Size            *int64 `json:"size,omitempty"`
	StorageProvider string `json:"storageProvider,omitempty"`
}
type URLArtifactSource struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	Provider  string `json:"provider,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
}
type HostedArtifactSource struct {
	Type       string `json:"type"`
	RecordType string `json:"recordType"`
	RecordID   string `json:"recordId"`
}
type ProviderArtifactSource struct {
	Type      string          `json:"type"`
	Provider  string          `json:"provider"`
	Reference json.RawMessage `json:"reference"`
}

type ArtifactIntegrity struct {
	Algorithm  string `json:"algorithm"`
	Hash       string `json:"hash"`
	Size       *int64 `json:"size,omitempty"`
	VerifiedAt string `json:"verifiedAt,omitempty"`
}
type ArtifactVersion struct {
	SchemaVersion string             `json:"schemaVersion"`
	ID            string             `json:"id"`
	ArtifactID    string             `json:"artifactId"`
	Version       int                `json:"version"`
	Source        json.RawMessage    `json:"source"`
	Integrity     *ArtifactIntegrity `json:"integrity,omitempty"`
	CreatedBy     ActorReference     `json:"createdBy"`
	CreatedAt     string             `json:"createdAt"`
	Note          string             `json:"note,omitempty"`
	Metadata      Metadata           `json:"metadata,omitempty"`
}
type ArtifactLink struct {
	SchemaVersion     string           `json:"schemaVersion"`
	ID                string           `json:"id"`
	ArtifactID        string           `json:"artifactId"`
	ArtifactVersionID string           `json:"artifactVersionId,omitempty"`
	Subject           SubjectReference `json:"subject"`
	Role              string           `json:"role"`
	Note              string           `json:"note,omitempty"`
	CreatedBy         ActorReference   `json:"createdBy"`
	CreatedAt         string           `json:"createdAt"`
	Metadata          Metadata         `json:"metadata,omitempty"`
}

// Condition is the recursive wire representation. Fields are interpreted according to Kind.
type Condition struct {
	Kind       string          `json:"kind"`
	Namespace  string          `json:"namespace,omitempty"`
	In         []string        `json:"in,omitempty"`
	Artifact   string          `json:"artifact,omitempty"`
	Operator   string          `json:"operator,omitempty"`
	Value      json.RawMessage `json:"value,omitempty"`
	Conditions []Condition     `json:"conditions,omitempty"`
	Condition  *Condition      `json:"condition,omitempty"`
}
type ProviderPolicy struct {
	Actors     []string `json:"actors"`
	Mode       string   `json:"mode,omitempty"`
	Delegation string   `json:"delegation,omitempty"`
}
type RequirementPolicy struct {
	Mode      string     `json:"mode"`
	Condition *Condition `json:"condition,omitempty"`
	Blocks    []string   `json:"blocks,omitempty"`
}
type LifecyclePolicy struct {
	CreateAt       string     `json:"createAt,omitempty"`
	EditableDuring []string   `json:"editableDuring,omitempty"`
	SubmitDuring   []string   `json:"submitDuring,omitempty"`
	LockAt         string     `json:"lockAt,omitempty"`
	InvalidateOn   []string   `json:"invalidateOn,omitempty"`
	Condition      *Condition `json:"condition,omitempty"`
}
type AccessRule struct {
	Actors    []string   `json:"actors"`
	Condition *Condition `json:"condition,omitempty"`
}
type AccessPolicy struct {
	Read   []AccessRule `json:"read,omitempty"`
	Write  []AccessRule `json:"write,omitempty"`
	Submit []AccessRule `json:"submit,omitempty"`
	Verify []AccessRule `json:"verify,omitempty"`
}
type RevealRule struct {
	Actors         []string   `json:"actors"`
	When           *Condition `json:"when,omitempty"`
	Representation string     `json:"representation"`
}
type MaskingPolicy struct {
	Strategy string         `json:"strategy"`
	Config   map[string]any `json:"config,omitempty"`
}
type EncryptionPolicy struct {
	Required bool   `json:"required"`
	Level    string `json:"level,omitempty"`
	KeyScope string `json:"keyScope,omitempty"`
}
type PrivacyPolicy struct {
	Classification string            `json:"classification"`
	Reveal         []RevealRule      `json:"reveal,omitempty"`
	Masking        *MaskingPolicy    `json:"masking,omitempty"`
	Encryption     *EncryptionPolicy `json:"encryption,omitempty"`
}
type ValidationRule struct {
	Type   string         `json:"type"`
	Config map[string]any `json:"config,omitempty"`
}
type ValidationPolicy struct {
	Mode  string           `json:"mode,omitempty"`
	Rules []ValidationRule `json:"rules,omitempty"`
}
type VerificationPolicy struct {
	Required  bool       `json:"required"`
	Methods   []string   `json:"methods,omitempty"`
	Actors    []string   `json:"actors,omitempty"`
	Condition *Condition `json:"condition,omitempty"`
}
type RetentionPolicy struct {
	Policy string `json:"policy"`
	Days   *int   `json:"days,omitempty"`
	Date   string `json:"date,omitempty"`
	Key    string `json:"key,omitempty"`
}
type PresentationPolicy struct {
	Label    string         `json:"label,omitempty"`
	HelpText string         `json:"helpText,omitempty"`
	Order    *float64       `json:"order,omitempty"`
	Display  string         `json:"display,omitempty"`
	Config   map[string]any `json:"config,omitempty"`
}

// ValueSchema preserves every standardized value-schema discriminator and its constraints.
type ValueSchema struct {
	ValueType          string         `json:"valueType"`
	MinLength          *int           `json:"minLength,omitempty"`
	MaxLength          *int           `json:"maxLength,omitempty"`
	Multiline          *bool          `json:"multiline,omitempty"`
	Pattern            string         `json:"pattern,omitempty"`
	Minimum            any            `json:"minimum,omitempty"`
	Maximum            any            `json:"maximum,omitempty"`
	Integer            *bool          `json:"integer,omitempty"`
	MultipleOf         *float64       `json:"multipleOf,omitempty"`
	Currencies         []string       `json:"currencies,omitempty"`
	MinimumMinorUnits  *int64         `json:"minimumMinorUnits,omitempty"`
	MaximumMinorUnits  *int64         `json:"maximumMinorUnits,omitempty"`
	Mode               string         `json:"mode,omitempty"`
	RequireCoordinates *bool          `json:"requireCoordinates,omitempty"`
	AllowManualEntry   *bool          `json:"allowManualEntry,omitempty"`
	MinFiles           *int           `json:"minFiles,omitempty"`
	MaxFiles           *int           `json:"maxFiles,omitempty"`
	AcceptedMimeTypes  []string       `json:"acceptedMimeTypes,omitempty"`
	MaxSizeBytes       *int64         `json:"maxSizeBytes,omitempty"`
	RequireTimestamp   *bool          `json:"requireTimestamp,omitempty"`
	RequireLocation    *bool          `json:"requireLocation,omitempty"`
	MaxDurationSeconds *float64       `json:"maxDurationSeconds,omitempty"`
	AllowedSchemes     []string       `json:"allowedSchemes,omitempty"`
	AllowedHosts       []string       `json:"allowedHosts,omitempty"`
	JSONSchema         map[string]any `json:"jsonSchema,omitempty"`
	Providers          []string       `json:"providers,omitempty"`
	ResourceTypes      []string       `json:"resourceTypes,omitempty"`
	Methods            []string       `json:"methods,omitempty"`
	ItemSchema         *ValueSchema   `json:"itemSchema,omitempty"`
	MinItems           *int           `json:"minItems,omitempty"`
	MaxItems           *int           `json:"maxItems,omitempty"`
	UniqueItems        *bool          `json:"uniqueItems,omitempty"`
}

type ArtifactSpec struct {
	SchemaVersion string              `json:"schemaVersion"`
	ID            string              `json:"id"`
	Key           string              `json:"key"`
	Name          string              `json:"name"`
	Description   string              `json:"description,omitempty"`
	Version       int                 `json:"version"`
	Kind          string              `json:"kind"`
	ValueType     string              `json:"valueType"`
	Config        *ValueSchema        `json:"config,omitempty"`
	Provider      *ProviderPolicy     `json:"provider,omitempty"`
	Requirement   *RequirementPolicy  `json:"requirement,omitempty"`
	Lifecycle     *LifecyclePolicy    `json:"lifecycle,omitempty"`
	Access        *AccessPolicy       `json:"access,omitempty"`
	Privacy       *PrivacyPolicy      `json:"privacy,omitempty"`
	Validation    *ValidationPolicy   `json:"validation,omitempty"`
	Verification  *VerificationPolicy `json:"verification,omitempty"`
	Retention     *RetentionPolicy    `json:"retention,omitempty"`
	Presentation  *PresentationPolicy `json:"presentation,omitempty"`
	Metadata      Metadata            `json:"metadata,omitempty"`
}
type ArtifactSpecSnapshot struct {
	SchemaVersion string              `json:"schemaVersion"`
	SourceSpecID  string              `json:"sourceSpecId"`
	SourceVersion int                 `json:"sourceVersion"`
	Key           string              `json:"key"`
	Name          string              `json:"name"`
	Description   string              `json:"description,omitempty"`
	Kind          string              `json:"kind"`
	ValueType     string              `json:"valueType"`
	Config        *ValueSchema        `json:"config,omitempty"`
	Provider      *ProviderPolicy     `json:"provider,omitempty"`
	Requirement   *RequirementPolicy  `json:"requirement,omitempty"`
	Lifecycle     *LifecyclePolicy    `json:"lifecycle,omitempty"`
	Access        *AccessPolicy       `json:"access,omitempty"`
	Privacy       *PrivacyPolicy      `json:"privacy,omitempty"`
	Validation    *ValidationPolicy   `json:"validation,omitempty"`
	Verification  *VerificationPolicy `json:"verification,omitempty"`
	Retention     *RetentionPolicy    `json:"retention,omitempty"`
	Presentation  *PresentationPolicy `json:"presentation,omitempty"`
	Metadata      Metadata            `json:"metadata,omitempty"`
}
type ArtifactRequirement struct {
	SchemaVersion     string   `json:"schemaVersion"`
	ID                string   `json:"id"`
	Key               string   `json:"key,omitempty"`
	AllowedKinds      []string `json:"allowedKinds,omitempty"`
	AllowedValueTypes []string `json:"allowedValueTypes,omitempty"`
	MinimumCount      *int     `json:"minimumCount,omitempty"`
	MaximumCount      *int     `json:"maximumCount,omitempty"`
	Required          bool     `json:"required"`
	SpecID            string   `json:"specId,omitempty"`
	Metadata          Metadata `json:"metadata,omitempty"`
}
type SubmissionContext struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	DeviceID  string   `json:"deviceId,omitempty"`
	IPAddress string   `json:"ipAddress,omitempty"`
	UserAgent string   `json:"userAgent,omitempty"`
}
type ArtifactSubmission struct {
	SchemaVersion     string             `json:"schemaVersion"`
	ID                string             `json:"id"`
	ArtifactID        string             `json:"artifactId"`
	ArtifactVersionID string             `json:"artifactVersionId,omitempty"`
	SubmittedBy       ActorReference     `json:"submittedBy"`
	Value             json.RawMessage    `json:"value,omitempty"`
	SubmittedAt       string             `json:"submittedAt"`
	Context           *SubmissionContext `json:"context,omitempty"`
	Metadata          Metadata           `json:"metadata,omitempty"`
}
type ArtifactVerification struct {
	SchemaVersion     string          `json:"schemaVersion"`
	ID                string          `json:"id"`
	ArtifactID        string          `json:"artifactId"`
	ArtifactVersionID string          `json:"artifactVersionId,omitempty"`
	SubmissionID      string          `json:"submissionId,omitempty"`
	Status            string          `json:"status"`
	Method            string          `json:"method,omitempty"`
	VerifiedBy        *ActorReference `json:"verifiedBy,omitempty"`
	Reason            string          `json:"reason,omitempty"`
	CreatedAt         string          `json:"createdAt"`
	VerifiedAt        string          `json:"verifiedAt,omitempty"`
	Metadata          Metadata        `json:"metadata,omitempty"`
}
