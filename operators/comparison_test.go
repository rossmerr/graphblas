package operators_test

// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import (
	"testing"

	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/operators"
)

func Test_Equal(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOpToBool[float32]
		in1    float32
		in2    float32
		result bool
	}{
		{
			name:   "1",
			s:      operators.Equal[float32](),
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      operators.Equal[float32](),
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      operators.Equal[float32](),
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_NotEqual(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOpToBool[float32]
		in1    float32
		in2    float32
		result bool
	}{
		{
			name:   "1",
			s:      operators.NotEqual[float32](),
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      operators.NotEqual[float32](),
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      operators.NotEqual[float32](),
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_GreaterThan(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOpToBool[float32]
		in1    float32
		in2    float32
		result bool
	}{
		{
			name:   "1",
			s:      operators.GreaterThan[float32](),
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      operators.GreaterThan[float32](),
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      operators.GreaterThan[float32](),
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_LessThan(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOpToBool[float32]
		in1    float32
		in2    float32
		result bool
	}{
		{
			name:   "1",
			s:      operators.LessThan[float32](),
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      operators.LessThan[float32](),
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      operators.LessThan[float32](),
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_GreaterThanOrEqual(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOpToBool[float32]
		in1    float32
		in2    float32
		result bool
	}{
		{
			name:   "1",
			s:      operators.GreaterThanOrEqual[float32](),
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      operators.GreaterThanOrEqual[float32](),
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      operators.GreaterThanOrEqual[float32](),
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}

func Test_LessThanOrEqual(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOpToBool[float32]
		in1    float32
		in2    float32
		result bool
	}{
		{
			name:   "1",
			s:      operators.LessThanOrEqual[float32](),
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      operators.LessThanOrEqual[float32](),
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      operators.LessThanOrEqual[float32](),
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
			}
		})
	}
}
