package stages

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/drone/drone/handler/api/render"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/api/request"
	"github.com/laidingqing/stackbuild/handler/api/response"
)

//HandleCreatePipelineStage 创建构建管道场景
func HandleCreatePipelineStage(stages core.StageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		repo, _ := request.RepoFrom(ctx)
		stage := new(core.Stage)
		err := json.NewDecoder(r.Body).Decode(stage)
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		stage.RepoID = repo.ID
		stage.ID = bson.NewObjectId().Hex()
		stage.Created = time.Now().Unix()
		err = stages.Create(ctx, stage)
		if err != nil {
			response.InternalError(w, err)
		}
		response.JSON(w, stage, 200)
	}
}

//HandleListPipelineStage find repo's stages
func HandleListPipelineStage(repos core.RepositoryStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, nil, 200)
	}
}
