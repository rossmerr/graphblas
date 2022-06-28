// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas

import (
	"log"
	"reflect"

	"context"

	"github.com/rossmerr/graphblas/constraints"
)

func init() {
	RegisterMatrix(reflect.TypeOf((*CSCMatrix[float64])(nil)).Elem())
}

// CSCMatrix compressed storage by columns (CSC)
type CSCMatrix[T constraints.Number] struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []T
	rows     []int
	colStart []int
}

// NewCSCMatrix returns a CSCMatrix
func NewCSCMatrix[T constraints.Number](r, c int) *CSCMatrix[T] {
	return newCSCMatrix[T](r, c, 0)
}

// NewCSCMatrixFromArray returns a CSCMatrix
func NewCSCMatrixFromArray[T constraints.Number](data [][]T) *CSCMatrix[T] {
	r := len(data)
	c := len(data[0])
	s := newCSCMatrix[T](r, c, 0)

	for i := 0; i < r; i++ {
		for k := 0; k < c; k++ {
			s.Set(i, k, data[i][k])
		}
	}

	return s
}

func newCSCMatrix[T constraints.Number](r, c int, l int) *CSCMatrix[T] {
	ss := make([]T, l)

	s := &CSCMatrix[T]{
		r:        r,
		c:        c,
		values:   ss,
		rows:     make([]int, l),
		colStart: make([]int, c+1),
	}

	return s
}

// Columns the number of columns of the matrix
func (s *CSCMatrix[T]) Columns() int {
	return s.c
}

// Rows the number of rows of the matrix
func (s *CSCMatrix[T]) Rows() int {
	return s.r
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *CSCMatrix[T]) Update(r, c int, f func(T) T) {
	s.UpdateReturnPointer(r, c, f)
}

func (s *CSCMatrix[T]) UpdateReturnPointer(r, c int, f func(T) T) (pointer int, start int) {
	if r < 0 || r >= s.r {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		log.Panicf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.rowIndex(r, c)

	if pointerStart < pointerEnd && s.rows[pointerStart] == r {
		value := f(s.values[pointerStart])
		if IsZero(value) {
			s.remove(pointerStart, c)
		} else {
			s.values[pointerStart] = value
		}

		return pointerStart, 0
	} else {
		zero := Zero[T]()
		col := s.insert(pointerStart, r, c, f(zero))
		return pointerStart, col
	}
}

// At returns the value of a matrix element at r-th, c-th
func (s *CSCMatrix[T]) At(r, c int) (value T) {
	s.Update(r, c, func(v T) T {
		value = v
		return v
	})

	return
}

// Set sets the value at r-th, c-th of the matrix
func (s *CSCMatrix[T]) Set(r, c int, value T) {
	s.Update(r, c, func(v T) T {
		return value
	})
}

func (s *CSCMatrix[T]) SetReturnPointer(r, c int, value T) (pointer int, start int) {
	return s.UpdateReturnPointer(r, c, func(v T) T {
		return value
	})
}

// ColumnsAt return the columns at c-th
func (s *CSCMatrix[T]) ColumnsAt(c int) Vector[T] {
	if c < 0 || c >= s.c {
		log.Panicf("Column '%+v' is invalid", c)
	}

	columns := NewSparseVector[T](s.r)

	start := s.colStart[c]
	end := s.colStart[c+1]

	for i := start; i < end; i++ {
		columns.SetVec(s.rows[i], s.values[i])
	}

	return columns
}

// RowsAt return the rows at r-th
func (s *CSCMatrix[T]) RowsAt(r int) Vector[T] {
	if r < 0 || r >= s.r {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector[T](s.c)

	for c := range s.colStart[:s.c] {
		pointerStart, pointerEnd := s.rowIndex(r, c)
		if pointerStart < pointerEnd && s.rows[pointerStart] == r {
			rows.SetVec(c, s.values[pointerStart])
		}
	}

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *CSCMatrix[T]) RowsAtToArray(r int) []T {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]T, s.c)

	for c := range s.colStart[:s.c] {
		pointerStart, pointerEnd := s.rowIndex(r, c)
		if pointerStart < pointerEnd && s.rows[pointerStart] == r {
			rows[c] = s.values[pointerStart]
		}
	}

	return rows
}

func (s *CSCMatrix[T]) insert(pointer, r, c int, value T) int {
	if IsZero(value) {
		return 0
	}

	s.rows = append(s.rows[:pointer], append([]int{r}, s.rows[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]T{value}, s.values[pointer:]...)...)

	for i := c + 1; i <= s.c; i++ {
		s.colStart[i]++
	}

	return s.c
}

