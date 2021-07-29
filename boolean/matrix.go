// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolean

import GraphBLAS "github.com/rossmerr/graphblas"

// Matrix interface
type Matrix interface {
	GraphBLAS.Mask

	// At returns the value of a matrix element at r-th, c-th
	At(r, c int) bool

	// Set sets the value at r-th, c-th of the matrix
	Set(r, c int, value bool)

	// Update does a At and Set on the matrix element at r-th, c-th
	Update(r, c int, f func(bool) bool)

	// ColumnsAt return the columns at c-th
	ColumnsAt(c int) Vector

	// RowsAt return the rows at r-th
	RowsAt(r int) Vector

	// RowsAtToArray return the rows at r-th
	RowsAtToArray(r int) []bool

	// Copy copies the matrix
	Copy() Matrix

	// Enumerate iterates through all non-zero elements, order is not guaranteed
	Enumerate() Enumerate

	// Map iterates and replace each element with the result of applying a function to its value
	Map() Map

	// Transpose swaps the rows and columns
	//  C ⊕= Aᵀ
	Transpose() Matrix

	// Equal the two matrices are equal
	Equal(m Matrix) bool

	// NotEqual the two matrices are not equal
	NotEqual(m Matrix) bool

	// Size of the matrix
	Size() int

	// The number of elements in the matrix (non-zero counted for dense matrices)
	Values() int

	// Clear removes all elements from a matrix
	Clear()
}
