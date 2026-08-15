# Declarative conditions

Conditions are recursive JSON data, not callbacks. Closed node kinds are
`state`, `actor`, `artifact_exists`, `artifact_value`, `and`, `or`, and `not`.
Value comparisons use `eq`, `neq`, `gt`, `gte`, `lt`, `lte`, `contains`, or
`in`.

Namespaces, states, actors, and artifact lookup keys remain host-defined. A
host evaluator must document lookup rules, comparison/coercion semantics, and
failure behavior. The core protocol validates the AST but intentionally does
not execute it.
