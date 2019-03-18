package request

import (
	"context"

	"github.com/laidingqing/stackbuild/core"
)

type key int

const (
	userKey key = iota
	permKey
	repoKey
)

// WithRepo returns a copy of repository into context
func WithRepo(parent context.Context, repo *core.Repository) context.Context {
	return context.WithValue(parent, repoKey, repo)
}

// WithUser returns a copy of user into context
func WithUser(parent context.Context, user *core.User) context.Context {
	return context.WithValue(parent, userKey, user)
}

// RepoFrom returns the value of the repo key on the context
func RepoFrom(ctx context.Context) (*core.Repository, bool) {
	repo, ok := ctx.Value(repoKey).(*core.Repository)
	return repo, ok
}
