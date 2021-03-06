// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolop

import "github.com/rossmerr/graphblas/unaryop"

// UnaryOpBool is a function that maps one input value to one output value
type UnaryOpBool interface {
	unaryop.UnaryOp
	Apply(bool) bool
}

type unaryOpBool struct {
	apply func(bool) bool
}

func (s *unaryOpBool) Operator() {}
func (s *unaryOpBool) UnaryOp()  {}

func (s *unaryOpBool) Apply(in bool) bool {
	return s.apply(in)
}

// LogicalInverse logical inverse f(x) = ¬x
var LogicalInverse = &unaryOpBool{apply: func(in bool) bool {
	return !in
}}
