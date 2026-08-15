# Conformance

Canonical valid fixtures live under `tests/fixtures` by protocol concept;
negative fixtures live under `tests/fixtures/invalid`. AJV tests compile every
core and extension schema in strict Draft 2020-12 mode and cover required
fields, all source/value discriminators, recursive conditions, policies, open
and closed vocabularies, null semantics, snapshots, submissions, verification,
GitHub references, and exact-version links.

Some invariants span records and therefore need conformance or host checks:
pinned versions must belong to the named artifact; a verification's submission
and version must agree; current version pointers must identify a real revision;
and definition snapshots must remain immutable after capture.

TypeScript compilation proves its type contracts. PHP and Go tests deserialize
and reserialize the same fixture corpus as generic JSON maps so unknown fields
and opaque provider references are preserved without semantic change.
