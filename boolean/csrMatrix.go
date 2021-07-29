// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolean

import (
	"context"
	"log"
	"reflect"
)

func init() {
	RegisterMatrix(reflect.TypeOf((*CSRMatrix)(nil)).Elem())
}

// CSRMatrix compressed storage by rows (CSR)
type CSRMatrix struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []bool
	cols     []int
	rowStart []int
}

// NewCSRMatrix returns a CSRMatrix
func NewCSRMatrix(r, c int) *CSRMatrix {
	return newCSRMatrix(r, c, 0)
}

// NewCSRMatrixFromArray returns a CSRMatrix
func NewCSRMatrixFromArray(data [][]bool) *CSRMatrix {
	r := len(data)
	c := len(data[0])
	s := newCSRMatrix(r, c, 0)
	for i := 0; i < r; i++ {
		for k := 0; k < c; k++ {
			s.Set(i, k, data[i][k])
		}
	}
	return s
}

func newCSRMatrix(r, c int, l int) *CSRMatrix {
	s := &CSRMatrix{
		r:        r,
		c:        c,
		values:   make([]bool, l),
		cols:     make([]int, l),
		rowStart: make([]int, r+1),
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
func (s *CSRMatrix) Update(r, c int, f func(bool) bool) {
	if r < 0 || r >= s.r {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		log.Panicf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.columnIndex(r, c)

	if pointerStart < pointerEnd && s.cols[pointerStart] == c {
		value := f(s.values[pointerStart])
		if value == false {
			s.remove(pointerStart, r)
		} else {
			s.values[pointerStart] = value
		}
	} else {
		s.insert(pointerStart, r, c, f(false))
	}
}

// At returns the value of a matrix element at r-th, c-th
func (s *CSRMatrix) At(r, c int) (value bool) {
	s.Update(r, c, func(v bool) bool {
		value = v
		return v
	})

	return
}

// Set sets the value at r-th, c-th of the matrix
func (s *CSRMatrix) Set(r, c int, value bool) {
	s.Update(r, c, func(v bool) bool {
		return value
	})
}

// ColumnsAt return the columns at c-th
func (s *CSRMatrix) ColumnsAt(c int) Vector {
	if c < 0 || c >= s.c {
		log.Panicf("Column '%+v' is invalid", c)
	}

	columns := NewSparseVector(s.r)

	for r := range s.rowStart[:s.r] {
		pointerStart, pointerEnd := s.columnIndex(r, c)
		if pointerStart < pointerEnd && s.cols[pointerStart] == c {
			columns.SetVec(r, s.values[pointerStart])
		}
	}

	return columns

}

// RowsAt return the rows at r-th
func (s *CSRMatrix) RowsAt(r int) Vector {
	if r < 0 || r >= s.r {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector(s.c)

	start := s.rowStart[r]
	end := s.rowStart[r+1]

	for i := start; i < end; i++ {
		rows.SetVec(s.cols[i], s.values[i])
	}

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *CSRMatrix) RowsAtToArray(r int) []bool {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]bool, s.c)

	start := s.rowStart[r]
	end := s.rowStart[r+1]

	for i := start; i < end; i++ {
		rows[s.cols[i]] = s.values[i]
	}

	return rows
}

func (s *CSRMatrix) insert(pointer, r, c int, value bool) {
	if value == false {
		return
	}

	s.cols = append(s.cols[:pointer], append([]int{c}, s.cols[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]bool{value}, s.values[pointer:]...)...)

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
	matrix := newCSRMatrix(s.r, s.c, len(s.values))

	for i := range s.values {
		matrix.values[i] = s.values[i]
		matrix.cols[i] = s.cols[i]
	}

	for i := range s.rowStart {
		matrix.rowStart[i] = s.rowStart[i]
	}

	return matrix
}

// Transpose swaps the rows and columns
func (s *CSRMatrix) Transpose() Matrix {
	matrix := newCSRMatrix(s.c, s.r, 0)
	Transpose(context.Background(), s, nil, matrix)
	return matrix
}

// Equal the two matrices are equal
func (s *CSRMatrix) Equal(m Matrix) bool {
	return Equal(context.Background(), s, m)
}

// NotEqual the two matrices are not equal
func (s *CSRMatrix) NotEqual(m Matrix) bool {
	return NotEqual(context.Background(), s, m)
}

// Size of the matrix
func (s *CSRMatrix) Size() int {
	return s.Rows() * s.Columns()
}

// Values the number of non-zero elements in the matrix
func (s *CSRMatrix) Values() int {
	return len(s.values)
}

// Clear removes all elements from a matrix
func (s *CSRMatrix) Clear() {
	s.values = make([]bool, 0)
	s.cols = make([]int, 0)
	s.rowStart = make([]int, s.r+1)
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *CSRMatrix) Enumerate() Enumerate {
	return s.iterator()
}

type cSRMatrixIterator struct {
	matrix       *CSRMatrix
	size         int
	last         int
	c            int
	r            int
	cIndex       int
	index        int
	pointerStart int
	pointerEnd   int
}

func (s *CSRMatrix) iterator() *cSRMatrixIterator {
	i := &cSRMatrixIterator{
		matrix: s,
		size:   len(s.values),
		r:      -1,
	}
	return i
}

func (s *cSRMatrixIterator) next() {

	for s.pointerStart == s.pointerEnd {
		s.r++
		s.pointerStart = s.matrix.rowStart[s.r]
		s.pointerEnd = s.matrix.rowStart[s.r+1]
		s.cIndex = s.matrix.cols[s.pointerStart]
	}

	for s.pointerStart < s.pointerEnd {
		if s.matrix.cols[s.pointerStart] == s.cIndex {
			s.index = s.pointerStart
			s.pointerStart++
			s.c = s.cIndex
			s.cIndex++
			s.last++
			return
		}
		s.cIndex++
	}
}

// HasNext checks the iterator has any more values
func (s *cSRMatrixIterator) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

// Next moves the iterator and returns the row, column and value
func (s *cSRMatrixIterator) Next() (int, int, bool) {
	s.next()
	return s.r, s.c, s.matrix.values[s.index]
}

// Map replace each element with the result of applying a function to its value
func (s *CSRMatrix) Map() Map {
	t := s.iterator()
	i := &cSRMatrixMap{t}
	return i
}

type cSRMatrixMap struct {
	*cSRMatrixIterator
}

// HasNext checks the iterator has any more values
func (s *cSRMatrixMap) HasNext() bool {
	return s.cSRMatrixIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *cSRMatrixMap) Map(f func(int, int, bool) bool) {
	s.next()
	value := f(s.r, s.c, s.matrix.values[s.index])
	if value != false {
		s.matrix.values[s.index] = value
	} else {
		s.matrix.remove(s.index, s.r)
	}
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *CSRMatrix) Element(r, c int) (b bool) {
	s.Update(r, c, func(v bool) bool {
		return b
	})

	return
}
