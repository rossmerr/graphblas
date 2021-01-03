// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float32op

import "github.com/rossmerr/graphblas/unaryop"

// UnaryOpFloat32 is a function that maps one input value to one output value
type UnaryOpFloat32 interface {
	unaryop.UnaryOp
	Apply(float32) float32
}

type unaryOpFloat32 struct {
	apply func(float32) float32
}

func (s *unaryOpFloat32) Operator() {}
func (s *unaryOpFloat32) UnaryOp()  {}

func (s *unaryOpFloat32) Apply(in float32) float32 {
	return s.apply(in)
}

// Identity f(x) = x
var Identity = &unaryOpFloat32{apply: func(in float32) float32 {
	return in
}}

// AdditiveInverse additive inverse f(x) = -x
var AdditiveInverse = &unaryOpFloat32{apply: func(in float32) float32 {
	return -in
}}

// MultiplicativeInverse multiplicative inverse f(x) = 1/x
var MultiplicativeInverse = &unaryOpFloat32{apply: func(in float32) float32 {
	return 1 / in
}}
