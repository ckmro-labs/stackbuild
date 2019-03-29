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
		ID        string `bson:"_id" json:"id,omitempty"`
		UID       string `bson:"uid" json:"uid,omitempty"`
		UserID    string `bson:"userId" json:"user_id,omitempty"`
		Provider  string `bson:"provider" json:"provider,omitempty"`
		Namespace string `bson:"nameSpace" json:"namespace,omitempty"`
		Name      string `bson:"name" json:"name,omitempty"`
		Slug      string `bson:"slug" json:"slug,omitempty"`
		HTTPURL   string `bson:"httpUrl" json:"git_http_url,omitempty"`
		SSHURL    string `bson:"sshUrl" json:"git_ssh_url,omitempty"`
		Link      string `bson:"link" json:"link,omitempty"`
		Branch    string `bson:"branch" json:"default_branch,omitempty"`
		Private   bool   `bson:"private" json:"private,omitempty"`
		Timeout   int64  `bson:"timeOut" json:"timeout,omitempty"`
		Created   int64  `bson:"created" json:"created,omitempty"`
		Updated   int64  `bson:"updated" json:"updated,omitempty"`
	}

	// RepositoryStore 仓库操作接口
	RepositoryStore interface {
		List(context.Context, string, *User) ([]*Repository, error)
		Find(context.Context, string) (*Repository, error)
		FindByProvider(context.Context, string, string) (*Repository, error)
		Create(context.Context, *Repository) error
		Delete(context.Context, *Repository) error
		Activate(context.Context, *Repository) error
	}

	//RepositoryService 提供远程仓库接口操作
	RepositoryService interface {
		// List returns a list of repositories.
		List(ctx context.Context, token *Token) ([]*Repository, error)
		// Find returns the named repository details.
		Find(ctx context.Context, token *Token, id string) (*Repository, error)
	}
)

//Syncer ..
type Syncer interface {
	Sync(context.Context, *Token) error
}
