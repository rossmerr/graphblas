// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package doublePrecision_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/doublePrecision"
)

func TestSparseMatrixRegistry_IsSparseMatrix(t *testing.T) {

	tests := []struct {
		name     string
		s        doublePrecision.Matrix
		isSparse bool
	}{
		{
			name:     "DenseMatrix",
			s:        doublePrecision.NewDenseMatrix(2, 2),
			isSparse: false,
		},
		{
			name:     "CSCMatrix",
			s:        doublePrecision.NewCSCMatrix(2, 2),
			isSparse: true,
		},
		{
			name:     "CSRMatrix",
			s:        doublePrecision.NewCSRMatrix(2, 2),
			isSparse: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := doublePrecision.IsSparseMatrix(tt.s)
			if tt.isSparse != v {
				t.Errorf("%+v IsSparseMatrix = %+v, want %+v", tt.name, v, tt.isSparse)
			}
		})
	}
}
