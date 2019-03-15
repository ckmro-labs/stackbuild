package web

import (
	"net/http"

	"github.com/laidingqing/stackbuild/core"
)

//HandleHook hook handle
func HandleHook(
	repos core.RepositoryStore,
	parser core.HookParser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
