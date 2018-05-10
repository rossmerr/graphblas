// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

type BinaryOp interface {
	operator()
	binaryOp()
}

type BinaryOpFloat64 interface {
	Operator(in1, in2 float64) float64
	binaryOp()
	operator()
}

type binaryOpFloat64 struct {
	op func(float64, float64) float64
}

func (s *binaryOpFloat64) binaryOp() {}
func (s *binaryOpFloat64) operator() {}

func (s *binaryOpFloat64) Operator(in1, in2 float64) float64 {
	return s.op(in1, in2)
}

type BinaryOpBool interface {
	Operator(in1, in2 bool) bool
	binaryOp()
	operator()
}

type binaryOpBool struct {
	op func(bool, bool) bool
}

func (s *binaryOpBool) binaryOp() {}
func (s *binaryOpBool) operator() {}

func (s *binaryOpBool) Operator(in1, in2 bool) bool {
	return s.op(in1, in2)
}

type BinaryOpFloat64ToBool interface {
	Operator(in1, in2 float64) bool
	binaryOp()
	operator()
}

type binaryOpFloat64ToBool struct {
	op func(float64, float64) bool
}

func (s *binaryOpFloat64ToBool) binaryOp() {}
func (s *binaryOpFloat64ToBool) operator() {}

func (s *binaryOpFloat64ToBool) Operator(in1, in2 float64) bool {
	return s.op(in1, in2)
}

// LOR logical OR f(x, y) = x ∨ y
var LOR = func() BinaryOp {
	return &binaryOpBool{op: func(in1, in2 bool) bool {
		return in1 || in2
	}}
}

// LAND logical AND f(x, y) = x ∧ y
var LAND = func() BinaryOp {
	return &binaryOpBool{op: func(in1, in2 bool) bool {
		return in1 && in2
	}}
}

// LXOR logical XOR f(x, y) = x ⊕ y
var LXOR = func() BinaryOp {
	return &binaryOpBool{op: func(in1, in2 bool) bool {
		return in1 != in2
	}}
}

// Equal f(x, y) = (x == y)
var Equal = func() BinaryOp {
	return &binaryOpFloat64ToBool{op: func(in1, in2 float64) bool {
		return in1 == in2
	}}
}

// NotEqual f(x, y) = (x != y)
var NotEqual = func() BinaryOp {
	return &binaryOpFloat64ToBool{op: func(in1, in2 float64) bool {
		return in1 != in2
	}}
}

// GreaterThan f(x, y) = (x > y)
var GreaterThan = func() BinaryOp {
	return &binaryOpFloat64ToBool{op: func(in1, in2 float64) bool {
		return in1 > in2
	}}
}

// LessThan f(x, y) = (x < y)
var LessThan = func() BinaryOp {
	return &binaryOpFloat64ToBool{op: func(in1, in2 float64) bool {
		return in1 < in2
	}}
}

// GreaterThanOrEqual f(x, y) = (x >= y)
var GreaterThanOrEqual = func() BinaryOp {
	return &binaryOpFloat64ToBool{op: func(in1, in2 float64) bool {
		return in1 >= in2
	}}
}

// LessThanOrEqual f(x, y) = (x <= y)
var LessThanOrEqual = func() BinaryOp {
	return &binaryOpFloat64ToBool{op: func(in1, in2 float64) bool {
		return in1 <= in2
	}}
}

// FirstArgument f(x, y) = x
var FirstArgument = func() BinaryOp {
	return &binaryOpFloat64{op: func(in1, in2 float64) float64 {
		return in1
	}}
}

// SecondArgument f(x, y) = y
var SecondArgument = func() BinaryOp {
	return &binaryOpFloat64{op: func(in1, in2 float64) float64 {
		return in2
	}}
}

// Minimum f(x, y) = (x < y) ? x : y
var Minimum = func() BinaryOp {
	return &binaryOpFloat64{op: func(in1, in2 float64) float64 {
		if in1 < in2 {
			return in1
		}

		return in2
	}}
}

// Maximum f(x, y) = (x > y) ? x : y
var Maximum = func() BinaryOp {
	return &binaryOpFloat64{op: func(in1, in2 float64) float64 {
		if in1 > in2 {
			return in1
		}

		return in2
	}}
}

// Addition f(x, y) = x + y
var Addition = func() BinaryOp {
	return &binaryOpFloat64{op: func(in1, in2 float64) float64 {
		return in1 + in2
	}}
}

// Subtraction f(x, y) = x - y
var Subtraction = func() BinaryOp {
	return &binaryOpFloat64{op: func(in1, in2 float64) float64 {
		return in1 - in2
	}}
}

// Multiplication f(x, y) = x * y
var Multiplication = func() BinaryOp {
	return &binaryOpFloat64{op: func(in1, in2 float64) float64 {
		return in1 * in2
	}}
}

// Division f(x, y) = x / y
var Division = func() BinaryOp {
	return &binaryOpFloat64{op: func(in1, in2 float64) float64 {
		return in1 / in2
	}}
}
