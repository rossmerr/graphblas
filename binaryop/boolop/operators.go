// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolop

import "github.com/rossmerr/graphblas/binaryop"

// BinaryOpBool is a function that maps two input value to one output value
type BinaryOpBool interface {
	binaryop.Semigroup
	Apply(in1, in2 bool) bool
}

type binaryOpBool struct {
	apply func(bool, bool) bool
}

func (s *binaryOpBool) Operator()  {}
func (s *binaryOpBool) BinaryOp()  {}
func (s *binaryOpBool) Semigroup() {}

func (s *binaryOpBool) Apply(in1, in2 bool) bool {
	return s.apply(in1, in2)
}

// LOR logical OR f(x, y) = x ∨ y
var LOR = &binaryOpBool{apply: func(in1, in2 bool) bool {
	return in1 || in2
}}

// LAND logical AND f(x, y) = x ∧ y
var LAND = &binaryOpBool{apply: func(in1, in2 bool) bool {
	return in1 && in2
}}

// LXOR logical XOR f(x, y) = x ⊕ y
var LXOR = &binaryOpBool{apply: func(in1, in2 bool) bool {
	return in1 != in2
}}
