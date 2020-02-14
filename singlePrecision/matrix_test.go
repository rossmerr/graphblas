// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package singlePrecision_test

import (
	"testing"

	singlePrecision "github.com/RossMerr/Caudex.GraphBLAS/singlePrecision"
)

func TestMatrix_Update(t *testing.T) {

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	tests := []struct {
		name  string
		s     singlePrecision.Matrix
		want  float32
		value float32
	}{
		{
			name:  "DenseMatrix",
			s:     singlePrecision.NewDenseMatrix(2, 2),
			want:  2,
			value: 2,
		},
		// {
		// 	name:  "CSCMatrix",
		// 	s:     singlePrecision.NewCSCMatrix(2, 2),
		// 	want:  2,
		// 	value: 2,
		// },
		// {
		// 	name:  "CSRMatrix",
		// 	s:     singlePrecision.NewCSRMatrix(2, 2),
		// 	want:  2,
		// 	value: 2,
		// },
		// // Checks values get removed for sparse matrix
		// {
		// 	name:  "CSCMatrix",
		// 	s:     GraphBLAS.NewCSCMatrix(2, 2),
		// 	want:  0,
		// 	value: 0,
		// },
		// {
		// 	name:  "CSRMatrix",
		// 	s:     GraphBLAS.NewCSRMatrix(2, 2),
		// 	want:  0,
		// 	value: 0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			tt.s.Update(0, 0, func(v float32) float32 {
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
	setup := func(m singlePrecision.Matrix) {
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

	dense := singlePrecision.NewDenseMatrix(3, 3)
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
		s    singlePrecision.Matrix
	}{
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(3, 3),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(3, 3),
		// },
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
	setup := func(m singlePrecision.Matrix) {
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

	dense := singlePrecision.NewDenseMatrix(3, 3)
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
		s    singlePrecision.Matrix
	}{
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(3, 3),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(3, 3),
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			count := 0
			for iterator := tt.s.Map(); iterator.HasNext(); {
				iterator.Map(func(r, c int, value float32) float32 {
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := singlePrecision.NewDenseVector(2)
	want.SetVec(0, 4)
	want.SetVec(1, 1)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 2),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 2),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 2),
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := singlePrecision.NewDenseVector(2)
	want.SetVec(0, 4)
	want.SetVec(1, 0)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 2),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 2),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 2),
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := singlePrecision.NewDenseMatrix(2, 2)
	want.Set(0, 0, 8)
	want.Set(0, 1, 0)
	want.Set(1, 0, 2)
	want.Set(1, 1, -18)

	tests := []struct {
		name  string
		s     singlePrecision.Matrix
		alpha float32
	}{
		{
			name:  "DenseMatrix",
			s:     singlePrecision.NewDenseMatrix(2, 2),
			alpha: 2,
		},
		// {
		// 	name:  "CSCMatrix",
		// 	s:     singlePrecision.NewCSCMatrix(2, 2),
		// 	alpha: 2,
		// },
		// {
		// 	name:  "CSRMatrix",
		// 	s:     singlePrecision.NewCSRMatrix(2, 2),
		// 	alpha: 2,
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 2)
		m.Set(0, 1, -4)
		m.Set(1, 0, 7)
		m.Set(1, 1, 10)
	}

	want := singlePrecision.NewDenseMatrix(2, 2)
	want.Set(0, 0, -2)
	want.Set(0, 1, 4)
	want.Set(1, 0, -7)
	want.Set(1, 1, -10)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 2),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 2),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 2),
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
		m.Set(1, 2, 8)
	}

	want := singlePrecision.NewDenseMatrix(3, 2)
	want.Set(0, 0, 6)
	want.Set(0, 1, 1)
	want.Set(1, 0, 4)
	want.Set(1, 1, -9)
	want.Set(2, 0, 24)
	want.Set(2, 1, 8)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 3),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 3),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 3),
		// },
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

