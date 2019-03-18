package wire

import (
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/runner/config"
	"github.com/laidingqing/stackbuild/server"
)

// wire set for loading the server.
var serverSet = wire.NewSet(
	provideServer,
)

func provideServer(config config.Config) *server.GrpcServer {
	return &server.GrpcServer{
		TLS:  config.Server.TLS,
		Port: config.Server.Port,
		Cert: config.Server.Cert,
		Key:  config.Server.Key,
	}
}
