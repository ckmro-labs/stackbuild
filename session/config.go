package session

import "time"

// Config provides the session configuration.
type Config struct {
	Secret  string
	Timeout time.Duration
}

// NewConfig returns a new session configuration.
func NewConfig(secret string, timeout time.Duration) Config {
	return Config{
		Secret:  secret,
		Timeout: timeout,
	}
}
