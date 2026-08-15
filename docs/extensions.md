# Provider extensions

Core `ProviderArtifactSource` contains an open provider identifier and opaque
reference. A consumer without an extension preserves that reference unchanged.
An extension adds typed reference schemas without changing core semantics.

`extensions/github` demonstrates pull request, commit, issue, discussion,
deployment, workflow-run, and check-run references in JSON Schema, TypeScript,
PHP, and Go. It intentionally contains no GitHub client or network behavior.
GitLab, Figma, Jira, or Drive can follow the same pattern with new provider and
resource identifiers; the core source union does not change.

Extension schemas are loaded by the same generator as core schemas. The GitHub
TypeScript wire union, PHP DTOs, and Go wire structs are therefore regenerated
whenever its canonical reference schemas change. Handwritten extension helpers
may select variants, but must not redefine their serialized fields.
