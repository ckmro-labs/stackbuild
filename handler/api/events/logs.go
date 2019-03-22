package events

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/laidingqing/stackbuild/core"
	"github.com/laidingqing/stackbuild/logger"
	"github.com/sirupsen/logrus"
)

//HandleLogStream log stream handler func.
func HandleLogStream(
	events core.Pubsub,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			name = chi.URLParam(r, "name")
		)

		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		logger := logger.FromRequest(r).WithFields(
			logrus.Fields{
				"name": name,
			},
		)

		h := w.Header()
		h.Set("Content-Type", "text/event-stream")
		h.Set("Cache-Control", "no-cache")
		h.Set("Connection", "keep-alive")
		h.Set("X-Accel-Buffering", "no")

		f, ok := w.(http.Flusher)
		if !ok {
			return
		}

		io.WriteString(w, ": ping\n\n")
		f.Flush()

		events, errc := events.Subscribe(ctx)
		logger.Debugln("events: stream opened")

	L:
		for {
			select {
			case <-ctx.Done():
				logger.Debugln("events: stream cancelled")
				break L
			case <-errc:
				logger.Debugln("events: stream error")
				break L
			case <-time.After(time.Hour):
				logger.Debugln("events: stream timeout")
				break L
			case <-time.After(pingInterval):
				io.WriteString(w, ": ping\n\n")
				f.Flush()
			case event := <-events:
				io.WriteString(w, "data: aa")
				w.Write(event.Data)
				io.WriteString(w, "\n\n")
				f.Flush()
			}
		}

		io.WriteString(w, "event: error\ndata: eof\n\n")
		f.Flush()

		logger.Debugln("events: stream closed")
	}
}
