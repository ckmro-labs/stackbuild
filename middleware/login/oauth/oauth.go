package oauth

import (
	"errors"
	"net/http"
	"time"

	"github.com/laidingqing/stackbuild/middleware/login"
)

// Handler returns the oauth2 authorization flow.
func Handler(h http.Handler, c *Config) http.Handler {
	return &handler{next: h, conf: c}
}

type handler struct {
	conf *Config
	next http.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// 检查url query paramater.
	if erro := r.FormValue("error"); erro != "" {
		ctx = login.WithError(ctx, errors.New(erro))
		h.next.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	// 检查是否有code
	code := r.FormValue("code")
	if len(code) == 0 {
		state := createState(w)
		http.Redirect(w, r, h.conf.authorizeRedirect(state), 303)
		return
	}

	// 检查state.
	state := r.FormValue("state")
	deleteState(w)
	if err := validateState(r, state); err != nil {
		ctx = login.WithError(ctx, err)
		h.next.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	// 转换access_token.
	source, err := h.conf.exchange(code, state)
	if err != nil {
		ctx = login.WithError(ctx, err)
		h.next.ServeHTTP(w, r.WithContext(ctx))
		return
	}

	// 转换为上下文token.
	ctx = login.WithToken(ctx, &login.Token{
		Access:  source.AccessToken,
		Refresh: source.RefreshToken,
		Expires: time.Now().UTC().Add(
			time.Duration(source.Expires) * time.Second,
		),
	})

	h.next.ServeHTTP(w, r.WithContext(ctx))
}
