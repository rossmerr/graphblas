// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package doublePrecision_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/doublePrecision"
	"golang.org/x/net/context"
)

func setupMatrix(m doublePrecision.Matrix) {
	m.Set(0, 0, 0)
	m.Set(0, 1, 1)
	m.Set(0, 2, 0)
	m.Set(0, 3, 1)
	m.Set(0, 4, 0)
	m.Set(0, 5, 0)
	m.Set(0, 6, 0)

	m.Set(1, 0, 0)
	m.Set(1, 1, 0)
	m.Set(1, 2, 0)
	m.Set(1, 3, 0)
	m.Set(1, 4, 0)
	m.Set(1, 5, 0)
	m.Set(1, 6, 1)

	m.Set(2, 0, 0)
	m.Set(2, 1, 0)
	m.Set(2, 2, 0)
	m.Set(2, 3, 0)
	m.Set(2, 4, 0)
	m.Set(2, 5, 0)
	m.Set(2, 6, 0)

	m.Set(3, 0, 1)
	m.Set(3, 1, 0)
	m.Set(3, 2, 0)
	m.Set(3, 3, 0)
	m.Set(3, 4, 0)
	m.Set(3, 5, 0)
	m.Set(3, 6, 0)

	m.Set(4, 0, 0)
	m.Set(4, 1, 0)
	m.Set(4, 2, 0)
	m.Set(4, 3, 0)
	m.Set(4, 4, 0)
	m.Set(4, 5, 0)
	m.Set(4, 6, 0)

	m.Set(5, 0, 0)
	m.Set(5, 1, 0)
	m.Set(5, 2, 0)
	m.Set(5, 3, 0)
	m.Set(5, 4, 0)
	m.Set(5, 5, 0)
	m.Set(5, 6, 0)

	m.Set(6, 0, 0)
	m.Set(6, 1, 0)
	m.Set(6, 2, 0)
	m.Set(6, 3, 1)
	m.Set(6, 4, 0)
	m.Set(6, 5, 0)
	m.Set(6, 6, 0)

}

