package core

import "context"

type (
	// Step 构建开始后步骤信息.
	Step struct {
		ID        int64  `json:"id"`
		StageID   int64  `json:"step_id"`
		Number    int    `json:"number"`
		Name      string `json:"name"`
		Status    string `json:"status"`
		Error     string `json:"error,omitempty"`
		ErrIgnore bool   `json:"errignore,omitempty"`
		ExitCode  int    `json:"exit_code"`
		Started   int64  `json:"started,omitempty"`
		Stopped   int64  `json:"stopped,omitempty"`
		Version   int64  `json:"version"`
	}

	// StepStore step's storage.
	StepStore interface {
		// List returns a build stage list from the datastore.
		List(context.Context, string) ([]*Step, error)

		// Find returns a build stage from the datastore by ID.
		Find(context.Context, string) (*Step, error)

		// FindNumber returns a stage from the datastore by number.
		FindNumber(context.Context, string, int) (*Step, error)

		// Create persists a new stage to the datastore.
		Create(context.Context, *Step) error

		// Update persists an updated stage to the datastore.
		Update(context.Context, *Step) error
	}
)

//IsDone 获取当前步骤是否完成了。
func (s *Step) IsDone() bool {
	switch s.Status {
	case StatusWaiting,
		StatusPending,
		StatusRunning,
		StatusBlocked:
		return false
	default:
		return true
	}
}
