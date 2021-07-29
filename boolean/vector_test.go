// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolean_test

import (
	"testing"

	"github.com/rossmerr/graphblas/boolean"
)

func TestVector_Update(t *testing.T) {

	setup := func(m boolean.Vector) {
		m.SetVec(0, false)
		m.SetVec(1, true)
	}

	tests := []struct {
		name  string
		s     boolean.Vector
		want  bool
		value bool
	}{
		{
			name:  "DenseVector",
			s:     boolean.NewDenseVector(2),
			want:  true,
			value: true,
		},
		{
			name:  "SparseVector",
			s:     boolean.NewSparseVector(2),
			want:  true,
			value: true,
		},
		// Checks values get removed for sparse matrix
		{
			name:  "DenseVector",
			s:     boolean.NewDenseVector(2),
			want:  true,
			value: true,
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

func TestVector_ColumnsAt(t *testing.T) {

	setup := func(m boolean.Vector) {
		m.SetVec(0, true)
		m.SetVec(1, false)
	}

	want := boolean.NewDenseVector(2)
	want.SetVec(0, true)
	want.SetVec(1, false)

	tests := []struct {
		name string
		s    boolean.Vector
	}{
		{
			name: "DenseVector",
			s:    boolean.NewDenseVector(2),
		},
		{
			name: "SparseVector",
			s:    boolean.NewSparseVector(2),
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

	setup := func(m boolean.Vector) {
		m.SetVec(0, true)
		m.SetVec(1, false)
	}

	want := boolean.NewDenseVector(1)
	want.SetVec(0, true)

	tests := []struct {
		name string
		s    boolean.Vector
	}{
		{
			name: "DenseVector",
			s:    boolean.NewDenseVector(2),
		},
		{
			name: "SparseVector",
			s:    boolean.NewSparseVector(2),
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

func TestVector_Transpose(t *testing.T) {

	setup := func(m boolean.Vector) {
		m.SetVec(0, true)
		m.SetVec(1, false)
		m.SetVec(2, true)
	}

	want := boolean.NewDenseMatrix(1, 3)
	want.Set(0, 0, true)
	want.Set(0, 1, false)
	want.Set(0, 2, true)

	tests := []struct {
		name string
		s    boolean.Vector
	}{
		{
			name: "DenseVector",
			s:    boolean.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    boolean.NewSparseVector(3),
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

	setup := func(m boolean.Vector) {
		m.SetVec(0, true)
		m.SetVec(1, false)
		m.SetVec(2, true)
	}

	want := boolean.NewDenseVector(3)
	want.SetVec(0, true)
	want.SetVec(1, false)
	want.SetVec(2, true)

	tests := []struct {
		name string
		s    boolean.Vector
	}{
		{
			name: "DenseVector",
			s:    boolean.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    boolean.NewSparseVector(3),
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

	setup := func(m boolean.Vector) {
		m.SetVec(0, false)
		m.SetVec(1, false)
		m.SetVec(2, false)
	}

	want := boolean.NewDenseVector(3)
	want.SetVec(0, true)
	want.SetVec(1, true)
	want.SetVec(2, true)

	tests := []struct {
		name string
		s    boolean.Vector
	}{
		{
			name: "DenseVector",
			s:    boolean.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    boolean.NewSparseVector(3),
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

	setup := func(m boolean.Vector) {
		m.SetVec(0, true)
		m.SetVec(1, false)
		m.SetVec(2, true)
	}

	want := boolean.NewDenseVector(3)
	want.SetVec(0, true)
	want.SetVec(1, false)
	want.SetVec(2, true)

	tests := []struct {
		name string
		s    boolean.Vector
	}{
		{
			name: "DenseVector",
			s:    boolean.NewDenseVector(3),
		},
		{
			name: "SparseVector",
			s:    boolean.NewSparseVector(3),
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

func TestVector_Size(t *testing.T) {

	setup := func(m boolean.Vector) {
		m.SetVec(0, true)
		m.SetVec(1, true)
		m.SetVec(2, true)
		m.SetVec(3, true)
		m.SetVec(4, false)
		m.SetVec(5, true)
	}

	tests := []struct {
		name string
		s    boolean.Vector
		size int
	}{
		{
			name: "DenseVector",
			s:    boolean.NewDenseVector(6),
			size: 6,
		},
		{
			name: "SparseVector",
			s:    boolean.NewSparseVector(6),
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
