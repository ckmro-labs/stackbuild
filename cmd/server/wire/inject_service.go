package wire

import (
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/service/repo"
	"github.com/laidingqing/stackbuild/service/syncer"
)

// wire set for loading the services.
var serviceSet = wire.NewSet(
	providerRepositoryService,
	provideSyncer,
)

func providerRepositoryService(config config.Config) core.RepositoryService {
	return repo.New(config)
}

func provideSyncer(repoz core.RepositoryService,
	repos core.RepositoryStore,
	users core.UserStore,
	config config.Config) core.Syncer {
	return syncer.New(repoz, repos, users)
}
