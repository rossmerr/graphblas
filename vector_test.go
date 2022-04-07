// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas_test

import (
	"testing"

	"github.com/rossmerr/graphblas"
)

func TestVector_Update(t *testing.T) {

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 3)
		m.SetVec(1, 8)
	}

	tests := []struct {
		name  string
		s     graphblas.Vector[float64]
		want  float64
		value float64
	}{
		{
			name:  "DenseVector",
			s:     graphblas.NewDenseVector[float64](2),
			want:  2,
			value: 2,
		},
		{
			name:  "SparseVector",
			s:     graphblas.NewSparseVector[float64](2),
			want:  2,
			value: 2,
		},
		// Checks values get removed for sparse matrix
		{
			name:  "DenseVector",
			s:     graphblas.NewDenseVector[float64](2),
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

func TestVector_ColumnsAt(t *testing.T) {

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 4)
		m.SetVec(1, 0)
	}

	want := graphblas.NewDenseVector[float64](2)
	want.SetVec(0, 4)
	want.SetVec(1, 0)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](2),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](2),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 4)
		m.SetVec(1, 0)
	}

	want := graphblas.NewDenseVector[float64](1)
	want.SetVec(0, 4)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](2),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](2),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 4)
		m.SetVec(1, 0)
	}

	want := graphblas.NewDenseVector[float64](2)
	want.SetVec(0, 8)
	want.SetVec(1, 0)

	tests := []struct {
		name  string
		s     graphblas.Vector[float64]
		alpha float64
	}{
		{
			name:  "DenseVector",
			s:     graphblas.NewDenseVector[float64](2),
			alpha: 2,
		},
		{
			name:  "SparseVector",
			s:     graphblas.NewSparseVector[float64](2),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 2)
		m.SetVec(1, -4)
	}

	want := graphblas.NewDenseVector[float64](2)
	want.SetVec(0, -2)
	want.SetVec(1, 4)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](2),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](2),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 6)
		m.SetVec(1, 4)
		m.SetVec(2, 24)
	}

	want := graphblas.NewDenseMatrix[float64](1, 3)
	want.Set(0, 0, 6)
	want.Set(0, 1, 4)
	want.Set(0, 2, 24)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](3),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](3),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 1)
		m.SetVec(1, 2)
		m.SetVec(2, 3)
	}

	want := graphblas.NewDenseVector[float64](3)
	want.SetVec(0, 1)
	want.SetVec(1, 2)
	want.SetVec(2, 3)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](3),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](3),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 1)
		m.SetVec(1, 2)
		m.SetVec(2, 3)
	}

	want := graphblas.NewDenseVector[float64](3)
	want.SetVec(0, 4)
	want.SetVec(1, 5)
	want.SetVec(2, 6)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](3),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](3),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 1)
		m.SetVec(1, 2)
		m.SetVec(2, 3)
	}

	want := graphblas.NewDenseVector[float64](3)
	want.SetVec(0, 1)
	want.SetVec(1, 2)
	want.SetVec(2, 3)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](3),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](3),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 3)
		m.SetVec(1, 4)
		m.SetVec(2, 3)
	}

	want := graphblas.NewDenseMatrix[float64](2, 1)
	want.Set(0, 0, 27)
	want.Set(1, 0, 41)

	matrix := graphblas.NewDenseMatrix[float64](2, 3)
	matrix.Set(0, 0, 0)
	matrix.Set(0, 1, 3)
	matrix.Set(0, 2, 5)
	matrix.Set(1, 0, 5)
	matrix.Set(1, 1, 5)
	matrix.Set(1, 2, 2)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](3),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			got := matrix.Multiply(tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Multiply = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestVector_Add(t *testing.T) {

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 3)
		m.SetVec(1, 8)
	}

	want := graphblas.NewDenseMatrix[float64](2, 1)
	want.Set(0, 0, 7)
	want.Set(1, 0, 8)

	matrix := graphblas.NewDenseVector[float64](2)
	matrix.SetVec(0, 4)
	matrix.SetVec(1, 0)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](2),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](2),
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

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 3)
		m.SetVec(1, 8)
	}

	want := graphblas.NewDenseMatrix[float64](2, 1)
	want.Set(0, 0, -1)
	want.Set(1, 0, 8)

	matrix := graphblas.NewDenseMatrix[float64](2, 1)
	matrix.Set(0, 0, 4)
	matrix.Set(1, 0, 0)

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](2),
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](2),
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

func TestVector_Size(t *testing.T) {

	setup := func(m graphblas.Vector[float64]) {
		m.SetVec(0, 6)
		m.SetVec(1, 4)
		m.SetVec(2, 24)
		m.SetVec(3, 1)
		m.SetVec(4, 0)
		m.SetVec(5, 8)
	}

	tests := []struct {
		name string
		s    graphblas.Vector[float64]
		size int
	}{
		{
			name: "DenseVector",
			s:    graphblas.NewDenseVector[float64](6),
			size: 6,
		},
		{
			name: "SparseVector",
			s:    graphblas.NewSparseVector[float64](6),
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
