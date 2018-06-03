// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// Mask is used to control how computed values are stored in the output from a method
type Mask interface {
	// Columns the number of columns of the mask
	Columns() int

	// Rows the number of rows of the mask
	Rows() int

	// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
	Element(r, c int) bool

	// EnumerateMask iterates through all non-zero elements, order is not guaranteed
	EnumerateMask() EnumerateMask
}

// maskEnumerate iterates over the mask
type EnumerateMask interface {
	// HasNext checks for the next element in the matrix
	HasNext() bool

	// Next move the iterator over the matrix
	Next() (r, c int, v bool)
}
