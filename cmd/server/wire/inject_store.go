package wire

import (
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/repos"
	"github.com/laidingqing/stackbuild/store/shared/db"
	"github.com/laidingqing/stackbuild/store/source"
	"github.com/laidingqing/stackbuild/store/user"
)

// wire set for loading the stores.
var storeSet = wire.NewSet(
	provideDatabase,
	provideUserStore,
	provideRepositoryStore,
	provideSourceStore,
)

// provideDatabase is a Wire provider
func provideDatabase(config config.Config) *db.SessionStore {
	return db.NewSessionStore(config.Database.Datasource, config.Database.Database)
}

// provideUserStore is a Wire provider
func provideUserStore(db *db.SessionStore) core.UserStore {
	users := user.New(db)
	return users
}

// provideUserStore is a Wire provider
func provideRepositoryStore(db *db.SessionStore) core.RepositoryStore {
	repository := repos.New(db)
	return repository
}

// provideSourceStore is a Wire provider
func provideSourceStore(db *db.SessionStore) core.SourceAuthStore {
	source := source.New(db)
	return source
}
