// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import "fmt"

// CSRMatrix compressed storage by rows (CSR)
type CSRMatrix struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []float64
	cols     []int
	rowStart []int
}

// NewCSRMatrix returns a GraphBLAS.CSRMatrix
func NewCSRMatrix(r, c int) *CSRMatrix {
	return newCSRMatrix(r, c, nil, 0)
}

// NewCSRMatrixFromArray returns a GraphBLAS.CSRMatrix
func NewCSRMatrixFromArray(data [][]float64) *CSRMatrix {
	r := len(data)
	c := len(data[0])
	return newCSRMatrix(r, c, data, 0)
}

func newCSRMatrix(r, c int, data [][]float64, l int) *CSRMatrix {

	s := &CSRMatrix{
		r:        r,
		c:        c,
		values:   make([]float64, l),
		cols:     make([]int, l),
		rowStart: make([]int, r+1),
	}

	if data != nil {
		for i := 0; i < r; i++ {
			for k := 0; k < c; k++ {
				s.Set(i, k, data[i][k])
			}
		}
	}

	return s
}

// Columns the number of columns of the matrix
func (s *CSRMatrix) Columns() int {
	return s.c
}

// Rows the number of rows of the matrix
func (s *CSRMatrix) Rows() int {
	return s.r
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *CSRMatrix) Update(r, c int, f func(float64) float64) error {
	if r < 0 || r >= s.r {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.columnIndex(r, c)

	if pointerStart < pointerEnd && s.cols[pointerStart] == c {
		value := f(s.values[pointerStart])
		if value == 0 {
			s.remove(pointerStart, r)
		} else {
			s.values[pointerStart] = value
		}
	} else {
		s.insert(pointerStart, r, c, f(0))
	}

	return nil
}

// At returns the value of a matrix element at r-th, c-th
func (s *CSRMatrix) At(r, c int) (float64, error) {
	value := 0.0
	err := s.Update(r, c, func(v float64) float64 {
		value = v
		return v
	})

	return value, err
}

// Set sets the value at r-th, c-th of the matrix
func (s *CSRMatrix) Set(r, c int, value float64) error {
	return s.Update(r, c, func(v float64) float64 {
		return value
	})
}

// ColumnsAt return the columns at c-th
func (s *CSRMatrix) ColumnsAt(c int) (Vector, error) {
	if c < 0 || c >= s.c {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	columns := NewSparseVector(s.r)

	for r := range s.rowStart[:s.r] {
		pointerStart, pointerEnd := s.columnIndex(r, c)
		if pointerStart < pointerEnd && s.cols[pointerStart] == c {
			columns.SetVec(r, s.values[pointerStart])
		}
	}

	return columns, nil

}

// RowsAt return the rows at r-th
func (s *CSRMatrix) RowsAt(r int) (Vector, error) {
	if r < 0 || r >= s.r {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	start := s.rowStart[r]
	end := s.rowStart[r+1]

	rows := NewSparseVector(s.c)
	for i := start; i < end; i++ {
		rows.SetVec(s.cols[i], s.values[i])
	}

	return rows, nil
}

func (s *CSRMatrix) insert(pointer, r, c int, value float64) {
	if value == 0 {
		return
	}

	s.cols = append(s.cols[:pointer], append([]int{c}, s.cols[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]float64{value}, s.values[pointer:]...)...)

	for i := r + 1; i <= s.r; i++ {
		s.rowStart[i]++
	}
}

func (s *CSRMatrix) remove(pointer, r int) {
	s.cols = append(s.cols[:pointer], s.cols[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)

	for i := r + 1; i <= s.r; i++ {
		s.rowStart[i]--
	}
}

func (s *CSRMatrix) columnIndex(r, c int) (int, int) {

	start := s.rowStart[r]
	end := s.rowStart[r+1]

	if start-end == 0 {
		return start, end
	}

	if c > s.cols[end-1] {
		return end, end
	}

	for start < end {
		p := (start + end) / 2
		if s.cols[p] > c {
			end = p
		} else if s.cols[p] < c {
			start = p + 1
		} else {
			return p, end
		}
	}

	return start, end
}

// Copy copies the matrix
func (s *CSRMatrix) Copy() Matrix {
	return s.CopyArithmetic(func(value float64) float64 {
		return value
	})
}

// CopyArithmetic copies the matrix and applies a arithmetic function through all non-zero elements, order is not guaranteed
func (s *CSRMatrix) CopyArithmetic(action func(float64) float64) Matrix {
	matrix := newCSRMatrix(s.r, s.c, nil, len(s.values))

	for i := range s.values {
		matrix.values[i] = action(s.values[i])
		matrix.cols[i] = s.cols[i]
	}

	for i := range s.rowStart {
		matrix.rowStart[i] = s.rowStart[i]
	}

	return matrix
}

// Iterator iterates through all non-zero elements, order is not guaranteed
func (s *CSRMatrix) Iterator(i func(r, c int, v float64) bool) bool {
	for r := 0; r < s.Rows(); r++ {
		pointerStart := s.rowStart[r]
		pointerEnd := s.rowStart[r+1]

		for c := pointerStart; c < pointerEnd; c++ {
			if i(r, s.cols[c], s.values[c]) == false {
				return false
			}
		}
	}

	return true
}

// Scalar multiplication of a matrix by alpha
func (s *CSRMatrix) Scalar(alpha float64) Matrix {
	return scalar(s, alpha)
}

// Multiply multiplies a matrix by another matrix
func (s *CSRMatrix) Multiply(m Matrix) (Matrix, error) {
	if s.Rows() != m.Columns() {
		return nil, fmt.Errorf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	matrix := newCSRMatrix(s.Rows(), m.Columns(), nil, 0)

	return multiply(s, m, matrix)
}

// Add addition of a matrix by another matrix
func (s *CSRMatrix) Add(m Matrix) (Matrix, error) {
	return add(s, m)
}

// Subtract subtracts one matrix from another matrix
func (s *CSRMatrix) Subtract(m Matrix) (Matrix, error) {
	return subtract(s, m)
}

// Negative the negative of a matrix
func (s *CSRMatrix) Negative() Matrix {
	return negative(s)
}

// Transpose swaps the rows and columns
func (s *CSRMatrix) Transpose() Matrix {
	matrix := newCSRMatrix(s.c, s.r, nil, len(s.values))

	return transpose(s, matrix)
}

// Equal the two matrices are equal
func (s *CSRMatrix) Equal(m Matrix) bool {
	return equal(s, m)
}

// NotEqual the two matrices are not equal
func (s *CSRMatrix) NotEqual(m Matrix) bool {
	return notEqual(s, m)
}
