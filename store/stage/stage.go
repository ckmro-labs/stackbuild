package stage

import (
	"context"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/shared/db"
)

//StageCollName stage db name
var StageCollName = "stages"

// New returns a new RepositoryStore.
func New(db *db.SessionStore) core.StageStore {
	return &stageStore{db}
}

type stageStore struct {
	db *db.SessionStore
}

func (s *stageStore) Create(ctx context.Context, stage *core.Stage) error {
	return nil
}

// Find returns a build stage from the datastore by ID.
func (s *stageStore) Find(ctx context.Context, id int64) (*core.Stage, error) {
	return nil, nil
}

// List returns a build stage list from the datastore, where the stage is incomplete (pending or running).
func (s *stageStore) ListIncomplete(ctx context.Context) ([]*core.Stage, error) {
	items := []*core.Stage{
		{ID: "3", Status: core.StatusRunning},
		{ID: "2", Status: core.StatusPassing},
		{ID: "1", Status: core.StatusPending},
	}
	return items, nil
}
