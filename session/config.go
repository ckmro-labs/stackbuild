package session

import (
	"net/http"
	"strings"
	"time"

	"github.com/dchest/authcookie"
	"github.com/laidingqing/stackbuild/core"
)

// Config provides the session configuration.
type Config struct {
	Secret  string
	Timeout time.Duration
}

type session struct {
	users   core.UserStore
	secret  []byte
	timeout time.Duration
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration) Config {
	return Config{
		Secret:  secret,
		Timeout: timeout,
	}
}

// New returns a new cookie-based session management.
func New(users core.UserStore, config Config) core.Session {
	return &session{
		secret:  []byte(config.Secret),
		timeout: config.Timeout,
		users:   users,
	}
}

//core.sessionService implement.

func (s *session) Create(w http.ResponseWriter, user *core.User) error {
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

func (s *session) Delete(w http.ResponseWriter) error {
	w.Header().Add("Set-Cookie", "_session_=deleted; Path=/; Max-Age=0")
	return nil
}

func (s *session) Get(r *http.Request) (*core.User, error) {
	switch {
	case isAuthorizationToken(r):
		return s.fromToken(r)
	case isAuthorizationParameter(r):
		return s.fromToken(r)
	default:
		return s.fromSession(r)
	}
}

func (s *session) fromSession(r *http.Request) (*core.User, error) {
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

func (s *session) fromToken(r *http.Request) (*core.User, error) {
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
