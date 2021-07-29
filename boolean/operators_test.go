// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolean_test

import (
	"testing"

	"github.com/rossmerr/graphblas/boolean"

	"golang.org/x/net/context"
)

func TestMatrix_Transpose_To_CSR(t *testing.T) {

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, false)
		m.Set(0, 2, true)
		m.Set(1, 0, false)
		m.Set(1, 1, false)
		m.Set(1, 2, true)
	}

	want := boolean.NewDenseMatrix(3, 2)
	want.Set(0, 0, true)
	want.Set(0, 1, false)
	want.Set(1, 0, false)
	want.Set(1, 1, false)
	want.Set(2, 0, true)
	want.Set(2, 1, true)

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
			got := boolean.TransposeToCSR(context.Background(), tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Transpose_To_CSC(t *testing.T) {

	setup := func(m boolean.Matrix) {
		m.Set(0, 0, true)
		m.Set(0, 1, false)
		m.Set(0, 2, true)
		m.Set(1, 0, false)
		m.Set(1, 1, false)
		m.Set(1, 2, true)
	}

	want := boolean.NewDenseMatrix(3, 2)
	want.Set(0, 0, true)
	want.Set(0, 1, false)
	want.Set(1, 0, false)
	want.Set(1, 1, false)
	want.Set(2, 0, true)
	want.Set(2, 1, true)

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
			got := boolean.TransposeToCSC(context.Background(), tt.s)
			if !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}
