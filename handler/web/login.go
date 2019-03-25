package web

import (
	"net/http"

	"github.com/laidingqing/stackbuild/core"
)

//HandleOAuthLogin A 3rd authentication and session initialization.
func HandleOAuthLogin(
	users core.UserStore,
	userz core.UserService,
	syncer core.Syncer,
	session core.Session,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从context中获取上下文
		// 获取token
		// 过期跳转登录表单/3rd platform
		// 查询用户
		// 更新context/cookies.
	}
}

//HandleFormLogin A system authentication and session initialization. include callback by 3rd.
func HandleFormLogin(
	users core.UserStore,
	userz core.UserService,
	syncer core.Syncer,
	session core.Session,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从context中获取上下文，分辩是OAuth回调登录或表单登录

	}
}
