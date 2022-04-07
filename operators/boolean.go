package operators

// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import (
	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/unaryop"
)

// LOR logical OR f(x, y) = x ∨ y
var LOR = binaryop.NewBinaryOp(func(in1, in2 bool) bool {
	return in1 || in2
})

// LAND logical AND f(x, y) = x ∧ y
var LAND = binaryop.NewBinaryOp(func(in1, in2 bool) bool {
	return in1 && in2
})

// LXOR logical XOR f(x, y) = x ⊕ y
var LXOR = binaryop.NewBinaryOp(func(in1, in2 bool) bool {
	return in1 != in2
})

// LogicalInverse logical inverse f(x) = ¬x
var LogicalInverse = unaryop.NewUnaryOp(func(in bool) bool {
	return !in
})
