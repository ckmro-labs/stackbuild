package wire

import (
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/pubsub"
	"github.com/laidingqing/stackbuild/service/repo"
	"github.com/laidingqing/stackbuild/service/syncer"
	"github.com/laidingqing/stackbuild/service/user"
	"github.com/laidingqing/stackbuild/session"
)

// wire set for loading the services.
var serviceSet = wire.NewSet(
	pubsub.New,
	user.New,
	providerRepositoryService,
	provideSyncer,
	provideSession,
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

func provideSession(store core.UserStore, config config.Config) core.Session {
	return session.New(store, session.NewConfig(
		config.Session.Secret,
		config.Session.Timeout),
	)
}
