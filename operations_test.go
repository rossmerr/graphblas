// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
	"github.com/RossMerr/Caudex.GraphBLAS/binaryOp/float64Op"
)

func setupMatrix(m GraphBLAS.Matrix) {
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
	want := GraphBLAS.NewDenseMatrixFromArray(array)

	array2 := [][]float64{
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 0},
	}
	matrix := GraphBLAS.NewDenseMatrixFromArray(array2)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
		got  func(t GraphBLAS.Matrix) GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(7, 7),
			got: func(t GraphBLAS.Matrix) GraphBLAS.Matrix {
				return GraphBLAS.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(7, 7),
			got: func(t GraphBLAS.Matrix) GraphBLAS.Matrix {
				return GraphBLAS.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(7, 7),
			got: func(t GraphBLAS.Matrix) GraphBLAS.Matrix {
				return GraphBLAS.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "Target is Source",
			s:    GraphBLAS.NewCSRMatrix(7, 7),
			got: func(t GraphBLAS.Matrix) GraphBLAS.Matrix {
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)
			got := tt.got(tt.s)
			GraphBLAS.ElementWiseMatrixMultiply(tt.s, matrix, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseMatrixMultiply = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

// func TestMatrix_VectorMatrixMultiply(t *testing.T) {
// 	want := GraphBLAS.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 0, 0})

// 	array := [][]float64{
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 1, 0, 1},
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 1, 0, 0},
// 	}
// 	matrix := GraphBLAS.NewDenseMatrixFromArray(array)

// 	tests := []struct {
// 		name string
// 		s    GraphBLAS.Matrix
// 	}{
// 		{
// 			name: "DenseMatrix",
// 			s:    GraphBLAS.NewDenseMatrix(7, 7),
// 		},
// 		{
// 			name: "CSCMatrix",
// 			s:    GraphBLAS.NewCSCMatrix(7, 7),
// 		},
// 		{
// 			name: "CSRMatrix",
// 			s:    GraphBLAS.NewCSRMatrix(7, 7),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			setupMatrix(tt.s)
// 			got := GraphBLAS.NewDenseVector(matrix.Columns())
// 			GraphBLAS.VectorMatrixMultiply(tt.s.ColumnsAt(6), matrix, got)
// 			if !got.Equal(want) {
// 				t.Errorf("%+v VectorMatrixMultiply = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
// 			}
// 		})
// 	}
// }

// func TestMatrix_MatrixVectorMultiply(t *testing.T) {
// 	vector := GraphBLAS.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 0, 0})

// 	array2 := [][]float64{
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 1, 0, 1},
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 0, 0, 0},
// 		[]float64{0, 0, 0, 0, 1, 0, 0},
// 	}
// 	want := GraphBLAS.NewDenseMatrixFromArray(array2)

// 	tests := []struct {
// 		name string
// 		s    GraphBLAS.Matrix
// 	}{
// 		// {
// 		// 	name: "DenseMatrix",
// 		// 	s:    GraphBLAS.NewDenseMatrix(7, 7),
// 		// },
// 		// {
// 		// 	name: "CSCMatrix",
// 		// 	s:    GraphBLAS.NewCSCMatrix(7, 7),
// 		// },
// 		// {
// 		// 	name: "CSRMatrix",
// 		// 	s:    GraphBLAS.NewCSRMatrix(7, 7),
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			setupMatrix(tt.s)
// 			got := GraphBLAS.NewDenseVector(7)
// 			GraphBLAS.MatrixVectorMultiply(tt.s, vector, got)
// 			if !got.Equal(want) {
// 				t.Errorf("%+v MatrixVectorMultiply = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
// 			}
// 		})
// 	}
// }

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
	want := GraphBLAS.NewDenseMatrixFromArray(array)

	array2 := [][]float64{
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 0},
	}
	matrix := GraphBLAS.NewDenseMatrixFromArray(array2)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
		got  func(t GraphBLAS.Matrix) GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(7, 7),
			got: func(t GraphBLAS.Matrix) GraphBLAS.Matrix {
				return GraphBLAS.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(7, 7),
			got: func(t GraphBLAS.Matrix) GraphBLAS.Matrix {
				return GraphBLAS.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(7, 7),
			got: func(t GraphBLAS.Matrix) GraphBLAS.Matrix {
				return GraphBLAS.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "Target is Source",
			s:    GraphBLAS.NewDenseMatrix(7, 7),
			got: func(t GraphBLAS.Matrix) GraphBLAS.Matrix {
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)
			got := tt.got(tt.s)
			GraphBLAS.ElementWiseMatrixAdd(tt.s, matrix, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseMatrixAdd = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

func TestMatrix_Transpose_To_CSR(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
		m.Set(1, 2, 8)
	}

	want := GraphBLAS.NewDenseMatrix(3, 2)
	want.Set(0, 0, 6)
	want.Set(0, 1, 1)
	want.Set(1, 0, 4)
	want.Set(1, 1, -9)
	want.Set(2, 0, 24)
	want.Set(2, 1, 8)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := GraphBLAS.TransposeToCSR(tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Transpose_To_CSC(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
		m.Set(1, 2, 8)
	}

	want := GraphBLAS.NewDenseMatrix(3, 2)
	want.Set(0, 0, 6)
	want.Set(0, 1, 1)
	want.Set(1, 0, 4)
	want.Set(1, 1, -9)
	want.Set(2, 0, 24)
	want.Set(2, 1, 8)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := GraphBLAS.TransposeToCSC(tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_ReduceMatrixToVector(t *testing.T) {
	want := GraphBLAS.NewDenseVectorFromArray([]float64{1, 1, 0, 1, 0, 0, 1})

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(7, 7),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(7, 7),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(7, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)

			monoID := float64Op.NewMonoIDFloat64(0, float64Op.Maximum)
			got := GraphBLAS.ReduceMatrixToVector(tt.s, monoID)

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
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(7, 7),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(7, 7),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(7, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)

			monoID := float64Op.NewMonoIDFloat64(0, float64Op.Addition)
			got := GraphBLAS.ReduceMatrixToScalar(tt.s, monoID)

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
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(7),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := GraphBLAS.NewDenseMatrix(7, 7)
			setupMatrix(matrix)
			tt.s = matrix.ColumnsAt(0)

			monoID := float64Op.NewMonoIDFloat64(0, float64Op.Addition)
			got := GraphBLAS.ReduceVectorToScalar(tt.s, monoID)

			if got != want {
				t.Errorf("%+v ReduceVectorToScalar = \nhave %+v, \nwant %+v", tt.name, got, want)
			}
		})
	}
}
