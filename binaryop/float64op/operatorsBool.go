// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64op

import "github.com/rossmerr/graphblas/binaryop"

// BinaryOpFloat64ToBool is a function that maps two input value to one output value
type BinaryOpFloat64ToBool interface {
	binaryop.BinaryOp
	Apply(in1, in2 float64) bool
}

type binaryOpFloat64ToBool struct {
	apply func(float64, float64) bool
}

func (s *binaryOpFloat64ToBool) BinaryOp() {}
func (s *binaryOpFloat64ToBool) Operator() {}

func (s *binaryOpFloat64ToBool) Apply(in1, in2 float64) bool {
	return s.apply(in1, in2)
}

// Equal f(x, y) = (x == y)
var Equal = &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
	return in1 == in2
}}

// NotEqual f(x, y) = (x != y)
var NotEqual = &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
	return in1 != in2
}}

// GreaterThan f(x, y) = (x > y)
var GreaterThan = &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
	return in1 > in2
}}

// LessThan f(x, y) = (x < y)
var LessThan = &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
	return in1 < in2
}}

// GreaterThanOrEqual f(x, y) = (x >= y)
var GreaterThanOrEqual = &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
	return in1 >= in2
}}

// LessThanOrEqual f(x, y) = (x <= y)
var LessThanOrEqual = &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
	return in1 <= in2
}}
