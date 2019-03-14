package wire

import (
	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/handler/api"
	"github.com/laidingqing/stackbuild/handler/web"
	"github.com/laidingqing/stackbuild/server"
)

// wire set for loading the server.
var serverSet = wire.NewSet(
	api.New,
	web.New,
	provideRouter,
	provideServer,
)

func provideRouter(api api.Server, web web.Server) *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/api", api.Handler())
	r.Mount("/", web.Handler())

	return r
}

func provideServer(handler *chi.Mux, config config.Config) *server.Server {
	return &server.Server{
		TLS:     config.Server.TLS,
		Addr:    config.Server.Port,
		Cert:    config.Server.Cert,
		Key:     config.Server.Key,
		Host:    config.Server.Host,
		Handler: handler,
	}
}
