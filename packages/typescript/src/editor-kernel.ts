export type EditorResult<T> =
  | { ok: true; value: T }
  | { ok: false; error: unknown };

export interface EditorHistory<T> {
  undo(): T | undefined;
  redo(): T | undefined;
  readonly canUndo: boolean;
  readonly canRedo: boolean;
  clear(): void;
}

type HistoryEntry<T> = {
  before: T;
  after: T;
};

function clone<T>(value: T): T {
  return structuredClone(value);
}

function equal<T>(left: T, right: T): boolean {
  return JSON.stringify(left) === JSON.stringify(right);
}

class EditorHistoryImpl<T> implements EditorHistory<T> {
  private readonly undoStack: HistoryEntry<T>[] = [];
  private readonly redoStack: HistoryEntry<T>[] = [];

  public constructor(
    private readonly read: () => T,
    private readonly write: (value: T) => void,
  ) {}

  public get canUndo(): boolean {
    return this.undoStack.length > 0;
  }

  public get canRedo(): boolean {
    return this.redoStack.length > 0;
  }

  public undo(): T | undefined {
    const entry = this.undoStack.pop();
    if (entry === undefined) return undefined;

    this.redoStack.push(entry);
    this.write(clone(entry.before));
    return clone(this.read());
  }

  public redo(): T | undefined {
    const entry = this.redoStack.pop();
    if (entry === undefined) return undefined;

    this.undoStack.push(entry);
    this.write(clone(entry.after));
    return clone(this.read());
  }

  public push(before: T, after: T): void {
    this.undoStack.push({ before: clone(before), after: clone(after) });
    this.redoStack.length = 0;
  }

  public clear(): void {
    this.undoStack.length = 0;
    this.redoStack.length = 0;
  }
}

export class EditorKernel<T> {
  private state: T;
  public readonly history: EditorHistoryImpl<T>;

  public constructor(initial: T) {
    this.state = clone(initial);
    this.history = new EditorHistoryImpl(
      () => this.state,
      (value) => {
        this.state = clone(value);
      },
    );
  }

  public value(): T {
    return clone(this.state);
  }

  public transaction(callback: (draft: T) => void): EditorResult<T> {
    const before = clone(this.state);
    const draft = clone(this.state);

    try {
      callback(draft);

      if (!equal(before, draft)) {
        this.state = draft;
        this.history.push(before, draft);
      }

      return { ok: true, value: clone(this.state) };
    } catch (error) {
      this.state = before;
      return { ok: false, error };
    }
  }
}
