// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas

import (
	"sync"

	"github.com/rossmerr/graphblas/constraints"
)

// MutexMatrix a matrix wrapper that has a mutex lock support
type MutexMatrix[T constraints.Number] struct {
	sync.RWMutex
	matrix Matrix[T]
}

// NewMutexMatrix returns a MutexMatrix
func NewMutexMatrix[T constraints.Number](matrix Matrix[T]) *MutexMatrix[T] {
	return &MutexMatrix[T]{
		matrix: matrix,
	}
}

// Columns the number of columns of the matrix
func (s *MutexMatrix[T]) Columns() int {
	return s.matrix.Columns()
}

// Rows the number of rows of the matrix
func (s *MutexMatrix[T]) Rows() int {
	return s.matrix.Rows()
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *MutexMatrix[T]) Update(r, c int, f func(T) T) {
	s.Lock()
	defer s.Unlock()

	s.matrix.Update(r, c, f)
}

// At returns the value of a matrix element at r-th, c-th
func (s *MutexMatrix[T]) At(r, c int) T {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.At(r, c)
}

// Set sets the value at r-th, c-th of the matrix
func (s *MutexMatrix[T]) Set(r, c int, value T) {
	s.Lock()
	defer s.Unlock()

	s.matrix.Set(r, c, value)
}

// ColumnsAt return the columns at c-th
func (s *MutexMatrix[T]) ColumnsAt(c int) VectorLogial[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.ColumnsAt(c)
}

// RowsAt return the rows at r-th
func (s *MutexMatrix[T]) RowsAt(r int) VectorLogial[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.RowsAt(r)
}

// RowsAtToArray return the rows at r-th
func (s *MutexMatrix[T]) RowsAtToArray(r int) []T {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.RowsAtToArray(r)
}

// Copy copies the matrix
func (s *MutexMatrix[T]) Copy() MatrixLogical[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Copy()
}

// Scalar multiplication of a matrix by alpha
func (s *MutexMatrix[T]) Scalar(alpha T) Matrix[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Scalar(alpha)
}

// Multiply multiplies a matrix by another matrix
func (s *MutexMatrix[T]) Multiply(m Matrix[T]) Matrix[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Multiply(m)
}

// Add addition of a matrix by another matrix
func (s *MutexMatrix[T]) Add(m Matrix[T]) Matrix[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Add(m)
}

// Subtract subtracts one matrix from another matrix
func (s *MutexMatrix[T]) Subtract(m Matrix[T]) Matrix[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Subtract(m)
}

// Negative the negative of a matrix
func (s *MutexMatrix[T]) Negative() MatrixLogical[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Negative()
}

// Transpose swaps the rows and columns
func (s *MutexMatrix[T]) Transpose() MatrixLogical[T] {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Transpose()
}

// Equal the two matrices are equal
func (s *MutexMatrix[T]) Equal(m MatrixLogical[T]) bool {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Equal(m)
}

// NotEqual the two matrices are not equal
func (s *MutexMatrix[T]) NotEqual(m MatrixLogical[T]) bool {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.NotEqual(m)
}

// Size of the matrix
func (s *MutexMatrix[T]) Size() int {
	return s.matrix.Size()
}

// Values the number of elements in the matrix
func (s *MutexMatrix[T]) Values() int {
	return s.matrix.Values()
}

// Clear removes all elements from a matrix
func (s *MutexMatrix[T]) Clear() {
	s.RLock()
	defer s.RUnlock()

	s.matrix.Clear()
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *MutexMatrix[T]) Enumerate() Enumerate[T] {
	return s.matrix.Enumerate()
}

// Map replace each element with the result of applying a function to its value
func (s *MutexMatrix[T]) Map() Map[T] {
	return s.matrix.Map()
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *MutexMatrix[T]) Element(r, c int) bool {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Element(r, c)
}
