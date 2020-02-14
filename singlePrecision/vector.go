// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package singlePrecision

// Vector interface
type Vector interface {
	Matrix

	// AtVec returns the value of a vector element at i-th
	AtVec(i int) float32

	// SetVec sets the value at i-th of the vector
	SetVec(i int, value float32)

	// Length of the vector
	Length() int
}
