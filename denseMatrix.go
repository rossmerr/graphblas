// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas

import (
	"context"
	"log"

	"github.com/rossmerr/graphblas/constraints"
)

// DenseMatrix a dense matrix
type DenseMatrix[T constraints.Number] struct {
	c    int // number of rows in the sparse matrix
	r    int // number of columns in the sparse matrix
	data [][]T
}

// NewDenseMatrix returns a DenseMatrix
func NewDenseMatrix[T constraints.Number](r, c int) *DenseMatrix[T] {
	return newMatrix[T](r, c, nil)
}

// NewDenseMatrixFromArray returns a DenseMatrix
func NewDenseMatrixFromArray[T constraints.Number](data [][]T) *DenseMatrix[T] {
	r := len(data)
	c := len(data[0])
	s := &DenseMatrix[T]{data: data, r: r, c: c}

	return s
}

func newMatrix[T constraints.Number](r, c int, initialise func([]T, int)) *DenseMatrix[T] {
	s := &DenseMatrix[T]{data: make([][]T, r), r: r, c: c}

	for i := 0; i < r; i++ {
		s.data[i] = make([]T, c)

		if initialise != nil {
			initialise(s.data[i], i)
		}
	}

	return s
}

// Columns the number of columns of the matrix
func (s *DenseMatrix[T]) Columns() int {
	return s.c
}

// Rows the number of rows of the matrix
func (s *DenseMatrix[T]) Rows() int {
	return s.r
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *DenseMatrix[T]) Update(r, c int, f func(T) T) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.data[r][c] = f(s.data[r][c])

	return
}

// At returns the value of a matrix element at r-th, c-th
func (s *DenseMatrix[T]) At(r, c int) T {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	return s.data[r][c]
}

// Set sets the value at r-th, c-th of the matrix
func (s *DenseMatrix[T]) Set(r, c int, value T) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.data[r][c] = value
}

// ColumnsAt return the columns at c-th
func (s *DenseMatrix[T]) ColumnsAt(c int) Vector[T] {
	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	columns := NewDenseVector[T](s.r)

	for r := 0; r < s.r; r++ {
		columns.SetVec(r, s.data[r][c])
	}

	return columns
}

// RowsAt return the rows at r-th
func (s *DenseMatrix[T]) RowsAt(r int) Vector[T] {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewDenseVector[T](s.c)

	for i := 0; i < s.c; i++ {
		rows.SetVec(i, s.data[r][i])
	}

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *DenseMatrix[T]) RowsAtToArray(r int) []T {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]T, s.c)

	for i := 0; i < s.c; i++ {
		rows[i] = s.data[r][i]
	}

	return rows
}

// Copy copies the matrix
func (s *DenseMatrix[T]) Copy() Matrix[T] {
	v := Default[T]()

	matrix := newMatrix(s.Rows(), s.Columns(), func(row []T, r int) {
		for c := 0; c < s.Columns(); c++ {
			v = s.data[r][c]
			if v != Default[T]() {
				row[c] = v
			} else {
				row[c] = v
			}
		}
	})

	return matrix
}

// Scalar multiplication of a matrix by alpha
func (s *DenseMatrix[T]) Scalar(alpha T) Matrix[T] {
	return Scalar[T](context.Background(), s, alpha)
}

// Multiply multiplies a matrix by another matrix
func (s *DenseMatrix[T]) Multiply(m Matrix[T]) Matrix[T] {
	matrix := newMatrix[T](s.Rows(), m.Columns(), nil)
	MatrixMatrixMultiply[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Add addition of a matrix by another matrix
func (s *DenseMatrix[T]) Add(m Matrix[T]) Matrix[T] {
	matrix := s.Copy()
	Add[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Subtract subtracts one matrix from another matrix
func (s *DenseMatrix[T]) Subtract(m Matrix[T]) Matrix[T] {
	matrix := m.Copy()
	Subtract[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Negative the negative of a matrix
func (s *DenseMatrix[T]) Negative() Matrix[T] {
	matrix := s.Copy()
	Negative[T](context.Background(), s, nil, matrix)
	return matrix
}

// Transpose swaps the rows and columns
func (s *DenseMatrix[T]) Transpose() Matrix[T] {
	matrix := newMatrix[T](s.Columns(), s.Rows(), nil)
	Transpose[T](context.Background(), s, nil, matrix)
	return matrix
}

// Equal the two matrices are equal
func (s *DenseMatrix[T]) Equal(m Matrix[T]) bool {
	return Equal[T](context.Background(), s, m)
}

// NotEqual the two matrices are not equal
func (s *DenseMatrix[T]) NotEqual(m Matrix[T]) bool {
	return NotEqual[T](context.Background(), s, m)
}

// Size of the matrix
func (s *DenseMatrix[T]) Size() int {
	return s.r * s.c
}

// Values the number of elements in the matrix
func (s *DenseMatrix[T]) Values() int {
	return s.r * s.c
}

// Clear removes all elements from a matrix
func (s *DenseMatrix[T]) Clear() {
	s.data = make([][]T, s.r)
	for i := 0; i < s.r; i++ {
		s.data[i] = make([]T, s.c)
	}
}

// RawMatrix returns the raw matrix
func (s *DenseMatrix[T]) RawMatrix() [][]T {
	return s.data
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *DenseMatrix[T]) Enumerate() Enumerate[T] {
	return s.iterator()
}

func (s *DenseMatrix[T]) iterator() *denseMatrixIterator[T] {
	i := &denseMatrixIterator[T]{
		matrix: s,
		size:   s.Values(),
		last:   0,
		c:      0,
		r:      0,
	}
	return i
}

type denseMatrixIterator[T constraints.Number] struct {
	matrix *DenseMatrix[T]
	size   int
	last   int
	c      int
	r      int
	cOld   int
	rOld   int
}

// HasNext checks the iterator has any more values
func (s *denseMatrixIterator[T]) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *denseMatrixIterator[T]) next() {
	if s.c == s.matrix.Columns() {
		s.c = 0
		s.r++
	}
	s.cOld = s.c
	s.c++
	s.last++
}

// Next moves the iterator and returns the row, column and value
func (s *denseMatrixIterator[T]) Next() (int, int, T) {
	s.next()

	return s.r, s.cOld, s.matrix.At(s.r, s.cOld)
}

// Map replace each element with the result of applying a function to its value
func (s *DenseMatrix[T]) Map() Map[T] {
	t := s.iterator()
	i := &denseMatrixMap[T]{t}
	return i
}

type denseMatrixMap[T constraints.Number] struct {
	*denseMatrixIterator[T]
}

// HasNext checks the iterator has any more values
func (s *denseMatrixMap[T]) HasNext() bool {
	return s.denseMatrixIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *denseMatrixMap[T]) Map(f func(int, int, T) T) {
	s.next()

	s.matrix.Set(s.r, s.cOld, f(s.r, s.cOld, s.matrix.At(s.r, s.cOld)))
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *DenseMatrix[T]) Element(r, c int) bool {
	return s.element(r, c)
}

func (s *DenseMatrix[T]) element(r, c int) bool {
	return s.At(r, c) > Default[T]()
}
