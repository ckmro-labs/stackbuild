package login

import (
	"context"
	"net/http"
	"time"
)

type key int

const (
	tokenKey key = iota
	errorKey
)

// LoginMiddleware provides login middleware.
type LoginMiddleware interface {
	// Handler returns a http.Handler
	Handler(h http.Handler) http.Handler
}

// Token represents an authorization token.
type Token struct {
	Access  string
	Refresh string
	Expires time.Time
}

// WithToken returns a context with the token.
func WithToken(ctx context.Context, token *Token) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

// WithError returns a parent context with the error.
func WithError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errorKey, err)
}

// TokenFrom returns the login token rom the context.
func TokenFrom(ctx context.Context) *Token {
	token, _ := ctx.Value(tokenKey).(*Token)
	return token
}

// ErrorFrom returns the login error from the context.
func ErrorFrom(ctx context.Context) error {
	err, _ := ctx.Value(errorKey).(error)
	return err
}
