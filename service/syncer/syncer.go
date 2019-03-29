package syncer

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/laidingqing/stackbuild/core"
)

//Synchronizer sync repository between remote and local.
type Synchronizer struct {
	repoz core.RepositoryService
	repos core.RepositoryStore
	users core.UserStore
}

// New returns a new Synchronizer.
func New(
	repoz core.RepositoryService,
	repos core.RepositoryStore,
	users core.UserStore,
) *Synchronizer {
	return &Synchronizer{
		repoz: repoz,
		repos: repos,
		users: users,
	}
}

//Sync sync remote repository to local
func (s *Synchronizer) Sync(ctx context.Context, token *core.Token) error {
	//
	// save local repository db. from remote repository.
	//
	{
		repos, err := s.repoz.List(ctx, token)
		if err != nil {
			return err
		}
		logrus.Infof("Repository size: %v", len(repos))
		for _, rep := range repos {
			err := s.repos.Create(ctx, rep)
			if err != nil {
				logrus.Errorf("save or update repository err: %v", err.Error())
			}
		}
	}

	return nil
}
