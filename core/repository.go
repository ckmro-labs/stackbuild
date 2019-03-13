package core

import "context"

type (
	// Repository represents a source code repository.
	Repository struct {
		ID         int64  `json:"id"`
		UID        string `json:"uid"`
		UserID     int64  `json:"user_id"`
		Namespace  string `json:"namespace"`
		Name       string `json:"name"`
		Slug       string `json:"slug"`
		SCM        string `json:"scm"`
		HTTPURL    string `json:"git_http_url"`
		SSHURL     string `json:"git_ssh_url"`
		Link       string `json:"link"`
		Branch     string `json:"default_branch"`
		Private    bool   `json:"private"`
		Visibility string `json:"visibility"`
		Timeout    int64  `json:"timeout"`
		Created    int64  `json:"created"`
		Updated    int64  `json:"updated"`
		Version    int64  `json:"version"`
	}

	// RepositoryStore 仓库操作接口
	RepositoryStore interface {
		List(context.Context, int64) ([]*Repository, error)
		Find(context.Context, int64) (*Repository, error)
		Create(context.Context, *Repository) error
		Delete(context.Context, *Repository) error
		Activate(context.Context, *Repository) error
	}

	//RepositoryService 提供远程仓库接口操作
	RepositoryService interface {
		// List returns a list of repositories.
		List(ctx context.Context, user *User) ([]*Repository, error)
		// Find returns the named repository details.
		Find(ctx context.Context, user *User, repo string) (*Repository, error)
	}
)
