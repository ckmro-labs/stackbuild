package runtime

import (
	"errors"
	"fmt"
)

var (
	// ErrSkip ...
	ErrSkip = errors.New("Skipped")
	// ErrCancel ...
	ErrCancel = errors.New("Cancelled")
	// ErrInterrupt .
	ErrInterrupt = errors.New("Interrupt")
)

// ExitError ...
type ExitError struct {
	Name string
	Code int
}

// Error ...
func (e *ExitError) Error() string {
	return fmt.Sprintf("%s : exit code %d", e.Name, e.Code)
}

// OomError ..
type OomError struct {
	Name string
	Code int
}

// Error ...
func (e *OomError) Error() string {
	return fmt.Sprintf("%s : received oom kill", e.Name)
}
