package core

import "context"

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
