// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package unaryOp

// UnaryOperatorFloat64 is a function that maps one input value to one output value
type UnaryOperatorFloat64 func(in float64) float64

// UnaryOperatorBool is a function that maps one input value to one output value
type UnaryOperatorBool func(in bool) bool

// Identity f(x) = x
var Identity = func(in float64) float64 {
	return in
}

// AdditiveInverse additive inverse f(x) = -x
var AdditiveInverse = func(in float64) float64 {
	return -in
}

// MultiplicativeInverse multiplicative inverse f(x) = 1/x
var MultiplicativeInverse = func(in float64) float64 {
	return 1 / in
}

// LogicalInverse logical inverse f(x) = ¬x
var LogicalInverse = func(in bool) bool {
	return !in
}
