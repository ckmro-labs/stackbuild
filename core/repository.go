package core

import "context"

// VcsProvider types.
type VcsProvider string

const (
	//VcsProviderGitHub github vcs provider
	VcsProviderGitHub VcsProvider = "github"
	//VcsProviderGitLab gitlab vcs provider
	VcsProviderGitLab VcsProvider = "gitlab"
	//VcsProviderGitee gitee vcs provider
	VcsProviderGitee VcsProvider = "gitee"
)

func (p VcsProvider) String() string {
	switch p {
	case VcsProviderGitHub:
		return "github"
	case VcsProviderGitLab:
		return "gitlab"
	case VcsProviderGitee:
		return "gitee"
	default:
		return "UNKNOWN"
	}
}

type (
	// Repository represents a source code repository.
	Repository struct {
		ID        int64       `json:"id"`
		UID       string      `json:"uid"`
		UserID    int64       `json:"user_id"`
		Provider  VcsProvider `json:"provider"`
		Namespace string      `json:"namespace"`
		Name      string      `json:"name"`
		Slug      string      `json:"slug"`
		SCM       string      `json:"scm"`
		HTTPURL   string      `json:"git_http_url"`
		SSHURL    string      `json:"git_ssh_url"`
		Link      string      `json:"link"`
		Branch    string      `json:"default_branch"`
		Private   bool        `json:"private"`
		Timeout   int64       `json:"timeout"`
		Created   int64       `json:"created"`
		Updated   int64       `json:"updated"`
		Version   int64       `json:"version"`
	}

	// RepositoryStore 仓库操作接口
	RepositoryStore interface {
		List(context.Context, string) ([]*Repository, error)
		Find(context.Context, string) (*Repository, error)
		Create(context.Context, *Repository) error
		Delete(context.Context, *Repository) error
		Activate(context.Context, *Repository) error
	}

	//RepositoryService 提供远程仓库接口操作
	RepositoryService interface {
		// List returns a list of repositories.
		List(ctx context.Context, user *User, provider VcsProvider) ([]*Repository, error)
		// Find returns the named repository details.
		Find(ctx context.Context, user *User, id string, provider VcsProvider) (*Repository, error)
	}
)

//Syncer ..
type Syncer interface {
	Sync(context.Context, *User, VcsProvider) error
}
