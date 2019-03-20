package runtime

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/laidingqing/stackbuild/runner/executor"
	"github.com/natessilva/dag"
	"golang.org/x/sync/errgroup"
)

// Runtime executes a pipeline configuration.
type Runtime struct {
	mu sync.Mutex

	executor executor.Executor
	config   *executor.Spec
	hook     *Hook
	start    int64
	error    error
}

// New returns a new runtime instance.
func New(opts ...Option) *Runtime {
	r := &Runtime{}
	r.hook = &Hook{}
	for _, opts := range opts {
		opts(r)
	}
	return r
}

// Run starts the pipeline
func (r *Runtime) Run(ctx context.Context) error {
	return r.Resume(ctx, 0)
}

// Resume starts the pipeline waits for it to complete.
func (r *Runtime) Resume(ctx context.Context, start int) error {
	defer func() {
		r.executor.Destroy(context.Background(), r.config)
	}()

	r.error = nil
	r.start = time.Now().Unix()

	if r.hook.Before != nil {
		state := snapshot(r, nil, nil)
		if err := r.hook.Before(state); err != nil {
			return err
		}
	}

	if err := r.executor.Setup(ctx, r.config); err != nil {
		return err
	}

	if isSerial(r.config) {
		for i, step := range r.config.Steps {
			steps := []*executor.Step{step}
			if i < start {
				continue
			}
			select {
			case <-ctx.Done():
				return ErrCancel
			case err := <-r.execAll(steps):
				if err != nil {
					r.error = err
				}
			}
		}
	} else {
		err := r.execGraph(ctx)
		if err != nil {
			return err
		}
	}

	if r.hook.After != nil {
		state := snapshot(r, nil, nil)
		if err := r.hook.After(state); err != nil {
			return err
		}
	}
	return r.error
}

func (r *Runtime) execGraph(ctx context.Context) error {
	var d dag.Runner
	for _, s := range r.config.Steps {
		step := s
		d.AddVertex(step.Metadata.Name, func() error {
			select {
			case <-ctx.Done():
				return ErrCancel
			default:
			}
			r.mu.Lock()
			skip := r.error == ErrInterrupt
			r.mu.Unlock()
			if skip {
				return nil
			}
			err := r.exec(step)
			if err != nil {
				r.mu.Lock()
				r.error = err
				r.mu.Unlock()
			}
			return nil
		})
	}
	for _, s := range r.config.Steps {
		for _, dep := range s.DependsOn {
			d.AddEdge(dep, s.Metadata.Name)
		}
	}
	return d.Run()
}

func (r *Runtime) execAll(group []*executor.Step) <-chan error {
	var g errgroup.Group
	done := make(chan error)
	if r.error == ErrInterrupt {
		close(done)
		return done
	}

	for _, step := range group {
		step := step
		g.Go(func() error {
			return r.exec(step)
		})
	}

	go func() {
		done <- g.Wait()
		close(done)
	}()
	return done
}

func (r *Runtime) exec(step *executor.Step) error {
	ctx := context.Background()

	switch {
	case step.RunPolicy == executor.RunNever:
		return nil
	case r.error != nil && step.RunPolicy == executor.RunOnSuccess:
		return nil
	case r.error == nil && step.RunPolicy == executor.RunOnFailure:
		return nil
	}

	if r.hook.BeforeEach != nil {
		state := snapshot(r, step, nil)
		if err := r.hook.BeforeEach(state); err == ErrSkip {
			return nil
		} else if err != nil {
			return err
		}
	}

	if err := r.executor.Create(ctx, r.config, step); err != nil {
		if r.hook.AfterEach != nil {
			r.hook.AfterEach(
				snapshot(r, step, &executor.State{
					ExitCode: 255, Exited: true,
				}),
			)
		}
		return err
	}
	log.Printf("do executor start...")
	if err := r.executor.Start(ctx, r.config, step); err != nil {
		if r.hook.AfterEach != nil {
			r.hook.AfterEach(
				snapshot(r, step, &executor.State{
					ExitCode: 255, Exited: true,
				}),
			)
		}
		return err
	}

	rc, err := r.executor.Tail(ctx, r.config, step)
	if err != nil {
		if r.hook.AfterEach != nil {
			r.hook.AfterEach(
				snapshot(r, step, &executor.State{
					ExitCode: 255, Exited: true,
				}),
			)
		}
		return err
	}

	var g errgroup.Group
	state := snapshot(r, step, nil)
	g.Go(func() error {
		return stream(state, rc)
	})

	if step.Detach {
		return nil
	}

	defer func() {
		g.Wait()
		rc.Close()
	}()

	wait, err := r.executor.Wait(ctx, r.config, step)
	if err != nil {
		return err
	}

	err = g.Wait()

	if wait.OOMKilled {
		err = &OomError{
			Name: step.Metadata.Name,
			Code: wait.ExitCode,
		}
	} else if wait.ExitCode == 78 {
		err = ErrInterrupt
	} else if wait.ExitCode != 0 {
		err = &ExitError{
			Name: step.Metadata.Name,
			Code: wait.ExitCode,
		}
	}

	if r.hook.AfterEach != nil {
		state := snapshot(r, step, wait)
		if err := r.hook.AfterEach(state); err != nil {
			return err
		}
	}

	if step.IgnoreErr {
		return nil
	}
	return err
}

func stream(state *State, rc io.ReadCloser) error {
	defer rc.Close()

	w := newWriter(state)
	io.Copy(w, rc)

	if state.hook.GotLogs != nil {
		return state.hook.GotLogs(state, w.lines)
	}
	return nil
}
