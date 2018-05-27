// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestSparseMatrixRegistry_IsSparseMatrix(t *testing.T) {

	tests := []struct {
		name     string
		s        GraphBLAS.Matrix
		isSparse bool
	}{
		{
			name:     "DenseMatrix",
			s:        GraphBLAS.NewDenseMatrix(2, 2),
			isSparse: false,
		},
		{
			name:     "CSCMatrix",
			s:        GraphBLAS.NewCSCMatrix(2, 2),
			isSparse: true,
		},
		{
			name:     "CSRMatrix",
			s:        GraphBLAS.NewCSRMatrix(2, 2),
			isSparse: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := GraphBLAS.IsSparseMatrix(tt.s)
			if tt.isSparse != v {
				t.Errorf("%+v IsSparseMatrix = %+v, want %+v", tt.name, v, tt.isSparse)
			}
		})
	}
}
