package acl

import (
	"net/http"

	"github.com/drone/drone/handler/api/render"
	"github.com/go-chi/chi"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/api/request"
	"github.com/laidingqing/stackbuild/handler/errors"
	"github.com/laidingqing/stackbuild/logger"
	"gopkg.in/mgo.v2/bson"
)

//InjectRepository set repos into context
func InjectRepository(
	repoz core.RepositoryService,
	repos core.RepositoryStore,
) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				ctx   = r.Context()
				owner = chi.URLParam(r, "owner")
				name  = chi.URLParam(r, "name")
			)
			repo, _ := repos.Query(ctx, map[string]interface{}{
				"query": bson.M{
					"nameSpace": owner,
					"name":      name,
				},
				"pagination": false,
			})
			if repo != nil && len(repo) > 0 {
				ctx = request.WithRepo(ctx, repo[0])
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				render.NotFound(w, errors.ErrNotFound)
				logger.FromRequest(r).Debugln("api inject: no found resource by repository")
			}
		})
	}
}
