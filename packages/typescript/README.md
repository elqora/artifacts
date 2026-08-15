# `@elqora/artifacts`

TypeScript SDK for the language-neutral Elqora Artifact Protocol. It provides
typed contracts for artifact identity, immutable versions, physical and
provider sources, contextual links, specifications, submissions, verification,
integrity, and declarative policies.

The SDK defines protocol data. Applications remain responsible for storage,
workflows, authorization, provider clients, and user interfaces.

## Install

```sh
npm install @elqora/artifacts
```

The package is ESM-only and supports Node.js 20 or newer.

## Use

```ts
import {
  ARTIFACT_PROTOCOL_VERSION,
  type Artifact,
  type ArtifactSpec,
  type ArtifactVersion,
} from "@elqora/artifacts";

const artifact: Artifact<"delivery_evidence", "image"> = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "artifact_1",
  kind: "delivery_evidence",
  valueType: "image",
  createdBy: { type: "user", id: "user_1" },
  createdAt: "2026-08-15T09:00:00Z",
  updatedAt: "2026-08-15T09:00:00Z",
};
```

The root export includes the ergonomic TypeScript interfaces, generated
`Wire*` contracts, and closed-vocabulary constants. Generated contracts follow
the canonical JSON Schemas in the
[source repository](https://github.com/elqora/artifacts).

## Compatibility

During `0.x`, changes are additive by default. Existing public exports and wire
serialization are preserved unless a documented correction is unavoidable.
The serialized protocol `schemaVersion` is independent of the npm package
version.

Licensed under AGPL-3.0-only.
