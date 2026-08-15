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
