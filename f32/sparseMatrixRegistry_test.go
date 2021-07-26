// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package f32_test

import (
	"testing"

	"github.com/rossmerr/graphblas/f32"
)

func TestSparseMatrixRegistry_IsSparseMatrix(t *testing.T) {

	tests := []struct {
		name     string
		s        f32.Matrix
		isSparse bool
	}{
		{
			name:     "DenseMatrix",
			s:        f32.NewDenseMatrix(2, 2),
			isSparse: false,
		},
		{
			name:     "CSCMatrix",
			s:        f32.NewCSCMatrix(2, 2),
			isSparse: true,
		},
		{
			name:     "CSRMatrix",
			s:        f32.NewCSRMatrix(2, 2),
			isSparse: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := f32.IsSparseMatrix(tt.s)
			if tt.isSparse != v {
				t.Errorf("%+v IsSparseMatrix = %+v, want %+v", tt.name, v, tt.isSparse)
			}
		})
	}
}
