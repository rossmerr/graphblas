package operators

// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import (
	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/constraints"
)

// Equal f(x, y) = (x == y)
func Equal[T constraints.Number]() binaryop.BinaryOpToBool[T] {
	return binaryop.NewBinaryOpToBool(func(in1, in2 T) bool {
		return in1 == in2
	})
}

// NotEqual f(x, y) = (x != y)
func NotEqual[T constraints.Number]() binaryop.BinaryOpToBool[T] {
	return binaryop.NewBinaryOpToBool(func(in1, in2 T) bool {
		return in1 != in2
	})
}

// GreaterThan f(x, y) = (x > y)
func GreaterThan[T constraints.Number]() binaryop.BinaryOpToBool[T] {
	return binaryop.NewBinaryOpToBool(func(in1, in2 T) bool {
		return in1 > in2
	})
}

// LessThan f(x, y) = (x < y)
func LessThan[T constraints.Number]() binaryop.BinaryOpToBool[T] {
	return binaryop.NewBinaryOpToBool(func(in1, in2 T) bool {
		return in1 < in2
	})
}

// GreaterThanOrEqual f(x, y) = (x >= y)
func GreaterThanOrEqual[T constraints.Number]() binaryop.BinaryOpToBool[T] {
	return binaryop.NewBinaryOpToBool(func(in1, in2 T) bool {
		return in1 >= in2
	})
}

// LessThanOrEqual f(x, y) = (x <= y)
func LessThanOrEqual[T constraints.Number]() binaryop.BinaryOpToBool[T] {
	return binaryop.NewBinaryOpToBool(func(in1, in2 T) bool {
		return in1 <= in2
	})
}
