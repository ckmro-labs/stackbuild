package web

import (
	"net/http"

	"github.com/laidingqing/stackbuild/core"
)

//HandleSignup creates an http.HandlerFunc that can registe account by form data.
func HandleSignup(
	users core.UserStore,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
