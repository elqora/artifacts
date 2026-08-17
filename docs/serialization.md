# Serialization and compatibility

These rules settle the wire-level choices required for Artifact Protocol 1.1.
The canonical JSON Schemas remain authoritative if this document and a schema
ever disagree.

## Protocol version

Every top-level artifact, specification, requirement, submission, and
verification record must
contain `"schemaVersion": "1.1"`. This version describes the serialized
protocol shape; it is independent of `ArtifactVersion.version` and future
artifact-specification versions. A breaking wire-format change requires a new
protocol major version. Compatible additions may use a new minor version and a
corresponding schema.

## JSON representation

- Records are JSON objects. Property names and standardized values are
  case-sensitive.
- Identifiers are non-empty, opaque, case-sensitive strings. The protocol does
  not assign UUID, ULID, numeric, or provider-specific semantics to them.
- Timestamps are RFC 3339 date-time strings. Producers should emit UTC with a
  `Z` suffix; consumers must accept valid RFC 3339 offsets.
- Optional properties are omitted when no value is present. `null` is not a
  substitute for omission unless the property's schema explicitly accepts an
  arbitrary JSON value, such as an inline value, semantic-specification value,
  or metadata entry.
- Numeric sizes are non-negative integers measured in bytes. Artifact version
  numbers are positive integers starting at 1. They increase within one
  artifact; gaps are permitted because allocation and persistence are host
  concerns.
- Integrity digests are lowercase hexadecimal strings with the length implied
  by `sha256`, `sha384`, or `sha512`.

## Open and closed vocabulary

Standardized values use the exact lowercase strings defined by the schemas.
This includes artifact value types, source and condition discriminators, local
synchronization states, integrity algorithms, verification status, privacy
classification and representation, and policy discriminators.

Application vocabulary remains open. In particular, artifact kinds, link
roles, subject types, actor types, provider identifiers, hosted record types,
and storage-provider identifiers are non-empty host-defined strings.

## Extensions and unknown properties

Protocol records permit unknown object properties so that additive protocol and
extension fields can pass through older implementations. Consumers should
preserve fields they do not understand when reading and re-emitting a record.
Unknown source discriminator values still fail 1.1 validation because a
consumer cannot safely infer that source's required structure.

`metadata` is an object whose members may contain any JSON value. It is for
host extensions and non-core information. A field needed for shared semantics
should be promoted into a future protocol schema rather than hidden in
metadata. Provider `reference` and inline `value` are deliberately opaque JSON
values for the same language-neutral reason.

## Identity, versions, and links

An `Artifact` is stable logical identity. An `ArtifactVersion` is an immutable
content revision and carries exactly one `ArtifactSource`. `currentVersionId`
is a host-maintained pointer; changing it does not mutate prior versions.

An `ArtifactLink` without `artifactVersionId` follows the logical artifact. A
link with `artifactVersionId` is pinned to that exact immutable revision. JSON
Schema validates identifiers and presence, but cross-record rules—such as a
pinned version belonging to the same artifact—must be checked by a conformance
or host validation layer that has access to both records.

## Instance semantic specifications

`ArtifactSpecification` is separate from `ArtifactSpec`: it identifies the
schema and revision governing semantic interpretation of an actual artifact.
Its `schema` is a non-empty case-sensitive host/domain identifier, `version`
is a positive integer, and `value` is any JSON value (including `null`). The
core protocol does not interpret or validate that value beyond its JSON shape.

An `Artifact.specification` applies to stable logical meaning. An
`ArtifactVersion.specification` applies only to that immutable content
revision, such as a page, time range, line range, or image region. The fields
are independent: neither overrides, merges with, nor is inherited from the
other; hosts set a new version's specification explicitly.

## Forward-compatible parsing

Implementations should select the schema by `schemaVersion`, validate known
fields, retain unknown fields, and keep provider references opaque unless an
extension understands them. Encountering an unsupported protocol major version
must not be reported as successful validation against 1.1.

JSON numbers follow JSON interoperability limits. Hosts needing exact decimal
or very large integer semantics should use a structured/string representation
defined by their value schema rather than relying on binary floating point.
Protocol enums and identifiers are case-sensitive. Missing optional fields and
explicit `null` are different; `null` is accepted only where a schema permits
it as an opaque JSON value. Provider references are opaque to core bindings and
must survive round trips unchanged.
