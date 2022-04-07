package operators_test

import (
	"testing"

	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/operators"
)

func Test_Addition(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[float32]
		in1    float32
		in2    float32
		result float32
	}{
		{
			name:   "1",
			s:      operators.Addition[float32](),
			in1:    2,
			in2:    2,
			result: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator want %+v", tt.name, tt.result)
			}
		})
	}
}

func Test_Subtraction(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[float32]
		in1    float32
		in2    float32
		result float32
	}{
		{
			name:   "1",
			s:      operators.Subtraction[float32](),
			in1:    4,
			in2:    3,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator want %+v", tt.name, tt.result)
			}
		})
	}
}

func Test_Multiplication(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[float32]
		in1    float32
		in2    float32
		result float32
	}{
		{
			name:   "1",
			s:      operators.Multiplication[float32](),
			in1:    2,
			in2:    3,
			result: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator want %+v", tt.name, tt.result)
			}
		})
	}
}

func Test_Division(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[float32]
		in1    float32
		in2    float32
		result float32
	}{
		{
			name:   "1",
			s:      operators.Division[float32](),
			in1:    6,
			in2:    2,
			result: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator want %+v", tt.name, tt.result)
			}
		})
	}
}
