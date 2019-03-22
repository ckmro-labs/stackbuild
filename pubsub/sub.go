package pubsub

import (
	"sync"

	"github.com/laidingqing/stackbuild/core"
)

type subscriber struct {
	sync.Mutex

	handler chan *core.Message
	quit    chan struct{}
	done    bool
}

func (s *subscriber) publish(event *core.Message) {
	select {
	case <-s.quit:
	case s.handler <- event:
	default:
	}
}

func (s *subscriber) close() {
	s.Lock()
	if s.done == false {
		close(s.quit)
		s.done = true
	}
	s.Unlock()
}
