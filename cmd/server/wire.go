// Copyright 2018 CAIKONG LTD.

//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/server/config"
)

//InitializeApplication ...
func InitializeApplication(config config.Config) (application, error) {
	wire.Build(
		serverSet,
		newApplication,
	)
	return application{}, nil
}
