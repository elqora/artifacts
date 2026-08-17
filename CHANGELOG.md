# Changelog

All notable changes to the TypeScript SDK will be documented in this file.

The project follows Semantic Versioning for package releases. During `0.x`,
changes remain additive by default and breaking corrections must be called out
explicitly.

## 0.2.0 - Unreleased

### Added

- Artifact Protocol 1.1 and `ArtifactSpecification`, a schema-identified
  semantic interpretation envelope for artifacts and immutable versions.
- Canonical JSON Schemas for the Artifact Protocol and GitHub extension.
- Generated TypeScript wire contracts and vocabulary constants.
- Ergonomic TypeScript contracts for artifacts, versions, sources, links,
  specifications, policies, submissions, and verification.
- Cross-language conformance fixtures and deterministic binding generation.

### Packaging

- Prepared the ESM-only `@elqora/artifacts` package for Node.js 20 and newer.
- Added self-contained package documentation, licensing, source maps, and
  clean-install verification.
