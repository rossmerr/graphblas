package operators

// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import (
	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/constraints"
)

// FirstArgument f(x, y) = x
func FirstArgument[T constraints.Number]() binaryop.BinaryOp[T] {
	return binaryop.NewBinaryOp(func(in1, in2 T) T {
		return in1
	})
}

// SecondArgument f(x, y) = y
func SecondArgument[T constraints.Number]() binaryop.BinaryOp[T] {
	return binaryop.NewBinaryOp(func(in1, in2 T) T {
		return in2
	})
}

// Minimum f(x, y) = (x < y) ? x : y
func Minimum[T constraints.Number]() binaryop.BinaryOp[T] {
	return binaryop.NewBinaryOp(func(in1, in2 T) T {
		if in1 < in2 {
			return in1
		}

		return in2
	})
}

// Maximum f(x, y) = (x > y) ? x : y
func Maximum[T constraints.Number]() binaryop.BinaryOp[T] {
	return binaryop.NewBinaryOp(func(in1, in2 T) T {
		if in1 > in2 {
			return in1
		}

		return in2
	})
}
