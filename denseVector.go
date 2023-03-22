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

// DenseVector a vector
type DenseVector[T constraints.Type] struct {
	l      int // length of the sparse vector
	values []T
}

type DenseVectorNumber[T constraints.Number] struct {
	DenseVector[T]
}

// NewDenseVector returns a DenseVector
func NewDenseVector[T constraints.Type](l int) *DenseVector[T] {
	return newDenseVectorType[T](l)
}

// NewDenseVectorFromArray returns a SparseVector
func NewDenseVectorFromArray[T constraints.Type](data []T) *DenseVector[T] {
	v := newDenseVectorType[T](len(data))
	v.values = data
	return v
}

func NewDenseVectorN[T constraints.Number](l int) *DenseVectorNumber[T] {
	return newDenseVectorNumber[T](l)
}

// NewDenseVectorFromArray returns a SparseVector
func NewDenseVectorFromArrayN[T constraints.Number](data []T) *DenseVectorNumber[T] {
	v := newDenseVectorNumber[T](len(data))
	v.DenseVector.values = data
	return v
}

func newDenseVector[T constraints.Type](l int) DenseVector[T] {
	return DenseVector[T]{l: l, values: make([]T, l)}
}

func newDenseVectorType[T constraints.Type](l int) *DenseVector[T] {
	s := newDenseVector[T](l)
	return &s
}

func newDenseVectorNumber[T constraints.Number](l int) *DenseVectorNumber[T] {
	s := &DenseVectorNumber[T]{
		DenseVector: newDenseVector[T](l),
	}
	return s
}

// AtVec returns the value of a vector element at i-th
func (s *DenseVector[T]) AtVec(i int) T {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	return s.values[i]
}

// SetVec sets the value at i-th of the vector
func (s *DenseVector[T]) SetVec(i int, value T) {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	s.values[i] = value
}

// Size of the vector
func (s *DenseVector[T]) Size() int {
	return s.l
}

// Re-size the vector, will cut if the new size is smaller than the old size
func (s *DenseVector[T]) ReSize(size int) int {
	if size > s.l {
		length := size - s.l
		arr := make([]T, length)
		s.values = append(s.values, arr...)
	} else if size < s.l {
		length := s.l - size
		s.values = append([]T(nil), s.values[:length]...)
	}

	s.l = len(s.values)
	return s.l
}

// Length of the vector
func (s *DenseVector[T]) Length() int {
	return s.l
}

// Columns the number of columns of the vector
func (s *DenseVector[T]) Columns() int {
	return 1
}

// Rows the number of rows of the vector
func (s *DenseVector[T]) Rows() int {
	return s.l
}

// Update does a At and Set on the vector element at r-th, c-th
func (s *DenseVector[T]) Update(r, c int, f func(T) T) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	v := s.AtVec(r)
	s.SetVec(r, f(v))
}

// At returns the value of a vector element at r-th, c-th
func (s *DenseVector[T]) At(r, c int) (value T) {
	return s.AtVec(r)
}

// Set sets the value at r-th, c-th of the vector
func (s *DenseVector[T]) Set(r, c int, value T) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.SetVec(r, value)
}

// ColumnsAt return the columns at c-th
func (s *DenseVector[T]) ColumnsAt(c int) VectorLogial[T] {
	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	return s.copy()
}

// RowsAt return the rows at r-th
func (s *DenseVector[T]) RowsAt(r int) VectorLogial[T] {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewDenseVector[T](1)

	v := s.AtVec(r)
	rows.SetVec(0, v)

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *DenseVector[T]) RowsAtToArray(r int) []T {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]T, 1)

	v := s.AtVec(r)
	rows[0] = v

	return rows
}

func (s *DenseVector[T]) copy() *DenseVector[T] {
	vector := NewDenseVector[T](s.l)

	for i, v := range s.values {
		if v != Default[T]() {
			vector.SetVec(i, v)
		} else {
			vector.SetVec(i, v)
		}
	}

	return vector
}

