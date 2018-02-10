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

	// Columns the number of columns of the matrix
	Columns() int

	// Rows the number of rows of the matrix
	Rows() int

	// Copy copies the matrix
	Copy() Matrix

	// CopyArithmetic copies the matrix and applies a arithmetic function through all non-zero elements, order is not guaranteed
	CopyArithmetic(i func(v float64) float64) Matrix

	// Iterator iterates through all non-zero elements, order is not guaranteed
	Iterator(i func(r, c int, v float64) bool) bool

	// Scalar multiplication of a matrix by alpha
	Scalar(alpha float64) Matrix

	// Multiply multiplies a matrix by another matrix
	Multiply(m Matrix) Matrix

	// Add addition of a matrix by another matrix
	Add(m Matrix) Matrix

	// Subtract subtracts one matrix from another matrix
	Subtract(m Matrix) Matrix

	// Negative the negative of a matrix
	Negative() Matrix

	// Transpose swaps the rows and columns
	Transpose() Matrix

	// Equal the two matrices are equal
	Equal(m Matrix) bool

	// NotEqual the two matrices are not equal
	NotEqual(m Matrix) bool

	// The number of elements in the matrix (non-zero only for sparse matrices)
	Size() int
}
