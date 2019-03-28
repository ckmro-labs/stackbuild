package web

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/errors"
	"github.com/laidingqing/stackbuild/logger"
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
		// 授权回调
		// * 根据Context Session判断是否是登录用户，是进行绑定第三方授权信息
		// * 新用户创建用户并绑定
		// * 创建会话
		ctx := r.Context()
		log := logger.FromContext(ctx)
		user, err := session.Get(r)
		source := core.TokenFrom(ctx)
		if err == nil && user != nil || user.ID != "" {
			log.Debugf("已登录用户: %v", user.ID)
			for _, auth := range user.Authentications {
				if auth.AuthName.String() == source.Provider {
					auth.Expired = source.Expires.Unix()
					auth.Token = source.Access
					auth.Refresh = source.Refresh
				}
			}

		} else {
			// TODO 非登录用户
		}

		if is := isSync(user.Synced); is {
			go synchornize(r.Context(), syncer, user)
		}

		session.Create(w, user)
		http.Redirect(w, r, "/healthz", 303)
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

		if is := isSync(user.Synced); is {
			go synchornize(r.Context(), syncer, user)
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

func isSync(synced int64) bool {
	return time.Unix(synced, 0).Add(time.Hour * 24 * 7).Before(time.Now())
}

//synchornize start sync repo info.
func synchornize(ctx context.Context, syncer core.Syncer, user *core.User) {
	log := logrus.WithField("login", user.Login)
	log.Debugf("begin synchronization")
	timeout, cancel := context.WithTimeout(context.Background(), time.Minute*30)
	timeout = logger.WithContext(timeout, log)
	defer cancel()

	for _, auth := range user.Authentications {
		err := syncer.Sync(timeout, &core.Token{
			Provider: auth.AuthName.String(),
			Access:   auth.Token,
			Refresh:  auth.Refresh,
		})
		if err != nil {
			log.Debugf("synchronization failed: %s", err)
		} else {
			log.Debugf("synchronization success")
		}
	}
}
