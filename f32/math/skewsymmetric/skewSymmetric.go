// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package skewsymmetric

import "github.com/rossmerr/graphblas/f32"

// SkewSymmetric (or antisymmetric or antimetric) matrix is a square matrix whose transpose equals its negative
func SkewSymmetric(s f32.Matrix) bool {
	r := s.Rows()
	c := s.Columns()
	if r != c {
		return false
	}

	t := s.Transpose()
	negativeTranspose := t.Negative()
	return negativeTranspose.Equal(s)
}
