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
	// STEP1: get the list of repositories from the remote
	//
	{
		repos, err := s.repoz.List(ctx, token)
		if err != nil {
			return err
		}
		logrus.Infof("Repository size: %v", len(repos))
	}
	//
	// STEP2: get the list of repositories stored in the local db
	//

	//
	// STEP3 no found in local. Insert.
	//

	//
	// STEP3 exist in local. Update.
	//

	//
	// STEP3 exist in local, but not exist in remote. delete.
	//

	//
	// STEP4 update the store.
	//

	return nil
}
