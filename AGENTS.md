AGENTS.md

Project

This repository implements a language-neutral Artifact Specification and SDK family.

The goal is to define a reusable protocol for artifacts, versions, sources, relationships, specifications, submissions, verification, integrity, policies, and related concepts that can be consumed by completely different applications and implemented across multiple programming languages.

The SDK defines the shared artifact language and contracts.

Host applications remain responsible for persistence, storage, workflows, UI, authorization, provider integrations, and business behavior.

Authoritative Design

Read and follow:

"./GOAL.md"

"GOAL.md" is the authoritative architecture and implementation brief for this repository.

Before making architectural decisions, adding domain concepts, changing schemas, or introducing new abstractions, verify that the change is consistent with "GOAL.md".

Do not duplicate the contents of "GOAL.md" into other design documents unnecessarily.

Core Rules

- Keep the specification language-neutral.
- TypeScript may be used as reference notation, but TypeScript is not the protocol authority.
- Prefer machine-readable canonical schemas for protocol definitions.
- Keep "Artifact", "ArtifactVersion", "ArtifactSource", "ArtifactLink", "ArtifactSpec", submissions, and verification as separate concepts.
- Do not mix artifact identity with physical content.
- Do not encode deliverable/evidence/review usage directly into "Artifact"; use relationships.
- Support exact artifact-version pinning where historical correctness matters.
- Keep application vocabulary extensible. Do not hardcode concepts such as errands, milestones, reviews, runners, projects, or workflow-specific state names into the core protocol.
- Do not place executable callbacks or language-specific behavior inside serialized definitions.
- Do not implement application-specific persistence, UI, workflows, authorization systems, file storage, or provider API clients in the core SDK.
- Keep provider-specific models outside the core where practical.
- Preserve forward compatibility and schema-versioning concerns from the beginning.
- Prefer small, composable contracts over giant interfaces.
- Add tests and canonical fixtures for protocol behavior and cross-language compatibility.

Implementation Approach

Work incrementally from the stable core outward.

Start with the core protocol models and schemas described in "GOAL.md", then add specifications, runtime records, policies, generated language bindings, and provider extensions in the order described there.

Avoid speculative abstractions that are not required by the design or demonstrated by real use cases.

When uncertain about scope or terminology, treat "GOAL.md" as the source of truth.