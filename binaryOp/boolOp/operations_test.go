// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolOp_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/binaryOp"
	"github.com/RossMerr/Caudex.GraphBLAS/binaryOp/boolOp"
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
			s:      boolOp.LOR,
			in1:    true,
			in2:    true,
			result: true,
		},
		{
			name:   "2",
			s:      boolOp.LOR,
			in1:    false,
			in2:    true,
			result: true,
		},
		{
			name:   "3",
			s:      boolOp.LOR,
			in1:    true,
			in2:    false,
			result: true,
		},
		{
			name:   "4",
			s:      boolOp.LOR,
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(boolOp.BinaryOpBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
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
			s:      boolOp.LAND,
			in1:    true,
			in2:    true,
			result: true,
		},
		{
			name:   "2",
			s:      boolOp.LAND,
			in1:    false,
			in2:    true,
			result: false,
		},
		{
			name:   "3",
			s:      boolOp.LAND,
			in1:    true,
			in2:    false,
			result: false,
		},
		{
			name:   "4",
			s:      boolOp.LAND,
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(boolOp.BinaryOpBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
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
			s:      boolOp.LXOR,
			in1:    true,
			in2:    true,
			result: false,
		},
		{
			name:   "2",
			s:      boolOp.LXOR,
			in1:    false,
			in2:    true,
			result: true,
		},
		{
			name:   "3",
			s:      boolOp.LXOR,
			in1:    true,
			in2:    false,
			result: true,
		},
		{
			name:   "4",
			s:      boolOp.LXOR,
			in1:    false,
			in2:    false,
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(boolOp.BinaryOpBool); ok {
				if tt.result != op.Apply(tt.in1, tt.in2) {
					t.Errorf("%+v Operator = %+v, want %+v", tt.name, !tt.result, tt.result)
				}
			} else {
				t.Errorf("%+v not a BinaryOpBool", tt.name)
			}
		})
	}
}
