// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package symmetric

import "github.com/rossmerr/graphblas/f64"

// Symmetric matrix is a square matrix that is equal to its transpose
func Symmetric(s f64.Matrix) bool {
	r := s.Rows()
	c := s.Columns()
	if r != c {
		return false
	}

	t := s.Transpose()
	return t.Equal(s)
}
