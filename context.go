package GraphBLAS

import (
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
	context.Context
	mode Mode
}

// NewContext returns a Context
func NewContext(ctx context.Context, mode Mode) context.Context {
	return &graphContect{ctx, mode}
}

func (s *graphContect) Mode() Mode {
	return s.mode
}

// BlockingMode returns the Mode from the context
func BlockingMode(ctx context.Context) (Mode, bool) {
	graph, ok := ctx.(*graphContect)
	if ok {
		return graph.Mode(), ok
	}
	return Blocking, false
}
