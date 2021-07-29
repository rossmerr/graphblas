// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolean_test

import (
	"testing"

	"github.com/rossmerr/graphblas/boolean"
)

func TestSparseMatrixRegistry_IsSparseMatrix(t *testing.T) {

	tests := []struct {
		name     string
		s        boolean.Matrix
		isSparse bool
	}{
		{
			name:     "DenseMatrix",
			s:        boolean.NewDenseMatrix(2, 2),
			isSparse: false,
		},
		{
			name:     "CSCMatrix",
			s:        boolean.NewCSCMatrix(2, 2),
			isSparse: true,
		},
		{
			name:     "CSRMatrix",
			s:        boolean.NewCSRMatrix(2, 2),
			isSparse: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := boolean.IsSparseMatrix(tt.s)
			if tt.isSparse != v {
				t.Errorf("%+v IsSparseMatrix = %+v, want %+v", tt.name, v, tt.isSparse)
			}
		})
	}
}
