package source

import (
	"context"
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/shared/db"
)

//SourceAuthCollKey source auth mongo collection
const SourceAuthCollKey = "source_auths"

// errs definition
var (
	ErrorProviderNameExisted = fmt.Errorf("provider name existed")
)

// New returns a new UserStore.
func New(db *db.SessionStore) core.SourceAuthStore {
	return &sourceStore{db}
}

type sourceStore struct {
	db *db.SessionStore
}

//Create add a source auth by user.
func (s *sourceStore) Create(ctx context.Context, sourceAuth *core.SourceAuth) error {
	source, _ := s.Find(ctx, sourceAuth.AuthName, sourceAuth.UID)
	if source.ID == "" {
		source = sourceAuth
		source.ID = bson.NewObjectId().Hex()
	}
	return s.db.C(SourceAuthCollKey).UpdateId(sourceAuth.ID, sourceAuth)
}

// Find find a source auth provider
func (s *sourceStore) Find(ctx context.Context, provider string, uid string) (*core.SourceAuth, error) {
	query := bson.M{
		"uid":  uid,
		"name": provider,
	}
	var source *core.SourceAuth
	err := s.db.C(SourceAuthCollKey).Find(query).One(&source)

	return source, err
}

// List find a source auth provider
func (s *sourceStore) List(ctx context.Context, userID string) ([]*core.SourceAuth, error) {
	query := bson.M{
		"userId": userID,
	}
	var sources []*core.SourceAuth
	err := s.db.C(SourceAuthCollKey).Find(query).All(&sources)

	return sources, err
}
