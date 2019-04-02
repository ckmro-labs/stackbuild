package core

import "context"

//Workflow ..
type Workflow string

const (
	//BuildWorkflow build workflow type.
	BuildWorkflow Workflow = "build"
	//UnitTestWorkflow test workflow type.
	UnitTestWorkflow Workflow = "test"
	//RegistryWorkflow push to registry
	RegistryWorkflow Workflow = "registry"
)

type (
	//Stage 场景集成构建管道定义
	Stage struct {
		ID        string        `json:"id"`
		UID       string        `json:"uid"`
		RepoID    string        `json:"repo_id"`
		Name      string        `json:"name"`
		Webhook   string        `json:"webhook"`
		Limit     int           `json:"limit,omitempty"`
		Status    string        `json:"status"`
		Variables []Variables   `json:"variables"`
		Steps     []interface{} `json:"steps"`
	}

	//Triggers trigger for vcs.
	Triggers struct {
		Type   string   `json:"type"`
		Events []string `json:"events"`
		Branch string   `json:"branch"`
	}

	//BuildingDockerImage pipeline's workflow for build.
	BuildingDockerImage struct {
		Title      string   `json:"title"`
		Type       Workflow `json:"type" default:"build"`
		ImageName  string   `json:"image_name"`
		Working    string   `json:"working_directory"`
		Dockerfile string   `json:"dockerfile"`
	}

	//RegistryImage pipeline's workflow for push registry.
	RegistryImage struct {
		Type Workflow `json:"type" default:"registry"`
	}

	//UnitTest pipeline's workflow for unit tests.
	UnitTest struct {
		Type   Workflow `json:"type" default:"test"`
		Script string   `json:"script"`
	}

	//Variables pipeline env vars.
	Variables struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	//StageStore pipline store for db.
	StageStore interface {
		// Create persists a new stage to the datastore.
		Create(context.Context, *Stage) error
		// Find returns a build stage from the datastore by ID.
		Find(context.Context, int64) (*Stage, error)
		// List returns a build stage list from the datastore, where the stage is incomplete (pending or running).
		ListIncomplete(context.Context) ([]*Stage, error)
	}
)
