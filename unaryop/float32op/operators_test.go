// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float32op_test

import (
	"testing"

	"github.com/rossmerr/graphblas/unaryop/float32op"
)

func Test_Identity(t *testing.T) {
	tests := []struct {
		name   string
		s      float32op.UnaryOpFloat32
		in     float32
		result float32
	}{
		{
			name:   "1",
			s:      float32op.Identity,
			in:     2,
			result: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float32op.UnaryOpFloat32); ok {
				result := op.Apply(tt.in)
				if tt.result != result {
					t.Errorf("%+v Identity = %+v, want %+v", tt.name, result, tt.result)
				}
			} else {
				t.Errorf("%+v not a UnaryOpFloat32", tt.name)
			}
		})
	}
}

func Test_AdditiveInverse(t *testing.T) {
	tests := []struct {
		name   string
		s      float32op.UnaryOpFloat32
		in     float32
		result float32
	}{
		{
			name:   "1",
			s:      float32op.AdditiveInverse,
			in:     2,
			result: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float32op.UnaryOpFloat32); ok {
				result := op.Apply(tt.in)
				if tt.result != result {
					t.Errorf("%+v AdditiveInverse = %+v, want %+v", tt.name, result, tt.result)
				}
			} else {
				t.Errorf("%+v not a UnaryOpFloat32", tt.name)
			}
		})
	}
}

func Test_MultiplicativeInverse(t *testing.T) {
	tests := []struct {
		name   string
		s      float32op.UnaryOpFloat32
		in     float32
		result float32
	}{
		{
			name:   "1",
			s:      float32op.MultiplicativeInverse,
			in:     2,
			result: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float32op.UnaryOpFloat32); ok {
				result := op.Apply(tt.in)
				if tt.result != result {
					t.Errorf("%+v MultiplicativeInverse = %+v, want %+v", tt.name, result, tt.result)
				}
			} else {
				t.Errorf("%+v not a UnaryOpFloat32", tt.name)
			}
		})
	}
}

func Test_Operator(t *testing.T) {
	float32op.Identity.Operator()
}

func Test_BinaryOp(t *testing.T) {
	float32op.Identity.UnaryOp()
}
