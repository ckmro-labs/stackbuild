package wire

import (
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/shared/db"
	"github.com/laidingqing/stackbuild/store/user"
)

// wire set for loading the stores.
var storeSet = wire.NewSet(
	provideDatabase,
	provideUserStore,
)

// provideDatabase is a Wire provider
func provideDatabase(config config.Config) (*db.SessionStore, error) {
	return nil, db.Connect(config.Database.Datasource)
}

// provideUserStore is a Wire provider
func provideUserStore(db *db.SessionStore) core.UserStore {
	users := user.New(db)
	return users
}
