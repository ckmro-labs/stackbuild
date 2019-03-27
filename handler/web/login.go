package web

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/errors"
	"golang.org/x/crypto/bcrypt"
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
		r.ParseForm()
		login := r.Form.Get("login")
		password := r.Form.Get("password")
		logrus.Infof("password: %v", password)

		user, err := users.FindLogin(r.Context(), login)
		if err != nil {
			logrus.Errorf("query user err: %v", err.Error())
			writeLoginError(w, r, errors.ErrUserExisted)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.EncryptPassword), []byte(password))
		if err != nil {
			writeLoginError(w, r, errors.ErrPasswordNotMatched)
			return
		}

		user.LastLogin = time.Now().Unix()
		// err = users.Update(ctx, user)
		// if err != nil {
		// 	logger.Errorf("cannot update user: %s", err)
		// }
		session.Create(w, user)
		http.Redirect(w, r, "/healthz", 303)
	}
}

func writeLoginError(w http.ResponseWriter, r *http.Request, err error) {
	http.Redirect(w, r, "/static/?err="+err.Error(), 303)
}
