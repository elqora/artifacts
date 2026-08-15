# Artifact SDK

This repository defines a language-neutral protocol for typed, identifiable,
versionable artifacts. It standardizes artifact identity, content revisions,
sources, integrity, and relationships while leaving persistence, storage,
workflow, authorization, provider APIs, and UI to host applications.

The repository contains:

- canonical JSON Schemas under `spec/`;
- TypeScript, PHP, and Go bindings under `packages/`;
- a typed GitHub provider extension under `extensions/github/`;
- cross-language fixtures under `tests/fixtures/`;
- schema conformance tests under `tests/conformance/`;
- interoperability rules and architecture documentation under `docs/`.

The architectural brief is [GOAL.md](./GOAL.md). JSON Schema is the protocol
authority; the TypeScript package is the first developer-facing binding.

## Development

The full suite requires Node.js 20+, PHP 8.2+, Composer, and Go 1.22+. If Go is
not installed locally, the test runner uses Docker with the Go 1.22 image.

```sh
npm install
composer install --working-dir=packages/php
npm test
npm run build
npm run generate
```

The canonical schemas implement Artifact Protocol `1.0` using JSON Schema
Draft 2020-12. Run `npm run check:generated` to prove generated binding
vocabularies match those schemas.

## TypeScript package

The release-ready TypeScript SDK is `@elqora/artifacts`:

```sh
npm install @elqora/artifacts
```

```ts
import type { Artifact, ArtifactSpec, ArtifactVersion } from "@elqora/artifacts";
```

It is an ESM-only package for Node.js 20 and newer. The package has not yet
been published; `npm run test:typescript-package` builds a tarball, installs it
in a clean temporary project, and verifies both its type and runtime exports.

`npm test` validates deterministic generation, TypeScript, every JSON Schema
fixture, generated PHP DTOs and static analysis, both Go modules, package
manifests, package contents, and clean installation of `@elqora/artifacts`.

Start with [architecture](./docs/architecture.md),
[terminology](./docs/terminology.md), and
[serialization](./docs/serialization.md). Policy declarations express host
intent; they do not enforce access, privacy, encryption, retention, validation,
or workflow behavior.
