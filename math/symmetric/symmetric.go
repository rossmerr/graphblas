// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package symmetric

import (
	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/constraints"
)

// Symmetric matrix is a square matrix that is equal to its transpose
func Symmetric[T constraints.Number](s graphblas.Matrix[T]) bool {
	r := s.Rows()
	c := s.Columns()
	if r != c {
		return false
	}

	t := s.Transpose()
	return t.Equal(s)
}
