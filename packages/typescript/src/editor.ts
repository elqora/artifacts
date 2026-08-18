import type {
  ActorReference,
  Artifact,
  ArtifactAccessPolicy,
  ArtifactCondition,
  ArtifactEncryptionPolicy,
  ArtifactKind,
  ArtifactLifecyclePolicy,
  ArtifactMetadata,
  ArtifactPresentationPolicy,
  ArtifactPrivacyPolicy,
  ArtifactProviderPolicy,
  ArtifactRequirementPolicy,
  ArtifactRole,
  ArtifactScopeReference,
  ArtifactSpec,
  ArtifactSpecId,
  ArtifactSpecification,
  ArtifactSubjectReference,
  ArtifactValidationPolicy,
  ArtifactValueType,
  ArtifactVerificationPolicy,
  ArtifactRetentionPolicy,
} from "./index.js";
import { EditorKernel, type EditorHistory, type EditorResult } from "./editor-kernel.js";

export type { EditorHistory, EditorResult } from "./editor-kernel.js";

export interface ArtifactEditorOptions {
  specs: readonly ArtifactSpec[];
}

export interface ArtifactEditTransaction<TSpecification = unknown> {
  setScope(scope?: ArtifactScopeReference): void;
  setTitle(title?: string): void;
  setDescription(description?: string): void;
  setSpec(specId: ArtifactSpecId): void;
  setSpecification<TValue>(specification?: ArtifactSpecification<TValue>): void;
  setMetadata(metadata?: ArtifactMetadata): void;
  updateMetadata(partial: ArtifactMetadata): void;
  removeMetadata(key: string): void;
  setArchivedAt(timestamp?: string): void;
}

export interface ArtifactSpecEditTransaction<TConfig extends object = Record<string, unknown>> {
  setKey(key: string): void;
  setName(name: string): void;
  setDescription(description?: string): void;
  setKind(kind: ArtifactKind): void;
  setValueType(valueType: ArtifactValueType): void;
  setConfig(config?: TConfig): void;
  setProvider(provider?: ArtifactProviderPolicy): void;
  setRequirement(requirement?: ArtifactRequirementPolicy): void;
  setLifecycle(lifecycle?: ArtifactLifecyclePolicy): void;
  setAccess(access?: ArtifactAccessPolicy): void;
  setPrivacy(privacy?: ArtifactPrivacyPolicy): void;
  setValidation(validation?: ArtifactValidationPolicy): void;
  setVerification(verification?: ArtifactVerificationPolicy): void;
  setRetention(retention?: ArtifactRetentionPolicy): void;
  setPresentation(presentation?: ArtifactPresentationPolicy): void;
  setMetadata(metadata?: ArtifactMetadata): void;
  updateMetadata(partial: ArtifactMetadata): void;
  removeMetadata(key: string): void;
}

function setOptional<T extends object, K extends keyof T>(draft: T, key: K, value: T[K] | undefined): void {
  if (value === undefined) delete draft[key];
  else draft[key] = value;
}

function updateMetadata<T extends { metadata?: ArtifactMetadata }>(draft: T, partial: ArtifactMetadata): void {
  draft.metadata = { ...(draft.metadata ?? {}), ...partial };
}

function removeMetadata<T extends { metadata?: ArtifactMetadata }>(draft: T, key: string): void {
  if (draft.metadata === undefined || !Object.hasOwn(draft.metadata, key)) return;
  const metadata = { ...draft.metadata };
  delete metadata[key];
  setOptional(draft, "metadata", Object.keys(metadata).length > 0 ? metadata : undefined);
}

function error(message: string): Error {
  return new Error(message);
}

export class ArtifactEditor<
  TKind extends string = string,
  TValueType extends ArtifactValueType = ArtifactValueType,
  TSpecification = unknown,
