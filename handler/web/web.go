package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/logger"
	m2 "github.com/laidingqing/stackbuild/middleware"
)

// Server is a http.Handler over HTTP.
type Server struct {
	Repos   core.RepositoryStore
	Session core.Session
	Users   core.UserStore
	Userz   core.UserService
	Syncer  core.Syncer
}

//New ...
func New(
	repos core.RepositoryStore,
	session core.Session,
	users core.UserStore,
	userz core.UserService,
	syncer core.Syncer,
) Server {
	return Server{
		Repos:   repos,
		Session: session,
		Users:   users,
		Userz:   userz,
		Syncer:  syncer,
	}
}

// Handler returns an http.Handler
func (s Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(logger.Middleware)

	// r.Route("/hook", func(r chi.Router) {
	// 	//来自版本仓库的hook请求
	// 	// r.Post("/", HandleHook(s.Repos, s.Hooks))
	// })
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.Get("/healthz", HandleHealthz())
	r.Route("/login", func(r chi.Router) {
		r.Post("/form", http.HandlerFunc(HandleFormLogin(
			s.Users,
			s.Userz,
			s.Syncer,
			s.Session,
		)))
		r.Route("/{provider}", func(r chi.Router) {
			r.Use(m2.OAuthMiddleware)
			r.Get("/", http.HandlerFunc(HandleOAuthLogin(
				s.Users,
				s.Userz,
				s.Syncer,
				s.Session,
			)))
		})
	})

	return r
}
