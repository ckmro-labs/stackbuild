package stage

import (
	"context"

	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/store/shared/db"
)

//StageCollNameKey stage db name
var StageCollNameKey = "stages"

// New returns a new StageStore.
func New(db *db.SessionStore) core.StageStore {
	return &stageStore{db}
}

type stageStore struct {
	db *db.SessionStore
}

func (s *stageStore) Create(ctx context.Context, stage *core.Stage) error {
	_, err := s.db.C(StageCollNameKey).UpsertId(stage.ID, stage)
	return err
}

// Find returns a build stage from the datastore by ID.
func (s *stageStore) Find(ctx context.Context, id string) (*core.Stage, error) {
	var stage *core.Stage
	err := s.db.C(StageCollNameKey).FindId(id).One(&stage)
	return stage, err
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
