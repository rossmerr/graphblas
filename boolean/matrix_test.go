// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolean_test

import (
	"testing"

	"github.com/rossmerr/graphblas/boolean"
)

func TestMatrix_Update(t *testing.T) {

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, false)
		m.Set(1, 0, true)
		m.Set(1, 1, false)
	}

	tests := []struct {
		name  string
		s     boolean.Matrix
		want  bool
		value bool
	}{
		{
			name:  "DenseMatrix",
			s:     boolean.NewDenseMatrix(2, 2),
			want:  true,
			value: true,
		},
		{
			name:  "CSCMatrix",
			s:     boolean.NewCSCMatrix(2, 2),
			want:  true,
			value: true,
		},
		{
			name:  "CSRMatrix",
			s:     boolean.NewCSRMatrix(2, 2),
			want:  true,
			value: true,
		},
		// Checks values get removed for sparse matrix
		{
			name:  "CSCMatrix",
			s:     boolean.NewCSCMatrix(2, 2),
			want:  false,
			value: false,
		},
		{
			name:  "CSRMatrix",
			s:     boolean.NewCSRMatrix(2, 2),
			want:  false,
			value: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			tt.s.Update(0, 0, func(v bool) bool {
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
	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, false)
		m.Set(0, 2, true)
		m.Set(1, 0, false)
		m.Set(1, 1, false)
		m.Set(1, 2, false)
		m.Set(2, 0, true)
		m.Set(2, 1, false)
		m.Set(2, 2, true)
	}

	dense := boolean.NewDenseMatrix(3, 3)
	setup(dense)
	denseCount := 0
	for iterator := dense.Enumerate(); iterator.HasNext(); {
		_, _, value := iterator.Next()
		if value != false {
			denseCount++
		}
	}

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(3, 3),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(3, 3),
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
	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, false)
		m.Set(0, 2, true)
		m.Set(1, 0, false)
		m.Set(1, 1, false)
		m.Set(1, 2, false)
		m.Set(2, 0, true)
		m.Set(2, 1, false)
		m.Set(2, 2, true)
	}

	dense := boolean.NewDenseMatrix(3, 3)
	setup(dense)
	denseCount := 0
	for iterator := dense.Enumerate(); iterator.HasNext(); {
		_, _, value := iterator.Next()
		if value != false {
			denseCount++
		}
	}

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(3, 3),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			count := 0
			for iterator := tt.s.Map(); iterator.HasNext(); {
				iterator.Map(func(r, c int, value bool) bool {
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

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, false)
		m.Set(1, 0, true)
		m.Set(1, 1, false)
	}

	want := boolean.NewDenseVector(2)
	want.SetVec(0, true)
	want.SetVec(1, true)

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    boolean.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(2, 2),
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

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, false)
		m.Set(1, 0, true)
		m.Set(1, 1, false)
	}

	want := boolean.NewDenseVector(2)
	want.SetVec(0, true)
	want.SetVec(1, false)

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    boolean.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(2, 2),
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

func TestMatrix_Transpose(t *testing.T) {

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, false)
		m.Set(0, 1, false)
		m.Set(0, 2, true)
		m.Set(1, 0, true)
		m.Set(1, 1, false)
		m.Set(1, 2, false)
	}

	want := boolean.NewDenseMatrix(3, 2)
	want.Set(0, 0, false)
	want.Set(0, 1, false)
	want.Set(1, 0, true)
	want.Set(1, 1, true)
	want.Set(2, 0, false)
	want.Set(2, 1, false)

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    boolean.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(2, 3),
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

func TestMatrix_Equal(t *testing.T) {

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, true)
		m.Set(0, 2, false)
		m.Set(1, 0, false)
		m.Set(1, 1, true)
		m.Set(1, 2, true)
	}

	want := boolean.NewDenseMatrix(2, 3)
	want.Set(0, 0, true)
	want.Set(0, 1, true)
	want.Set(0, 2, false)
	want.Set(1, 0, false)
	want.Set(1, 1, true)
	want.Set(1, 2, true)

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    boolean.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(2, 3),
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

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, true)
		m.Set(0, 2, false)
		m.Set(1, 0, false)
		m.Set(1, 1, true)
		m.Set(1, 2, true)
	}

	want := boolean.NewDenseMatrix(2, 3)
	want.Set(0, 0, true)
	want.Set(0, 1, true)
	want.Set(0, 2, false)
	want.Set(1, 0, false)
	want.Set(1, 1, true)
	want.Set(1, 2, true)

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    boolean.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(2, 3),
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
		s    boolean.Matrix
		want boolean.Matrix
	}{
		{
			name: "DenseMatrix Row",
			s:    boolean.NewDenseMatrix(2, 2),
			want: boolean.NewDenseMatrix(3, 2),
		},
		{
			name: "DenseMatrix Column",
			s:    boolean.NewDenseMatrix(2, 2),
			want: boolean.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix Row",
			s:    boolean.NewCSCMatrix(2, 2),
			want: boolean.NewDenseMatrix(3, 2),
		},
		{
			name: "CSCMatrix Column",
			s:    boolean.NewCSCMatrix(2, 2),
			want: boolean.NewDenseMatrix(2, 3),
		},
		{
			name: "CSRMatrix Row",
			s:    boolean.NewCSRMatrix(2, 2),
			want: boolean.NewDenseMatrix(3, 2),
		},
		{
			name: "CSRMatrix Column",
			s:    boolean.NewCSRMatrix(2, 2),
			want: boolean.NewDenseMatrix(2, 3),
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

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, false)
		m.Set(0, 1, false)
		m.Set(0, 2, true)
		m.Set(1, 0, true)
		m.Set(1, 1, false)
		m.Set(1, 2, false)
	}

	want := boolean.NewDenseMatrix(2, 3)
	want.Set(0, 0, false)
	want.Set(0, 1, false)
	want.Set(0, 2, true)
	want.Set(1, 0, true)
	want.Set(1, 1, false)
	want.Set(1, 2, false)

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    boolean.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(2, 3),
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

func TestMatrix_Size(t *testing.T) {

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, true)
		m.Set(0, 2, true)
		m.Set(1, 0, true)
		m.Set(1, 1, false)
		m.Set(1, 2, true)
	}

	tests := []struct {
		name string
		s    boolean.Matrix
		size int
	}{
		{
			name: "DenseMatrix",
			s:    boolean.NewDenseMatrix(2, 3),
			size: 6,
		},
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrix(2, 3),
			size: 5,
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrix(2, 3),
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

	want := boolean.NewDenseMatrix(2, 3)
	want.Set(0, 0, true)
	want.Set(0, 1, true)
	want.Set(0, 2, true)
	want.Set(1, 0, true)
	want.Set(1, 1, true)
	want.Set(1, 2, true)

	setup := want.RawMatrix()

	tests := []struct {
		name string
		s    boolean.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    boolean.NewDenseMatrixFromArray(setup),
		},
		{
			name: "CSCMatrix",
			s:    boolean.NewCSCMatrixFromArray(setup),
		},
		{
			name: "CSRMatrix",
			s:    boolean.NewCSRMatrixFromArray(setup),
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
