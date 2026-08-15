# Artifact specifications and value schemas

An `ArtifactSpec` describes an expected artifact and declarative policy. Its
`key`, `kind`, policy actors, stages, operations, and methods remain host-defined.
Its `valueType` and value-schema discriminator use the protocol vocabulary.
When `config` is present, its `valueType` must match the containing spec.

Canonical value schemas cover text, number, boolean, currency, date, datetime,
time, location, file, image, video, audio, link, structured data, references,
signatures, and recursive collections. These are representation constraints,
not an exhaustive catalogue of domain meanings.

`ArtifactRequirement` supports allowed kinds/value types, counts, requiredness,
and an optional spec reference. Count consistency across multiple records is a
host/conformance concern because JSON Schema cannot inspect a host collection.
