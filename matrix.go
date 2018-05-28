// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// Matrix interface
type Matrix interface {
	// At returns the value of a matrix element at r-th, c-th
	At(r, c int) float64

	// Set sets the value at r-th, c-th of the matrix
	Set(r, c int, value float64)

	// Update does a At and Set on the matrix element at r-th, c-th
	Update(r, c int, f func(float64) float64)

	// ColumnsAt return the columns at c-th
	ColumnsAt(c int) Vector

	// RowsAt return the rows at r-th
	RowsAt(r int) Vector

	// RowsAtToArray return the rows at r-th
	RowsAtToArray(r int) []float64

	// Columns the number of columns of the matrix
	Columns() int

	// Rows the number of rows of the matrix
	Rows() int

	// Copy copies the matrix
	Copy() Matrix

	// Enumerate iterates through all non-zero elements, order is not guaranteed
	Enumerate() Enumerate

	// Map iterates and replace each element with the result of applying a function to its value
	Map() Map

	// Scalar multiplication of a matrix by alpha
	Scalar(alpha float64) Matrix

	// Multiply multiplies a matrix by another matrix
	// C = AB
	Multiply(m Matrix) Matrix

	// Add addition of a matrix by another matrix
	Add(m Matrix) Matrix

	// Subtract subtracts one matrix from another matrix
	Subtract(m Matrix) Matrix

	// Negative the negative of a matrix
	Negative() Matrix

	// Transpose swaps the rows and columns
	// C ⊕= Aᵀ
	Transpose() Matrix

	// Equal the two matrices are equal
	Equal(m Matrix) bool

	// NotEqual the two matrices are not equal
	NotEqual(m Matrix) bool

	// Size of the matrix
	Size() int

	// The number of elements in the matrix (non-zero counted for dense matrices)
	Values() int
}
