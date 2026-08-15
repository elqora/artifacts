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

Core schema and TypeScript development requires Node.js 20 or newer. The full
local suite also exercises the PHP binding when PHP 8.2+ is available; Go
conformance can be run separately with Go 1.22+.

```sh
npm install
npm test
npm run build
npm run generate

# when Go is installed
cd packages/go && go test ./...
```

The canonical schemas implement Artifact Protocol `1.0` using JSON Schema
Draft 2020-12. Run `npm run check:generated` to prove generated binding
vocabularies match those schemas.

Start with [architecture](./docs/architecture.md),
[terminology](./docs/terminology.md), and
[serialization](./docs/serialization.md). Policy declarations express host
intent; they do not enforce access, privacy, encryption, retention, validation,
or workflow behavior.
