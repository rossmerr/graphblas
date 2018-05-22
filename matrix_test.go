// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestMatrix_Update(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	tests := []struct {
		name  string
		s     GraphBLAS.Matrix
		want  float64
		value float64
	}{
		{
			name:  "DenseMatrix",
			s:     GraphBLAS.NewDenseMatrix(2, 2),
			want:  2,
			value: 2,
		},
		{
			name:  "CSCMatrix",
			s:     GraphBLAS.NewCSCMatrix(2, 2),
			want:  2,
			value: 2,
		},
		{
			name:  "CSRMatrix",
			s:     GraphBLAS.NewCSRMatrix(2, 2),
			want:  2,
			value: 2,
		},
		// Checks values get removed for sparse matrix
		{
			name:  "CSCMatrix",
			s:     GraphBLAS.NewCSCMatrix(2, 2),
			want:  0,
			value: 0,
		},
		{
			name:  "CSRMatrix",
			s:     GraphBLAS.NewCSRMatrix(2, 2),
			want:  0,
			value: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			tt.s.Update(0, 0, func(v float64) float64 {
				return tt.value
			})
			v := tt.s.At(0, 0)
			if tt.want != v {
				t.Errorf("%+v Update = %+v, want %+v", tt.name, v, tt.want)
			}
		})
	}
}

func TestMatrix_SparseEnumerate(t *testing.T) {
	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 9)
		m.Set(0, 1, 0)
		m.Set(0, 2, 7)
		m.Set(1, 0, 0)
		m.Set(1, 1, 0)
		m.Set(1, 2, 0)
		m.Set(2, 0, 3)
		m.Set(2, 1, 0)
		m.Set(2, 2, 1)
	}

	dense := GraphBLAS.NewDenseMatrix(3, 3)
	setup(dense)
	denseCount := 0
	for iterator := dense.Enumerate(); iterator.HasNext(); {
		_, _, value := iterator.Next()
		if value != 0 {
			denseCount++
		}
	}

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(3, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			count := 0
			for iterator := tt.s.Enumerate(); iterator.HasNext(); {
				r, c, value := iterator.Next()
				v := dense.At(r, c)
				if v != value {
					t.Errorf("%+v Sparse Enumerate = %+v, want %+v, (r %+v, c %+v)", tt.name, value, v, r, c)
				} else {
					count++
				}
			}
			if denseCount != count {
				t.Errorf("%+v Length miss match = %+v, want %+v", tt.name, count, denseCount)
			}
		})
	}
}

func TestMatrix_SparseMap(t *testing.T) {
	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 9)
		m.Set(0, 1, 0)
		m.Set(0, 2, 7)
		m.Set(1, 0, 0)
		m.Set(1, 1, 0)
		m.Set(1, 2, 0)
		m.Set(2, 0, 3)
		m.Set(2, 1, 0)
		m.Set(2, 2, 1)
	}

	dense := GraphBLAS.NewDenseMatrix(3, 3)
	setup(dense)
	denseCount := 0
	for iterator := dense.Enumerate(); iterator.HasNext(); {
		_, _, value := iterator.Next()
		if value != 0 {
			denseCount++
		}
	}

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(3, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			count := 0
			for iterator := tt.s.Map(); iterator.HasNext(); {
				iterator.Map(func(r, c int, value float64) float64 {
					v := dense.At(r, c)
					if v != value {
						t.Errorf("%+v Sparse Enumerate = %+v, want %+v, (r %+v, c %+v)", tt.name, value, v, r, c)
					} else {
						count++
					}
					return value
				})

			}
			if denseCount != count {
				t.Errorf("%+v Length miss match = %+v, want %+v", tt.name, count, denseCount)
			}
		})
	}
}

func TestMatrix_ColumnsAt(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := GraphBLAS.NewDenseVector(2)
	want.SetVec(0, 4)
	want.SetVec(1, 1)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := tt.s.ColumnsAt(0)
			if !got.Equal(want) {
				t.Errorf("%+v ColumnsAt = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_RowAt(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := GraphBLAS.NewDenseVector(2)
	want.SetVec(0, 4)
	want.SetVec(1, 0)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.RowsAt(0); !got.Equal(want) {
				t.Errorf("%+v RowsAt = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Scalar(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, 8)
	want.Set(0, 1, 0)
	want.Set(1, 0, 2)
	want.Set(1, 1, -18)

	tests := []struct {
		name  string
		s     GraphBLAS.Matrix
		alpha float64
	}{
		{
			name:  "DenseMatrix",
			s:     GraphBLAS.NewDenseMatrix(2, 2),
			alpha: 2,
		},
		{
			name:  "CSCMatrix",
			s:     GraphBLAS.NewCSCMatrix(2, 2),
			alpha: 2,
		},
		{
			name:  "CSRMatrix",
			s:     GraphBLAS.NewCSRMatrix(2, 2),
			alpha: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := tt.s.Scalar(tt.alpha)
			if !got.Equal(want) {
				t.Errorf("%+v Scalar = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Negative(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 2)
		m.Set(0, 1, -4)
		m.Set(1, 0, 7)
		m.Set(1, 1, 10)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, -2)
	want.Set(0, 1, 4)
	want.Set(1, 0, -7)
	want.Set(1, 1, -10)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := tt.s.Negative()
			if !got.Equal(want) {
				t.Errorf("%+v Negative = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Transpose(t *testing.T) {

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
			got := tt.s.Transpose()
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
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

func TestMatrix_Equal(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 3)
	want.Set(0, 0, 1)
	want.Set(0, 1, 2)
	want.Set(0, 2, 3)
	want.Set(1, 0, 4)
	want.Set(1, 1, 5)
	want.Set(1, 2, 6)

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
			if !tt.s.Equal(want) {
				t.Errorf("%+v Equal = %+v, want %+v", tt.name, tt.s, want)
			}
		})
	}
}

func TestMatrix_NotEqual(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 3)
	want.Set(0, 0, 2)
	want.Set(0, 1, 3)
	want.Set(0, 2, 4)
	want.Set(1, 0, 5)
	want.Set(1, 1, 6)
	want.Set(1, 2, 7)

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
			if !tt.s.NotEqual(want) {
				t.Errorf("%+v NotEqual = %+v, want %+v", tt.name, tt.s, want)
			}
		})
	}
}

func TestMatrix_NotEqual_Size(t *testing.T) {

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
		want GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix Row",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
			want: GraphBLAS.NewDenseMatrix(3, 2),
		},
		{
			name: "DenseMatrix Column",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
			want: GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix Row",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
			want: GraphBLAS.NewDenseMatrix(3, 2),
		},
		{
			name: "CSCMatrix Column",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
			want: GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSRMatrix Row",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
			want: GraphBLAS.NewDenseMatrix(3, 2),
		},
		{
			name: "CSRMatrix Column",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
			want: GraphBLAS.NewDenseMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.s.NotEqual(tt.want) {
				t.Errorf("%+v NotEqual = %+v, want %+v", tt.name, tt.s, tt.want)
			}
		})
	}
}

