// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float32op

import "github.com/rossmerr/graphblas/binaryop"

// BinaryOpFloat32 is a function that maps two input value to one output value
type BinaryOpFloat32 interface {
	binaryop.Semigroup
	Apply(in1, in2 float32) float32
}

type binaryOpFloat32 struct {
	apply func(float32, float32) float32
}

func (s *binaryOpFloat32) Operator()  {}
func (s *binaryOpFloat32) BinaryOp()  {}
func (s *binaryOpFloat32) Semigroup() {}

func (s *binaryOpFloat32) Apply(in1, in2 float32) float32 {
	return s.apply(in1, in2)
}

// FirstArgument f(x, y) = x
var FirstArgument = &binaryOpFloat32{apply: func(in1, in2 float32) float32 {
	return in1
}}

// SecondArgument f(x, y) = y
var SecondArgument = &binaryOpFloat32{apply: func(in1, in2 float32) float32 {
	return in2
}}

// Minimum f(x, y) = (x < y) ? x : y
var Minimum = &binaryOpFloat32{apply: func(in1, in2 float32) float32 {
	if in1 < in2 {
		return in1
	}

	return in2
}}

// Maximum f(x, y) = (x > y) ? x : y
var Maximum = &binaryOpFloat32{apply: func(in1, in2 float32) float32 {
	if in1 > in2 {
		return in1
	}

	return in2
}}

// Addition f(x, y) = x + y
var Addition = &binaryOpFloat32{apply: func(in1, in2 float32) float32 {
	return in1 + in2
}}

// Subtraction f(x, y) = x - y
var Subtraction = &binaryOpFloat32{apply: func(in1, in2 float32) float32 {
	return in1 - in2
}}

// Multiplication f(x, y) = x * y
var Multiplication = &binaryOpFloat32{apply: func(in1, in2 float32) float32 {
	return in1 * in2
}}

// Division f(x, y) = x / y
var Division = &binaryOpFloat32{apply: func(in1, in2 float32) float32 {
	return in1 / in2
}}
