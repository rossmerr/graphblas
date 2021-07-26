// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package f64_test

import (
	"github.com/rossmerr/graphblas/f64"
	"testing"

	"golang.org/x/net/context"
)

func setupMatrix(m f64.Matrix) {
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

	setup := func(m f64.Matrix) {
		m.Set(0, 0, 2)
		m.Set(0, 1, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(2, 0, -1)
		m.Set(2, 1, 6)
	}

	want := f64.NewDenseVector(3)
	want.SetVec(0, 29)
	want.SetVec(1, 51)
	want.SetVec(2, 38)

	vector := f64.NewDenseVector(2)
	vector.SetVec(0, 4)
	vector.SetVec(1, 7)

	tests := []struct {
		name string
		s    f64.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(3, 2),
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(3, 2),
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(3, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := f64.NewDenseVector(3)
			f64.VectorMatrixMultiply(context.Background(), vector, tt.s, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v VectorMatrixMultiply = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_MatrixVectorMultiply(t *testing.T) {

	setup := func(m f64.Matrix) {
		m.Set(0, 0, 2)
		m.Set(0, 1, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(2, 0, -1)
		m.Set(2, 1, 6)
	}

	want := f64.NewDenseVector(3)
	want.SetVec(0, 29)
	want.SetVec(1, 51)
	want.SetVec(2, 38)

	vector := f64.NewDenseVector(2)
	vector.SetVec(0, 4)
	vector.SetVec(1, 7)

	tests := []struct {
		name string
		s    f64.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(3, 2),
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(3, 2),
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(3, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := f64.NewDenseVector(3)
			f64.MatrixVectorMultiply(context.Background(), tt.s, vector, nil, got)
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
	want := f64.NewDenseMatrixFromArray(array)

	array2 := [][]float64{
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 0},
	}
	matrix := f64.NewDenseMatrixFromArray(array2)

	tests := []struct {
		name string
		s    f64.Matrix
		got  func(t f64.Matrix) f64.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(7, 7),
			got: func(t f64.Matrix) f64.Matrix {
				return f64.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(7, 7),
			got: func(t f64.Matrix) f64.Matrix {
				return f64.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(7, 7),
			got: func(t f64.Matrix) f64.Matrix {
				return f64.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "Target is Source",
			s:    f64.NewCSRMatrix(7, 7),
			got: func(t f64.Matrix) f64.Matrix {
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)
			got := tt.got(tt.s)
			f64.ElementWiseMatrixMultiply(context.Background(), tt.s, matrix, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseMatrixMultiply = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

func TestMatrix_ElementWiseVectorMultiply(t *testing.T) {
	vector := f64.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 1, 0})

	want := f64.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 0, 0})

	setup := []float64{0, 1, 0, 0, 0, 0, 1}

	tests := []struct {
		name string
		s    f64.Vector
	}{
		{
			name: "DenseVector",
			s:    f64.NewDenseVectorFromArray(setup),
		},
		{
			name: "SparseVector",
			s:    f64.NewSparseVectorFromArray(setup),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f64.NewDenseVector(7)
			f64.ElementWiseVectorMultiply(context.Background(), tt.s, vector, nil, got)
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
	want := f64.NewDenseMatrixFromArray(array)

	array2 := [][]float64{
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 1},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 0},
	}
	matrix := f64.NewDenseMatrixFromArray(array2)

	tests := []struct {
		name string
		s    f64.Matrix
		got  func(t f64.Matrix) f64.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(7, 7),
			got: func(t f64.Matrix) f64.Matrix {
				return f64.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(7, 7),
			got: func(t f64.Matrix) f64.Matrix {
				return f64.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(7, 7),
			got: func(t f64.Matrix) f64.Matrix {
				return f64.NewDenseMatrix(7, 7)
			},
		},
		{
			name: "Target is Source",
			s:    f64.NewDenseMatrix(7, 7),
			got: func(t f64.Matrix) f64.Matrix {
				return t
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)
			got := tt.got(tt.s)
			f64.ElementWiseMatrixAdd(context.Background(), tt.s, matrix, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseMatrixAdd = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

func TestMatrix_ElementWiseVectorAdd(t *testing.T) {
	vector := f64.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 1, 0})

	want := f64.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 1, 1})

	setup := []float64{0, 1, 0, 0, 0, 0, 1}

	tests := []struct {
		name string
		s    f64.Vector
	}{
		{
			name: "DenseVector",
			s:    f64.NewDenseVectorFromArray(setup),
		},
		{
			name: "SparseVector",
			s:    f64.NewSparseVectorFromArray(setup),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f64.NewDenseVector(7)
			f64.ElementWiseVectorAdd(context.Background(), tt.s, vector, nil, got)
			if !got.Equal(want) {
				t.Errorf("%+v ElementWiseVectorAdd = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
			}
		})
	}
}

func TestMatrix_Transpose_To_CSR(t *testing.T) {

	setup := func(m f64.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
		m.Set(1, 2, 8)
	}

	want := f64.NewDenseMatrix(3, 2)
	want.Set(0, 0, 6)
	want.Set(0, 1, 1)
	want.Set(1, 0, 4)
	want.Set(1, 1, -9)
	want.Set(2, 0, 24)
	want.Set(2, 1, 8)

	tests := []struct {
		name string
		s    f64.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := f64.TransposeToCSR(context.Background(), tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Transpose_To_CSC(t *testing.T) {

	setup := func(m f64.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
		m.Set(1, 2, 8)
	}

	want := f64.NewDenseMatrix(3, 2)
	want.Set(0, 0, 6)
	want.Set(0, 1, 1)
	want.Set(1, 0, 4)
	want.Set(1, 1, -9)
	want.Set(2, 0, 24)
	want.Set(2, 1, 8)

	tests := []struct {
		name string
		s    f64.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := f64.TransposeToCSC(context.Background(), tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_ReduceMatrixToVector(t *testing.T) {
	want := f64.NewDenseVectorFromArray([]float64{1, 1, 0, 1, 0, 0, 1})

	tests := []struct {
		name string
		s    f64.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(7, 7),
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(7, 7),
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(7, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)

			got := f64.ReduceMatrixToVector(context.Background(), tt.s)

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
		s    f64.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(7, 7),
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(7, 7),
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(7, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)

			got := f64.ReduceMatrixToScalar(context.Background(), tt.s, nil)

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
		s    f64.Vector
	}{
		{
			name: "DenseVector",
			s:    f64.NewDenseVector(7),
		},
		{
			name: "SparseVector",
			s:    f64.NewSparseVector(7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := f64.NewDenseMatrix(7, 7)
			setupMatrix(matrix)
			tt.s = matrix.ColumnsAt(0)

			got := f64.ReduceVectorToScalar(context.Background(), tt.s, nil)

			if got != want {
				t.Errorf("%+v ReduceVectorToScalar = \nhave %+v, \nwant %+v", tt.name, got, want)
			}
		})
	}
}
