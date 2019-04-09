package builds

import (
	"net/http"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/api/response"
)

//HandleTryBuild try build
func HandleTryBuild(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, nil, 200)
	}
}
