// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryop

import (
	"github.com/rossmerr/graphblas/constraints"
)

// MonoIDFloat64 is a set of float64's that closed under an associative binary operation
type MonoID[T constraints.None] interface {
	Zero() T
	Reduce(done <-chan struct{}, slice <-chan T) <-chan T
}

type monoID[T constraints.None] struct {
	BinaryOp[T]
	unit T
}

// Zero the identity element
func (s *monoID[T]) Zero() T {
	return s.unit
}

// NewMonoID retun a MonoIDFloat64
func NewMonoID[T constraints.None](zero T, operator BinaryOp[T]) MonoID[T] {
	return &monoID[T]{unit: zero, BinaryOp: operator}
}

// Reduce left folding over the monoID
func (s *monoID[T]) Reduce(done <-chan struct{}, slice <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		result := s.unit
		for {
			select {
			case value := <-slice:
				result = s.BinaryOp.Apply(result, value)
			case <-done:
				out <- result
				close(out)
				return
			}
		}
	}()
	return out
}
