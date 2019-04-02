package queue

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/mock"
)

func TestQueue(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	items := []*core.Stage{
		{ID: "3"},
		{ID: "2"},
		{ID: "1"},
	}

	ctx := context.Background()
	store := mock.NewMockStageStore(controller)
	store.EXPECT().ListIncomplete(ctx).Return(items, nil).Times(1)
	store.EXPECT().ListIncomplete(ctx).Return(items[1:], nil).Times(1)
	store.EXPECT().ListIncomplete(ctx).Return(items[2:], nil).Times(1)

	q := newQueue(store)
	for _, item := range items {
		next, err := q.Request(ctx)
		if err != nil {
			t.Error(err)
			return
		}
		if got, want := next, item; got != want {
			t.Errorf("Want build %s, got %s", item.ID, item.ID)
		}
	}
}
func TestQueuePush(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	item1 := &core.Stage{
		ID: "1",
	}
	item2 := &core.Stage{
		ID: "2",
	}

	ctx := context.Background()
	store := mock.NewMockStageStore(controller)

	q := &queue{
		store: store,
		ready: make(chan struct{}, 1),
	}
	q.Schedule(ctx, item1)
	q.Schedule(ctx, item2)
	select {
	case <-q.ready:
	case <-time.After(time.Millisecond):
		t.Errorf("Expect queue signaled on push")
	}
}
