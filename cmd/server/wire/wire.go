// Copyright 2018 CAIKONG LTD.

//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package wire

import (
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/server/application"
	"github.com/laidingqing/stackbuild/cmd/server/config"
)

//InitializeApplication ...
func InitializeApplication(config config.Config) (application.Application, error) {
	wire.Build(
		serverSet,
		storeSet,
		application.NewApplication,
	)
	return application.Application{}, nil
}
