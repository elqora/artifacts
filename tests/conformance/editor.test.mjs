import assert from "node:assert/strict";
import test from "node:test";
import {
  ARTIFACT_PROTOCOL_VERSION,
  ArtifactEditor,
  ArtifactSpecEditor,
} from "../../packages/typescript/dist/index.js";

const imageSpec = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "artspec_image_v1",
  key: "image_evidence",
  name: "Image evidence",
  version: 1,
  kind: "image_evidence",
  valueType: "image",
};

const referenceSpec = {
  schemaVersion: ARTIFACT_PROTOCOL_VERSION,
  id: "artspec_reference_v1",
  key: "reference_evidence",
  name: "Reference evidence",
  version: 1,
  kind: "reference_evidence",
  valueType: "reference",
};

function createArtifact() {
  return {
    schemaVersion: ARTIFACT_PROTOCOL_VERSION,
    id: "artifact_1",
    specId: imageSpec.id,
    kind: imageSpec.kind,
    valueType: imageSpec.valueType,
    title: "Evidence",
    specification: {
      schema: "image-evidence",
      version: 1,
      value: { subject: "Entrance" },
    },
    metadata: { source: "host" },
    createdBy: { type: "user", id: "user_1" },
    createdAt: "2026-08-15T09:00:00Z",
    updatedAt: "2026-08-15T09:00:00Z",
  };
}

test("ArtifactEditor validates its spec collection and synchronizes selected spec projections", () => {
  const artifact = createArtifact();
  const opened = ArtifactEditor.open(artifact, { specs: [imageSpec, referenceSpec] });
  assert.equal(opened.ok, true);
  if (!opened.ok) return;

  const editor = opened.value;
  const before = editor.value();
  const result = editor.transaction((tx) => {
    tx.setSpec(referenceSpec.id);
  });

  assert.equal(result.ok, true);
  assert.deepEqual(editor.value().specification, before.specification);
  assert.deepEqual(editor.value().metadata, before.metadata);
  assert.equal(editor.value().specId, referenceSpec.id);
  assert.equal(editor.value().kind, referenceSpec.kind);
  assert.equal(editor.value().valueType, referenceSpec.valueType);
  assert.equal(artifact.specId, imageSpec.id);
  assert.equal(artifact.kind, imageSpec.kind);
});

test("ArtifactEditor transactions commit atomically, rollback errors, and compose with history", () => {
  const artifact = createArtifact();
  const opened = ArtifactEditor.open(artifact, { specs: [imageSpec, referenceSpec] });
  assert.equal(opened.ok, true);
  if (!opened.ok) return;

  const editor = opened.value;
  const committed = editor.transaction((tx) => {
    tx.setTitle("Updated evidence");
    tx.updateMetadata({ reviewed: true });
  });
  assert.equal(committed.ok, true);
  assert.equal(editor.value().title, "Updated evidence");
  assert.equal(editor.value().metadata?.reviewed, true);
  assert.equal(editor.history.canUndo, true);

  const beforeFailure = editor.value();
  const failure = new Error("edit failed");
  const failed = editor.transaction((tx) => {
    tx.setTitle("Should be rolled back");
    throw failure;
  });
  assert.equal(failed.ok, false);
  if (failed.ok) return;
  assert.equal(failed.error, failure);
  assert.deepEqual(editor.value(), beforeFailure);

  const undone = editor.history.undo();
  assert.equal(undone?.title, "Evidence");
  assert.equal(editor.history.canRedo, true);

  const noOp = editor.transaction(() => {});
  assert.equal(noOp.ok, true);
  assert.equal(editor.history.canUndo, false);
  assert.equal(editor.history.canRedo, true);
  assert.equal(editor.history.redo()?.title, "Updated evidence");
});

test("ArtifactEditor rejects invalid initialization state", () => {
  const artifact = createArtifact();
  assert.equal(ArtifactEditor.open({ ...artifact, specId: "missing" }, { specs: [imageSpec] }).ok, false);
  assert.equal(ArtifactEditor.open(artifact, { specs: [imageSpec, imageSpec] }).ok, false);
  assert.equal(ArtifactEditor.open({ ...artifact, kind: "wrong_kind" }, { specs: [imageSpec] }).ok, false);
  assert.equal(ArtifactEditor.open({ ...artifact, specId: undefined }, { specs: [imageSpec] }).ok, false);
});

test("ArtifactSpecEditor edits definition content without changing identity or version", () => {
  const editor = ArtifactSpecEditor.open(imageSpec);
  const result = editor.transaction((tx) => {
    tx.setName("Updated image evidence");
    tx.setConfig({ valueType: "image", minFiles: 1 });
    tx.setValidation({ mode: "strict", schema: { subject: "required|string" } });
    tx.updateMetadata({ owner: "host" });
  });

  assert.equal(result.ok, true);
  assert.equal(editor.value().id, imageSpec.id);
  assert.equal(editor.value().version, imageSpec.version);
  assert.equal(editor.value().name, "Updated image evidence");
  assert.equal(editor.value().metadata?.owner, "host");

  const failure = new Error("spec edit failed");
  const beforeFailure = editor.value();
  const failed = editor.transaction(() => { throw failure; });
  assert.equal(failed.ok, false);
  if (failed.ok) return;
  assert.equal(failed.error, failure);
  assert.deepEqual(editor.value(), beforeFailure);
  assert.equal(editor.history.undo()?.name, "Image evidence");
  assert.equal(editor.history.redo()?.name, "Updated image evidence");
});
