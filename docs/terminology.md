# Terminology

`Artifact` is logical identity. `ArtifactVersion` is an immutable content
revision and owns one `ArtifactSource`. `ArtifactLink` describes a usage or
relationship and can follow the logical artifact or pin one exact version.

`Artifact.kind` is open domain meaning, such as `delivery_evidence` or
`implementation`. `Artifact.valueType` is the closed representation vocabulary,
such as `image` or `reference`. Meaning and representation are intentionally
independent.

`ArtifactSpec` is a versioned declarative definition, never an actual artifact.
`ArtifactSpecSnapshot` is the immutable definition captured for historical
behavior. `ArtifactRequirement` can declare a need without owning a full spec.

`ArtifactSubmission` records who submitted what and when. `ArtifactVerification`
records an auditable outcome. Neither is a boolean field on `Artifact`.

Actor, scope, subject, kind, role, provider, workflow-stage, operation, and
verification-method identifiers are host vocabulary. Opaque identifiers are
case-sensitive non-empty strings; the protocol assigns no UUID or database
meaning to them.
