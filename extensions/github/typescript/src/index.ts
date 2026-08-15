export interface GitHubNumberedReference {
  repositoryId: number;
  number: number;
  nodeId?: string;
}
export interface GitHubPullRequestReference extends GitHubNumberedReference { resource: "pull_request"; }
export interface GitHubIssueReference extends GitHubNumberedReference { resource: "issue"; }
export interface GitHubDiscussionReference extends GitHubNumberedReference { resource: "discussion"; }
export interface GitHubCommitReference { resource: "commit"; repositoryId: number; sha: string; }
export interface GitHubDeploymentReference { resource: "deployment"; repositoryId: number; deploymentId: number; }
export interface GitHubWorkflowRunReference { resource: "workflow_run"; repositoryId: number; runId: number; }
export interface GitHubCheckRunReference { resource: "check_run"; repositoryId: number; checkRunId: number; }
export type GitHubArtifactReference = GitHubPullRequestReference | GitHubCommitReference | GitHubIssueReference | GitHubDiscussionReference | GitHubDeploymentReference | GitHubWorkflowRunReference | GitHubCheckRunReference;
export interface GitHubArtifactSource { type: "provider"; provider: "github"; reference: GitHubArtifactReference; }
