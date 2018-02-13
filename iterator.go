// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// Iterator iterates over the matrix
type Iterator interface {
	// HasNext checks for the next element in the matrix
	HasNext() bool

	// Next move the iterator over the matrix
	Next() (r, c int, v float64)

	// Update changes the elements current value
	Update(v float64)
}
