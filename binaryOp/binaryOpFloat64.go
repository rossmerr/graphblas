// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

type BinaryOpFloat64 interface {
	Semigroup
	Apply(in1, in2 float64) float64
}

type binaryOpFloat64 struct {
	apply func(float64, float64) float64
}

func (s *binaryOpFloat64) operator()  {}
func (s *binaryOpFloat64) binaryOp()  {}
func (s *binaryOpFloat64) semigroup() {}

func (s *binaryOpFloat64) Apply(in1, in2 float64) float64 {
	return s.apply(in1, in2)
}

// FirstArgument f(x, y) = x
var FirstArgument = func() BinaryOp {
	return &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
		return in1
	}}
}

// SecondArgument f(x, y) = y
var SecondArgument = func() BinaryOp {
	return &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
		return in2
	}}
}

// Minimum f(x, y) = (x < y) ? x : y
var Minimum = func() BinaryOp {
	return &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
		if in1 < in2 {
			return in1
		}

		return in2
	}}
}

// Maximum f(x, y) = (x > y) ? x : y
var Maximum = func() BinaryOp {
	return &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
		if in1 > in2 {
			return in1
		}

		return in2
	}}
}

// Addition f(x, y) = x + y
var Addition = func() BinaryOp {
	return &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
		return in1 + in2
	}}
}

// Subtraction f(x, y) = x - y
var Subtraction = func() BinaryOp {
	return &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
		return in1 - in2
	}}
}

// Multiplication f(x, y) = x * y
var Multiplication = func() BinaryOp {
	return &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
		return in1 * in2
	}}
}

// Division f(x, y) = x / y
var Division = func() BinaryOp {
	return &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
		return in1 / in2
	}}
}
