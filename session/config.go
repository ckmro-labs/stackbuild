package session

import (
	"time"

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