> {
  private constructor(
    private readonly kernel: EditorKernel<Artifact<TKind, TValueType, TSpecification>>,
    private readonly specs: ReadonlyMap<ArtifactSpecId, ArtifactSpec>,
  ) {}

  public static open<
    TKind extends string = string,
    TValueType extends ArtifactValueType = ArtifactValueType,
    TSpecification = unknown,
  >(
    artifact: Artifact<TKind, TValueType, TSpecification>,
    options: ArtifactEditorOptions,
  ): EditorResult<ArtifactEditor<TKind, TValueType, TSpecification>> {
    const specs = new Map<ArtifactSpecId, ArtifactSpec>();
    for (const spec of options.specs) {
      if (specs.has(spec.id)) {
        return { ok: false, error: error(`Duplicate artifact specification ID: ${spec.id}`) };
      }
      specs.set(spec.id, structuredClone(spec));
    }

    const selected = specs.get(artifact.specId);
    if (selected === undefined) {
      return { ok: false, error: error(`Artifact specification not found: ${artifact.specId}`) };
    }
    if (selected.kind !== artifact.kind || selected.valueType !== artifact.valueType) {
      return {
        ok: false,
        error: error(`Artifact ${artifact.id} does not match specification ${selected.id}`),
      };
    }

    return { ok: true, value: new ArtifactEditor(new EditorKernel(artifact), specs) };
  }

  public value(): Artifact<TKind, TValueType, TSpecification> {
    return this.kernel.value();
  }

  public get history(): EditorHistory<Artifact<TKind, TValueType, TSpecification>> {
    return this.kernel.history;
  }

  public transaction(
    callback: (transaction: ArtifactEditTransaction<TSpecification>) => void,
  ): EditorResult<Artifact<TKind, TValueType, TSpecification>> {
    return this.kernel.transaction((draft) => {
      const transaction: ArtifactEditTransaction<TSpecification> = {
        setScope: (scope) => setOptional(draft, "scope", scope),
        setTitle: (title) => setOptional(draft, "title", title),
        setDescription: (description) => setOptional(draft, "description", description),
        setSpec: (specId) => {
          const spec = this.specs.get(specId);
          if (spec === undefined) throw error(`Artifact specification not found: ${specId}`);
          draft.specId = spec.id;
          draft.kind = spec.kind as TKind;
          draft.valueType = spec.valueType as TValueType;
        },
        setSpecification: (specification) => setOptional(draft, "specification", specification as ArtifactSpecification<TSpecification> | undefined),
        setMetadata: (metadata) => setOptional(draft, "metadata", metadata),
        updateMetadata: (partial) => updateMetadata(draft, partial),
        removeMetadata: (key) => removeMetadata(draft, key),
        setArchivedAt: (timestamp) => setOptional(draft, "archivedAt", timestamp),
      };
      callback(transaction);
    });
  }
}

export class ArtifactSpecEditor<
  TKind extends string = string,
  TValueType extends ArtifactValueType = ArtifactValueType,
  TConfig extends object = Record<string, unknown>,
> {
  private constructor(private readonly kernel: EditorKernel<ArtifactSpec<TKind, TValueType, TConfig>>) {}

  public static open<
    TKind extends string = string,
    TValueType extends ArtifactValueType = ArtifactValueType,
    TConfig extends object = Record<string, unknown>,
  >(spec: ArtifactSpec<TKind, TValueType, TConfig>): ArtifactSpecEditor<TKind, TValueType, TConfig> {
    return new ArtifactSpecEditor(new EditorKernel(spec));
  }

  public value(): ArtifactSpec<TKind, TValueType, TConfig> {
    return this.kernel.value();
  }

  public get history(): EditorHistory<ArtifactSpec<TKind, TValueType, TConfig>> {
    return this.kernel.history;
  }

  public transaction(
    callback: (transaction: ArtifactSpecEditTransaction<TConfig>) => void,
  ): EditorResult<ArtifactSpec<TKind, TValueType, TConfig>> {
    return this.kernel.transaction((draft) => {
      const transaction: ArtifactSpecEditTransaction<TConfig> = {
        setKey: (key) => { draft.key = key; },
        setName: (name) => { draft.name = name; },
        setDescription: (description) => setOptional(draft, "description", description),
        setKind: (kind) => { draft.kind = kind as TKind; },
        setValueType: (valueType) => { draft.valueType = valueType as TValueType; },
        setConfig: (config) => setOptional(draft, "config", config),
        setProvider: (provider) => setOptional(draft, "provider", provider),
        setRequirement: (requirement) => setOptional(draft, "requirement", requirement),
        setLifecycle: (lifecycle) => setOptional(draft, "lifecycle", lifecycle),
        setAccess: (access) => setOptional(draft, "access", access),
        setPrivacy: (privacy) => setOptional(draft, "privacy", privacy),
        setValidation: (validation) => setOptional(draft, "validation", validation),
        setVerification: (verification) => setOptional(draft, "verification", verification),
        setRetention: (retention) => setOptional(draft, "retention", retention),
        setPresentation: (presentation) => setOptional(draft, "presentation", presentation),
        setMetadata: (metadata) => setOptional(draft, "metadata", metadata),
        updateMetadata: (partial) => updateMetadata(draft, partial),
        removeMetadata: (key) => removeMetadata(draft, key),
      };
      callback(transaction);
    });
  }
}
