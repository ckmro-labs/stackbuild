package core

type (
	//Build 一次构建信息, 它包含Step信息。
	Build struct {
		ID        int64             `json:"id"`
		RepoID    int64             `json:"repo_id"`
		Trigger   string            `json:"trigger"`
		Status    string            `json:"status"`
		Error     string            `json:"error,omitempty"`
		Link      string            `json:"link"`
		Timestamp int64             `json:"timestamp"`
		Title     string            `json:"title,omitempty"`
		Message   string            `json:"message"`
		Before    string            `json:"before"`
		After     string            `json:"after"`
		Ref       string            `json:"ref"`
		Target    string            `json:"target"` //master or branch.
		Author    string            `json:"author"`
		Sender    string            `json:"sender"`
		Params    map[string]string `json:"params,omitempty"`
		Started   int64             `json:"started"`
		Finished  int64             `json:"finished"`
		Created   int64             `json:"created"`
		Updated   int64             `json:"updated"`
	}

	//BuildStore build for storage.
	BuildStore interface {
	}
)
