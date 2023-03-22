// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas

import (
	"context"
	"log"
	"reflect"

	"github.com/rossmerr/graphblas/constraints"
)

func init() {
	RegisterMatrix(reflect.TypeOf((*CSRMatrix[float64])(nil)).Elem())
}

// CSRMatrix compressed storage by rows (CSR)
type CSRMatrix[T constraints.Number] struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []T
	cols     []int
	rowStart []int
}

// NewCSRMatrix returns a CSRMatrix
func NewCSRMatrix[T constraints.Number](r, c int) *CSRMatrix[T] {
	return newCSRMatrix[T](r, c, 0)
}

// NewCSRMatrixFromArray returns a CSRMatrix
func NewCSRMatrixFromArray[T constraints.Number](data [][]T) *CSRMatrix[T] {
	r := len(data)
	c := len(data[0])
	s := newCSRMatrix[T](r, c, 0)
	for i := 0; i < r; i++ {
		for k := 0; k < c; k++ {
			s.Set(i, k, data[i][k])
		}
	}
	return s
}

func newCSRMatrix[T constraints.Number](r, c int, l int) *CSRMatrix[T] {
	s := &CSRMatrix[T]{
		r:        r,
		c:        c,
		values:   make([]T, l),
		cols:     make([]int, l),
		rowStart: make([]int, r+1),
	}
	return s
}

// Columns the number of columns of the matrix
func (s *CSRMatrix[T]) Columns() int {
	return s.c
}

// Rows the number of rows of the matrix
func (s *CSRMatrix[T]) Rows() int {
	return s.r
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *CSRMatrix[T]) Update(r, c int, f func(T) T) {
	s.UpdateReturnPointer(r, c, f)
}

func (s *CSRMatrix[T]) UpdateReturnPointer(r, c int, f func(T) T) (pointer int, start int) {
	if r < 0 || r >= s.r {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		log.Panicf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.columnIndex(r, c)

	if pointerStart < pointerEnd && s.cols[pointerStart] == c {
		value := f(s.values[pointerStart])
		if value == Default[T]() {
			s.remove(pointerStart, r)
		} else {
			s.values[pointerStart] = value
		}
		return pointerStart, 0
	} else {
		row := s.insert(pointerStart, r, c, f(Default[T]()))
		return pointerStart, row
	}
}

// At returns the value of a matrix element at r-th, c-th
func (s *CSRMatrix[T]) At(r, c int) (value T) {
	s.Update(r, c, func(v T) T {
		value = v
		return v
	})

	return
}

// Set sets the value at r-th, c-th of the matrix
func (s *CSRMatrix[T]) Set(r, c int, value T) {
	s.Update(r, c, func(v T) T {
		return value
	})
}

func (s *CSRMatrix[T]) SetReturnPointer(r, c int, value T) (pointer int, start int) {
	return s.UpdateReturnPointer(r, c, func(v T) T {
		return value
	})
}

// ColumnsAt return the columns at c-th
func (s *CSRMatrix[T]) ColumnsAt(c int) VectorLogial[T] {
	if c < 0 || c >= s.c {
		log.Panicf("Column '%+v' is invalid", c)
	}

	columns := NewSparseVector[T](s.r)

	for r := range s.rowStart[:s.r] {
		pointerStart, pointerEnd := s.columnIndex(r, c)
		if pointerStart < pointerEnd && s.cols[pointerStart] == c {
			columns.SetVec(r, s.values[pointerStart])
		}
	}

	return columns

}

// RowsAt return the rows at r-th
func (s *CSRMatrix[T]) RowsAt(r int) VectorLogial[T] {
	if r < 0 || r >= s.r {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector[T](s.c)

	start := s.rowStart[r]
	end := s.rowStart[r+1]

	for i := start; i < end; i++ {
		rows.SetVec(s.cols[i], s.values[i])
	}

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *CSRMatrix[T]) RowsAtToArray(r int) []T {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]T, s.c)

	start := s.rowStart[r]
	end := s.rowStart[r+1]

	for i := start; i < end; i++ {
		rows[s.cols[i]] = s.values[i]
	}

	return rows
}

func (s *CSRMatrix[T]) insert(pointer, r, c int, value T) int {
	if value == Default[T]() {
		return 0
	}

	s.cols = append(s.cols[:pointer], append([]int{c}, s.cols[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]T{value}, s.values[pointer:]...)...)

	for i := r + 1; i <= s.r; i++ {
		s.rowStart[i]++
	}

	return s.r
}

