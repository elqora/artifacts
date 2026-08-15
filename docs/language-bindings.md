# Language bindings and generation

JSON Schema under `spec/` is authoritative. TypeScript exposes discriminated
unions and generics; Go exposes JSON-tagged structs and raw opaque values; PHP
uses lossless associative-array wire records with PHPStan array-shape contracts.
These ergonomic differences do not change serialization.

`npm run generate` reads protocol-owned vocabularies from canonical schemas and
recreates `packages/go/vocabulary_gen.go` and
`packages/php/src/Vocabulary.php`. Generated files say `DO NOT EDIT`.
`npm run check:generated` fails when they drift and is part of `npm test`.
Record-shape changes originate in schemas, are reflected in binding contracts,
and are proved against shared fixtures.

TypeScript and its GitHub extension build with `npm run build`. PHP conformance
runs in the root suite. Go users run `go test ./...` inside `packages/go` when
Go 1.22+ is installed.
