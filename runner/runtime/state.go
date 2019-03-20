package runtime

import "github.com/laidingqing/stackbuild/runner/executor"

//State runtime's process state for pipeline.
type State struct {
	hook     *Hook
	config   *executor.Spec
	executor executor.Executor

	Runtime struct {
		Time  int64
		Error error
	}
	Step  *executor.Step
	State *executor.State
}

func snapshot(r *Runtime, step *executor.Step, state *executor.State) *State {
	s := new(State)
	s.Runtime.Error = r.error
	s.Runtime.Time = r.start
	s.config = r.config
	s.hook = r.hook
	s.executor = r.executor
	s.Step = step
	s.State = state
	return s
}
