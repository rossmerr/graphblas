// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"fmt"
	"log"
	"sync"
)

// DenseVector a vector
type DenseVector struct {
	sync.RWMutex
	l      int // length of the sparse vector
	values []float64
}

func (s DenseVector) String() string {
	return fmt.Sprintf("{l:%+v, values:%+v}", s.l, s.values)
}

// NewDenseVector returns a GraphBLAS.DenseVector
func NewDenseVector(l int) *DenseVector {
	return &DenseVector{l: l, values: make([]float64, l)}
}

// NewDenseVectorFromArray returns a GraphBLAS.SparseVector
func NewDenseVectorFromArray(data []float64) *DenseVector {
	arr := make([]float64, 0)
	arr = append(arr, data...)
	return &DenseVector{l: len(data), values: arr}
}

// AtVec returns the value of a vector element at i-th
func (s *DenseVector) AtVec(i int) float64 {
	s.RLock()
	defer s.RUnlock()

	return s.atVec(i)
}

func (s *DenseVector) atVec(i int) float64 {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	return s.values[i]
}

// SetVec sets the value at i-th of the vector
func (s *DenseVector) SetVec(i int, value float64) {
	s.Lock()
	defer s.Unlock()

	s.setVec(i, value)
}

func (s *DenseVector) setVec(i int, value float64) {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	s.values[i] = value
}

// Size of the vector
func (s *DenseVector) Size() int {
	return s.l
}

// Length of the vector
func (s *DenseVector) Length() int {
	return s.l
}

// Columns the number of columns of the vector
func (s *DenseVector) Columns() int {
	return 1
}

// Rows the number of rows of the vector
func (s *DenseVector) Rows() int {
	return s.l
}

// Update does a At and Set on the vector element at r-th, c-th
func (s *DenseVector) Update(r, c int, f func(float64) float64) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.Lock()
	defer s.Unlock()

	v := s.atVec(r)
	s.setVec(r, f(v))
}

// At returns the value of a vector element at r-th, c-th
func (s *DenseVector) At(r, c int) (value float64) {
	s.RLock()
	defer s.RUnlock()

	return s.atVec(r)
}

// Set sets the value at r-th, c-th of the vector
func (s *DenseVector) Set(r, c int, value float64) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.Lock()
	defer s.Unlock()

	s.setVec(r, value)
}

// ColumnsAt return the columns at c-th
func (s *DenseVector) ColumnsAt(c int) Vector {
	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.RLock()
	defer s.RUnlock()

	return s.copy()
}

// RowsAt return the rows at r-th
func (s *DenseVector) RowsAt(r int) Vector {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewDenseVector(1)

	s.RLock()
	defer s.RUnlock()

	v := s.atVec(r)
	rows.setVec(0, v)

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *DenseVector) RowsAtToArray(r int) []float64 {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]float64, 1)

	s.RLock()
	defer s.RUnlock()

	v := s.atVec(r)
	rows[0] = v

	return rows
}

func (s *DenseVector) copy() *DenseVector {
	vector := NewDenseVector(s.l)

	for i, v := range s.values {
		vector.setVec(i, v)
	}

	return vector
}

// Copy copies the vector
func (s *DenseVector) Copy() Matrix {
	vector := NewDenseVector(s.l)

	s.RLock()
	defer s.RUnlock()

	for i, v := range s.values {
		if v != 0.0 {
			vector.setVec(i, v)
		} else {
			vector.setVec(i, v)
		}
	}

	return vector
}

// Scalar multiplication of a vector by alpha
func (s *DenseVector) Scalar(alpha float64) Matrix {
	return Scalar(s, alpha)
}

// Multiply multiplies a vector by another vector
func (s *DenseVector) Multiply(m Matrix) Matrix {
	matrix := newMatrix(m.Rows(), s.Columns(), nil)
	MatrixMatrixMultiply(s, m, matrix)
	return matrix
}

// Add addition of a vector by another vector
func (s *DenseVector) Add(m Matrix) Matrix {
	matrix := s.Copy()
	Add(s, m, matrix)
	return matrix
}

// Subtract subtracts one vector from another vector
func (s *DenseVector) Subtract(m Matrix) Matrix {
	matrix := m.Copy()
	Subtract(s, m, matrix)
	return matrix
}

// Negative the negative of a metrix
func (s *DenseVector) Negative() Matrix {
	matrix := s.Copy()
	Negative(s, matrix)
	return matrix
}

// Transpose swaps the rows and columns
func (s *DenseVector) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), nil)
	Transpose(s, matrix)
	return matrix
}

// Equal the two vectors are equal
func (s *DenseVector) Equal(m Matrix) bool {
	return Equal(s, m)
}

// NotEqual the two vectors are not equal
func (s *DenseVector) NotEqual(m Matrix) bool {
	return NotEqual(s, m)
}

// Values the number of elements in the Vector
func (s *DenseVector) Values() int {
	return s.l
}

// Apply modifies edge weights by the UnaryOperator
// C ⊕= f(A)
func (s *DenseVector) Apply(u UnaryOperator) {
	Apply(s, s, u)
}

// ReduceToScalar perform's a reduction on the Vector
func (s *DenseVector) ReduceToScalar() int {
	return 0
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *DenseVector) Enumerate() Enumerate {
	return s.iterator()
}

func (s *DenseVector) iterator() *denseVectorIterator {
	i := &denseVectorIterator{
		matrix: s,
		size:   s.Values(),
		last:   0,
		c:      0,
		r:      0,
	}
	return i
}

type denseVectorIterator struct {
	matrix *DenseVector
	size   int
	last   int
	c      int
	r      int
	rOld   int
}

// HasNext checks the iterator has any more values
func (s *denseVectorIterator) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *denseVectorIterator) next() {
	if s.r == s.matrix.Rows() {
		s.r = 0
		s.c++
	}
	s.rOld = s.r
	s.r++
	s.last++
}

// Next moves the iterator and returns the row, column and value
func (s *denseVectorIterator) Next() (int, int, float64) {
	s.next()

	s.matrix.RLock()
	defer s.matrix.RUnlock()

	return s.rOld, 0, s.matrix.atVec(s.rOld)
}

// Map replace each element with the result of applying a function to its value
func (s *DenseVector) Map() Map {
	t := s.iterator()
	i := &denseVectorMap{t}
	return i
}

type denseVectorMap struct {
	*denseVectorIterator
}

// HasNext checks the iterator has any more values
func (s *denseVectorMap) HasNext() bool {
	return s.denseVectorIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *denseVectorMap) Map(f func(int, int, float64) float64) {
	s.next()

	s.matrix.Lock()
	defer s.matrix.Unlock()

	s.matrix.setVec(s.rOld, f(s.rOld, 0, s.matrix.atVec(s.rOld)))
}
