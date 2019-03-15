package middleware

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"
)

//OAuthMiddleware ...
func OAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		q := r.URL.Query()
		provider := chi.URLParam(r, "provider")
		q.Add("provider", provider)
		r.URL.RawQuery = q.Encode()
		gothic.BeginAuthHandler(w, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
