// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolOp

type MonoIDBool interface {
	BinaryOpBool
	Zero() bool
	Reduce(done <-chan bool, slice <-chan bool) <-chan bool
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

func (s *monoIDBool) Reduce(done <-chan bool, slice <-chan bool) <-chan bool {
	out := make(chan bool)
	go func() {
		result := s.unit
		for {
			select {
			case value := <-slice:
				result = s.BinaryOpBool.Apply(result, value)
				out <- result
			case <-done:
				close(out)
				return
			}
		}
	}()
	return out
}
