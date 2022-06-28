// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas

import "github.com/rossmerr/graphblas/constraints"

type MatrixCompressed[T constraints.Number] interface {
	UpdateReturnPointer(r, c int, value T) (pointer int, start int)
}

// Matrix interface
type Matrix[T constraints.Number] interface {
	Mask

	// At returns the value of a matrix element at r-th, c-th
	At(r, c int) T

	// Set sets the value at r-th, c-th of the matrix
	Set(r, c int, value T)

	// Update does a At and Set on the matrix element at r-th, c-th
	Update(r, c int, f func(T) T)

	// ColumnsAt return the columns at c-th
	ColumnsAt(c int) Vector[T]

	// RowsAt return the rows at r-th
	RowsAt(r int) Vector[T]

	// RowsAtToArray return the rows at r-th
	RowsAtToArray(r int) []T

	// Copy copies the matrix
	Copy() Matrix[T]

	// Enumerate iterates through all non-zero elements, order is not guaranteed
	Enumerate() Enumerate[T]

	// Map iterates and replace each element with the result of applying a function to its value
	Map() Map[T]

	// Scalar multiplication of a matrix by alpha
	Scalar(alpha T) Matrix[T]

	// Multiply multiplies a matrix by another matrix
	//  C = AB
	Multiply(m Matrix[T]) Matrix[T]

	// Add addition of a matrix by another matrix
	Add(m Matrix[T]) Matrix[T]

	// Subtract subtracts one matrix from another matrix
	Subtract(m Matrix[T]) Matrix[T]

	// Negative the negative of a matrix
	Negative() Matrix[T]

	// Transpose swaps the rows and columns
	//  C ⊕= Aᵀ
	Transpose() Matrix[T]

	// Equal the two matrices are equal
	Equal(m Matrix[T]) bool

	// NotEqual the two matrices are not equal
	NotEqual(m Matrix[T]) bool

	// Size of the matrix
	Size() int

	// The number of elements in the matrix (non-zero counted for dense matrices)
	Values() int

	// Clear removes all elements from a matrix
	Clear()
}
