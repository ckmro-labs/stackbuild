package repos

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/drone/drone/handler/api/render"
	"github.com/go-chi/chi"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/api/request"
	"github.com/laidingqing/stackbuild/handler/api/response"
	"github.com/laidingqing/stackbuild/handler/errors"
	"gopkg.in/mgo.v2/bson"
)

// HandleFind returns an http.HandlerFunc
// response is repostory list
func HandleFind(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx   = r.Context()
			owner = chi.URLParam(r, "owner")
			name  = chi.URLParam(r, "name")
		)

		if repo, _ := repos.Query(ctx, map[string]interface{}{
			"query": bson.M{
				"nameSpace": owner,
				"name":      name,
			},
			"pagination": false,
		}); repo != nil && len(repo) > 0 {
			response.JSON(w, repo[0], 200)
		}
		response.JSON(w, nil, 200)
	}
}

//HandleListRepos list remote repository.
func HandleListRepos(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var ctx = r.Context()
		var user *core.User
		if user, ok := request.UserFrom(ctx); !ok {
			logrus.Infof("context user: %v", user.ID)
			render.InternalError(w, errors.ErrUnauthorized)
		}
		repos, err := repos.Query(ctx, map[string]interface{}{
			"query": bson.M{
				"userId": user.ID,
			},
			"pagination": false,
		})
		if err != nil {
			render.InternalError(w, err)
		}
		response.JSON(w, repos, 200)
	}
}
