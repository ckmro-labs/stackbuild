package web

import (
	"net/http"
)

//HandleLogin login handler. include authentication and session initialization.
//callback auto login.
func HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从context中获取上下文
		// 获取token
		// 过期跳转登录表单/3rd platform
		// 查询用户
		// 更新context/cookies.
	}
}
