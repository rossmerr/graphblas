// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import "fmt"

// DenseMatrix a dense matrix
type DenseMatrix struct {
	c    int // number of rows in the sparse matrix
	r    int // number of columns in the sparse matrix
	data [][]float64
}

// NewDenseMatrix returns a GraphBLAS.DenseMatrix
func NewDenseMatrix(r, c int) *DenseMatrix {
	return newMatrix(r, c, nil, nil)
}

// NewDenseMatrixFromArray returns a GraphBLAS.DenseMatrix
func NewDenseMatrixFromArray(r, c int, data [][]float64) *DenseMatrix {
	return newMatrix(r, c, data, nil)
}

func newMatrix(r, c int, data [][]float64, initialise func([]float64, int)) *DenseMatrix {
	s := &DenseMatrix{data: make([][]float64, r), r: r, c: c}

	for i := 0; i < r; i++ {
		s.data[i] = make([]float64, c)

		if data != nil {
			for k := 0; k < c; k++ {
				s.data[i][k] = data[i][k]
			}
		}

		if initialise != nil {
			initialise(s.data[i], i)
		}
	}

	return s
}

// Columns the number of columns of the matrix
func (s *DenseMatrix) Columns() int {
	return s.c
}

// Rows the number of rows of the matrix
func (s *DenseMatrix) Rows() int {
	return s.r
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *DenseMatrix) Update(r, c int, f func(float64) float64) error {
	if r < 0 || r >= s.Rows() {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	s.data[r][c] = f(s.data[r][c])

	return nil
}

// At returns the value of a matrix element at r-th, c-th
func (s *DenseMatrix) At(r, c int) (float64, error) {
	if r < 0 || r >= s.Rows() {
		return 0, fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		return 0, fmt.Errorf("Column '%+v' is invalid", c)
	}

	return s.data[r][c], nil
}

// Set sets the value at r-th, c-th of the matrix
func (s *DenseMatrix) Set(r, c int, value float64) error {
	if r < 0 || r >= s.Rows() {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	s.data[r][c] = value

	return nil
}

// ColumnsAt return the columns at c-th
func (s *DenseMatrix) ColumnsAt(c int) (Vector, error) {
	if c < 0 || c >= s.Columns() {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	columns := NewSparseVector(s.r)

	for r := 0; r < s.r; r++ {
		columns.SetVec(r, s.data[r][c])
	}

	return columns, nil
}

// RowsAt return the rows at r-th
func (s *DenseMatrix) RowsAt(r int) (Vector, error) {
	if r < 0 || r >= s.Rows() {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector(s.c)
	for i := 0; i < s.c; i++ {
		rows.SetVec(i, s.data[r][i])
	}

	return rows, nil
}

// Copy copies the matrix
func (s *DenseMatrix) Copy() Matrix {
	return s.CopyArithmetic(func(value float64) float64 {
		return value
	})
}

// CopyArithmetic copies the matrix and applies a arithmetic function through all non-zero elements, order is not guaranteed
func (s *DenseMatrix) CopyArithmetic(action func(float64) float64) Matrix {
	v := 0.0
	matrix := newMatrix(s.Rows(), s.Columns(), nil, func(row []float64, r int) {
		for c := 0; c < s.Columns(); c++ {
			v = s.data[r][c]
			if v != 0.0 {
				row[c] = action(v)
			} else {
				row[c] = v
			}
		}
	})

	return matrix
}

// Iterator iterates through all non-zero elements, order is not guaranteed
func (s *DenseMatrix) Iterator(i func(r, c int, v float64) bool) bool {
	for c := 0; c < s.Columns(); c++ {
		for r := 0; r < s.Rows(); r++ {
			v, _ := s.At(r, c)
			if v != 0.0 {
				if i(r, c, v) == false {
					return false
				}
			}
		}
	}

	return true
}

// Scalar multiplication of a matrix by alpha
func (s *DenseMatrix) Scalar(alpha float64) Matrix {
	return scalar(s, alpha)
}

// Multiply multiplies a matrix by another matrix
func (s *DenseMatrix) Multiply(m Matrix) (Matrix, error) {
	return multiply(s, m)
}

// Add addition of a matrix by another matrix
func (s *DenseMatrix) Add(m Matrix) (Matrix, error) {
	return add(s, m)
}

// Subtract subtracts one matrix from another matrix
func (s *DenseMatrix) Subtract(m Matrix) (Matrix, error) {
	return subtract(s, m)
}

// Negative the negative of a matrix
func (s *DenseMatrix) Negative() Matrix {
	return negative(s)
}

// Transpose swaps the rows and columns
func (s *DenseMatrix) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), nil, nil)

	return transpose(s, matrix)
}

// Equal the two matrices are equal
func (s *DenseMatrix) Equal(m Matrix) bool {
	return equal(s, m)
}

// NotEqual the two matrices are not equal
func (s *DenseMatrix) NotEqual(m Matrix) bool {
	return notEqual(s, m)
}
