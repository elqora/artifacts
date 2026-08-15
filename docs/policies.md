# Declarative policies

Specification policy contracts describe provider expectations, requiredness
and blocking operations, lifecycle stages, access rules, privacy/reveal/masking/
encryption intent, validation rules, verification expectations, retention, and
presentation hints. Host actor names, stages, operations, validation rule types,
masking strategies, encryption levels, verification methods, and display hints
remain extensible strings.

Access answers whether an actor may access an artifact. Privacy answers which
representation that actor may receive. Presentation hints are optional and have
no dependency on a UI framework.

> Policy records are declarations. Security exists only when a host correctly
> enforces them.

`classification: restricted` does not restrict data, `encryption.required`
does not encrypt it, and a retention declaration does not delete anything.
The SDK provides no authorization, cryptography, policy evaluator, storage
cleanup, or workflow engine.
