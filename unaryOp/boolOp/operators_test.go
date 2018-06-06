// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolOp_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/unaryOp/boolOp"
)

func Test_LogicalInverse(t *testing.T) {
	tests := []struct {
		name   string
		s      boolOp.UnaryOpBool
		in     bool
		result bool
	}{
		{
			name:   "1",
			s:      boolOp.LogicalInverse,
			in:     true,
			result: false,
		},
		{
			name:   "2",
			s:      boolOp.LogicalInverse,
			in:     false,
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if op, ok := tt.s.(boolOp.UnaryOpBool); ok {
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
