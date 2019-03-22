package pubsub

import (
	"context"
	"sync"

	"github.com/laidingqing/stackbuild/core"
)

type hub struct {
	sync.Mutex

	subs map[*subscriber]struct{}
}

// New creates a new publish subscriber.
func New() core.Pubsub {
	return &hub{
		subs: map[*subscriber]struct{}{},
	}
}

func (h *hub) Publish(ctx context.Context, e *core.Message) error {
	h.Lock()
	for s := range h.subs {
		s.publish(e)
	}
	h.Unlock()
	return nil
}

func (h *hub) Subscribe(ctx context.Context) (<-chan *core.Message, <-chan error) {
	h.Lock()
	s := &subscriber{
		handler: make(chan *core.Message, 100),
		quit:    make(chan struct{}),
	}
	h.subs[s] = struct{}{}
	h.Unlock()
	errc := make(chan error)
	go func() {
		defer close(errc)
		select {
		case <-ctx.Done():
			h.Lock()
			delete(h.subs, s)
			h.Unlock()
			s.close()
		}
	}()
	return s.handler, errc
}

func (h *hub) Subscribers() int {
	h.Lock()
	c := len(h.subs)
	h.Unlock()
	return c
}
