// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64Op

import "github.com/RossMerr/Caudex.GraphBLAS/binaryOp"

// BinaryOpFloat64 is a function that maps two input value to one output value
type BinaryOpFloat64 interface {
	binaryOp.Semigroup
	Apply(in1, in2 float64) float64
}

type binaryOpFloat64 struct {
	apply func(float64, float64) float64
}

func (s *binaryOpFloat64) Operator()  {}
func (s *binaryOpFloat64) BinaryOp()  {}
func (s *binaryOpFloat64) Semigroup() {}

func (s *binaryOpFloat64) Apply(in1, in2 float64) float64 {
	return s.apply(in1, in2)
}

// FirstArgument f(x, y) = x
var FirstArgument = &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
	return in1
}}

// SecondArgument f(x, y) = y
var SecondArgument = &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
	return in2
}}

// Minimum f(x, y) = (x < y) ? x : y
var Minimum = &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
	if in1 < in2 {
		return in1
	}

	return in2
}}

// Maximum f(x, y) = (x > y) ? x : y
var Maximum = &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
	if in1 > in2 {
		return in1
	}

	return in2
}}

// Addition f(x, y) = x + y
var Addition = &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
	return in1 + in2
}}

// Subtraction f(x, y) = x - y
var Subtraction = &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
	return in1 - in2
}}

// Multiplication f(x, y) = x * y
var Multiplication = &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
	return in1 * in2
}}

// Division f(x, y) = x / y
var Division = &binaryOpFloat64{apply: func(in1, in2 float64) float64 {
	return in1 / in2
}}
