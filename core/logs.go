package core

import (
	"context"
	"io"
)

// Line represents a line in the logs.
type Line struct {
	Number    int    `json:"pos"`
	Message   string `json:"out"`
	Timestamp int64  `json:"time"`
}

// LogStore  output to storage.
type LogStore interface {
	// Find returns a log stream from the datastore.
	Find(ctx context.Context, pipeline string) (io.ReadCloser, error)

	// Create writes copies the log stream from Reader r to the datastore.
	Create(ctx context.Context, pipeline string, r io.Reader) error

	// Update writes copies the log stream from Reader r to the datastore.
	Update(ctx context.Context, pipeline string, r io.Reader) error

	// Delete purges the log stream from the datastore.
	Delete(ctx context.Context, pipeline string) error
}

// LogStream manages a live stream of logs.
type LogStream interface {
	// Create creates the log stream for the pipline's step ID.
	Create(context.Context, string) error

	// Delete deletes the log stream for the pipline's ID.
	Delete(context.Context, string) error

	// Writes writes to the log stream.
	Write(context.Context, string, *Line) error

	// Tail tails the log stream.
	Tail(context.Context, string) (<-chan *Line, <-chan error)

	// Info returns internal stream information.
	Info(context.Context) *LogStreamInfo
}

//LogStreamInfo stream's log info.
type LogStreamInfo struct {
	// Streams is a key-value pair where the key is the step
	Streams map[int64]int `json:"streams"`
}
