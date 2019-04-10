package core

import (
	"context"
)

//Workflow ..
type Workflow string

const (
	//BuildWorkflow build workflow type.
	BuildWorkflow Workflow = "build"
	//UnitTestWorkflow test workflow type.
	UnitTestWorkflow Workflow = "test"
	//PackageWorkflow package type.
	PackageWorkflow Workflow = "package"
	//RegistryWorkflow push to registry
	RegistryWorkflow Workflow = "registry"
)

type (
	//Stage 场景集成构建管道定义
	Stage struct {
		ID        string        `bson:"_id" json:"id"`
		Name      string        `bson:"name" json:"name"`
		UID       string        `bson:"uid" json:"uid"` //user id
		RepoID    string        `bson:"repoId" json:"repo_id"`
		Branch    string        `bson:"branch" json:"branch"`
		Ref       string        `bson:"ref" json:"ref"`
		Webhook   string        `bson:"webhook" json:"webhook"`
		Limit     int           `bson:"limit" json:"limit,omitempty"`
		Created   int64         `bson:"created" json:"created_at"`
		Updated   int64         `bson:"updated" json:"updated_at"`
		Status    string        `bson:"status" json:"status"`        ////最后一次构建状态
		LastBuild int64         `bson:"lastBuild" json:"last_build"` //最后一次构建时间
		Variables []Variables   `bson:"variables" json:"variables"`
		Steps     []PiplineStep `bson:"steps" json:"steps"`
	}

	//Triggers trigger for vcs.
	Triggers struct {
		Type   string   `json:"type"`
		Events []string `json:"events"`
		Branch string   `json:"branch"`
	}

	//PiplineStep pipeline's workflow for build.
	PiplineStep struct {
		Title     string   `bson:"title" json:"title"`
		Type      Workflow `bson:"type" json:"type" default:"build"`
		Command   string   `bson:"command" json:"command"`
		Host      string   `bson:"host" json:"host"`
		UserName  string   `bson:"userName" json:"user_name"`
		Password  string   `bson:"password" json:"password"`
		Cert      string   `bson:"cert" json:"cert"`
		ImageName string   `bson:"imageName"  json:"image_name"`
		Working   string   `bson:"workDir"  json:"working_directory"`
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
		Find(context.Context, string) (*Stage, error)
		// List returns a build stage list from the datastore, where the stage is incomplete (pending or running).
		ListIncomplete(context.Context) ([]*Stage, error)
	}
)
