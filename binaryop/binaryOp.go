// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryop

import (
	"github.com/rossmerr/graphblas/constraints"
)

type IBinaryOp interface {
	binaryOp()
//	Apply(interface{}, interface{}) interface{}
}

// BinaryOpFloat64 is a function that maps two input value to one output value
type BinaryOp[T constraints.None] interface {
	Semigroup
	IBinaryOp
	Apply(T, T) T
	BinaryOp()
	Operator()
}

func NewBinaryOp[T constraints.None](apply func(T, T) T) BinaryOp[T] {
	return &binaryOp[T]{apply: apply}
}

type binaryOp[T constraints.None] struct {
	IBinaryOp
	apply func(T, T) T
}

func (s *binaryOp[T]) Operator()  {}
func (s *binaryOp[T]) BinaryOp()  {}
func (s *binaryOp[T]) binaryOp()  {}
func (s *binaryOp[T]) Semigroup() {}

func (s *binaryOp[T]) Apply(in1, in2 T) T {
	return s.apply(in1, in2)
}
