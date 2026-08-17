# Versioning and snapshots

Four version axes are independent:

- `schemaVersion` identifies the wire protocol (`1.1` here).
- `ArtifactSpec.version` revises a definition.
- `ArtifactVersion.version` revises actual artifact content.
- `ArtifactSpecification.version` revises the schema governing semantic data.

A compatible additive protocol change may be introduced within a major line;
a breaking wire change requires a new protocol major version. Consumers must
not validate an unsupported major as 1.1.

Hosts capture an `ArtifactSpecSnapshot` when historical behavior matters. A
workflow that captured `delivery_evidence` v3 keeps its v3 constraints even if
the live spec advances to v4. Likewise, a link, submission, or verification
that contains `artifactVersionId` remains attached to exact content when the
artifact later gains a new version.
