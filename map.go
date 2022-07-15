// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas

import "github.com/rossmerr/graphblas/constraints"

// Map replace each element with the result of applying a function to its value
type Map[T constraints.Scaler] interface {
	// HasNext checks for the next element in the matrix
	HasNext() bool

	// Map move the iterator and uses a higher order function to changes the elements current value
	Map(func(r, c int, v T) T)
}
