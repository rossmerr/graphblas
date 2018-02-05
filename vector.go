// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// Vector interface
type Vector interface {
	Matrix

	// AtVec returns the value of a vector element at i-th
	AtVec(i int) (float64, error)

	// SetVec sets the value at i-th of the vector
	SetVec(i int, value float64) error

	// Length of the vector
	Length() int
}
