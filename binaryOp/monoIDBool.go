// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

type MonoIDBool interface {
	BinaryOpBool
	Zero() bool
	Reduce(done <-chan struct{}, slice <-chan bool) <-chan bool
}

type monoIDBool struct {
	BinaryOpBool
	unit bool
}

func (s *monoIDBool) Zero() bool {
	return s.unit
}

// NewMonoIDBool retuns a MonoIDBool
func NewMonoIDBool(zero bool, operator BinaryOpBool) MonoIDBool {
	return &monoIDBool{unit: zero, BinaryOpBool: operator}
}

func (s *monoIDBool) Reduce(done <-chan struct{}, slice <-chan bool) <-chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)
		for value := range slice {
			select {
			case out <- s.BinaryOpBool.Apply(s.unit, value):
			case <-done:
				return
			}
		}
	}()
	return out
}
