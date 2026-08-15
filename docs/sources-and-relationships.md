# Sources, relationships, and version pinning

An `ArtifactVersion` owns exactly one source. Inline sources carry opaque JSON;
local sources track offline synchronization; object sources name storage-neutral
objects; URL sources name external resources; hosted sources name internal host
records; provider sources carry an open provider identifier and opaque reference.
Provider-specific typing belongs in an extension.

Artifact usage is always an `ArtifactLink`. Its subject type and role are open
host vocabulary, allowing one artifact to participate in several contexts
without duplication. A link without `artifactVersionId` follows the logical
artifact and therefore its current revision. A link with `artifactVersionId`
pins immutable content. Reviews, evidence, approvals, and handovers should pin a
version whenever historical correctness matters.

Integrity digests belong to versions, not logical artifacts. The host computes
and verifies bytes; the protocol only records algorithm, digest, optional size,
and verification time.
