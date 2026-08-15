export type * from "./generated/wire.js";
import type * as Generated from "./generated/wire.js";

export interface GitHubNumberedReference extends Generated.WireGitHubArtifactReferenceNumbered {
  repositoryId: number;
  number: number;
  nodeId?: string;
}
export interface GitHubPullRequestReference extends Generated.WireGitHubArtifactReferencePullRequest { resource: "pull_request"; }
export interface GitHubIssueReference extends Generated.WireGitHubArtifactReferenceIssue { resource: "issue"; }
export interface GitHubDiscussionReference extends Generated.WireGitHubArtifactReferenceDiscussion { resource: "discussion"; }
export interface GitHubCommitReference extends Generated.WireGitHubArtifactReferenceCommit { resource: "commit"; }
export interface GitHubDeploymentReference extends Generated.WireGitHubArtifactReferenceDeployment { resource: "deployment"; }
export interface GitHubWorkflowRunReference extends Generated.WireGitHubArtifactReferenceWorkflowRun { resource: "workflow_run"; }
export interface GitHubCheckRunReference extends Generated.WireGitHubArtifactReferenceCheckRun { resource: "check_run"; }
export type GitHubArtifactReference = GitHubPullRequestReference | GitHubCommitReference | GitHubIssueReference | GitHubDiscussionReference | GitHubDeploymentReference | GitHubWorkflowRunReference | GitHubCheckRunReference;
export interface GitHubArtifactSource extends Omit<Generated.WireGitHubArtifactSource, "reference"> { type: "provider"; provider: "github"; reference: GitHubArtifactReference; }
