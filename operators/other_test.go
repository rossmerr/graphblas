package operators_test

import (
	"testing"

	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/operators"
)

func Test_FirstArgument(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[float32]
		in1    float32
		in2    float32
		result float32
	}{
		{
			name:   "1",
			s:      operators.FirstArgument[float32](),
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      operators.FirstArgument[float32](),
			in1:    1,
			in2:    2,
			result: 1,
		},
		{
			name:   "3",
			s:      operators.FirstArgument[float32](),
			in1:    2,
			in2:    1,
			result: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator want %+v", tt.name, tt.in1)
			}

		})
	}
}

func Test_SecondArgument(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[float32]
		in1    float32
		in2    float32
		result float32
	}{
		{
			name:   "1",
			s:      operators.SecondArgument[float32](),
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      operators.SecondArgument[float32](),
			in1:    1,
			in2:    2,
			result: 2,
		},
		{
			name:   "3",
			s:      operators.SecondArgument[float32](),
			in1:    2,
			in2:    1,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.result != tt.s.Apply(tt.in1, tt.in2) {
				t.Errorf("%+v Operator want %+v", tt.name, tt.in2)
			}
		})
	}
}

func Test_Minimum(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[float32]
		in1    float32
		in2    float32
		result float32
	}{
		{
			name:   "1",
			s:      operators.Minimum[float32](),
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      operators.Minimum[float32](),
			in1:    1,
			in2:    2,
			result: 1,
		},
		{
			name:   "3",
			s:      operators.Minimum[float32](),
			in1:    2,
			in2:    1,
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

func Test_Maximum(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp[float32]
		in1    float32
		in2    float32
		result float32
	}{
		{
			name:   "1",
			s:      operators.Maximum[float32](),
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      operators.Maximum[float32](),
			in1:    1,
			in2:    2,
			result: 2,
		},
		{
			name:   "3",
			s:      operators.Maximum[float32](),
			in1:    2,
			in2:    1,
			result: 2,
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