func (s *DenseVectorNumber[T]) Copy() Matrix[T] {
	vector := newDenseVectorNumber[T](s.l)

	for i, v := range s.values {
		if v != Default[T]() {
			vector.SetVec(i, v)
		} else {
			vector.SetVec(i, v)
		}
	}

	return vector
}

// Copy copies the vector
func (s *DenseVector[T]) CopyLogical() MatrixLogical[T] {
	return s.copy()
}

// Scalar multiplication of a vector by alpha
func (s *DenseVectorNumber[T]) Scalar(alpha T) Matrix[T] {
	return Scalar[T](context.Background(), s, alpha)
}

// Multiply multiplies a vector by another vector
func (s *DenseVectorNumber[T]) Multiply(m Matrix[T]) Matrix[T] {
	matrix := newMatrixNumber[T](m.Rows(), s.Columns(), nil)
	MatrixMatrixMultiply[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Add addition of a vector by another vector
func (s *DenseVectorNumber[T]) Add(m Matrix[T]) Matrix[T] {
	matrix := s.Copy()
	Add[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Subtract subtracts one vector from another vector
func (s *DenseVectorNumber[T]) Subtract(m Matrix[T]) Matrix[T] {
	matrix := m.Copy()
	Subtract[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Negative the negative of a metrix
func (s *DenseVectorNumber[T]) Negative() MatrixLogical[T] {
	matrix := s.Copy()
	Negative[T](context.Background(), s, nil, matrix)
	return matrix
}

// Transpose swaps the rows and columns
func (s *DenseVector[T]) Transpose() MatrixLogical[T] {
	matrix := newMatrix[T](s.Columns(), s.Rows(), nil)
	Transpose[T](context.Background(), s, nil, &matrix)
	return &matrix
}

// Equal the two vectors are equal
func (s *DenseVector[T]) Equal(m MatrixLogical[T]) bool {
	return Equal[T](context.Background(), s, m)
}

// NotEqual the two vectors are not equal
func (s *DenseVector[T]) NotEqual(m MatrixLogical[T]) bool {
	return NotEqual[T](context.Background(), s, m)
}

// Values the number of elements in the vector
func (s *DenseVector[T]) Values() int {
	return s.l
}

// Clear removes all elements from a vector
func (s *DenseVector[T]) Clear() {
	s.values = make([]T, s.l)
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *DenseVector[T]) Enumerate() Enumerate[T] {
	return s.iterator()
}

func (s *DenseVector[T]) iterator() *denseVectorIterator[T] {
	i := &denseVectorIterator[T]{
		matrix: s,
		size:   s.Values(),
		last:   0,
		c:      0,
		r:      0,
	}
	return i
}

type denseVectorIterator[T constraints.Type] struct {
	matrix *DenseVector[T]
	size   int
	last   int
	c      int
	r      int
	rOld   int
}

// HasNext checks the iterator has any more values
func (s *denseVectorIterator[T]) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *denseVectorIterator[T]) next() {
	if s.r == s.matrix.Rows() {
		s.r = 0
		s.c++
	}
	s.rOld = s.r
	s.r++
	s.last++
}

// Next moves the iterator and returns the row, column and value
func (s *denseVectorIterator[T]) Next() (int, int, T) {
	s.next()

	return s.rOld, 0, s.matrix.AtVec(s.rOld)
}

// Map replace each element with the result of applying a function to its value
func (s *DenseVectorNumber[T]) Map() Map[T] {
	t := s.iterator()
	i := &denseVectorMap[T]{t}
	return i
}

type denseVectorMap[T constraints.Number] struct {
	*denseVectorIterator[T]
}

// HasNext checks the iterator has any more values
func (s *denseVectorMap[T]) HasNext() bool {
	return s.denseVectorIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *denseVectorMap[T]) Map(f func(int, int, T) T) {
	s.next()

	s.matrix.SetVec(s.rOld, f(s.rOld, 0, s.matrix.AtVec(s.rOld)))
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *DenseVectorNumber[T]) Element(r, c int) bool {
	return s.AtVec(r) > Default[T]()
}
