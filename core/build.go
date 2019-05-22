package core

import "context"

type (
	//BuildStep build step...
	BuildStep struct {
		Number   int    `bson:"num" json:"num"`
		Name     string `bson:"name" json:"name"`
		Status   string `bson:"status" json:"status"`
		Error    string `bson:"err" json:"error,omitempty"`
		ExitCode int    `bson:"exitCode" json:"exit_code"`
		Started  int64  `bson:"nstarted" json:"started,omitempty"`
		Stopped  int64  `bson:"stopped" json:"stopped,omitempty"`
	}

	//Build 一次管道构建信息, 来自Stage触发
	Build struct {
		ID         int64             `bson:"_id" json:"id"`
		RepoID     int64             `bson:"repoId" json:"repo_id"`
		Status     string            `bson:"status" json:"status"`
		Error      string            `bson:"err" json:"error,omitempty"`
		Link       string            `bson:"link" json:"link"`
		Timestamp  int64             `json:"timestamp"`
		Title      string            `bson:"title" json:"title,omitempty"`
		Ref        string            `bson:"ref" json:"ref"`       // like this: refs/heads/master
		Target     string            `bson:"target" json:"target"` //master or branch.
		Params     map[string]string `bson:"params" json:"params,omitempty"`
		Started    int64             `bson:"started" json:"started,omitempty"`
		Finished   int64             `bson:"finished" json:"finished,omitempty"`
		Created    int64             `bson:"created" json:"created,omitempty"`
		Updated    int64             `bson:"updated" json:"updated,omitempty"`
		BuildSteps []BuildStep       `bson:"steps" json:"steps,omitempty"`
	}

	//BuildStore build for storage.
	BuildStore interface {
		Create(context.Context, *Build) error
	}
)
