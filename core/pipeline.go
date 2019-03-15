package core

type (
	//Pipeline 集成构建管道定义
	Pipeline struct {
		ID   string `json:"id"`
		UID  string `json:"uid"`
		Name string `json:"name"`
	}

	//PipelineStore pipline store for db.
	PipelineStore interface {
	}
)
