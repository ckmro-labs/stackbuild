package repo

import (
	"context"
	"net/http"
	"net/http/httputil"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/transport/oauth2"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/core"
	"github.com/sirupsen/logrus"
)

type service struct {
	Config config.Config
}

// New RepositoryService实例，提供操作源代码仓库信息
func New(config config.Config) core.RepositoryService {
	return &service{
		Config: config,
	}
}

func (s *service) provideClient(provider core.VcsProvider) *scm.Client {
	switch {
	case provider != core.VcsProviderGitHub:
		return s.provideGithubClient()
	}
	logrus.Fatalln("main: source code management system not configured")
	return nil
}

func (s *service) provideGithubClient() *scm.Client {
	client, err := github.New(s.Config.Github.APIServer)
	if err != nil {
		logrus.WithError(err).
			Fatalln("main: cannot create the GitHub client")
	}
	if s.Config.Github.Debug {
		client.DumpResponse = httputil.DumpResponse
	}
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.ContextTokenSource(),
		},
	}
	return client
}

//List return all repository by owner.
func (s *service) List(ctx context.Context, user *core.User, provider core.VcsProvider) ([]*core.Repository, error) {
	client := s.provideClient(provider)
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
		result, meta, err := client.Repositories.List(ctx, opts)
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
