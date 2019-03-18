package core

type (
	//Pipeline 集成构建管道定义
	Pipeline struct {
		ID        string      `json:"id"`
		UID       string      `json:"uid"`
		RepoID    string      `json:"repo_id"`
		Name      string      `json:"name"`
		Webhook   string      `json:"webhook"`
		Variables []Variables `json:"variables"`
	}

	//Configuration pipeline's general config
	Configuration struct {
	}

	//Variables pipeline env vars.
	Variables struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	//PipelineStore pipline store for db.
	PipelineStore interface {
	}
)
