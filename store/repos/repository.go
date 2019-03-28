package repos

import (
	"context"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/shared/db"
)

// New returns a new UserStore.
func New(db *db.SessionStore) core.RepositoryStore {
	return &repositoryStore{db}
}

type repositoryStore struct {
	db *db.SessionStore
}

// Find returns a user from the datastore.
func (s *repositoryStore) Find(ctx context.Context, id string) (*core.Repository, error) {

	return nil, nil
}

// FindLogin returns a user from the datastore by repo name.
func (s *repositoryStore) List(ctx context.Context, name string, user *core.User) ([]*core.Repository, error) {
	return nil, nil
}

// FindToken returns a user from the datastore by token.
func (s *repositoryStore) Create(ctx context.Context, repo *core.Repository) error {

	return nil
}

//Delete delete repository
func (s *repositoryStore) Delete(ctx context.Context, repo *core.Repository) error {

	return nil
}

//Activate ..
func (s *repositoryStore) Activate(ctx context.Context, repo *core.Repository) error {

	return nil
}
