package runtime

import (
	"strings"
	"time"

	"github.com/laidingqing/stackbuild/runner/executor"
)

// Line 容器执行日志，对接主应用Logs.
type Line struct {
	Number    int    `json:"pos,omitempty"`
	Message   string `json:"out,omitempty"`
	Timestamp int64  `json:"time,omitempty"`
}

type lineWriter struct {
	num   int
	now   time.Time
	rep   *strings.Replacer
	state *State
	lines []*Line
	size  int
	limit int
}

func newWriter(state *State) *lineWriter {
	w := &lineWriter{}
	w.num = 0
	w.now = time.Now().UTC()
	w.state = state
	w.rep = newReplacer(state.config.Secrets)
	w.limit = 5242880 // 5MB max log size
	return w
}

func (w *lineWriter) Write(p []byte) (n int, err error) {
	if w.size >= w.limit {
		return len(p), nil
	}

	out := string(p)
	if w.rep != nil {
		out = w.rep.Replace(out)
	}

	parts := []string{out}

	if strings.Contains(strings.TrimSuffix(out, "\n"), "\n") {
		parts = strings.SplitAfter(out, "\n")
	}

	for _, part := range parts {
		line := &Line{
			Number:    w.num,
			Message:   part,
			Timestamp: int64(time.Since(w.now).Seconds()),
		}

		if w.state.hook.GotLine != nil {
			w.state.hook.GotLine(w.state, line)
		}
		w.size = w.size + len(part)
		w.num++

		w.lines = append(w.lines, line)
	}

	if w.size >= w.limit {
		w.lines = append(w.lines, &Line{
			Number:    w.num,
			Message:   "warning: maximum output exceeded",
			Timestamp: int64(time.Since(w.now).Seconds()),
		})
	}

	return len(p), nil
}

func newReplacer(secrets []*executor.Secret) *strings.Replacer {
	var oldnew []string
	for _, secret := range secrets {
		oldnew = append(oldnew, secret.Data)
		oldnew = append(oldnew, "********")
	}
	if len(oldnew) == 0 {
		return nil
	}
	return strings.NewReplacer(oldnew...)
}
