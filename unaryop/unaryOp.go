// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package unaryop

import "github.com/rossmerr/graphblas/constraints"

// UnaryOpFloat64 is a function that maps one input value to one output value
type UnaryOp[T constraints.None] interface {
	Apply(T) T
	UnaryOp()
	Operator()
}

func NewUnaryOp[T constraints.None](apply func(T) T) UnaryOp[T] {
	return &unaryOp[T]{apply: apply}
}

type unaryOp[T constraints.None] struct {
	apply func(T) T
}

func (s *unaryOp[T]) Operator() {}
func (s *unaryOp[T]) UnaryOp()  {}

func (s *unaryOp[T]) Apply(in T) T {
	return s.apply(in)
}
