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
func Multiply(s, m, matrix Matrix) {
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
}

// Add addition of a matrix by another matrix
func Add(s, m, matrix Matrix) {
	if s.Columns() != m.Columns() {
		log.Panicf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	var iterator Enumerate

	if SparseMatrix(s) {
		iterator = s.Enumerate()
	} else {
		iterator = m.Enumerate()
	}

	for iterator.HasNext() {
		r, c, value := iterator.Next()
		matrix.Update(r, c, func(v float64) float64 {
			return value + v
		})
	}
}

// Subtract subtracts one matrix from another matrix
func Subtract(s, m, matrix Matrix) {
	if s.Columns() != m.Columns() {
		log.Panicf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	for iterator := s.Enumerate(); iterator.HasNext(); {
		r, c, value := iterator.Next()
		matrix.Update(r, c, func(v float64) float64 {
			return value - v
		})
	}
}

// Apply modifies edge weights by the UnaryOperator
// C ⊕= f(A)
func Apply(in, out Matrix, u UnaryOperator) {
	if in == out {
		for iterator := in.Map(); iterator.HasNext(); {
			iterator.Map(func(r, c int, value float64) (result float64) {
				u(value, result)
				return
			})
		}

		return
	}

	for iterator := in.Enumerate(); iterator.HasNext(); {
		r, c, value := iterator.Next()
		var result float64
		u(value, result)
		out.Set(c, r, result)
	}
}

// Negative the negative of a matrix
func Negative(s, matrix Matrix) {
	for iterator := matrix.Map(); iterator.HasNext(); {
		iterator.Map(func(r, c int, v float64) float64 {
			return -v
		})
	}
}

// Transpose swaps the rows and columns
// C ⊕= Aᵀ
func Transpose(s, m Matrix) {
	for iterator := s.Enumerate(); iterator.HasNext(); {
		r, c, value := iterator.Next()
		m.Set(c, r, value)
	}
}

// TransposeToCSR swaps the rows and columns and returns a compressed storage by rows (CSR) matrix
func TransposeToCSR(s Matrix) Matrix {
	matrix := NewCSRMatrix(s.Columns(), s.Rows())

	Transpose(s, matrix)
	return matrix
}

// TransposeToCSC swaps the rows and columns and returns a compressed storage by columns (CSC) matrix
func TransposeToCSC(s Matrix) Matrix {
	matrix := NewCSCMatrix(s.Columns(), s.Rows())

	Transpose(s, matrix)
	return matrix
}

// Equal the two matrices are equal
func Equal(s, m Matrix) bool {
	if s == nil && m == nil {
		return true
	}

	if s != nil && m == nil {
		return false
	}

	if m != nil && s == nil {
		return false
	}

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
		if s.Values() != m.Values() {
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

// NotEqual the two matrices are not equal
func NotEqual(s, m Matrix) bool {
	return !s.Equal(m)
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

// Reduced row echelon form of matrix (Gauss-Jordan elimination)
// rref
func Reduced(s Matrix) Matrix {
	m := s.Copy()
	lead := 0
	rowCount := m.Rows()
	columnCount := m.Columns()

	for r := 0; r < rowCount; r++ {
		if lead >= columnCount {
			return m
		}
		i := r
		for m.At(i, lead) == 0 {
			i++
			if rowCount == i {
				i = r
				lead++
				if columnCount == lead {
					return m
				}
			}
		}

		if i != r {
			v1 := m.RowsAtToArray(i)
			v2 := m.RowsAtToArray(r)

			for c := 0; c < len(v1); c++ {
				m.Set(r, c, v1[c])
			}

			for c := 0; c < len(v2); c++ {
				m.Set(i, c, v2[c])
			}
		}

		f := 1 / m.At(r, lead)

		vector := m.RowsAtToArray(r)
		for c := 0; c < len(vector); c++ {
			value := vector[c]
			value *= f
			m.Set(r, c, value)
		}

		for i = 0; i < rowCount; i++ {
			if i != r {
				f = m.At(i, lead)
				vector := m.RowsAtToArray(r)
				for c := 0; c < len(vector); c++ {
					value := vector[c]
					m.Update(i, c, func(v float64) float64 {
						v -= value * f
						return v
					})

				}
			}
		}
		lead++
	}

	return m
}
