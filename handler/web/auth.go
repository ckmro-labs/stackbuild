package web

import (
	"net/http"

	"github.com/go-chi/chi"
)

//OAuthHandler ...
func OAuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		provider := chi.URLParam(r, "provider")
		if provider != "" {

		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
