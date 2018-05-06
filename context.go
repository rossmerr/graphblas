package GraphBLAS

import (
	"time"

	"golang.org/x/net/context"
)

// Mode for blocking
type Mode = int

const (
	// Blocking until each method has completed
	Blocking Mode = iota
	// NonBlocking may execute methods in any order
	NonBlocking
)

// Context exends the standard context to include Mode
type Context interface {
	context.Context
	Mode() Mode
}

// graphContect context with Mode support
type graphContect struct {
	mode Mode
}

// NewContext returns a Context
func NewContext(mode Mode) Context {
	return &graphContect{
		mode: mode,
	}
}

// Done returns a channel that is closed when this Context is canceled
// or times out.
func (s *graphContect) Done() <-chan struct{} {
	return nil
}

// Err indicates why this context was canceled, after the Done channel
// is closed.
func (s *graphContect) Err() error {
	return nil
}

// Deadline returns the time when this Context will be canceled, if any.
func (s *graphContect) Deadline() (deadline time.Time, ok bool) {
	return time.Now(), false
}

// Value returns the value associated with key or nil if none.
func (s *graphContect) Value(key interface{}) interface{} {
	return nil
}

func (s *graphContect) Mode() Mode {
	return s.mode
}
