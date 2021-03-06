package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/api/acl"
	"github.com/laidingqing/stackbuild/handler/api/events"
	"github.com/laidingqing/stackbuild/handler/api/repos"
	"github.com/laidingqing/stackbuild/handler/api/stages"
	"github.com/laidingqing/stackbuild/handler/api/users"
	"github.com/laidingqing/stackbuild/logger"
)

var corsOpts = cors.Options{
	AllowedOrigins:   []string{"*"},
	AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: true,
	MaxAge:           300,
}

// Server is a http.Handler which exposes drone functionality over HTTP.
type Server struct {
	Repos   core.RepositoryStore
	Repoz   core.RepositoryService
	Syncer  core.Syncer
	Events  core.Pubsub
	Users   core.UserStore
	Stages  core.StageStore
	Session core.Session
}

//New new api server
func New(
	repos core.RepositoryStore,
	repoz core.RepositoryService,
	syncer core.Syncer,
	events core.Pubsub,
	users core.UserStore,
	stages core.StageStore,
	session core.Session,
) Server {
	return Server{
		Repos:   repos,
		Repoz:   repoz,
		Syncer:  syncer,
		Events:  events,
		Users:   users,
		Stages:  stages,
		Session: session,
	}
}

// Handler returns an http.Handler
func (s Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(logger.Middleware)
	cors := cors.New(corsOpts)
	r.Use(cors.Handler)

	r.Route("/repos/{owner}/{name}", func(r chi.Router) {
		r.Use(acl.AuthorizeSessionUser(s.Session))
		r.Use(acl.InjectRepository(s.Repoz, s.Repos))

		r.Get("/", repos.HandleFind(s.Repos))
		// r.Post("/building", builds.HandleBuild(s.Repos))
		r.Post("/stages", stages.HandleCreatePipelineStage(s.Stages))
		r.Get("/stages", stages.HandleListPipelineStage(s.Repos))
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", users.HandleCreateUser(s.Users))
		r.Get("/repos", repos.HandleListRepos(s.Repos))
		r.Post("/login", users.HandleUserLogin(s.Users))
	})

	r.Route("/stream", func(r chi.Router) {
		r.Get("/", events.HandleLogStream(s.Events))
	})

	return r
}
