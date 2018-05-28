// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64Op

import "github.com/RossMerr/Caudex.GraphBLAS/unaryOp"

// UnaryOpFloat64 is a function that maps one input value to one output value
type UnaryOpFloat64 interface {
	unaryOp.UnaryOp
	Apply(float64) float64
}

type unaryOpFloat64 struct {
	apply func(float64) float64
}

func (s *unaryOpFloat64) Operator() {}
func (s *unaryOpFloat64) BinaryOp() {}

func (s *unaryOpFloat64) Apply(in float64) float64 {
	return s.apply(in)
}

// Identity f(x) = x
var Identity = &unaryOpFloat64{apply: func(in float64) float64 {
	return in
}}

// AdditiveInverse additive inverse f(x) = -x
var AdditiveInverse = &unaryOpFloat64{apply: func(in float64) float64 {
	return -in
}}

// MultiplicativeInverse multiplicative inverse f(x) = 1/x
var MultiplicativeInverse = &unaryOpFloat64{apply: func(in float64) float64 {
	return 1 / in
}}
