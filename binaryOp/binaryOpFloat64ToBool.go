// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

type BinaryOpFloat64ToBool interface {
	Apply(in1, in2 float64) bool
	binaryOp()
	operator()
}

type binaryOpFloat64ToBool struct {
	apply func(float64, float64) bool
}

func (s *binaryOpFloat64ToBool) binaryOp() {}
func (s *binaryOpFloat64ToBool) operator() {}

func (s *binaryOpFloat64ToBool) Apply(in1, in2 float64) bool {
	return s.apply(in1, in2)
}

// Equal f(x, y) = (x == y)
var Equal = func() BinaryOp {
	return &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
		return in1 == in2
	}}
}

// NotEqual f(x, y) = (x != y)
var NotEqual = func() BinaryOp {
	return &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
		return in1 != in2
	}}
}

// GreaterThan f(x, y) = (x > y)
var GreaterThan = func() BinaryOp {
	return &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
		return in1 > in2
	}}
}

// LessThan f(x, y) = (x < y)
var LessThan = func() BinaryOp {
	return &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
		return in1 < in2
	}}
}

// GreaterThanOrEqual f(x, y) = (x >= y)
var GreaterThanOrEqual = func() BinaryOp {
	return &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
		return in1 >= in2
	}}
}

// LessThanOrEqual f(x, y) = (x <= y)
var LessThanOrEqual = func() BinaryOp {
	return &binaryOpFloat64ToBool{apply: func(in1, in2 float64) bool {
		return in1 <= in2
	}}
}
