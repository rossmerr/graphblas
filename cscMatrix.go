// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"fmt"
)

// CSCMatrix compressed storage by columns (CSC)
type CSCMatrix struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []float64
	rows     []int
	colStart []int
}

// NewCSCMatrix returns a GraphBLAS.CSCMatrix
func NewCSCMatrix(r, c int) *CSCMatrix {
	return newCSCMatrix(r, c, 0)
}

func newCSCMatrix(r, c, l int) *CSCMatrix {
	s := &CSCMatrix{
		r:        r,
		c:        c,
		values:   make([]float64, l),
		rows:     make([]int, l),
		colStart: make([]int, c+1),
	}

	return s
}

// Columns the number of columns of the matrix
func (s *CSCMatrix) Columns() int {
	return s.c
}

// Rows the number of rows of the matrix
func (s *CSCMatrix) Rows() int {
	return s.r
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *CSCMatrix) Update(r, c int, f func(float64) float64) error {
	if r < 0 || r >= s.r {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.rowIndex(r, c)

	if pointerStart < pointerEnd && s.rows[pointerStart] == r {
		value := f(s.values[pointerStart])
		if value == 0 {
			s.remove(pointerStart, c)
		} else {
			s.values[pointerStart] = value
		}
	} else {
		s.insert(pointerStart, r, c, f(0))
	}

	return nil
}

// At returns the value of a matrix element at r-th, c-th
func (s *CSCMatrix) At(r, c int) (float64, error) {
	value := 0.0
	err := s.Update(r, c, func(v float64) float64 {
		value = v
		return v
	})

	return value, err
}

// Set sets the value at r-th, c-th of the matrix
func (s *CSCMatrix) Set(r, c int, value float64) error {
	return s.Update(r, c, func(v float64) float64 {
		return value
	})
}

// ColumnsAt return the columns at c-th
func (s *CSCMatrix) ColumnsAt(c int) (Vector, error) {
	if c < 0 || c >= s.c {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	start := s.colStart[c]
	end := s.colStart[c+1]
	columns := NewSparseVector(s.r)

	for i := start; i < end; i++ {
		columns.Set(s.rows[i], s.values[i])
	}

	return columns, nil
}

// RowsAt return the rows at r-th
func (s *CSCMatrix) RowsAt(r int) (Vector, error) {
	if r < 0 || r >= s.r {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector(s.c)

	for c := range s.colStart[:s.c] {
		pointerStart, _ := s.rowIndex(r, c)
		rows.Set(c, s.values[pointerStart])
	}

	return rows, nil
}

func (s *CSCMatrix) insert(pointer, r, c int, value float64) {
	if value == 0 {
		return
	}

	s.rows = append(s.rows[:pointer], append([]int{r}, s.rows[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]float64{value}, s.values[pointer:]...)...)

	for i := c + 1; i <= s.c; i++ {
		s.colStart[i]++
	}
}

func (s *CSCMatrix) remove(pointer, c int) {
	s.rows = append(s.rows[:pointer], s.rows[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)

	for i := c + 1; i <= s.c; i++ {
		s.colStart[i]--
	}
}

func (s *CSCMatrix) rowIndex(r, c int) (int, int) {

	start := s.colStart[c]
	end := s.colStart[c+1]

	if start-end == 0 {
		return start, end
	}

	if r > s.rows[end-1] {
		return end, end
	}

	for start < end {
		p := (start + end) / 2
		if s.rows[p] > r {
			end = p
		} else if s.rows[p] < r {
			start = p + 1
		} else {
			return p, end
		}
	}

	return start, end
}

// Copy copies the matrix
func (s *CSCMatrix) Copy() Matrix {
	return s.copy(func(value float64) float64 {
		return value
	})
}

func (s *CSCMatrix) copy(action func(float64) float64) *CSCMatrix {
	matrix := newCSCMatrix(s.r, s.c, len(s.values))

	for i := range s.values {
		matrix.values[i] = action(s.values[i])
		matrix.rows[i] = s.rows[i]
	}

	for i := range s.colStart {
		matrix.colStart[i] = s.colStart[i]
	}

	return matrix
}

// Scalar multiplication of a matrix by alpha
func (s *CSCMatrix) Scalar(alpha float64) Matrix {
	return s.copy(func(value float64) float64 {
		return alpha * value
	})
}

// Multiply multiplies a matrix by another matrix
func (s *CSCMatrix) Multiply(m Matrix) (Matrix, error) {
	if s.Rows() != m.Columns() {
		return nil, fmt.Errorf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	matrix := newCSCMatrix(s.Rows(), m.Columns(), 0)

	for r := 0; r < s.Rows(); r++ {
		rows, _ := s.RowsAt(r)

		for c := 0; c < m.Columns(); c++ {
			column, _ := m.ColumnsAt(c)

			sum := 0.0
			for l := 0; l < rows.Length(); l++ {
				vC, _ := column.At(l)
				vR, _ := rows.At(l)
				sum += vR * vC
			}

			matrix.Set(r, c, sum)
		}

	}

	return matrix, nil
}

// Add addition of a matrix by another matrix
func (s *CSCMatrix) Add(m Matrix) (Matrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := m.Copy()

	s.iterator(func(r, c int, value float64) {
		matrix.Update(r, c, func(v float64) float64 {
			return value + v
		})
	})

	return matrix, nil
}

// Subtract subtracts one matrix from another matrix
func (s *CSCMatrix) Subtract(m Matrix) (Matrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := m.Copy()

	s.iterator(func(r, c int, value float64) {
		matrix.Update(r, c, func(v float64) float64 {
			return value - v
		})
	})

	return matrix, nil
}

// Negative the negative of a matrix
func (s *CSCMatrix) Negative() Matrix {
	return s.copy(func(value float64) float64 {
		return -value
	})
}

// Transpose swaps the rows and columns
func (s *CSCMatrix) Transpose() Matrix {
	matrix := newCSCMatrix(s.c, s.r, len(s.values))

	s.iterator(func(r, c int, value float64) {
		matrix.Set(c, r, value)
	})

	return matrix
}

func (s *CSCMatrix) iterator(i func(r, c int, v float64)) bool {
	for c := 0; c < s.Columns(); c++ {
		pointerStart := s.colStart[c]
		pointerEnd := s.colStart[c+1]

		for r := pointerStart; r < pointerEnd; r++ {
			i(s.rows[r], c, s.values[r])
		}
	}

	return false
}
