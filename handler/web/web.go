package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/laidingqing/stackbuild/logger"
)

// Server is a http.Handler over HTTP.
type Server struct {
}

//New ...
func New() Server {
	return Server{}
}

// Handler returns an http.Handler
func (s Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(logger.Middleware)

	return r
}
