// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64Op_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/unaryOp/float64Op"
)

func Test_Identity(t *testing.T) {
	tests := []struct {
		name   string
		s      float64Op.UnaryOpFloat64
		in     float64
		result float64
	}{
		{
			name:   "1",
			s:      float64Op.Identity,
			in:     2,
			result: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64Op.UnaryOpFloat64); ok {
				result := op.Apply(tt.in)
				if tt.result != result {
					t.Errorf("%+v Identity = %+v, want %+v", tt.name, result, tt.result)
				}
			} else {
				t.Errorf("%+v not a UnaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_AdditiveInverse(t *testing.T) {
	tests := []struct {
		name   string
		s      float64Op.UnaryOpFloat64
		in     float64
		result float64
	}{
		{
			name:   "1",
			s:      float64Op.AdditiveInverse,
			in:     2,
			result: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64Op.UnaryOpFloat64); ok {
				result := op.Apply(tt.in)
				if tt.result != result {
					t.Errorf("%+v AdditiveInverse = %+v, want %+v", tt.name, result, tt.result)
				}
			} else {
				t.Errorf("%+v not a UnaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_MultiplicativeInverse(t *testing.T) {
	tests := []struct {
		name   string
		s      float64Op.UnaryOpFloat64
		in     float64
		result float64
	}{
		{
			name:   "1",
			s:      float64Op.MultiplicativeInverse,
			in:     2,
			result: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64Op.UnaryOpFloat64); ok {
				result := op.Apply(tt.in)
				if tt.result != result {
					t.Errorf("%+v MultiplicativeInverse = %+v, want %+v", tt.name, result, tt.result)
				}
			} else {
				t.Errorf("%+v not a UnaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_Operator(t *testing.T) {
	float64Op.Identity.Operator()
}

func Test_BinaryOp(t *testing.T) {
	float64Op.Identity.UnaryOp()
}