func TestMatrix_VectorMatrixMultiply(t *testing.T) {

	setup := func(m doublePrecision.Matrix) {
		m.Set(0, 0, 2)
		m.Set(0, 1, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(2, 0, -1)
		m.Set(2, 1, 6)
	}

	want := doublePrecision.NewDenseVector(3)
	want.SetVec(0, 29)
	want.SetVec(1, 51)
	want.SetVec(2, 38)

	vector := doublePrecision.NewDenseVector(2)
	vector.SetVec(0, 4)
	vector.SetVec(1, 7)

	tests := []struct {
		name string
		s    doublePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    doublePrecision.NewDenseMatrix(3, 2),
		},
		{
			name: "CSCMatrix",
			s:    doublePrecision.NewCSCMatrix(3, 2),
		},
		{
			name: "CSRMatrix",
			s:    doublePrecision.NewCSRMatrix(3, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := doublePrecision.NewDenseVector(3)
			doublePrecision.VectorMatrixMultiply(context.Background(), vector, tt.s, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v VectorMatrixMultiply = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_MatrixVectorMultiply(t *testing.T) {

	setup := func(m doublePrecision.Matrix) {
		m.Set(0, 0, 2)
		m.Set(0, 1, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(2, 0, -1)
		m.Set(2, 1, 6)
	}

	want := doublePrecision.NewDenseVector(3)
	want.SetVec(0, 29)
	want.SetVec(1, 51)
	want.SetVec(2, 38)

	vector := doublePrecision.NewDenseVector(2)
	vector.SetVec(0, 4)
	vector.SetVec(1, 7)

	tests := []struct {
		name string
		s    doublePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    doublePrecision.NewDenseMatrix(3, 2),
		},
		{
			name: "CSCMatrix",
			s:    doublePrecision.NewCSCMatrix(3, 2),
		},
		{
			name: "CSRMatrix",
			s:    doublePrecision.NewCSRMatrix(3, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := doublePrecision.NewDenseVector(3)
			doublePrecision.MatrixVectorMultiply(context.Background(), tt.s, vector, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v MatrixVectorMultiply = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_ElementWiseMatrixMultiply(t *testing.T) {
	array := [][]float64{
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
	}
	want := doublePrecision.NewDenseMatrixFromArray(array)

	array2 := [][]float64{
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 0},
	}
	matrix := doublePrecision.NewDenseMatrixFromArray(array2)

	tests := []struct {
		name string
		s    doublePrecision.Matrix
		got  func(t doublePrecision.Matrix) doublePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    doublePrecision.NewDenseMatrix(7, 7),
			got: func(t doublePrecision.Matrix) doublePrecision.Matrix {
				return doublePrecision.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSCMatrix",
			s:    doublePrecision.NewCSCMatrix(7, 7),
			got: func(t doublePrecision.Matrix) doublePrecision.Matrix {
				return doublePrecision.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSRMatrix",
			s:    doublePrecision.NewCSRMatrix(7, 7),
			got: func(t doublePrecision.Matrix) doublePrecision.Matrix {
				return doublePrecision.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "Target is Source",
			s:    doublePrecision.NewCSRMatrix(7, 7),
			got: func(t doublePrecision.Matrix) doublePrecision.Matrix {
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)
			got := tt.got(tt.s)
			doublePrecision.ElementWiseMatrixMultiply(context.Background(), tt.s, matrix, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseMatrixMultiply = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

func TestMatrix_ElementWiseVectorMultiply(t *testing.T) {
	vector := doublePrecision.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 1, 0})

	want := doublePrecision.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 0, 0})

	setup := []float64{0, 1, 0, 0, 0, 0, 1}

	tests := []struct {
		name string
		s    doublePrecision.Vector
	}{
		{
			name: "DenseVector",
			s:    doublePrecision.NewDenseVectorFromArray(setup),
		},
		{
			name: "SparseVector",
			s:    doublePrecision.NewSparseVectorFromArray(setup),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doublePrecision.NewDenseVector(7)
			doublePrecision.ElementWiseVectorMultiply(context.Background(), tt.s, vector, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseVectorMultiply = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

func TestMatrix_ElementWiseMatrixAdd(t *testing.T) {
	array := [][]float64{
		[]float64{0, 1, 0, 1, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{1, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 1, 1, 0, 0},
	}
	want := doublePrecision.NewDenseMatrixFromArray(array)

	array2 := [][]float64{
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 0},
	}
	matrix := doublePrecision.NewDenseMatrixFromArray(array2)

	tests := []struct {
		name string
		s    doublePrecision.Matrix
		got  func(t doublePrecision.Matrix) doublePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    doublePrecision.NewDenseMatrix(7, 7),
			got: func(t doublePrecision.Matrix) doublePrecision.Matrix {
				return doublePrecision.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSCMatrix",
			s:    doublePrecision.NewCSCMatrix(7, 7),
			got: func(t doublePrecision.Matrix) doublePrecision.Matrix {
				return doublePrecision.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSRMatrix",
			s:    doublePrecision.NewCSRMatrix(7, 7),
			got: func(t doublePrecision.Matrix) doublePrecision.Matrix {
				return doublePrecision.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "Target is Source",
			s:    doublePrecision.NewDenseMatrix(7, 7),
			got: func(t doublePrecision.Matrix) doublePrecision.Matrix {
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)
			got := tt.got(tt.s)
			doublePrecision.ElementWiseMatrixAdd(context.Background(), tt.s, matrix, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseMatrixAdd = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

func TestMatrix_ElementWiseVectorAdd(t *testing.T) {
	vector := doublePrecision.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 1, 0})

	want := doublePrecision.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 1, 1})

	setup := []float64{0, 1, 0, 0, 0, 0, 1}

	tests := []struct {
		name string
		s    doublePrecision.Vector
	}{
		{
			name: "DenseVector",
			s:    doublePrecision.NewDenseVectorFromArray(setup),
		},
		{
			name: "SparseVector",
			s:    doublePrecision.NewSparseVectorFromArray(setup),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doublePrecision.NewDenseVector(7)
			doublePrecision.ElementWiseVectorAdd(context.Background(), tt.s, vector, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseVectorAdd = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

func TestMatrix_Transpose_To_CSR(t *testing.T) {

	setup := func(m doublePrecision.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
		m.Set(1, 2, 8)
	}

	want := doublePrecision.NewDenseMatrix(3, 2)
	want.Set(0, 0, 6)
	want.Set(0, 1, 1)
	want.Set(1, 0, 4)
	want.Set(1, 1, -9)
	want.Set(2, 0, 24)
	want.Set(2, 1, 8)

	tests := []struct {
		name string
		s    doublePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    doublePrecision.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    doublePrecision.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    doublePrecision.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := doublePrecision.TransposeToCSR(context.Background(), tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Transpose_To_CSC(t *testing.T) {

	setup := func(m doublePrecision.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
		m.Set(1, 2, 8)
	}

	want := doublePrecision.NewDenseMatrix(3, 2)
	want.Set(0, 0, 6)
	want.Set(0, 1, 1)
	want.Set(1, 0, 4)
	want.Set(1, 1, -9)
	want.Set(2, 0, 24)
	want.Set(2, 1, 8)

	tests := []struct {
		name string
		s    doublePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    doublePrecision.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    doublePrecision.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    doublePrecision.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := doublePrecision.TransposeToCSC(context.Background(), tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_ReduceMatrixToVector(t *testing.T) {
	want := doublePrecision.NewDenseVectorFromArray([]float64{1, 1, 0, 1, 0, 0, 1})

	tests := []struct {
		name string
		s    doublePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    doublePrecision.NewDenseMatrix(7, 7),
		},
		{
			name: "CSCMatrix",
			s:    doublePrecision.NewCSCMatrix(7, 7),
		},
		{
			name: "CSRMatrix",
			s:    doublePrecision.NewCSRMatrix(7, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)

			got := doublePrecision.ReduceMatrixToVector(context.Background(), tt.s)

			if !got.Equal(want) {
				t.Errorf("%+v ReduceMatrixToVector = \nhave %+v, \nwant %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_ReduceMatrixToScalar(t *testing.T) {
	want := float64(5)

	tests := []struct {
		name string
		s    doublePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    doublePrecision.NewDenseMatrix(7, 7),
		},
		{
			name: "CSCMatrix",
			s:    doublePrecision.NewCSCMatrix(7, 7),
		},
		{
			name: "CSRMatrix",
			s:    doublePrecision.NewCSRMatrix(7, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)

			got := doublePrecision.ReduceMatrixToScalar(context.Background(), tt.s, nil)

			if got != want {
				t.Errorf("%+v ReduceMatrixToScalar = \nhave %+v, \nwant %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_ReduceVectorToScalar(t *testing.T) {
	want := float64(1)

	tests := []struct {
		name string
		s    doublePrecision.Vector
	}{
		{
			name: "DenseVector",
			s:    doublePrecision.NewDenseVector(7),
		},
		{
			name: "SparseVector",
			s:    doublePrecision.NewSparseVector(7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := doublePrecision.NewDenseMatrix(7, 7)
			setupMatrix(matrix)
			tt.s = matrix.ColumnsAt(0)

			got := doublePrecision.ReduceVectorToScalar(context.Background(), tt.s, nil)

			if got != want {
				t.Errorf("%+v ReduceVectorToScalar = \nhave %+v, \nwant %+v", tt.name, got, want)
			}
		})
	}
}
