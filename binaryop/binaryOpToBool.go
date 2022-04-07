// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryop

import "github.com/rossmerr/graphblas/constraints"

type BinaryOpToBool[T constraints.None] interface {
	Apply(T, T) bool
}

type binaryOpToBool[T constraints.None] struct {
	apply func(T, T) bool
}

func NewBinaryOpToBool[T constraints.None](apply func(T, T) bool) BinaryOpToBool[T] {
	return &binaryOpToBool[T]{apply: apply}
}

func (s *binaryOpToBool[T]) BinaryOp() {}
func (s *binaryOpToBool[T]) Operator() {}

func (s *binaryOpToBool[T]) Apply(in1, in2 T) bool {
	return s.apply(in1, in2)
}
