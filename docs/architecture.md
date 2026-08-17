# Architecture

The canonical dependency direction is:

```text
Artifact (logical identity)
  -> ArtifactSpecification (optional stable semantic interpretation)
  -> ArtifactVersion (immutable revision)
       -> ArtifactSpecification (optional revision-specific interpretation)
       -> ArtifactSource (inline, local, object, URL, hosted, or provider)
       -> ArtifactIntegrity (optional digest)

ArtifactLink (usage relationship)
  -> Artifact
  -> ArtifactVersion (optional exact pin)
  -> ArtifactSubjectReference

ArtifactSpec (versioned definition)
  -> ArtifactSpecSnapshot (immutable historical definition)
  -> ArtifactValueSchema
  -> declarative policies and conditions

ArtifactSubmission (provenance-bearing operational event)
  -> Artifact / optional exact ArtifactVersion

ArtifactVerification (auditable decision record)
  -> Artifact / optional exact ArtifactVersion / optional Submission
```

These concepts remain separate. Context such as evidence, approval, attachment,
or deliverable is an extensible link role, never a flag on the artifact.
Provider-specific reference contracts belong in extension packages. The GitHub
extension demonstrates that boundary without an API client.

The JSON Schemas under `spec/` define the wire contract. Language bindings
expose the same contract ergonomically but do not own or redefine it.

The SDK contains no persistence, HTTP API, provider client, authorization
engine, encryption implementation, policy evaluator, workflow, storage, or UI.
Hosts resolve open identifiers and enforce declared policy.
