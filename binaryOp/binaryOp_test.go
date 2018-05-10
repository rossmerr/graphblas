// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/binaryOp"
)

func Test_LOR(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryOp.BinaryOp
		in1    bool
		in2    bool
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.LOR(),
			in1:    true,
			in2:    true,
			result: true,
		},
		{
			name:   "2",
			s:      binaryOp.LOR(),
			in1:    false,
			in2:    true,
			result: true,
		},
		{
			name:   "3",
			s:      binaryOp.LOR(),
			in1:    true,
			in2:    false,
			result: true,
		},
		{
			name:   "4",
			s:      binaryOp.LOR(),
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpBool", tt.name)
			}
		})
	}
}

func Test_LAND(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryOp.BinaryOp
		in1    bool
		in2    bool
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.LAND(),
			in1:    true,
			in2:    true,
			result: true,
		},
		{
			name:   "2",
			s:      binaryOp.LAND(),
			in1:    false,
			in2:    true,
			result: false,
		},
		{
			name:   "3",
			s:      binaryOp.LAND(),
			in1:    true,
			in2:    false,
			result: false,
		},
		{
			name:   "4",
			s:      binaryOp.LAND(),
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpBool", tt.name)
			}
		})
	}
}

func Test_LXOR(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryOp.BinaryOp
		in1    bool
		in2    bool
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.LXOR(),
			in1:    true,
			in2:    true,
			result: false,
		},
		{
			name:   "2",
			s:      binaryOp.LXOR(),
			in1:    false,
			in2:    true,
			result: true,
		},
		{
			name:   "3",
			s:      binaryOp.LXOR(),
			in1:    true,
			in2:    false,
			result: true,
		},
		{
			name:   "4",
			s:      binaryOp.LXOR(),
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpBool", tt.name)
			}
		})
	}
}

func Test_Equal(t *testing.T) {
	tests := []struct {
		name   string
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.Equal(),
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      binaryOp.Equal(),
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      binaryOp.Equal(),
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.NotEqual(),
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      binaryOp.NotEqual(),
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      binaryOp.NotEqual(),
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.GreaterThan(),
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      binaryOp.GreaterThan(),
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      binaryOp.GreaterThan(),
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.LessThan(),
			in1:    1,
			in2:    1,
			result: false,
		},
		{
			name:   "2",
			s:      binaryOp.LessThan(),
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      binaryOp.LessThan(),
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.GreaterThanOrEqual(),
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      binaryOp.GreaterThanOrEqual(),
			in1:    1,
			in2:    2,
			result: false,
		},
		{
			name:   "3",
			s:      binaryOp.GreaterThanOrEqual(),
			in1:    2,
			in2:    1,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result bool
	}{
		{
			name:   "1",
			s:      binaryOp.LessThanOrEqual(),
			in1:    1,
			in2:    1,
			result: true,
		},
		{
			name:   "2",
			s:      binaryOp.LessThanOrEqual(),
			in1:    1,
			in2:    2,
			result: true,
		},
		{
			name:   "3",
			s:      binaryOp.LessThanOrEqual(),
			in1:    2,
			in2:    1,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64ToBool); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      binaryOp.FirstArgument(),
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      binaryOp.FirstArgument(),
			in1:    1,
			in2:    2,
			result: 1,
		},
		{
			name:   "3",
			s:      binaryOp.FirstArgument(),
			in1:    2,
			in2:    1,
			result: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      binaryOp.SecondArgument(),
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      binaryOp.SecondArgument(),
			in1:    1,
			in2:    2,
			result: 2,
		},
		{
			name:   "3",
			s:      binaryOp.SecondArgument(),
			in1:    2,
			in2:    1,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      binaryOp.Minimum(),
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      binaryOp.Minimum(),
			in1:    1,
			in2:    2,
			result: 1,
		},
		{
			name:   "3",
			s:      binaryOp.Minimum(),
			in1:    2,
			in2:    1,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      binaryOp.Maximum(),
			in1:    1,
			in2:    1,
			result: 1,
		},
		{
			name:   "2",
			s:      binaryOp.Maximum(),
			in1:    1,
			in2:    2,
			result: 2,
		},
		{
			name:   "3",
			s:      binaryOp.Maximum(),
			in1:    2,
			in2:    1,
			result: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      binaryOp.Addition(),
			in1:    2,
			in2:    2,
			result: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      binaryOp.Subtraction(),
			in1:    4,
			in2:    3,
			result: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      binaryOp.Multiplication(),
			in1:    2,
			in2:    3,
			result: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
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
		s      binaryOp.BinaryOp
		in1    float64
		in2    float64
		result float64
	}{
		{
			name:   "1",
			s:      binaryOp.Division(),
			in1:    6,
			in2:    2,
			result: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(binaryOp.BinaryOpFloat64); ok {
				if tt.result != op.Operator(tt.in1, tt.in2) {
					t.Errorf("%+v Operator want %+v", tt.name, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpFloat64", tt.name)
			}
		})
	}
}
