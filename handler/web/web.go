package web

import (
	"net/http"

	"github.com/drone/go-scm/scm"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/logger"
)

// Server is a http.Handler over HTTP.
type Server struct {
	Client *scm.Client
	Hooks  core.HookParser
	Repos  core.RepositoryStore
}

//New ...
func New(
	client *scm.Client,
	hooks core.HookParser,
	repos core.RepositoryStore,
) Server {
	return Server{
		Client: client,
		Hooks:  hooks,
		Repos:  repos,
	}
}

// Handler returns an http.Handler
func (s Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(logger.Middleware)

	r.Route("/hook", func(r chi.Router) {
		//来自版本仓库的hook请求
		r.Post("/", HandleHook(s.Repos, s.Hooks))
	})
	r.Get("/healthz", HandleHealthz())

	r.Handle("/login", nil)

	return r
}
