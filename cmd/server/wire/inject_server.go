package wire

import (
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/laidingqing/stackbuild/auth/provider"
	"github.com/laidingqing/stackbuild/cmd/server/config"
	"github.com/laidingqing/stackbuild/handler/api"
	"github.com/laidingqing/stackbuild/handler/web"
	"github.com/laidingqing/stackbuild/server"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
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
	provideOauthProvider(config)
	return &server.Server{
		TLS:     config.Server.TLS,
		Addr:    config.Server.Port,
		Cert:    config.Server.Cert,
		Key:     config.Server.Key,
		Host:    config.Server.Host,
		Handler: handler,
	}
}

func provideOauthProvider(config config.Config) {
	providerMap := make(map[string]*provider.Provider)
	providers := make([]provider.Provider, 0)
	if config.Github.ClientID != "" {
		p, _ := provider.New("github", config.Github.ClientSecret, config.Github.ClientID, config.Github.CallbackURL+"/"+"github")
		providers = append(providers, *p)
	}
	for i := range providers {
		providerMap[providers[i].Name] = &providers[i]
		goth.UseProviders(*providers[i].Implementation)
	}

	store := cookie.NewStore([]byte(""))
	store.Options(sessions.Options{
		MaxAge: int(time.Second * time.Duration(21600)),
		Path:   "/",
	})
	gothic.Store = store
}
