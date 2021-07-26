// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package f32

import (
	"sync"
)

// MutexMatrix a matrix wrapper that has a mutex lock support
type MutexMatrix struct {
	sync.RWMutex
	matrix Matrix
}

// NewMutexMatrix returns a MutexMatrix
func NewMutexMatrix(matrix Matrix) *MutexMatrix {
	return &MutexMatrix{
		matrix: matrix,
	}
}

// Columns the number of columns of the matrix
func (s *MutexMatrix) Columns() int {
	return s.matrix.Columns()
}

// Rows the number of rows of the matrix
func (s *MutexMatrix) Rows() int {
	return s.matrix.Rows()
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *MutexMatrix) Update(r, c int, f func(float32) float32) {
	s.Lock()
	defer s.Unlock()

	s.matrix.Update(r, c, f)
}

// At returns the value of a matrix element at r-th, c-th
func (s *MutexMatrix) At(r, c int) float32 {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.At(r, c)
}

// Set sets the value at r-th, c-th of the matrix
func (s *MutexMatrix) Set(r, c int, value float32) {
	s.Lock()
	defer s.Unlock()

	s.matrix.Set(r, c, value)
}

// ColumnsAt return the columns at c-th
func (s *MutexMatrix) ColumnsAt(c int) Vector {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.ColumnsAt(c)
}

// RowsAt return the rows at r-th
func (s *MutexMatrix) RowsAt(r int) Vector {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.RowsAt(r)
}

// RowsAtToArray return the rows at r-th
func (s *MutexMatrix) RowsAtToArray(r int) []float32 {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.RowsAtToArray(r)
}

// Copy copies the matrix
func (s *MutexMatrix) Copy() Matrix {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Copy()
}

// Scalar multiplication of a matrix by alpha
func (s *MutexMatrix) Scalar(alpha float32) Matrix {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Scalar(alpha)
}

// Multiply multiplies a matrix by another matrix
func (s *MutexMatrix) Multiply(m Matrix) Matrix {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Multiply(m)
}

// Add addition of a matrix by another matrix
func (s *MutexMatrix) Add(m Matrix) Matrix {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Add(m)
}

// Subtract subtracts one matrix from another matrix
func (s *MutexMatrix) Subtract(m Matrix) Matrix {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Subtract(m)
}

// Negative the negative of a matrix
func (s *MutexMatrix) Negative() Matrix {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Negative()
}

// Transpose swaps the rows and columns
func (s *MutexMatrix) Transpose() Matrix {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Transpose()
}

// Equal the two matrices are equal
func (s *MutexMatrix) Equal(m Matrix) bool {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Equal(m)
}

// NotEqual the two matrices are not equal
func (s *MutexMatrix) NotEqual(m Matrix) bool {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.NotEqual(m)
}

// Size of the matrix
func (s *MutexMatrix) Size() int {
	return s.matrix.Size()
}

// Values the number of elements in the matrix
func (s *MutexMatrix) Values() int {
	return s.matrix.Values()
}

// Clear removes all elements from a matrix
func (s *MutexMatrix) Clear() {
	s.RLock()
	defer s.RUnlock()

	s.matrix.Clear()
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *MutexMatrix) Enumerate() Enumerate {
	return s.matrix.Enumerate()
}

// Map replace each element with the result of applying a function to its value
func (s *MutexMatrix) Map() Map {
	return s.matrix.Map()
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *MutexMatrix) Element(r, c int) bool {
	s.RLock()
	defer s.RUnlock()

	return s.matrix.Element(r, c)
}
