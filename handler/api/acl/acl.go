package acl

import (
	"net/http"

	"github.com/laidingqing/stackbuild/handler/api/request"
	"github.com/laidingqing/stackbuild/handler/api/response"
	"github.com/laidingqing/stackbuild/handler/errors"
	"github.com/laidingqing/stackbuild/logger"
)

// AuthorizeUser returns an http.Handler middleware that protect resources
func AuthorizeUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, ok := request.UserFrom(r.Context())
		if !ok {
			response.Unauthorized(w, errors.ErrUnauthorized)
			logger.FromRequest(r).
				Debugln("api: authentication required")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// AuthorizeAdmin returns an http.Handler middleware that admin chain.
func AuthorizeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := request.UserFrom(r.Context())
		if !ok {
			response.Unauthorized(w, errors.ErrUnauthorized)
			logger.FromRequest(r).
				Debugln("api: authentication required")
		} else if !user.Admin {
			response.Forbidden(w, errors.ErrForbidden)
			logger.FromRequest(r).
				Debugln("api: administrative access required")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