func (s *CSRMatrix[T]) remove(pointer, r int) {
	s.cols = append(s.cols[:pointer], s.cols[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)

	for i := r + 1; i <= s.r; i++ {
		s.rowStart[i]--
	}
}

func (s *CSRMatrix[T]) columnIndex(r, c int) (int, int) {

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
func (s *CSRMatrix[T]) CopyLogical() MatrixLogical[T] {
	return s.Copy()
}

func (s *CSRMatrix[T]) Copy() Matrix[T] {
	matrix := newCSRMatrix[T](s.r, s.c, len(s.values))

	for i := range s.values {
		matrix.values[i] = s.values[i]
		matrix.cols[i] = s.cols[i]
	}

	for i := range s.rowStart {
		matrix.rowStart[i] = s.rowStart[i]
	}

	return matrix
}

// Scalar multiplication of a matrix by alpha
func (s *CSRMatrix[T]) Scalar(alpha T) Matrix[T] {
	return Scalar[T](context.Background(), s, alpha)
}

// Multiply multiplies a matrix by another matrix
func (s *CSRMatrix[T]) Multiply(m Matrix[T]) Matrix[T] {
	matrix := newCSRMatrix[T](s.Rows(), m.Columns(), 0)
	MatrixMatrixMultiply[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Add addition of a matrix by another matrix
func (s *CSRMatrix[T]) Add(m Matrix[T]) Matrix[T] {
	matrix := s.Copy()
	Add[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Subtract subtracts one matrix from another matrix
func (s *CSRMatrix[T]) Subtract(m Matrix[T]) Matrix[T] {
	matrix := m.Copy()
	Subtract[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Negative the negative of a matrix
func (s *CSRMatrix[T]) Negative() MatrixLogical[T] {
	matrix := s.Copy()
	Negative[T](context.Background(), s, nil, matrix)
	return matrix
}

// Transpose swaps the rows and columns
func (s *CSRMatrix[T]) TransposeLogical() MatrixLogical[T] {
	return s.Transpose()
}

// Transpose swaps the rows and columns
func (s *CSRMatrix[T]) Transpose() Matrix[T] {
	matrix := newCSRMatrix[T](s.c, s.r, 0)
	Transpose[T](context.Background(), s, nil, matrix)
	return matrix
}

// Equal the two matrices are equal
func (s *CSRMatrix[T]) Equal(m MatrixLogical[T]) bool {
	return Equal[T](context.Background(), s, m)
}

// NotEqual the two matrices are not equal
func (s *CSRMatrix[T]) NotEqual(m MatrixLogical[T]) bool {
	return NotEqual[T](context.Background(), s, m)
}

// Size of the matrix
func (s *CSRMatrix[T]) Size() int {
	return s.Rows() * s.Columns()
}

// Values the number of non-zero elements in the matrix
func (s *CSRMatrix[T]) Values() int {
	return len(s.values)
}

// Clear removes all elements from a matrix
func (s *CSRMatrix[T]) Clear() {
	s.values = make([]T, 0)
	s.cols = make([]int, 0)
	s.rowStart = make([]int, s.r+1)
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *CSRMatrix[T]) Enumerate() Enumerate[T] {
	return s.iterator()
}

type cSRMatrixIterator[T constraints.Number] struct {
	matrix       *CSRMatrix[T]
	size         int
	last         int
	c            int
	r            int
	cIndex       int
	index        int
	pointerStart int
	pointerEnd   int
}

func (s *CSRMatrix[T]) iterator() *cSRMatrixIterator[T] {
	i := &cSRMatrixIterator[T]{
		matrix: s,
		size:   len(s.values),
		r:      -1,
	}
	return i
}

func (s *cSRMatrixIterator[T]) next() {

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
func (s *cSRMatrixIterator[T]) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

// Next moves the iterator and returns the row, column and value
func (s *cSRMatrixIterator[T]) Next() (int, int, T) {
	s.next()
	return s.r, s.c, s.matrix.values[s.index]
}

// Map replace each element with the result of applying a function to its value
func (s *CSRMatrix[T]) Map() Map[T] {
	t := s.iterator()
	i := &cSRMatrixMap[T]{t}
	return i
}

type cSRMatrixMap[T constraints.Number] struct {
	*cSRMatrixIterator[T]
}

// HasNext checks the iterator has any more values
func (s *cSRMatrixMap[T]) HasNext() bool {
	return s.cSRMatrixIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *cSRMatrixMap[T]) Map(f func(int, int, T) T) {
	s.next()
	value := f(s.r, s.c, s.matrix.values[s.index])
	if value != Default[T]() {
		s.matrix.values[s.index] = value
	} else {
		s.matrix.remove(s.index, s.r)
	}
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *CSRMatrix[T]) Element(r, c int) (b bool) {
	s.Update(r, c, func(v T) T {
		b = v > Default[T]()
		return v
	})

	return
}
