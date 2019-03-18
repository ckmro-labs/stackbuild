// Copyright 2018 CAIKONG LTD.

//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package wire

import (
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/runner/application"
	"github.com/laidingqing/stackbuild/cmd/runner/config"
)

//InitializeApplication ...
func InitializeApplication(config config.Config) (application.Application, error) {
	wire.Build(
		serverSet,
		application.NewApplication,
	)
	return application.Application{}, nil
}
