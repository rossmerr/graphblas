// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas_test

import (
	"testing"

	"github.com/rossmerr/graphblas"
)

func TestSparseMatrixRegistry_IsSparseMatrix(t *testing.T) {

	tests := []struct {
		name     string
		s        graphblas.MatrixLogical[float64]
		isSparse bool
	}{
		{
			name:     "DenseMatrix",
			s:        graphblas.NewDenseMatrix[float64](2, 2),
			isSparse: false,
		},
		{
			name:     "CSCMatrix",
			s:        graphblas.NewCSCMatrix[float64](2, 2),
			isSparse: true,
		},
		{
			name:     "CSRMatrix",
			s:        graphblas.NewCSRMatrix[float64](2, 2),
			isSparse: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := graphblas.IsSparseMatrix(tt.s)
			if tt.isSparse != v {
				t.Errorf("%+v IsSparseMatrix = %+v, want %+v", tt.name, v, tt.isSparse)
			}
		})
	}
}
