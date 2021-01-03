// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolop

// MonoIDBool is a set of bool's that closed under an associative binary operation
type MonoIDBool interface {
	Zero() bool
	Reduce(done <-chan struct{}, slice <-chan bool) <-chan bool
}

type monoIDBool struct {
	BinaryOpBool
	unit bool
}

// Zero the identity element
func (s *monoIDBool) Zero() bool {
	return s.unit
}

// NewMonoIDBool retun a MonoIDBool
func NewMonoIDBool(zero bool, operator BinaryOpBool) MonoIDBool {
	return &monoIDBool{unit: zero, BinaryOpBool: operator}
}

// Reduce left folding over the monoID
func (s *monoIDBool) Reduce(done <-chan struct{}, slice <-chan bool) <-chan bool {
	out := make(chan bool)
	go func() {
		result := s.unit
		for {
			select {
			case value := <-slice:
				result = s.BinaryOpBool.Apply(result, value)
			case <-done:
				out <- result
				close(out)
				return
			}
		}
	}()
	return out
}
