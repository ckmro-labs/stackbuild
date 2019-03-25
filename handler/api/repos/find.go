package repos

import (
	"net/http"

	"github.com/drone/drone/handler/api/render"
	"github.com/go-chi/chi"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/api/request"
	"github.com/laidingqing/stackbuild/handler/api/response"
)

// HandleFind returns an http.HandlerFunc
// response is repostory list
func HandleFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		repo, _ := request.RepoFrom(ctx)
		response.JSON(w, repo, 200)
	}
}

//HandleListRepos list remote repository.
func HandleListRepos(repoz core.RepositoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		provider := chi.URLParam(r, "provider")
		user, _ := request.UserFrom(ctx)
		repos, err := repoz.List(ctx, user, request.ProviderFrom(provider))
		if err != nil {
			render.InternalError(w, err)
		}
		response.JSON(w, repos, 200)
	}
}
