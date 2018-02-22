// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"log"
)

//StrassenMultiply multiplies a matrix by another matrix using the Strassen algorithm
func StrassenMultiply(s, m Matrix) Matrix {
	if s.Columns() != m.Rows() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", m.Rows(), s.Columns())
	}

	n := m.Rows()
	if n <= 48 {
		return NormalMultiply(s, m, nil)
	}

	halfN := n / 2

	a11 := subMatrix(s, 0, halfN, 0, halfN)
	a12 := subMatrix(s, 0, halfN, halfN, n)
	a21 := subMatrix(s, halfN, n, 0, halfN)
	a22 := subMatrix(s, halfN, n, halfN, n)

	b11 := subMatrix(m, 0, halfN, 0, halfN)
	b12 := subMatrix(m, 0, halfN, halfN, n)
	b21 := subMatrix(m, halfN, n, 0, halfN)
	b22 := subMatrix(m, halfN, n, halfN, n)

	mm := [7]Matrix{
		StrassenMultiply(a11.Add(a22), b11.Add(b22)),      // m1
		StrassenMultiply(a21.Add(a22), b11),               // m2
		StrassenMultiply(a11, b12.Subtract(b22)),          // m3
		StrassenMultiply(a22, b21.Subtract(b11)),          // m4
		StrassenMultiply(a11.Add(a12), b22),               // m5
		StrassenMultiply(a21.Subtract(a11), b11.Add(b12)), // m6
		StrassenMultiply(a12.Subtract(a22), b21.Add(b22)), // m7
	}

	c11 := mm[0].Add(mm[3]).Subtract(mm[4]).Add(mm[6])
	c12 := mm[2].Add(mm[4])
	c21 := mm[1].Add(mm[3])
	c22 := mm[0].Subtract(mm[1]).Add(mm[2]).Add(mm[5])

	return combineSubMatrices(c11, c12, c21, c22)
}

func subMatrix(s Matrix, rowFrom, rowTo, colFrom, colTo int) Matrix {
	result := NewDenseMatrix(rowTo-rowFrom, colTo-colFrom)
	i := 0
	for row := rowFrom; row < rowTo; row++ {
		i++
		j := 0
		for col := colFrom; col < colTo; col++ {
			j++
			result.Set(i, j, s.At(row, col))
		}
	}

	return result
}

func combineSubMatrices(a11, a12, a21, a22 Matrix) Matrix {
	result := NewDenseMatrix(a11.Rows()*2, a11.Rows()*2)
	shift := a11.Rows()
	for row := 0; row < a11.Rows(); row++ {
		for col := 0; col < a11.Columns(); col++ {
			result.Set(row, col, a11.At(row, col))
			result.Set(row, col+shift, a12.At(row, col))
			result.Set(row+shift, col, a21.At(row, col))
			result.Set(row+shift, col+shift, a22.At(row, col))
		}
	}
	return result
}

// NormalMultiply multiplies a matrix by another matrix
func NormalMultiply(s, m, matrix Matrix) Matrix {
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

	for iterator := s.Iterator(); iterator.HasNext(); {
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

	for iterator := s.Iterator(); iterator.HasNext(); {
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
	for iterator := matrix.Iterator(); iterator.HasNext(); {
		_, _, v := iterator.Next()
		iterator.Update(-v)
	}
	return matrix
}

// Scalar multiplication of a matrix by alpha
func Scalar(s Matrix, alpha float64) Matrix {
	matrix := s.Copy()
	for iterator := matrix.Iterator(); iterator.HasNext(); {
		_, _, v := iterator.Next()
		iterator.Update(alpha * v)
	}
	return matrix
}

// Transpose swaps the rows and columns
func Transpose(s, m Matrix) Matrix {
	for iterator := s.Iterator(); iterator.HasNext(); {
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

	sIterator := s.Iterator()
	mIterator := m.Iterator()

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
