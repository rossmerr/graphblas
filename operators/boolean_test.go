// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package operators_test

import (
	"testing"

	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/operators"
	"github.com/rossmerr/graphblas/unaryop"
)

func Test_LOR(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[bool]
		in1    bool
		in2    bool
		result bool
	}{
		{
			name:   "1",
			s:      operators.LOR,
			in1:    true,
			in2:    true,
			result: true,
		},
		{
			name:   "2",
			s:      operators.LOR,
			in1:    false,
			in2:    true,
			result: true,
		},
		{
			name:   "3",
			s:      operators.LOR,
			in1:    true,
			in2:    false,
			result: true,
		},
		{
			name:   "4",
			s:      operators.LOR,
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v LOR = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_LAND(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[bool]
		in1    bool
		in2    bool
		result bool
	}{
		{
			name:   "1",
			s:      operators.LAND,
			in1:    true,
			in2:    true,
			result: true,
		},
		{
			name:   "2",
			s:      operators.LAND,
			in1:    false,
			in2:    true,
			result: false,
		},
		{
			name:   "3",
			s:      operators.LAND,
			in1:    true,
			in2:    false,
			result: false,
		},
		{
			name:   "4",
			s:      operators.LAND,
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v LAND = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_LXOR(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[bool]
		in1    bool
		in2    bool
		result bool
	}{
		{
			name:   "1",
			s:      operators.LXOR,
			in1:    true,
			in2:    true,
			result: false,
		},
		{
			name:   "2",
			s:      operators.LXOR,
			in1:    false,
			in2:    true,
			result: true,
		},
		{
			name:   "3",
			s:      operators.LXOR,
			in1:    true,
			in2:    false,
			result: true,
		},
		{
			name:   "4",
			s:      operators.LXOR,
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v LXOR = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_Associative(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[bool]
		a      bool
		b      bool
		c      bool
		result bool
	}{
		{
			name:   "1",
			s:      operators.LOR,
			a:      true,
			b:      true,
			c:      true,
			result: true,
		},
		{
			name:   "2",
			s:      operators.LOR,
			a:      false,
			b:      true,
			c:      true,
			result: true,
		},
		{
			name:   "3",
			s:      operators.LOR,
			a:      true,
			b:      false,
			c:      true,
			result: true,
		},
		{
			name:   "4",
			s:      operators.LOR,
			a:      true,
			b:      true,
			c:      false,
			result: true,
		},
		{
			name:   "5",
			s:      operators.LOR,
			a:      false,
			b:      false,
			c:      false,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.s.Apply(tt.s.Apply(tt.a, tt.b), tt.c) == tt.s.Apply(tt.a, tt.s.Apply(tt.b, tt.c))
			if tt.result != result {
				t.Errorf("%+v Associative = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_Commutative(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[bool]
		a      bool
		b      bool
		c      bool
		result bool
	}{
		{
			name:   "1",
			s:      operators.LOR,
			a:      true,
			b:      true,
			result: false,
		},
		{
			name:   "2",
			s:      operators.LOR,
			a:      false,
			b:      true,
			result: false,
		},
		{
			name:   "3",
			s:      operators.LOR,
			a:      true,
			b:      false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.s.Apply(tt.a, tt.b) != tt.s.Apply(tt.b, tt.a)
			if tt.result != result {
				t.Errorf("%+v Commutative = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_LogicalInverse(t *testing.T) {
	tests := []struct {
		name   string
		s      unaryop.UnaryOp[bool]
		in     bool
		result bool
	}{
		{
			name:   "1",
			s:      operators.LogicalInverse,
			in:     true,
			result: false,
		},
		{
			name:   "2",
			s:      operators.LogicalInverse,
			in:     false,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.s.Apply(tt.in)
			if tt.result != result {
				t.Errorf("%+v LogicalInverse = %+v, want %+v", tt.name, result, tt.result)
			}
		})
	}
}
