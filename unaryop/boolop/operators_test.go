// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolop_test

import (
	"testing"

	"github.com/rossmerr/graphblas/unaryop/boolop"
)

func Test_LogicalInverse(t *testing.T) {
	tests := []struct {
		name   string
		s      boolop.UnaryOpBool
		in     bool
		result bool
	}{
		{
			name:   "1",
			s:      boolop.LogicalInverse,
			in:     true,
			result: false,
		},
		{
			name:   "2",
			s:      boolop.LogicalInverse,
			in:     false,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(boolop.UnaryOpBool); ok {
				result := op.Apply(tt.in)
				if tt.result != result {
					t.Errorf("%+v LogicalInverse = %+v, want %+v", tt.name, result, tt.result)
				}
			} else {
				t.Errorf("%+v not a UnaryOpBool", tt.name)
			}
		})
	}
}

func Test_Operator(t *testing.T) {
	boolop.LogicalInverse.Operator()
}

func Test_BinaryOp(t *testing.T) {
	boolop.LogicalInverse.UnaryOp()
}
