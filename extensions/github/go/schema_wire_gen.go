// Code generated from canonical JSON Schemas; DO NOT EDIT.
package githubartifact

import "encoding/json"

type WireGitHubArtifactReference = json.RawMessage

type WireGitHubArtifactReferenceNumbered struct {
	RepositoryID int64 `json:"repositoryId"`
	Number int64 `json:"number"`
	NodeID *string `json:"nodeId,omitempty"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireGitHubArtifactReferencePullRequest struct {
	RepositoryID int64 `json:"repositoryId"`
	Number int64 `json:"number"`
	NodeID *string `json:"nodeId,omitempty"`
	Resource string `json:"resource"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireGitHubArtifactReferenceIssue struct {
	RepositoryID int64 `json:"repositoryId"`
	Number int64 `json:"number"`
	NodeID *string `json:"nodeId,omitempty"`
	Resource string `json:"resource"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireGitHubArtifactReferenceDiscussion struct {
	RepositoryID int64 `json:"repositoryId"`
	Number int64 `json:"number"`
	NodeID *string `json:"nodeId,omitempty"`
	Resource string `json:"resource"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireGitHubArtifactReferenceCommit struct {
	Resource string `json:"resource"`
	RepositoryID int64 `json:"repositoryId"`
	SHA string `json:"sha"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireGitHubArtifactReferenceDeployment struct {
	Resource string `json:"resource"`
	RepositoryID int64 `json:"repositoryId"`
	DeploymentID int64 `json:"deploymentId"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireGitHubArtifactReferenceWorkflowRun struct {
	Resource string `json:"resource"`
	RepositoryID int64 `json:"repositoryId"`
	RunID int64 `json:"runId"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireGitHubArtifactReferenceCheckRun struct {
	Resource string `json:"resource"`
	RepositoryID int64 `json:"repositoryId"`
	CheckRunID int64 `json:"checkRunId"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

type WireGitHubArtifactSource struct {
	Type string `json:"type"`
	Provider string `json:"provider"`
	Reference WireGitHubArtifactReference `json:"reference"`
	UnknownFields map[string]json.RawMessage `json:"-"`
}

func (value *WireGitHubArtifactReferenceNumbered) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactReferenceNumbered
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactReferenceNumbered(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"repositoryId": {}, "number": {}, "nodeId": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactReferenceNumbered) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactReferenceNumbered
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"repositoryId": {}, "number": {}, "nodeId": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireGitHubArtifactReferencePullRequest) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactReferencePullRequest
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactReferencePullRequest(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"repositoryId": {}, "number": {}, "nodeId": {}, "resource": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactReferencePullRequest) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactReferencePullRequest
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"repositoryId": {}, "number": {}, "nodeId": {}, "resource": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireGitHubArtifactReferenceIssue) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactReferenceIssue
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactReferenceIssue(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"repositoryId": {}, "number": {}, "nodeId": {}, "resource": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactReferenceIssue) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactReferenceIssue
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"repositoryId": {}, "number": {}, "nodeId": {}, "resource": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireGitHubArtifactReferenceDiscussion) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactReferenceDiscussion
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactReferenceDiscussion(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"repositoryId": {}, "number": {}, "nodeId": {}, "resource": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactReferenceDiscussion) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactReferenceDiscussion
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"repositoryId": {}, "number": {}, "nodeId": {}, "resource": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireGitHubArtifactReferenceCommit) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactReferenceCommit
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactReferenceCommit(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"resource": {}, "repositoryId": {}, "sha": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactReferenceCommit) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactReferenceCommit
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"resource": {}, "repositoryId": {}, "sha": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireGitHubArtifactReferenceDeployment) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactReferenceDeployment
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactReferenceDeployment(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"resource": {}, "repositoryId": {}, "deploymentId": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactReferenceDeployment) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactReferenceDeployment
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"resource": {}, "repositoryId": {}, "deploymentId": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireGitHubArtifactReferenceWorkflowRun) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactReferenceWorkflowRun
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactReferenceWorkflowRun(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"resource": {}, "repositoryId": {}, "runId": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactReferenceWorkflowRun) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactReferenceWorkflowRun
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"resource": {}, "repositoryId": {}, "runId": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireGitHubArtifactReferenceCheckRun) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactReferenceCheckRun
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactReferenceCheckRun(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"resource": {}, "repositoryId": {}, "checkRunId": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactReferenceCheckRun) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactReferenceCheckRun
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"resource": {}, "repositoryId": {}, "checkRunId": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}

func (value *WireGitHubArtifactSource) UnmarshalJSON(data []byte) error {
	type plain WireGitHubArtifactSource
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil { return err }
	*value = WireGitHubArtifactSource(decoded)
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil { return err }
	known := map[string]struct{}{"type": {}, "provider": {}, "reference": {}}
	for key := range known { delete(fields, key) }
	if len(fields) > 0 { value.UnknownFields = fields }
	return nil
}

func (value WireGitHubArtifactSource) MarshalJSON() ([]byte, error) {
	type plain WireGitHubArtifactSource
	knownBytes, err := json.Marshal(plain(value)); if err != nil { return nil, err }
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(knownBytes, &fields); err != nil { return nil, err }
	reserved := map[string]struct{}{"type": {}, "provider": {}, "reference": {}}
	for key, raw := range value.UnknownFields { if _, knownKey := reserved[key]; !knownKey { fields[key] = raw } }
	return json.Marshal(fields)
}
