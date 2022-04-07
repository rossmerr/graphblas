package operators

import (
	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/constraints"
)

// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// Addition f(x, y) = x + y
func Addition[T constraints.Number]() binaryop.BinaryOp[T] {
	return binaryop.NewBinaryOp(func(in1, in2 T) T {
		return in1 + in2
	})
}

// Subtraction f(x, y) = x - y
func Subtraction[T constraints.Number]() binaryop.BinaryOp[T] {
	return binaryop.NewBinaryOp(func(in1, in2 T) T {
		return in1 - in2
	})
}

// Multiplication f(x, y) = x * y
func Multiplication[T constraints.Number]() binaryop.BinaryOp[T] {
	return binaryop.NewBinaryOp(func(in1, in2 T) T {
		return in1 * in2
	})
}

// Division f(x, y) = x / y
func Division[T constraints.Number]() binaryop.BinaryOp[T] {
	return binaryop.NewBinaryOp(func(in1, in2 T) T {
		return in1 / in2
	})
}
