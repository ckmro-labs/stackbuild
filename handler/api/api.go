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
	Repos  core.RepositoryStore
	Repoz  core.RepositoryService
	Syncer core.Syncer
	Events core.Pubsub
}

//New new api server
func New(
	repos core.RepositoryStore,
	repoz core.RepositoryService,
	syncer core.Syncer,
	events core.Pubsub,
) Server {
	return Server{
		Repos:  repos,
		Repoz:  repoz,
		Syncer: syncer,
		Events: events,
	}
}

// Handler returns an http.Handler
func (s Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(logger.Middleware)
	// r.Use(auth.HandleAuthentication(s.Session))
	cors := cors.New(corsOpts)
	r.Use(cors.Handler)

	r.Route("/repos/{owner}/{name}", func(r chi.Router) {
		// r.Use(acl.InjectRepository)
		r.Get("/", repos.HandleFind())
	})
	r.Route("/users", func(r chi.Router) {
		r.Use(acl.AuthorizeAdmin)
	})

	r.Route("/stream", func(r chi.Router) {
		r.Get("/", events.HandleLogStream(s.Events))
	})

	return r
}
