package repos

import (
	"context"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/shared/db"
	"gopkg.in/mgo.v2/bson"
)

//RepositoryCollName repos db name
var RepositoryCollName = "repos"

// New returns a new RepositoryStore.
func New(db *db.SessionStore) core.RepositoryStore {
	return &repositoryStore{db}
}

type repositoryStore struct {
	db *db.SessionStore
}

// Find returns a user from the datastore.
func (s *repositoryStore) Find(ctx context.Context, id string) (*core.Repository, error) {
	var repository *core.Repository
	err := s.db.C(RepositoryCollName).FindId(id).One(&repository)

	return repository, err
}

//FindByProvider find repository by uid and provider
func (s *repositoryStore) FindByProvider(ctx context.Context, uid string, provider string) (*core.Repository, error) {
	query := bson.M{
		"provider": provider,
		"uid":      uid,
	}
	var repository *core.Repository
	err := s.db.C(RepositoryCollName).Find(query).One(&repository)
	return repository, err
}

// FindLogin returns a user from the datastore by repo name.
func (s *repositoryStore) List(ctx context.Context, name string, user *core.User) ([]*core.Repository, error) {
	return nil, nil
}

// FindToken returns a user from the datastore by token.
func (s *repositoryStore) Create(ctx context.Context, repo *core.Repository) error {

	repoEntry, _ := s.FindByProvider(ctx, repo.UID, repo.Provider)
	if repoEntry == nil {
		repoEntry = repo
		repoEntry.ID = bson.NewObjectId().Hex()
	}
	_, err := s.db.C(RepositoryCollName).UpsertId(repoEntry.ID, repoEntry)
	return err
}

//Delete delete repository
func (s *repositoryStore) Delete(ctx context.Context, repo *core.Repository) error {

	return nil
}

//Activate ..
func (s *repositoryStore) Activate(ctx context.Context, repo *core.Repository) error {

	return nil
}
