# Contributing

Protocol changes begin in the canonical JSON Schemas. Do not edit files marked
`Code generated` or `DO NOT EDIT`.

## Schema workflow

1. Update the relevant core or provider-extension schema.
2. Add valid and invalid canonical fixtures for the changed behavior.
3. Run `npm run generate`.
4. Update only handwritten ergonomic facades or decoders that need new behavior.
5. Run `npm test` and confirm `npm run check:generated` succeeds afterward.

The generator must fail on structural schema constructs it cannot represent.
Add explicit generator support and tests instead of falling back silently to a
weak type. Validation-only constraints such as formats, ranges, and patterns
remain enforced by JSON Schema.

## Compatibility

The `0.x` policy is additive by default. Preserve existing exported TypeScript
and Go names and PHP array helpers. Wire-level breaking changes require a new
protocol major version. Unknown fields and opaque provider references must
survive decoding and re-encoding.

Generated wire contracts are authoritative for binding structure. Handwritten
facades provide generics, constructors, factories, or discriminator helpers and
must be assignable to or round-trip through their generated wire counterparts.

## Local requirements

Install Node.js 20+, PHP 8.2+, Composer, and Go 1.22+, then run:

```sh
npm install
composer install --working-dir=packages/php
npm test
```

When Go is unavailable, the root Go check uses Docker automatically.
