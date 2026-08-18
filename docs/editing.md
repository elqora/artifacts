# Programmatic editing

The TypeScript package provides small in-memory editors for generic artifact
and artifact-spec content. Editors do not persist changes, execute policies,
advance versions, publish definitions, or implement host workflows.

## Artifacts

An artifact editor is opened with the specifications available to the host:

```typescript
const opened = ArtifactEditor.open(artifact, { specs });

if (!opened.ok) {
  throw opened.error;
}

const editor = opened.value;
```

Opening validates that the artifact has a `specId`, that the supplied
specifications have unique IDs, that the referenced specification exists, and
that its `kind` and `valueType` agree with the artifact's cached projections.

`setSpec()` accepts a specification ID and synchronizes only `specId`, `kind`,
and `valueType`. It does not replace the artifact's specification value or
metadata.

## Transactions and history

Edits are grouped with a callback transaction:

```typescript
const result = editor.transaction((tx) => {
  tx.setTitle("Updated title");
  tx.setSpec("delivery_evidence_v3");
  tx.updateMetadata({ reviewed: true });
});
```

The result is either `{ ok: true, value: updatedArtifact }` or
`{ ok: false, error }`. Callback and edit errors automatically restore the
exact pre-transaction state and create no history entry. A successful no-op
also creates no entry.

`editor.history.undo()` and `.redo()` operate on editor snapshots. They are
separate from `ArtifactVersion` history. The original object passed to `open`
is never mutated.

Metadata supports complete replacement with `setMetadata()`, shallow merging
with `updateMetadata()`, and single-key removal with `removeMetadata()`.

## Artifact specifications

`ArtifactSpecEditor.open(spec)` provides the same transaction and history
behavior for specification definition content. It can edit fields such as
`name`, `description`, `kind`, `valueType`, `config`, provider and policy
definitions, presentation, and metadata. It does not change `id`, `version`,
or `schemaVersion`; hosts are responsible for version advancement and
publishing.
