// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package skewsymmetric_test

import (
	"testing"

	"github.com/rossmerr/graphblas/f64"
	"github.com/rossmerr/graphblas/f64/math/skewsymmetric"
)

func TestMatrix_SkewSymmetric(t *testing.T) {

	setup := func(m f64.Matrix) {
		m.Set(0, 0, 0)
		m.Set(0, 1, 2)
		m.Set(0, 2, -1)
		m.Set(1, 0, -2)
		m.Set(1, 1, 0)
		m.Set(1, 2, -4)
		m.Set(2, 0, 1)
		m.Set(2, 1, 4)
		m.Set(2, 2, 0)
	}

	tests := []struct {
		name string
		s    f64.Matrix
		want bool
	}{
		{
			name: "DenseMatrix",
			s:    f64.NewDenseMatrix(3, 3),
			want: true,
		},
		{
			name: "CSCMatrix",
			s:    f64.NewCSCMatrix(3, 3),
			want: true,
		},
		{
			name: "CSRMatrix",
			s:    f64.NewCSRMatrix(3, 3),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := skewsymmetric.SkewSymmetric(tt.s); got != tt.want {
				t.Errorf("%+v SkewSymmetric = %+v, want %+v", tt.name, got, tt.want)
			}
		})
	}
}
