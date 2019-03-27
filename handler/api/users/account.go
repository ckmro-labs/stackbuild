package users

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/handler/api/response"
	"github.com/laidingqing/stackbuild/logger"
	"golang.org/x/crypto/bcrypt"
)

//HandleCreateUser create a user.
func HandleCreateUser(users core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		in := new(core.User)
		err := json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			response.BadRequest(w, err)
			logger.FromRequest(r).WithError(err).
				Debugln("api: cannot unmarshal request body")
			return
		}
		logrus.Infof("pass:%v", in.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		logrus.Infof("hashedPassword:%v", string(hashedPassword))
		user := &core.User{
			Login:           in.Login,
			EncryptPassword: string(hashedPassword),
			Active:          true, //TODO 确认激活账号
			Admin:           false,
			Created:         time.Now().Unix(),
			Updated:         time.Now().Unix(),
		}

		err = users.Create(r.Context(), user)
		if err != nil {
			response.InternalError(w, err)
			logger.FromRequest(r).WithError(err).
				Warnln("api: cannot create user")
			return
		}
		var out interface{} = user
		response.JSON(w, out, 200)
	}
}
