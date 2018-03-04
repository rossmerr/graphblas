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
	if m.Rows() != s.Columns() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", m.Rows(), s.Columns())
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

// Add addition of a matrix by another matrix
func Add(s, m Matrix) Matrix {
	if s.Columns() != m.Columns() {
		log.Panicf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := m.Copy()

	for iterator := s.Enumerate(); iterator.HasNext(); {
		r, c, value := iterator.Next()
		matrix.Update(r, c, func(v float64) float64 {
			return value + v
		})
	}

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

	for iterator := s.Enumerate(); iterator.HasNext(); {
		r, c, value := iterator.Next()
		matrix.Update(r, c, func(v float64) float64 {
			return value - v
		})
	}
	return matrix
}

// Negative the negative of a matrix
func Negative(s Matrix) Matrix {
	matrix := s.Copy()
	for iterator := matrix.Map(); iterator.HasNext(); {
		iterator.Map(func(r, c int, v float64) float64 {
			return -v
		})
	}
	return matrix
}

// Scalar multiplication of a matrix by alpha
func Scalar(s Matrix, alpha float64) Matrix {
	matrix := s.Copy()
	for iterator := matrix.Map(); iterator.HasNext(); {
		iterator.Map(func(r, c int, v float64) float64 {
			return alpha * v
		})
	}
	return matrix
}

// Transpose swaps the rows and columns
func Transpose(s, m Matrix) Matrix {
	for iterator := s.Enumerate(); iterator.HasNext(); {
		r, c, value := iterator.Next()
		m.Set(c, r, value)
	}
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

	sIterator := s.Enumerate()
	mIterator := m.Enumerate()

	if sIterator.HasNext() && mIterator.HasNext() {
		sR, sC, sV := sIterator.Next()
		mR, mC, mV := mIterator.Next()

		if sR != mR || sC != mC || sV != mV {
			return false
		}
	}

	return true
}

// NotEqual the two matrices are not equal
func NotEqual(s, m Matrix) bool {
	return !s.Equal(m)
}
