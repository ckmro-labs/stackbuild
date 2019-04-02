package queue

import (
	"context"
	"sync"
	"time"

	"github.com/laidingqing/stackbuild/core"
	"github.com/sirupsen/logrus"
)

type Queue struct {
	sync.Mutex
	Ready    chan struct{}
	Paused   bool
	Interval time.Duration
	Store    core.StageStore
	Workers  map[*worker]struct{}
	Ctx      context.Context
}

type worker struct {
	labels  map[string]string
	channel chan *core.Stage
}

type counter struct {
	counts map[string]int
}

// NewQueue returns a new Queue build stage datastore.
func NewQueue(store core.StageStore) *Queue {
	q := &Queue{
		Store:    store,
		Ready:    make(chan struct{}, 1),
		Workers:  map[*worker]struct{}{},
		Interval: time.Minute,
		Ctx:      context.Background(),
	}
	go q.start()
	return q
}

func (q *Queue) Schedule(ctx context.Context, stage *core.Stage) error {
	select {
	case q.Ready <- struct{}{}:
	default:
	}
	return nil
}

func (q *Queue) Pause(ctx context.Context) error {
	q.Lock()
	q.Paused = true
	q.Unlock()
	return nil
}

func (q *Queue) paused(ctx context.Context) (bool, error) {
	q.Lock()
	paused := q.Paused
	q.Unlock()
	return paused, nil
}

func (q *Queue) Resume(ctx context.Context) error {
	q.Lock()
	q.Paused = false
	q.Unlock()

	select {
	case q.Ready <- struct{}{}:
	default:
	}
	return nil
}

func (q *Queue) start() error {
	for {
		select {
		case <-q.Ctx.Done():
			return q.Ctx.Err()
		case <-q.Ready:
			q.signal(q.Ctx)
		case <-time.After(q.Interval):
			q.signal(q.Ctx)
		}
	}
}

func (q *Queue) signal(ctx context.Context) error {
	q.Lock()
	count := len(q.Workers)
	pause := q.Paused
	q.Unlock()
	if pause {
		return nil
	}
	if count == 0 {
		return nil
	}
	items, err := q.Store.ListIncomplete(ctx)
	if err != nil {
		return err
	}

	q.Lock()
	defer q.Unlock()
	for _, item := range items {
		logrus.Infof("worker item : %v", item.ID)
		if item.Status == core.StatusRunning {
			continue
		}

		if withinLimits(item, items) == false {
			continue
		}

	loop:
		for w := range q.Workers {
			select {
			case w.channel <- item:
				delete(q.Workers, w)
				break loop
			}
		}
	}
	return nil
}

func (q *Queue) Request(ctx context.Context) (*core.Stage, error) {
	w := &worker{
		channel: make(chan *core.Stage),
	}
	q.Lock()
	q.Workers[w] = struct{}{}
	q.Unlock()

	select {
	case q.Ready <- struct{}{}:
	default:
	}

	select {
	case <-ctx.Done():
		q.Lock()
		delete(q.Workers, w)
		q.Unlock()
		return nil, ctx.Err()
	case b := <-w.channel:
		return b, nil
	}
}

func withinLimits(stage *core.Stage, siblings []*core.Stage) bool {
	if stage.Limit == 0 {
		return true
	}
	count := 0
	for _, sibling := range siblings {
		if sibling.RepoID != stage.RepoID {
			continue
		}
		if sibling.ID == stage.ID {
			continue
		}
		if sibling.Name != stage.Name {
			continue
		}
		if sibling.ID < stage.ID {
			count++
		}
	}
	return count < stage.Limit
}
