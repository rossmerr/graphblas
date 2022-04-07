// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package skewsymmetric

import (
	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/constraints"
)

// SkewSymmetric (or antisymmetric or antimetric) matrix is a square matrix whose transpose equals its negative
func SkewSymmetric[T constraints.Number](s graphblas.Matrix[T]) bool {
	r := s.Rows()
	c := s.Columns()
	if r != c {
		return false
	}

	t := s.Transpose()
	negativeTranspose := t.Negative()
	return negativeTranspose.Equal(s)
}
