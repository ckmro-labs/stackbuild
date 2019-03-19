package core

type (
	//Build 构建信息, 它包含Step信息。
	Build struct {
		ID         string  `json:"id"`
		RepoID     string  `json:"repo_id"`
		PipelineID string  `json:"pipeline_id"`
		Status     string  `json:"status"`
		Error      string  `json:"error,omitempty"`
		Started    int64   `json:"started"`
		Finished   int64   `json:"finished"`
		Created    int64   `json:"created"`
		Updated    int64   `json:"updated"`
		Steps      []*Step `json:"steps"`
	}

	//BuildStore build for storage.
	BuildStore interface {
	}
)