func (s *CSCMatrix[T]) remove(pointer, c int) {
	s.rows = append(s.rows[:pointer], s.rows[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)

	for i := c + 1; i <= s.c; i++ {
		s.colStart[i]--
	}
}

func (s *CSCMatrix[T]) rowIndex(r, c int) (int, int) {

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
func (s *CSCMatrix[T]) Copy() Matrix[T] {
	matrix := newCSCMatrix[T](s.r, s.c, len(s.values))

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
func (s *CSCMatrix[T]) Scalar(alpha T) Matrix[T] {
	return Scalar[T](context.Background(), s, alpha)
}

// Multiply multiplies a matrix by another matrix
func (s *CSCMatrix[T]) Multiply(m Matrix[T]) Matrix[T] {
	matrix := newCSCMatrix[T](s.Rows(), m.Columns(), 0)
	MatrixMatrixMultiply[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Add addition of a matrix by another matrix
func (s *CSCMatrix[T]) Add(m Matrix[T]) Matrix[T] {
	matrix := s.Copy()
	Add[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Subtract subtracts one matrix from another matrix
func (s *CSCMatrix[T]) Subtract(m Matrix[T]) Matrix[T] {
	matrix := m.Copy()
	Subtract[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Negative the negative of a matrix
func (s *CSCMatrix[T]) Negative() Matrix[T] {
	matrix := s.Copy()
	Negative[T](context.Background(), s, nil, matrix)
	return matrix
}

// Transpose swaps the rows and columns
func (s *CSCMatrix[T]) Transpose() Matrix[T] {
	matrix := newCSCMatrix[T](s.c, s.r, 0)

	Transpose[T](context.Background(), s, nil, matrix)
	return matrix
}

// Equal the two matrices are equal
func (s *CSCMatrix[T]) Equal(m Matrix[T]) bool {
	return Equal[T](context.Background(), s, m)
}

// NotEqual the two matrices are not equal
func (s *CSCMatrix[T]) NotEqual(m Matrix[T]) bool {
	return NotEqual[T](context.Background(), s, m)
}

// Size of the matrix
func (s *CSCMatrix[T]) Size() int {
	return s.Rows() * s.Columns()
}

// Values the number of non-zero elements in the matrix
func (s *CSCMatrix[T]) Values() int {
	return len(s.values)
}

// Clear removes all elements from a matrix
func (s *CSCMatrix[T]) Clear() {
	s.values = make([]T, 0)
	s.rows = make([]int, 0)
	s.colStart = make([]int, s.c+1)
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *CSCMatrix[T]) Enumerate() Enumerate[T] {
	return s.iterator()
}

func (s *CSCMatrix[T]) iterator() *cSCMatrixIterator[T] {
	i := &cSCMatrixIterator[T]{
		matrix: s,
		size:   len(s.values),
		c:      -1,
	}
	return i
}

type cSCMatrixIterator[T constraints.Number] struct {
	matrix       *CSCMatrix[T]
	size         int
	last         int
	c            int
	r            int
	rIndex       int
	index        int
	pointerStart int
	pointerEnd   int
}

// HasNext checks the iterator has any more values
func (s *cSCMatrixIterator[T]) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *cSCMatrixIterator[T]) next() {

	for s.pointerStart == s.pointerEnd {
		s.c++
		s.pointerStart = s.matrix.colStart[s.c]
		s.pointerEnd = s.matrix.colStart[s.c+1]
		s.rIndex = s.matrix.rows[s.pointerStart]
	}

	for s.pointerStart < s.pointerEnd {
		if s.matrix.rows[s.pointerStart] == s.rIndex {
			s.index = s.pointerStart
			s.pointerStart++
			s.r = s.rIndex
			s.rIndex++
			s.last++
			return
		}
		s.rIndex++
	}
}

// Next moves the iterator and returns the row, column and value
func (s *cSCMatrixIterator[T]) Next() (int, int, T) {
	s.next()
	return s.r, s.c, s.matrix.values[s.index]
}

// Map replace each element with the result of applying a function to its value
func (s *CSCMatrix[T]) Map() Map[T] {
	t := s.iterator()
	i := &cSCMatrixMap[T]{t}
	return i
}

type cSCMatrixMap[T constraints.Number] struct {
	*cSCMatrixIterator[T]
}

// HasNext checks the iterator has any more values
func (s *cSCMatrixMap[T]) HasNext() bool {
	return s.cSCMatrixIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *cSCMatrixMap[T]) Map(f func(int, int, T) T) {
	s.next()
	value := f(s.r, s.c, s.matrix.values[s.index])
	if value != Default[T]() {
		s.matrix.values[s.index] = value
	} else {
		s.matrix.remove(s.index, s.c)
	}
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *CSCMatrix[T]) Element(r, c int) (b bool) {
	s.Update(r, c, func(v T) T {
		b = v > Default[T]()
		return v
	})

	return
}
