package github

import (
	"net/http"
	"strings"

	"github.com/laidingqing/stackbuild/middleware/login/oauth"
)

// Config github oauth config.
type Config struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	Server       string
	Scope        []string
}

// Handler returns a http.Handler
func (c *Config) Handler(h http.Handler) http.Handler {
	server := normalizeAddress(c.Server)
	return oauth.Handler(h, &oauth.Config{
		BasicAuthOff:     true,
		Client:           c.Client,
		ClientID:         c.ClientID,
		ClientSecret:     c.ClientSecret,
		AccessTokenURL:   server + "/login/oauth/access_token",
		AuthorizationURL: server + "/login/oauth/authorize",
		Scope:            c.Scope,
	})
}

func normalizeAddress(address string) string {
	if address == "" {
		return "https://github.com"
	}
	return strings.TrimSuffix(address, "/")
}
