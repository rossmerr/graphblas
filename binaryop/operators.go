// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryop

import (
	"github.com/rossmerr/graphblas/constraints"
)

// FirstArgument f(x, y) = x
func FirstArgument[T constraints.Number]() BinaryOp[T] {
	return NewBinaryOp(func(in1, in2 T) T {
		return in1
	})
}

// SecondArgument f(x, y) = y
func SecondArgument[T constraints.Number]() BinaryOp[T] {
	return NewBinaryOp(func(in1, in2 T) T {
		return in2
	})
}

// Minimum f(x, y) = (x < y) ? x : y
func Minimum[T constraints.Number]() BinaryOp[T] {
	return NewBinaryOp(func(in1, in2 T) T {
		if in1 < in2 {
			return in1
		}
		return in2
	})
}

// Maximum f(x, y) = (x > y) ? x : y
func Maximum[T constraints.Number]() BinaryOp[T] {
	return NewBinaryOp(func(in1, in2 T) T {
		if in1 > in2 {
			return in1
		}

		return in2
	})
}

// Addition f(x, y) = x + y
func Addition[T constraints.Number]() BinaryOp[T] {
	return NewBinaryOp(func(in1, in2 T) T {
		return in1 + in2
	})
}

// Subtraction f(x, y) = x - y
func Subtraction[T constraints.Number]() BinaryOp[T] {
	return NewBinaryOp(func(in1, in2 T) T {
		return in1 - in2
	})
}

// Multiplication f(x, y) = x * y
func Multiplication[T constraints.Number]() BinaryOp[T] {
	return NewBinaryOp(func(in1, in2 T) T {
		return in1 * in2
	})
}

// Division f(x, y) = x / y
func Division[T constraints.Number]() BinaryOp[T] {
	return NewBinaryOp(func(in1, in2 T) T {
		return in1 / in2
	})
}
