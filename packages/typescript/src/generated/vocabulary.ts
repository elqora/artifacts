// Code generated from canonical JSON Schemas; DO NOT EDIT.
export const ARTIFACT_VALUE_TYPES = [
  "text",
  "number",
  "boolean",
  "currency",
  "date",
  "datetime",
  "time",
  "location",
  "file",
  "image",
  "video",
  "audio",
  "link",
  "structured",
  "reference",
  "signature",
  "collection"
] as const;
export const ARTIFACT_SOURCE_TYPES = [
  "inline",
  "local",
  "object",
  "url",
  "hosted",
  "provider"
] as const;
export const LOCAL_ARTIFACT_SYNC_STATES = [
  "local_only",
  "pending_upload",
  "uploading",
  "uploaded",
  "failed"
] as const;
export const ARTIFACT_INTEGRITY_ALGORITHMS = [
  "sha256",
  "sha384",
  "sha512"
] as const;
export const ARTIFACT_CONDITION_KINDS = [
  "state",
  "actor",
  "artifact_exists",
  "artifact_value",
  "and",
  "or",
  "not"
] as const;
export const ARTIFACT_VALUE_OPERATORS = [
  "eq",
  "neq",
  "gt",
  "gte",
  "lt",
  "lte",
  "contains",
  "in"
] as const;
export const ARTIFACT_PRIVACY_CLASSIFICATIONS = [
  "public",
  "internal",
  "private",
  "sensitive",
  "restricted"
] as const;
export const ARTIFACT_PRIVACY_REPRESENTATIONS = [
  "hidden",
  "masked",
  "approximate",
  "full"
] as const;
export const ARTIFACT_VERIFICATION_STATUSES = [
  "pending",
  "verified",
  "rejected",
  "waived"
] as const;
