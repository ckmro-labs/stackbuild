package acl

import (
	"net/http"

	"github.com/laidingqing/stackbuild/core"
)

//InjectRepository set repos into context
func InjectRepository(
	repoz core.RepositoryService,
	repos core.RepositoryStore,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
