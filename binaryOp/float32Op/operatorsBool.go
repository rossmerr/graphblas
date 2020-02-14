// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float32Op

import "github.com/RossMerr/Caudex.GraphBLAS/binaryOp"

// BinaryOpFloat32ToBool is a function that maps two input value to one output value
type BinaryOpFloat32ToBool interface {
	binaryOp.BinaryOp
	Apply(in1, in2 float32) bool
}

type binaryOpFloat32ToBool struct {
	apply func(float32, float32) bool
}

func (s *binaryOpFloat32ToBool) BinaryOp() {}
func (s *binaryOpFloat32ToBool) Operator() {}

func (s *binaryOpFloat32ToBool) Apply(in1, in2 float32) bool {
	return s.apply(in1, in2)
}

// Equal f(x, y) = (x == y)
var Equal = &binaryOpFloat32ToBool{apply: func(in1, in2 float32) bool {
	return in1 == in2
}}

// NotEqual f(x, y) = (x != y)
var NotEqual = &binaryOpFloat32ToBool{apply: func(in1, in2 float32) bool {
	return in1 != in2
}}

// GreaterThan f(x, y) = (x > y)
var GreaterThan = &binaryOpFloat32ToBool{apply: func(in1, in2 float32) bool {
	return in1 > in2
}}

// LessThan f(x, y) = (x < y)
var LessThan = &binaryOpFloat32ToBool{apply: func(in1, in2 float32) bool {
	return in1 < in2
}}

// GreaterThanOrEqual f(x, y) = (x >= y)
var GreaterThanOrEqual = &binaryOpFloat32ToBool{apply: func(in1, in2 float32) bool {
	return in1 >= in2
}}

// LessThanOrEqual f(x, y) = (x <= y)
var LessThanOrEqual = &binaryOpFloat32ToBool{apply: func(in1, in2 float32) bool {
	return in1 <= in2
}}
