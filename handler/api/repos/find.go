package repos

import (
	"net/http"

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
