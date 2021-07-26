// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package strassen_test

import (
	"testing"

	"context"

	"github.com/rossmerr/graphblas/f32"
	"github.com/rossmerr/graphblas/f32/math/strassen"
)

func TestMatrix_Multiply(t *testing.T) {
	setup := func(m f32.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(0, 3, 4)
		m.Set(1, 0, 1)
		m.Set(1, 1, 2)
		m.Set(1, 2, 3)
		m.Set(1, 3, 4)
		m.Set(2, 0, 1)
		m.Set(2, 1, 2)
		m.Set(2, 2, 3)
		m.Set(2, 3, 4)
		m.Set(3, 0, 1)
		m.Set(3, 1, 2)
		m.Set(3, 2, 3)
		m.Set(3, 3, 4)
	}

	matrix := f32.NewDenseMatrix(4, 4)
	matrix.Set(0, 0, 1)
	matrix.Set(0, 1, 2)
	matrix.Set(0, 2, 3)
	matrix.Set(0, 3, 4)
	matrix.Set(1, 0, 1)
	matrix.Set(1, 1, 2)
	matrix.Set(1, 2, 3)
	matrix.Set(1, 3, 4)
	matrix.Set(2, 0, 1)
	matrix.Set(2, 1, 2)
	matrix.Set(2, 2, 3)
	matrix.Set(2, 3, 4)
	matrix.Set(3, 0, 1)
	matrix.Set(3, 1, 2)
	matrix.Set(3, 2, 3)
	matrix.Set(3, 3, 4)

	want := f32.NewDenseMatrix(4, 4)
	want.Set(0, 0, 10)
	want.Set(0, 1, 20)
	want.Set(0, 2, 30)
	want.Set(0, 3, 40)
	want.Set(1, 0, 10)
	want.Set(1, 1, 20)
	want.Set(1, 2, 30)
	want.Set(1, 3, 40)
	want.Set(2, 0, 10)
	want.Set(2, 1, 20)
	want.Set(2, 2, 30)
	want.Set(2, 3, 40)
	want.Set(3, 0, 10)
	want.Set(3, 1, 20)
	want.Set(3, 2, 30)
	want.Set(3, 3, 40)

	tests := []struct {
		name string
		s    f32.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    f32.NewDenseMatrix(4, 4),
		},
		{
			name: "CSCMatrix",
			s:    f32.NewCSCMatrix(4, 4),
		},
		{
			name: "CSRMatrix",
			s:    f32.NewCSRMatrix(4, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := strassen.MultiplyCrossoverPoint(context.Background(), tt.s, matrix, 2); !got.Equal(want) {
				t.Errorf("%+v Multiply = got %+v, want %+v", tt.name, got, want)
			}
		})
	}
}
