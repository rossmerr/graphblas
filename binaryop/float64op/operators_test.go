// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64op_test

import (
	"testing"

	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/binaryop/float64op"
)

func Test_Equal(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      float64op.Equal,
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      float64op.Equal,
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      float64op.Equal,
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64ToBool", tt.name)
			}
		})
	}
}

func Test_NotEqual(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      float64op.NotEqual,
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      float64op.NotEqual,
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      float64op.NotEqual,
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64ToBool", tt.name)
			}
		})
	}
}

func Test_GreaterThan(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      float64op.GreaterThan,
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      float64op.GreaterThan,
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      float64op.GreaterThan,
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64ToBool", tt.name)
			}
		})
	}
}

func Test_LessThan(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      float64op.LessThan,
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      float64op.LessThan,
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      float64op.LessThan,
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64ToBool", tt.name)
			}
		})
	}
}

func Test_GreaterThanOrEqual(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      float64op.GreaterThanOrEqual,
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      float64op.GreaterThanOrEqual,
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      float64op.GreaterThanOrEqual,
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64ToBool", tt.name)
			}
		})
	}
}

func Test_LessThanOrEqual(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      float64op.LessThanOrEqual,
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      float64op.LessThanOrEqual,
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      float64op.LessThanOrEqual,
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64ToBool", tt.name)
			}
		})
	}
}

func Test_FirstArgument(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      float64op.FirstArgument,
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      float64op.FirstArgument,
			in1:    1,
			in2:    2,
			result: 1,
		},
		{
			name:   "3",
			s:      float64op.FirstArgument,
			in1:    2,
			in2:    1,
			result: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.in1)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_SecondArgument(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      float64op.SecondArgument,
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      float64op.SecondArgument,
			in1:    1,
			in2:    2,
			result: 2,
		},
		{
			name:   "3",
			s:      float64op.SecondArgument,
			in1:    2,
			in2:    1,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.in2)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_Minimum(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      float64op.Minimum,
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      float64op.Minimum,
			in1:    1,
			in2:    2,
			result: 1,
		},
		{
			name:   "3",
			s:      float64op.Minimum,
			in1:    2,
			in2:    1,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_Maximum(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      float64op.Maximum,
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      float64op.Maximum,
			in1:    1,
			in2:    2,
			result: 2,
		},
		{
			name:   "3",
			s:      float64op.Maximum,
			in1:    2,
			in2:    1,
			result: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_Addition(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      float64op.Addition,
			in1:    2,
			in2:    2,
			result: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_Subtraction(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      float64op.Subtraction,
			in1:    4,
			in2:    3,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_Multiplication(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      float64op.Multiplication,
			in1:    2,
			in2:    3,
			result: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_Division(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryop.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      float64op.Division,
			in1:    6,
			in2:    2,
			result: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(float64op.BinaryOpFloat64); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}

func Test_Operator(t *testing.T) {
	float64op.Equal.Operator()
	float64op.FirstArgument.Operator()
}

func Test_BinaryOp(t *testing.T) {
	float64op.Equal.BinaryOp()
	float64op.FirstArgument.BinaryOp()
}

func Test_Semigroup(t *testing.T) {
	float64op.FirstArgument.Semigroup()
}
