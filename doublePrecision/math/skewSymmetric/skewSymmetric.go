// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package skewSymmetric

import "github.com/RossMerr/Caudex.GraphBLAS/doublePrecision"

// SkewSymmetric (or antisymmetric or antimetric) matrix is a square matrix whose transpose equals its negative
func SkewSymmetric(s doublePrecision.Matrix) bool {
	r := s.Rows()
	c := s.Columns()
	if r != c {
		return false
	}

	t := s.Transpose()
	negativeTranspose := t.Negative()
	return negativeTranspose.Equal(s)
}