func TestMatrix_Equal(t *testing.T) {

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := singlePrecision.NewDenseMatrix(2, 3)
	want.Set(0, 0, 1)
	want.Set(0, 1, 2)
	want.Set(0, 2, 3)
	want.Set(1, 0, 4)
	want.Set(1, 1, 5)
	want.Set(1, 2, 6)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 3),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 3),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 3),
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := singlePrecision.NewDenseMatrix(2, 3)
	want.Set(0, 0, 2)
	want.Set(0, 1, 3)
	want.Set(0, 2, 4)
	want.Set(1, 0, 5)
	want.Set(1, 1, 6)
	want.Set(1, 2, 7)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 3),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 3),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 3),
		// },
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
		s    singlePrecision.Matrix
		want singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix Row",
			s:    singlePrecision.NewDenseMatrix(2, 2),
			want: singlePrecision.NewDenseMatrix(3, 2),
		},
		{
			name: "DenseMatrix Column",
			s:    singlePrecision.NewDenseMatrix(2, 2),
			want: singlePrecision.NewDenseMatrix(2, 3),
		},
		// {
		// 	name: "CSCMatrix Row",
		// 	s:    singlePrecision.NewCSCMatrix(2, 2),
		// 	want: singlePrecision.NewDenseMatrix(3, 2),
		// },
		// {
		// 	name: "CSCMatrix Column",
		// 	s:    singlePrecision.NewCSCMatrix(2, 2),
		// 	want: singlePrecision.NewDenseMatrix(2, 3),
		// },
		// {
		// 	name: "CSRMatrix Row",
		// 	s:    singlePrecision.NewCSRMatrix(2, 2),
		// 	want: singlePrecision.NewDenseMatrix(3, 2),
		// },
		// {
		// 	name: "CSRMatrix Column",
		// 	s:    singlePrecision.NewCSRMatrix(2, 2),
		// 	want: singlePrecision.NewDenseMatrix(2, 3),
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := singlePrecision.NewDenseMatrix(2, 3)
	want.Set(0, 0, 1)
	want.Set(0, 1, 2)
	want.Set(0, 2, 3)
	want.Set(1, 0, 4)
	want.Set(1, 1, 5)
	want.Set(1, 2, 6)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 3),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 3),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 3),
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := singlePrecision.NewDenseMatrix(2, 2)
	want.Set(0, 0, 58)
	want.Set(0, 1, 64)
	want.Set(1, 0, 139)
	want.Set(1, 1, 154)

	matrix := singlePrecision.NewDenseMatrix(3, 2)
	matrix.Set(0, 0, 7)
	matrix.Set(0, 1, 8)
	matrix.Set(1, 0, 9)
	matrix.Set(1, 1, 10)
	matrix.Set(2, 0, 11)
	matrix.Set(2, 1, 12)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 3),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 3),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 3),
		// },
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

func TestMatrix_Add(t *testing.T) {
	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 3)
		m.Set(0, 1, 8)
		m.Set(1, 0, 4)
		m.Set(1, 1, 6)
	}

	want := singlePrecision.NewDenseMatrix(2, 2)
	want.Set(0, 0, 7)
	want.Set(0, 1, 8)
	want.Set(1, 0, 5)
	want.Set(1, 1, -3)

	matrix := singlePrecision.NewDenseMatrix(2, 2)
	matrix.Set(0, 0, 4)
	matrix.Set(0, 1, 0)
	matrix.Set(1, 0, 1)
	matrix.Set(1, 1, -9)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 2),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 2),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 2),
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 3)
		m.Set(0, 1, 8)
		m.Set(1, 0, 4)
		m.Set(1, 1, 6)
	}

	want := singlePrecision.NewDenseMatrix(2, 2)
	want.Set(0, 0, -1)
	want.Set(0, 1, 8)
	want.Set(1, 0, 3)
	want.Set(1, 1, 15)

	matrix := singlePrecision.NewDenseMatrix(2, 2)
	matrix.Set(0, 0, 4)
	matrix.Set(0, 1, 0)
	matrix.Set(1, 0, 1)
	matrix.Set(1, 1, -9)

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 2),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 2),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 2),
		// },
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

	setup := func(m singlePrecision.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, 0)
		m.Set(1, 2, 8)
	}

	tests := []struct {
		name string
		s    singlePrecision.Matrix
		size int
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrix(2, 3),
			size: 6,
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrix(2, 3),
		// 	size: 5,
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrix(2, 3),
		// 	size: 5,
		// },
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

	want := singlePrecision.NewDenseMatrix(2, 3)
	want.Set(0, 0, 1)
	want.Set(0, 1, 2)
	want.Set(0, 2, 3)
	want.Set(1, 0, 4)
	want.Set(1, 1, 5)
	want.Set(1, 2, 6)

	setup := want.RawMatrix()

	tests := []struct {
		name string
		s    singlePrecision.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    singlePrecision.NewDenseMatrixFromArray(setup),
		},
		// {
		// 	name: "CSCMatrix",
		// 	s:    singlePrecision.NewCSCMatrixFromArray(setup),
		// },
		// {
		// 	name: "CSRMatrix",
		// 	s:    singlePrecision.NewCSRMatrixFromArray(setup),
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.s.NotEqual(want) {
				t.Errorf("%+v From Array = want %+v", tt.name, want)
			}
		})
	}
}
