// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestVector_ColumnsAt(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 4)
		m.SetVec(1, 0)
	}

	want := GraphBLAS.NewDenseVector(2)
	want.SetVec(0, 4)
	want.SetVec(1, 0)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(2),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.ColumnsAt(0); !got.Equal(want) {
				t.Errorf("%+v ColumnsAt = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestVector_RowAt(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 4)
		m.SetVec(1, 0)
	}

	want := GraphBLAS.NewDenseVector(1)
	want.SetVec(0, 4)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(2),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := tt.s.RowsAt(0)
			if !got.Equal(want) {
				t.Errorf("%+v RowsAt = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestVector_Scalar(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 4)
		m.SetVec(1, 0)
	}

	want := GraphBLAS.NewDenseVector(2)
	want.SetVec(0, 8)
	want.SetVec(1, 0)

	tests := []struct {
		name  string
		s     GraphBLAS.Vector
		alpha float64
	}{
		{
			name:  "DenseVector",
			s:     GraphBLAS.NewDenseVector(2),
			alpha: 2,
		},
		{
			name:  "SparseVector",
			s:     GraphBLAS.NewSparseVector(2),
			alpha: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Scalar(tt.alpha); !got.Equal(want) {
				t.Errorf("%+v Scalar = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestVector_Negative(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 2)
		m.SetVec(1, -4)
	}

	want := GraphBLAS.NewDenseVector(2)
	want.SetVec(0, -2)
	want.SetVec(1, 4)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(2),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Negative(); !got.Equal(want) {
				t.Errorf("%+v Negative = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestVector_Transpose(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 6)
		m.SetVec(1, 4)
		m.SetVec(2, 24)
	}

	want := GraphBLAS.NewDenseMatrix(1, 3)
	want.Set(0, 0, 6)
	want.Set(0, 1, 4)
	want.Set(0, 2, 24)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Transpose(); !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestVector_Equal(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 1)
		m.SetVec(1, 2)
		m.SetVec(2, 3)
	}

	want := GraphBLAS.NewDenseVector(3)
	want.SetVec(0, 1)
	want.SetVec(1, 2)
	want.SetVec(2, 3)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(3),
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

func TestVector_NotEqual(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 1)
		m.SetVec(1, 2)
		m.SetVec(2, 3)
	}

	want := GraphBLAS.NewDenseVector(3)
	want.SetVec(0, 4)
	want.SetVec(1, 5)
	want.SetVec(2, 6)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(3),
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

func TestVector_Copy(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 1)
		m.SetVec(1, 2)
		m.SetVec(2, 3)
	}

	want := GraphBLAS.NewDenseVector(3)
	want.SetVec(0, 1)
	want.SetVec(1, 2)
	want.SetVec(2, 3)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(3),
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

func TestVector_Multiply(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 3)
		m.SetVec(1, 4)
		m.SetVec(2, 3)
	}

	want := GraphBLAS.NewDenseMatrix(2, 1)
	want.Set(0, 0, 27)
	want.Set(1, 0, 41)

	matrix := GraphBLAS.NewDenseMatrix(2, 3)
	matrix.Set(0, 0, 0)
	matrix.Set(0, 1, 3)
	matrix.Set(0, 2, 5)
	matrix.Set(1, 0, 5)
	matrix.Set(1, 1, 5)
	matrix.Set(1, 2, 2)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(3),
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

func TestVector_Add(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 3)
		m.SetVec(1, 8)
	}

	want := GraphBLAS.NewDenseMatrix(2, 1)
	want.Set(0, 0, 7)
	want.Set(1, 0, 8)

	matrix := GraphBLAS.NewDenseVector(2)
	matrix.SetVec(0, 4)
	matrix.SetVec(1, 0)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(2),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(2),
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

func TestVector_Subtract(t *testing.T) {

	setup := func(m GraphBLAS.Vector) {
		m.SetVec(0, 3)
		m.SetVec(1, 8)
	}

	want := GraphBLAS.NewDenseMatrix(2, 1)
	want.Set(0, 0, -1)
	want.Set(1, 0, 8)

	matrix := GraphBLAS.NewDenseMatrix(2, 1)
	matrix.Set(0, 0, 4)
	matrix.Set(1, 0, 0)

	tests := []struct {
		name string
		s    GraphBLAS.Vector
	}{
		{
			name: "DenseVector",
			s:    GraphBLAS.NewDenseVector(2),
		},
		{
			name: "SparseVector",
			s:    GraphBLAS.NewSparseVector(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := tt.s.Subtract(matrix)
			if !got.Equal(want) {
				t.Errorf("%+v Subtract = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}
