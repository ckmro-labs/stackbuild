package runner

import (
	"context"

	"golang.org/x/sync/errgroup"
)

//Runner ..an executor.
type Runner struct {
}

//Run run job.
// env build.
//use dock pull repo.
//use docker build
func (r *Runner) Run(ctx context.Context, id string) error {

	return nil
}

// Start starts N build runner processes. Each process polls
// the server for pednding builds to execute.
func (r *Runner) Start(ctx context.Context, n int) error {
	var g errgroup.Group
	for i := 0; i < n; i++ {
		g.Go(func() error {
			return r.start(ctx)
		})
	}
	return g.Wait()
}

func (r *Runner) start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			r.poll(ctx)
		}
	}
}

func (r *Runner) poll(ctx context.Context) error {

}