func TestMatrix_Copy(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 3)
	want.Set(0, 0, 1)
	want.Set(0, 1, 2)
	want.Set(0, 2, 3)
	want.Set(1, 0, 4)
	want.Set(1, 1, 5)
	want.Set(1, 2, 6)

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
			if got := tt.s.Copy(); !got.Equal(want) {
				t.Errorf("%+v Copy = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Multiply(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, 58)
	want.Set(0, 1, 64)
	want.Set(1, 0, 139)
	want.Set(1, 1, 154)

	matrix := GraphBLAS.NewDenseMatrix(3, 2)
	matrix.Set(0, 0, 7)
	matrix.Set(0, 1, 8)
	matrix.Set(1, 0, 9)
	matrix.Set(1, 1, 10)
	matrix.Set(2, 0, 11)
	matrix.Set(2, 1, 12)

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
			if got := tt.s.Multiply(matrix); !got.Equal(want) {
				t.Errorf("%+v Multiply = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

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

func TestMatrix_VectorMatrixMultiply(t *testing.T) {
	want := GraphBLAS.NewDenseVectorFromArray([]float64{0, 1, 0, 0, 0, 0, 0})

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
	}{
		// {
		// 	name: "DenseMatrix",
		// 	s:    GraphBLAS.NewDenseMatrix(7, 7),
		// },
		// {
		// 	name: "CSCMatrix",
		// 	s:    GraphBLAS.NewCSCMatrix(7, 7),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    GraphBLAS.NewCSRMatrix(7, 7),
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupMatrix(tt.s)
			got := GraphBLAS.NewDenseVector(matrix.Columns())
			GraphBLAS.VectorMatrixMultiply(tt.s.ColumnsAt(6), matrix, got)
			if !got.Equal(want) {
				t.Errorf("%+v VectorMatrixMultiply = \n%+v, \nwant %+v, \nhave %+v", tt.name, got, want, tt.s)
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

func TestMatrix_Add(t *testing.T) {
	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 3)
		m.Set(0, 1, 8)
		m.Set(1, 0, 4)
		m.Set(1, 1, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, 7)
	want.Set(0, 1, 8)
	want.Set(1, 0, 5)
	want.Set(1, 1, -3)

	matrix := GraphBLAS.NewDenseMatrix(2, 2)
	matrix.Set(0, 0, 4)
	matrix.Set(0, 1, 0)
	matrix.Set(1, 0, 1)
	matrix.Set(1, 1, -9)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Add(matrix); !got.Equal(want) {
				t.Errorf("%+v Add = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Subtract(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 3)
		m.Set(0, 1, 8)
		m.Set(1, 0, 4)
		m.Set(1, 1, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, -1)
	want.Set(0, 1, 8)
	want.Set(1, 0, 3)
	want.Set(1, 1, 15)

	matrix := GraphBLAS.NewDenseMatrix(2, 2)
	matrix.Set(0, 0, 4)
	matrix.Set(0, 1, 0)
	matrix.Set(1, 0, 1)
	matrix.Set(1, 1, -9)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Subtract(matrix); !got.Equal(want) {
				t.Errorf("%+v Subtract = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Size(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, 0)
		m.Set(1, 2, 8)
	}

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
		size int
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 3),
			size: 6,
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 3),
			size: 5,
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 3),
			size: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := tt.s.Values()
			if got != tt.size {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, tt.size)
			}
		})
	}
}

func TestMatrix_FromArray(t *testing.T) {

	want := GraphBLAS.NewDenseMatrix(2, 3)
	want.Set(0, 0, 1)
	want.Set(0, 1, 2)
	want.Set(0, 2, 3)
	want.Set(1, 0, 4)
	want.Set(1, 1, 5)
	want.Set(1, 2, 6)

	setup := want.RawMatrix()

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrixFromArray(setup),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrixFromArray(setup),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrixFromArray(setup),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.s.NotEqual(want) {
				t.Errorf("%+v From Array = want %+v", tt.name, want)
			}
		})
	}
}
