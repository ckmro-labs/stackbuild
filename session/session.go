package session

import (
	"net/http"
	"strings"
	"time"

	"github.com/dchest/authcookie"
	"github.com/laidingqing/stackbuild/core"
)

//Session ...
type Session struct {
	users         core.UserStore
	secret        []byte
	timeout       time.Duration
	administrator string // administrator account
}

// New returns a new cookie-based session management.
func New(users core.UserStore, config Config) core.Session {
	return &Session{
		secret:  []byte(config.Secret),
		timeout: config.Timeout,
		users:   users,
	}
}

//Create write session to http response.
func (s *Session) Create(w http.ResponseWriter, user *core.User) error {
	cookie := &http.Cookie{
		Name:     "_session_",
		Path:     "/",
		MaxAge:   2147483647,
		HttpOnly: true,
		Value: authcookie.NewSinceNow(
			user.Login,
			s.timeout,
			s.secret,
		),
	}
	w.Header().Add("Set-Cookie", cookie.String()+"; SameSite=lax")
	return nil
}

func (s *Session) Delete(w http.ResponseWriter) error {
	w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
	return nil
}

func (s *Session) Get(r *http.Request) (*core.User, error) {
	switch {
	case isAuthorizationToken(r):
		return s.fromToken(r)
	case isAuthorizationParameter(r):
		return s.fromToken(r)
	default:
		return s.fromSession(r)
	}
}

func (s *Session) fromSession(r *http.Request) (*core.User, error) {
	cookie, err := r.Cookie("_session_")
	if err != nil {
		return nil, nil
	}
	login := authcookie.Login(cookie.Value, s.secret)
	if login == "" {
		return nil, nil
	}
	return s.users.FindLogin(r.Context(), login)
}

func (s *Session) fromToken(r *http.Request) (*core.User, error) {
	return s.users.FindToken(r.Context(),
		extractToken(r),
	)
}

func isAuthorizationToken(r *http.Request) bool {
	return r.Header.Get("Authorization") != ""
}

func isAuthorizationParameter(r *http.Request) bool {
	return r.FormValue("access_token") != ""
}

func extractToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		bearer = r.FormValue("access_token")
	}
	return strings.TrimPrefix(bearer, "Bearer ")
}
