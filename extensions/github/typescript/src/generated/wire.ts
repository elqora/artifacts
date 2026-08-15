// Code generated from canonical JSON Schemas; DO NOT EDIT.

export type WireGitHubArtifactReference = WireGitHubArtifactReferencePullRequest | WireGitHubArtifactReferenceCommit | WireGitHubArtifactReferenceIssue | WireGitHubArtifactReferenceDiscussion | WireGitHubArtifactReferenceDeployment | WireGitHubArtifactReferenceWorkflowRun | WireGitHubArtifactReferenceCheckRun;

export type WireGitHubArtifactReferenceNumbered = {
  "repositoryId": number;
  "number": number;
  "nodeId"?: string;
  [key: string]: unknown;
};

export type WireGitHubArtifactReferencePullRequest = WireGitHubArtifactReferenceNumbered & {
  "resource": "pull_request";
  [key: string]: unknown;
};

export type WireGitHubArtifactReferenceIssue = WireGitHubArtifactReferenceNumbered & {
  "resource": "issue";
  [key: string]: unknown;
};

export type WireGitHubArtifactReferenceDiscussion = WireGitHubArtifactReferenceNumbered & {
  "resource": "discussion";
  [key: string]: unknown;
};

export type WireGitHubArtifactReferenceCommit = {
  "resource": "commit";
  "repositoryId": number;
  "sha": string;
  [key: string]: unknown;
};

export type WireGitHubArtifactReferenceDeployment = {
  "resource": "deployment";
  "repositoryId": number;
  "deploymentId": number;
  [key: string]: unknown;
};

export type WireGitHubArtifactReferenceWorkflowRun = {
  "resource": "workflow_run";
  "repositoryId": number;
  "runId": number;
  [key: string]: unknown;
};

export type WireGitHubArtifactReferenceCheckRun = {
  "resource": "check_run";
  "repositoryId": number;
  "checkRunId": number;
  [key: string]: unknown;
};

export type WireGitHubArtifactSource = {
  "type": "provider";
  "provider": "github";
  "reference": WireGitHubArtifactReference;
  [key: string]: unknown;
};
