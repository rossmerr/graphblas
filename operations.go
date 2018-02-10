// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"log"
)

// Multiply multiplies a matrix by another matrix
func Multiply(s, m, matrix Matrix) Matrix {
	if s.Rows() != m.Columns() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	for r := 0; r < s.Rows(); r++ {
		rows := s.RowsAt(r)

		for c := 0; c < m.Columns(); c++ {
			column := m.ColumnsAt(c)

			sum := 0.0
			for l := 0; l < rows.Length(); l++ {
				vC := column.AtVec(l)
				vR := rows.AtVec(l)
				sum += vR * vC
			}

			matrix.Set(r, c, sum)
		}

	}

	return matrix
}

// multiplyVector multiplies a vector by another matrix
func multiplyVector(s, m, matrix Matrix) Matrix {
	if s.Rows() != m.Columns() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	for r := 0; r < m.Rows(); r++ {
		rows := m.RowsAt(r)
		for c := 0; c < s.Columns(); c++ {
			column := s.ColumnsAt(c)
			sum := 0.0
			for l := 0; l < rows.Length(); l++ {
				vC := column.AtVec(l)
				vR := rows.AtVec(l)
				sum += vR * vC
			}

			matrix.Set(r, c, sum)
		}

	}
	return matrix
}

// Add addition of a matrix by another matrix
func Add(s, m Matrix) Matrix {
	if s.Columns() != m.Columns() {
		log.Panicf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := m.Copy()

	s.Iterator(func(r, c int, value float64) bool {
		matrix.Update(r, c, func(v float64) float64 {
			return value + v
		})
		return true
	})

	return matrix
}

// Subtract subtracts one matrix from another matrix
func Subtract(s, m Matrix) Matrix {
	if s.Columns() != m.Columns() {
		log.Panicf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := m.Copy()

	s.Iterator(func(r, c int, value float64) bool {
		matrix.Update(r, c, func(v float64) float64 {
			return value - v
		})
		return true
	})

	return matrix
}

// Negative the negative of a matrix
func Negative(s Matrix) Matrix {
	return s.CopyArithmetic(func(v float64) float64 {
		return -v
	})

}

// Scalar multiplication of a matrix by alpha
func Scalar(s Matrix, alpha float64) Matrix {
	return s.CopyArithmetic(func(v float64) float64 {
		return alpha * v
	})
}

// Transpose swaps the rows and columns
func Transpose(s, m Matrix) Matrix {
	s.Iterator(func(r, c int, value float64) bool {
		m.Set(c, r, value)
		return true
	})

	return m
}

// TransposeToCSR swaps the rows and columns and returns a compressed storage by rows (CSR) matrix
func TransposeToCSR(s Matrix) Matrix {
	matrix := NewCSRMatrix(s.Columns(), s.Rows())

	return Transpose(s, matrix)
}

// TransposeToCSC swaps the rows and columns and returns a compressed storage by columns (CSC) matrix
func TransposeToCSC(s Matrix) Matrix {
	matrix := NewCSCMatrix(s.Columns(), s.Rows())

	return Transpose(s, matrix)
}

// Equal the two matrices are equal
func Equal(s, m Matrix) bool {
	if s.Columns() != m.Columns() {
		return false
	}

	if s.Rows() != m.Rows() {
		return false
	}

	return s.Iterator(func(r, c int, v float64) bool {
		value := m.At(r, c)
		if v != value {
			return false
		}
		return true
	})
}

// NotEqual the two matrices are not equal
func NotEqual(s, m Matrix) bool {
	return !s.Equal(m)
}
