package core

import "context"

// Status types.
const (
	StatusSkipped  = "skipped"
	StatusBlocked  = "blocked"
	StatusDeclined = "declined"
	StatusWaiting  = "waiting_on_dependencies"
	StatusPending  = "pending"
	StatusRunning  = "running"
	StatusPassing  = "success"
	StatusFailing  = "failure"
	StatusKilled   = "killed"
	StatusError    = "error"
)

type (
	// Status ..
	Status struct {
		State  string
		Label  string
		Desc   string
		Target string
	}

	// StatusInput set the commit or deployment status's meta
	StatusInput struct {
		Repo  *Repository
		Build *Build
	}

	// StatusService sends the commit status
	StatusService interface {
		Send(ctx context.Context, user *User, req *StatusInput) error
	}
)
