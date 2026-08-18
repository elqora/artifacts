# Declarative policies

Specification policy contracts describe provider expectations, requiredness
and blocking operations, lifecycle stages, access rules, privacy/reveal/masking/
encryption intent, specification validation, verification expectations, retention,
and presentation hints. Host actor names, stages, operations, validation rule
types, masking strategies, encryption levels, verification methods, and display
hints remain extensible strings.

`validation.schema` is a serializable Laravel-style rule map applied by the host
to `Artifact.specification.value` or `ArtifactVersion.specification.value`:

```ts
const policy = {
  schema: {
    locations: "required|object",
    "locations.name": "required|string",
  },
};
```

Rule values may also use Bunwire's canonical array form, such as
`["required", "string", "min:3"]`. Paths use exact dot notation; wildcard and
executable callback rules are not part of the serialized protocol. Hosts choose
which specification context to validate and are responsible for invoking the
validation library.

Access answers whether an actor may access an artifact. Privacy answers which
representation that actor may receive. Presentation hints are optional and have
no dependency on a UI framework.

> Policy records are declarations. Security exists only when a host correctly
> enforces them.

`classification: restricted` does not restrict data, `encryption.required`
does not encrypt it, and a retention declaration does not delete anything.
The SDK provides no authorization, cryptography, policy evaluator, storage
cleanup, or workflow engine.
