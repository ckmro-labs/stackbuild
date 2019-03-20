package runtime

import "github.com/laidingqing/stackbuild/runner/executor"

// Option configures a Runtime option.
type Option func(*Runtime)

// WithExecutor sets the Runtime engine.
func WithExecutor(e executor.Executor) Option {
	return func(r *Runtime) {
		r.executor = e
	}
}

// WithConfig sets the Runtime configuration.
func WithConfig(c *executor.Spec) Option {
	return func(r *Runtime) {
		r.config = c
	}
}

// WithHooks sets the Runtime tracer.
func WithHooks(h *Hook) Option {
	return func(r *Runtime) {
		if h != nil {
			r.hook = h
		}
	}
}
