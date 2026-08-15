// Package githubartifact provides typed GitHub provider references only; it does not call GitHub APIs.
package githubartifact

type NumberedReference struct {
	RepositoryID int64  `json:"repositoryId"`
	Number       int    `json:"number"`
	NodeID       string `json:"nodeId,omitempty"`
}
type PullRequestReference struct {
	Resource string `json:"resource"`
	NumberedReference
}
type IssueReference struct {
	Resource string `json:"resource"`
	NumberedReference
}
type DiscussionReference struct {
	Resource string `json:"resource"`
	NumberedReference
}
type CommitReference struct {
	Resource     string `json:"resource"`
	RepositoryID int64  `json:"repositoryId"`
	SHA          string `json:"sha"`
}
type DeploymentReference struct {
	Resource     string `json:"resource"`
	RepositoryID int64  `json:"repositoryId"`
	DeploymentID int64  `json:"deploymentId"`
}
type WorkflowRunReference struct {
	Resource     string `json:"resource"`
	RepositoryID int64  `json:"repositoryId"`
	RunID        int64  `json:"runId"`
}
type CheckRunReference struct {
	Resource     string `json:"resource"`
	RepositoryID int64  `json:"repositoryId"`
	CheckRunID   int64  `json:"checkRunId"`
}
