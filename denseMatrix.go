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
	return newMatrix(r, c, nil)
}

func newMatrix(r, c int, initialise func([]float64, int)) *DenseMatrix {
	s := &DenseMatrix{data: make([][]float64, r), r: r, c: c}

	for i := 0; i < r; i++ {
		s.data[i] = make([]float64, c)
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
func (s *DenseMatrix) ColumnsAt(c int) ([]float64, error) {
	if c < 0 || c >= s.Columns() {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	columns := make([]float64, s.c)

	for r := 0; r < s.r; r++ {
		columns[r] = s.data[r][c]
	}

	return columns, nil
}

// RowsAt return the rows at r-th
func (s *DenseMatrix) RowsAt(r int) ([]float64, error) {
	if r < 0 || r >= s.Rows() {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	return s.data[r], nil
}

// Copy copies the matrix
func (s *DenseMatrix) Copy() Matrix {
	matrix := newMatrix(s.Rows(), s.Columns(), func(row []float64, r int) {
		for c := 0; c < s.Columns(); c++ {
			row[c] = s.data[r][c]
		}
	})

	return matrix
}

// Scalar multiplication of a matrix by alpha
func (s *DenseMatrix) Scalar(alpha float64) Matrix {
	matrix := newMatrix(s.Rows(), s.Columns(), func(rows []float64, r int) {
		for c := 0; c < s.Columns(); c++ {
			rows[c] = alpha * s.data[r][c]
		}
	})

	return matrix
}

// Multiply multiplies a matrix by another matrix
func (s *DenseMatrix) Multiply(m Matrix) (Matrix, error) {
	if s.Rows() != m.Columns() {
		return nil, fmt.Errorf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, r int) {
		for c := 0; c < m.Columns(); c++ {
			total := 0.0
			for k := 0; k < s.Columns(); k++ {
				v, _ := m.At(k, c)
				total += v * s.data[r][k]
			}
			row[c] = total
		}
	})

	return matrix, nil
}

// Add addition of a matrix by another matrix
func (s *DenseMatrix) Add(m Matrix) (Matrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, r int) {
		for c := 0; c < m.Columns(); c++ {
			v, _ := m.At(r, c)
			row[c] = s.data[r][c] + v
		}
	})

	return matrix, nil
}

// Subtract subtracts one matrix from another matrix
func (s *DenseMatrix) Subtract(m Matrix) (Matrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, r int) {
		for c := 0; c < m.Columns(); c++ {
			v, _ := m.At(r, c)
			row[c] = s.data[r][c] - v
		}
	})

	return matrix, nil
}

// Negative the negative of a matrix
func (s *DenseMatrix) Negative() Matrix {
	matrix := newMatrix(s.Rows(), s.Columns(), func(row []float64, r int) {
		for c := 0; c < s.Columns(); c++ {
			row[c] = -s.data[r][c]
		}
	})

	return matrix
}

// Transpose swaps the rows and columns
func (s *DenseMatrix) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), func(row []float64, c int) {
		for r := 0; r < s.Rows(); r++ {
			row[r] = s.data[r][c]
		}
	})

	return matrix
}
