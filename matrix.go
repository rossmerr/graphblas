// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// Matrix interface
type Matrix interface {
	At(r, c int) (float64, error)
	Set(r, c int, value float64) error
	Update(r, c int, f func(float64) float64) error
	ColumnsAt(c int) (Vector, error)
	RowsAt(r int) (Vector, error)
	Columns() int
	Rows() int
	Iterator(i func(r, c int, v float64) bool) bool

	Copy() Matrix
	CopyArithmetic(i func(v float64) float64) Matrix

	Scalar(alpha float64) Matrix
	Multiply(m Matrix) (Matrix, error)
	Add(m Matrix) (Matrix, error)
	Subtract(m Matrix) (Matrix, error)
	Negative() Matrix
	Transpose() Matrix
	Equal(m Matrix) bool
	NotEqual(m Matrix) bool
}
