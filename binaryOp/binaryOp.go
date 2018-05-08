// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

// BinaryOpFloat64 is a function that maps two input values to one output value
type BinaryOpFloat64 func(in1, in2 float64) float64

// BinaryOpFloat64ToBool is a function that maps two input values to one output value
type BinaryOpFloat64ToBool func(in1, in2 float64) bool

// BinaryOpBool is a function that maps two input values to one output value
type BinaryOpBool func(in1, in2 bool) bool

// LOR logical OR f(x, y) = x ∨ y
var LOR = func(in1, in2 bool) bool {
	return in1 || in2
}

// LAND logical AND f(x, y) = x ∧ y
var LAND = func(in1, in2 bool) bool {
	return in1 && in2
}

// LXOR logical XOR f(x, y) = x ⊕ y
var LXOR = func(in1, in2 bool) bool {
	return in1 != in2
}

// Equal f(x, y) = (x == y)
var Equal = func(in1, in2 float64) bool {
	return in1 == in2
}

// NotEqual f(x, y) = (x != y)
var NotEqual = func(in1, in2 float64) bool {
	return in1 != in2
}

// GreaterThan f(x, y) = (x > y)
var GreaterThan = func(in1, in2 float64) bool {
	return in1 > in2
}

// LessThan f(x, y) = (x < y)
var LessThan = func(in1, in2 float64) bool {
	return in1 < in2
}

// GreaterThanOrEqual f(x, y) = (x >= y)
var GreaterThanOrEqual = func(in1, in2 float64) bool {
	return in1 >= in2
}

// LessThanOrEqual f(x, y) = (x <= y)
var LessThanOrEqual = func(in1, in2 float64) bool {
	return in1 <= in2
}

// FirstArgument f(x, y) = x
var FirstArgument = func(in1, in2 float64) float64 {
	return in1
}

// SecondArgument f(x, y) = y
var SecondArgument = func(in1, in2 float64) float64 {
	return in2
}

// Minimum f(x, y) = (x < y) ? x : y
var Minimum = func(in1, in2 float64) float64 {
	if in1 < in2 {
		return in1
	}

	return in2
}

// Maximum f(x, y) = (x > y) ? x : y
var Maximum = func(in1, in2 float64) float64 {
	if in1 > in2 {
		return in1
	}

	return in2
}

// Addition f(x, y) = x + y
var Addition = func(in1, in2 float64) float64 {
	return in1 + in2
}

// Subtraction f(x, y) = x - y
var Subtraction = func(in1, in2 float64) float64 {
	return in1 - in2
}

// Multiplication f(x, y) = x * y
var Multiplication = func(in1, in2 float64) float64 {
	return in1 * in2
}

// Division f(x, y) = x / y
var Division = func(in1, in2 float64) float64 {
	return in1 / in2
}
