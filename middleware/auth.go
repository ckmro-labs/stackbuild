package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"
)

//OAuthMiddleware ...
func OAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		q := r.URL.Query()
		provider := chi.URLParam(r, "provider")
		if provider == "" {
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		q.Add("provider", provider)
		r.URL.RawQuery = q.Encode()
		logrus.Infof("Raw: %v", r.URL.RawQuery)
		if user, err := gothic.CompleteUserAuth(w, r); err != nil {
			gothic.BeginAuthHandler(w, r)
		} else {
			logrus.Infof("user: %v", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
