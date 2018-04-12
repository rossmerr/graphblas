// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"log"
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

// NewCSCMatrixFromArray returns a GraphBLAS.CSCMatrix
func NewCSCMatrixFromArray(data [][]float64) *CSCMatrix {
	r := len(data)
	c := len(data[0])
	s := newCSCMatrix(r, c, 0)

	for i := 0; i < r; i++ {
		for k := 0; k < c; k++ {
			s.Set(i, k, data[i][k])
		}
	}

	return s
}

func newCSCMatrix(r, c int, l int) *CSCMatrix {
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
func (s *CSCMatrix) Update(r, c int, f func(float64) float64) {
	if r < 0 || r >= s.r {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		log.Panicf("Column '%+v' is invalid", c)
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
}

// At returns the value of a matrix element at r-th, c-th
func (s *CSCMatrix) At(r, c int) (value float64) {
	s.Update(r, c, func(v float64) float64 {
		value = v
		return v
	})

	return
}

// Set sets the value at r-th, c-th of the matrix
func (s *CSCMatrix) Set(r, c int, value float64) {
	s.Update(r, c, func(v float64) float64 {
		return value
	})
}

// ColumnsAt return the columns at c-th
func (s *CSCMatrix) ColumnsAt(c int) Vector {
	if c < 0 || c >= s.c {
		log.Panicf("Column '%+v' is invalid", c)
	}

	start := s.colStart[c]
	end := s.colStart[c+1]
	columns := NewSparseVector(s.r)

	for i := start; i < end; i++ {
		columns.SetVec(s.rows[i], s.values[i])
	}

	return columns
}

// RowsAt return the rows at r-th
func (s *CSCMatrix) RowsAt(r int) Vector {
	if r < 0 || r >= s.r {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector(s.c)

	for c := range s.colStart[:s.c] {
		pointerStart, pointerEnd := s.rowIndex(r, c)
		if pointerStart < pointerEnd && s.rows[pointerStart] == r {
			rows.SetVec(c, s.values[pointerStart])
		}
	}

	return rows
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
	matrix := newCSCMatrix(s.r, s.c, len(s.values))

	for i := range s.values {
		matrix.values[i] = s.values[i]
		matrix.rows[i] = s.rows[i]
	}

	for i := range s.colStart {
		matrix.colStart[i] = s.colStart[i]
	}

	return matrix
}

// Scalar multiplication of a matrix by alpha
func (s *CSCMatrix) Scalar(alpha float64) Matrix {
	return Scalar(s, alpha)
}

// Multiply multiplies a matrix by another matrix
func (s *CSCMatrix) Multiply(m Matrix) Matrix {
	matrix := newCSCMatrix(s.Rows(), m.Columns(), 0)

	return Multiply(s, m, matrix)
}

// Add addition of a matrix by another matrix
func (s *CSCMatrix) Add(m Matrix) Matrix {
	return Add(s, m)
}

// Subtract subtracts one matrix from another matrix
func (s *CSCMatrix) Subtract(m Matrix) Matrix {
	return Subtract(s, m)
}

// Negative the negative of a matrix
func (s *CSCMatrix) Negative() Matrix {
	return Negative(s)
}

// Transpose swaps the rows and columns
func (s *CSCMatrix) Transpose() Matrix {
	matrix := newCSCMatrix(s.c, s.r, 0)

	return Transpose(s, matrix)
}

// Equal the two matrices are equal
func (s *CSCMatrix) Equal(m Matrix) bool {
	return Equal(s, m)
}

// NotEqual the two matrices are not equal
func (s *CSCMatrix) NotEqual(m Matrix) bool {
	return NotEqual(s, m)
}

// Size the number of non-zero elements in the matrix
func (s *CSCMatrix) Size() int {
	return len(s.values)
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *CSCMatrix) Enumerate() Enumerate {
	return s.iterator()
}

func (s *CSCMatrix) iterator() *cSCMatrixIterator {
	i := &cSCMatrixIterator{
		matrix:     s,
		size:       len(s.values),
		last:       0,
		c:          0,
		r:          s.colStart[0],
		pointerEnd: s.colStart[0+1],
		rOld:       0,
	}
	i.rOld = i.r
	return i
}

type cSCMatrixIterator struct {
	matrix     *CSCMatrix
	size       int
	last       int
	c          int
	r          int
	rOld       int
	pointerEnd int
}

// HasNext checks the iterator has any more values
func (s *cSCMatrixIterator) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *cSCMatrixIterator) next() {
	if s.r == s.pointerEnd {
		s.c++
		s.r = s.pointerEnd
		s.pointerEnd = s.matrix.colStart[s.c+1]
	}

	s.rOld = s.r
	s.r++
	s.last++
}

// Next moves the iterator and returns the row, column and value
func (s *cSCMatrixIterator) Next() (int, int, float64) {
	s.next()
	return s.matrix.rows[s.rOld], s.c, s.matrix.values[s.rOld]
}

// Map replace each element with the result of applying a function to its value
func (s *CSCMatrix) Map() Map {
	t := s.iterator()
	i := &cSCMatrixMap{t}
	return i
}

type cSCMatrixMap struct {
	*cSCMatrixIterator
}

// HasNext checks the iterator has any more values
func (s *cSCMatrixMap) HasNext() bool {
	return s.cSCMatrixIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *cSCMatrixMap) Map(f func(int, int, float64) float64) {
	s.next()
	value := f(s.matrix.rows[s.rOld], s.c, s.matrix.values[s.rOld])
	if value != 0 {
		s.matrix.values[s.rOld] = value
	} else {
		s.matrix.remove(s.rOld, s.c)
	}
}

// EnumerateZeros iterates through all elements including zero elements, order is not guaranteed
func (s *CSCMatrix) EnumerateZeros() Enumerate {
	t := s.iterator()
	t.size = s.Rows() * s.Columns()
	i := &cSCMatrixIteratorWithZero{cSCMatrixIterator: t}
	return i
}

type cSCMatrixIteratorWithZero struct {
	*cSCMatrixIterator
	row int
}

// HasNext checks the iterator has any more values
func (s *cSCMatrixIteratorWithZero) HasNext() bool {
	return s.cSCMatrixIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *cSCMatrixIteratorWithZero) Next() (int, int, float64) {
	return s.nextWithZero()
}

func (s *cSCMatrixIteratorWithZero) nextWithZero() (int, int, float64) {
	if s.r == s.pointerEnd {
		s.c++
		s.r = s.pointerEnd
		s.pointerEnd = s.matrix.colStart[s.c+1]
	}

	s.rOld = s.r
	s.r++
	s.last++
	s.row = s.matrix.rows[s.rOld]
	return s.row, s.c, s.matrix.values[s.rOld]

}
