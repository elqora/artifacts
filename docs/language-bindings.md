# Language bindings and generation

JSON Schema under `spec/` is authoritative. TypeScript exposes discriminated
unions and generics; Go exposes JSON-tagged structs and raw opaque values; PHP
uses lossless associative-array wire records with PHPStan array-shape contracts.
These ergonomic differences do not change serialization.

`npm run generate` loads all core and GitHub schemas, resolves their references,
builds an in-memory type graph, and generates:

- TypeScript wire contracts and vocabulary constants;
- Go `Wire*` records with lossless unknown-field codecs;
- immutable PHP DTOs, accessors, and backed enums;
- the equivalent GitHub-extension contracts.

The in-memory graph is derived on every run and is not another protocol source.
Generation supports references, recursive definitions, unions, composition,
constants, enums, arrays, maps, optional fields, and open properties. Validation
keywords remain the responsibility of JSON Schema. Unsupported structural
keywords fail generation rather than degrading silently.

Generated files say `DO NOT EDIT`. `npm run check:generated` fails when they
drift, while generator tests prove deterministic output and demonstrate that a
schema field addition changes TypeScript, Go, and PHP output.

The existing generic TypeScript API and unprefixed Go structs remain compatible
facades. New Go code can use generated `Wire*` records and discriminator-aware
decoders. PHP retains associative-array helpers while adding generated DTOs and
factories. Generated and facade contracts share the same JSON representation.

## TypeScript package surface

`@elqora/artifacts` is the supported TypeScript package. It is ESM-only,
requires Node.js 20 or newer, and exposes a single root entry point containing:

- ergonomic generic contracts such as `Artifact`, `ArtifactVersion`, and
  `ArtifactSpec`;
- schema-generated `Wire*` contracts;
- protocol and closed-vocabulary constants.

The generated contracts and handwritten facade are both public during `0.x`.
Changes are additive by default: a facade must not introduce wire fields or
semantics absent from the canonical schemas, and public names are not removed
without an explicitly documented breaking correction. The npm version is
independent of the serialized protocol `schemaVersion`.

Before publication, `npm run test:typescript-package` packs the workspace,
checks its file allowlist and embedded source-map content, installs the tarball
into a clean temporary project, compiles strict consumer code, and imports its
runtime vocabulary constants.
