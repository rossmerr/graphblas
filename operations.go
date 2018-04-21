// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"log"
	"reflect"
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

	var iterator Enumerate
	var matrix Matrix

	if SparseMatrix(s) {
		iterator = s.Enumerate()
		matrix = m.Copy()
	} else {
		iterator = m.Enumerate()
		matrix = s.Copy()
	}

	for iterator.HasNext() {
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

	// Using two Enumerators is the fastest way to iterate over matrices for a
	// equality checks of all values.
	// However that only work's when they are of the same type as the order is
	// not guaranteed and zero values are not returned for sparse matrices
	if reflect.TypeOf(s) == reflect.TypeOf(m) {
		// Because they are the same type they have the same storage method
		// so should or should not have the same size
		if s.Size() != m.Size() {
			return false
		}

		sIterator := s.Enumerate()
		mIterator := m.Enumerate()

		for {
			if sIterator.HasNext() && mIterator.HasNext() {
				sR, sC, sV := sIterator.Next()
				mR, mC, mV := mIterator.Next()

				if sR != mR || sC != mC || sV != mV {
					return false
				}
			} else {
				break
			}

		}

		return true
	}

	// If not the same type we can only enumerate over one matrix as order is not guaranteed
	var iterator Enumerate
	var matrix Matrix

	// Check for a sparse matrix as we want to use its Enumerate operation
	// Because the use of the At operation on a sparse matrix is expensive
	if SparseMatrix(s) {
		iterator = s.Enumerate()
		matrix = m
	} else {
		iterator = m.Enumerate()
		matrix = s
	}

	for {
		if iterator.HasNext() {
			sR, sC, sV := iterator.Next()
			mV := matrix.At(sR, sC)
			if sV != mV {
				return false
			}
		} else {
			break
		}

	}

	return true
}

// SparseMatrix is 's' a sparse matrix
func SparseMatrix(s Matrix) bool {
	if _, ok := s.(*CSCMatrix); ok {
		return true
	}

	if _, ok := s.(*CSRMatrix); ok {
		return true
	}

	if _, ok := s.(*SparseVector); ok {
		return true
	}

	return false
}

// NotEqual the two matrices are not equal
func NotEqual(s, m Matrix) bool {
	return !s.Equal(m)
}

// SkewSymmetric (or antisymmetric or antimetric) matrix is a square matrix whose transpose equals its negative
func SkewSymmetric(s Matrix) bool {
	r := s.Rows()
	c := s.Columns()
	if r != c {
		return false
	}

	t := s.Transpose()
	negativeTranspose := t.Negative()
	return negativeTranspose.Equal(s)
}

// Symmetric matrix is a square matrix that is equal to its transpose
func Symmetric(s Matrix) bool {
	r := s.Rows()
	c := s.Columns()
	if r != c {
		return false
	}

	t := s.Transpose()
	return t.Equal(s)
}

// // Hermitian
// func Hermitian(s Matrix) bool {
// 	r := s.Rows()
// 	c := s.Columns()
// 	if r != c {
// 		return false
// 	}

// 	t := s.Transpose()
// 	return t.Equal(s)
// }
