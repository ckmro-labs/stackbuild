package repo

import (
	"context"

	"github.com/drone/go-scm/scm"
	"github.com/laidingqing/stackbuild/core"
)

type service struct {
	client *scm.Client
}

// New RepositoryService实例，提供操作源代码仓库信息
func New(client *scm.Client) core.RepositoryService {
	return &service{
		client: client,
	}
}

//List return all repository by owner.
func (s *service) List(ctx context.Context, user *core.User, provider core.VcsProvider) ([]*core.Repository, error) {
	token, refresh, err := userToken(user, provider)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, scm.TokenKey{}, &scm.Token{
		Token:   token,
		Refresh: refresh,
	})
	repos := []*core.Repository{}
	opts := scm.ListOptions{Size: 100}
	for {
		result, meta, err := s.client.Repositories.List(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, src := range result {
			repos = append(repos, convertRepository(src))
		}
		opts.Page = meta.Page.Next
		opts.URL = meta.Page.NextURL

		if opts.Page == 0 && opts.URL == "" {
			break
		}
	}
	return repos, nil
}

//Find
func (s *service) Find(ctx context.Context, user *core.User, repo string, provider core.VcsProvider) (*core.Repository, error) {
	return nil, nil
}
