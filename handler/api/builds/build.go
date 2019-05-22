package builds

import (
	"net/http"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/api/response"
)

//HandleBuild  build
//generator all build step from repo's stage.
func HandleBuild(builds core.BuildStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		build := &core.Build{}

		//transform build step.. from service build.

		err := builds.Create(r.Context(), build)
		if err != nil {

		}
		response.JSON(w, nil, 200)
	}
}
