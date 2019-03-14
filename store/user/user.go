package user

import (
	"context"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/shared/db"
)

// New returns a new UserStore.
func New(db *db.SessionStore) core.UserStore {
	return &userStore{db}
}

type userStore struct {
	db *db.SessionStore
}

// Find returns a user from the datastore.
func (s *userStore) Find(ctx context.Context, id int64) (*core.User, error) {

	return nil, nil
}

// FindLogin returns a user from the datastore by username.
func (s *userStore) FindLogin(ctx context.Context, login string) (*core.User, error) {
	return nil, nil
}

// FindToken returns a user from the datastore by token.
func (s *userStore) FindToken(ctx context.Context, token string) (*core.User, error) {

	return nil, nil
}
