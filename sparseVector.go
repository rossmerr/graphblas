// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

func init() {
	RegisterMatrix(reflect.TypeOf((*SparseVector)(nil)).Elem())
}

// SparseVector compressed storage by indices
type SparseVector struct {
	sync.RWMutex
	l       int // length of the sparse vector
	values  []float64
	indices []int
}

func (s SparseVector) String() string {
	return fmt.Sprintf("{l:%+v, values:%+v, indices:%+v}", s.l, s.values, s.indices)
}

// NewSparseVector returns a GraphBLAS.SparseVector
func NewSparseVector(l int) *SparseVector {
	return newSparseVector(l, 0)
}

// NewSparseVectorFromArray returns a GraphBLAS.SparseVector
func NewSparseVectorFromArray(data []float64) *SparseVector {
	l := len(data)
	s := newSparseVector(l, 0)

	for i := 0; i < l; i++ {
		s.SetVec(i, data[i])
	}

	return s
}

func newSparseVector(l, s int) *SparseVector {
	return &SparseVector{l: l, values: make([]float64, s), indices: make([]int, s)}
}

// Length of the vector
func (s *SparseVector) Length() int {
	return s.l
}

// AtVec returns the value of a vector element at i-th
func (s *SparseVector) AtVec(i int) float64 {
	s.RLock()
	defer s.RUnlock()

	return s.atVec(i)
}

func (s *SparseVector) atVec(i int) float64 {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	pointer, length, _ := s.index(i)

	if pointer < length && s.indices[pointer] == i {
		return s.values[pointer]
	}

	return 0
}

// SetVec sets the value at i-th of the vector
func (s *SparseVector) SetVec(i int, value float64) {
	s.Lock()
	defer s.Unlock()

	s.setVec(i, value)
}

func (s *SparseVector) setVec(i int, value float64) {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	pointer, length, _ := s.index(i)

	if pointer < length && s.indices[pointer] == i {
		if value == 0 {
			s.remove(pointer)
		} else {
			s.values[pointer] = value
		}
	} else {
		s.insert(pointer, i, value)
	}
}

// Columns the number of columns of the vector
func (s *SparseVector) Columns() int {
	return 1
}

// Rows the number of rows of the vector
func (s *SparseVector) Rows() int {
	return s.l
}

// Update does a At and Set on the vector element at r-th, c-th
func (s *SparseVector) Update(r, c int, f func(float64) float64) {
	s.Lock()
	defer s.Unlock()

	s.update(r, c, f)
}

func (s *SparseVector) update(r, c int, f func(float64) float64) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	v := s.atVec(r)
	s.setVec(r, f(v))
}

// At returns the value of a vector element at r-th, c-th
func (s *SparseVector) At(r, c int) (value float64) {
	s.RLock()
	defer s.RUnlock()

	s.update(r, c, func(v float64) float64 {
		value = v
		return v
	})

	return
}

// Set sets the value at r-th, c-th of the vector
func (s *SparseVector) Set(r, c int, value float64) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.setVec(r, value)
}

// ColumnsAt return the columns at c-th
func (s *SparseVector) ColumnsAt(c int) Vector {
	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.RLock()
	defer s.RUnlock()

	return s.copy()
}

// RowsAt return the rows at r-th
func (s *SparseVector) RowsAt(r int) Vector {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector(1)

	s.RLock()
	defer s.RUnlock()

	v := s.atVec(r)
	rows.setVec(0, v)

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *SparseVector) RowsAtToArray(r int) []float64 {
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

func (s *SparseVector) insert(pointer, i int, value float64) {
	if value == 0 {
		return
	}

	s.indices = append(s.indices[:pointer], append([]int{i}, s.indices[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]float64{value}, s.values[pointer:]...)...)
}

func (s *SparseVector) remove(pointer int) {
	s.indices = append(s.indices[:pointer], s.indices[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)
}

func (s *SparseVector) index(i int) (int, int, error) {
	length := len(s.indices)
	if i > length {
		return length, length, nil
	}

	start := 0
	end := length

	for start < end {
		p := (start + end) / 2
		if s.indices[p] > i {
			end = p
		} else if s.indices[p] < i {
			start = p + 1
		} else {
			return p, length, nil
		}
	}

	return start, length, nil
}

func (s *SparseVector) copy() *SparseVector {
	vector := newSparseVector(s.l, len(s.indices))

	for i := range s.values {
		vector.values[i] = s.values[i]
		vector.indices[i] = s.indices[i]
	}

	return vector
}

// Copy copies the vector
func (s *SparseVector) Copy() Matrix {
	s.RLock()
	defer s.RUnlock()

	return s.copy()
}

// Scalar multiplication of a vector by alpha
func (s *SparseVector) Scalar(alpha float64) Matrix {
	return Scalar(s, alpha)
}

// Multiply multiplies a vector by another vector
func (s *SparseVector) Multiply(m Matrix) Matrix {
	matrix := newMatrix(m.Rows(), s.Columns(), nil)
	MatrixMatrixMultiply(s, m, matrix)
	return matrix
}

// Add addition of a metrix by another metrix
func (s *SparseVector) Add(m Matrix) Matrix {
	matrix := s.Copy()
	Add(s, m, matrix)
	return matrix
}

// Subtract subtracts one metrix from another metrix
func (s *SparseVector) Subtract(m Matrix) Matrix {
	matrix := m.Copy()
	Subtract(s, m, matrix)
	return matrix
}

// Negative the negative of a metrix
func (s *SparseVector) Negative() Matrix {
	matrix := s.Copy()
	Negative(s, matrix)
	return matrix
}

// Transpose swaps the rows and columns
func (s *SparseVector) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), nil)
	Transpose(s, matrix)
	return matrix
}

// Equal the two metrics are equal
func (s *SparseVector) Equal(m Matrix) bool {
	return Equal(s, m)
}

// NotEqual the two metrix are not equal
func (s *SparseVector) NotEqual(m Matrix) bool {
	return NotEqual(s, m)
}

// Size of the vector
func (s *SparseVector) Size() int {
	return s.l
}

// Values the number of non-zero elements in the Vector
func (s *SparseVector) Values() int {
	return len(s.values)
}

// Apply modifies edge weights by the UnaryOperator
// C ⊕= f(A)
func (s *SparseVector) Apply(u UnaryOperator) {
	Apply(s, s, u)
}

// ReduceToScalar perform's a reduction on the Vector
func (s *SparseVector) ReduceToScalar() int {
	return 0
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *SparseVector) Enumerate() Enumerate {
	return s.iterator()
}

func (s *SparseVector) iterator() *sparseVectorIterator {
	i := &sparseVectorIterator{
		matrix: s,
		size:   len(s.values),
		last:   0,
	}
	return i
}

type sparseVectorIterator struct {
	matrix *SparseVector
	size   int
	last   int
	old    int
}

// HasNext checks the iterator has any more values
func (s *sparseVectorIterator) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *sparseVectorIterator) next() {
	s.old = s.last
	s.last++
}

// Next moves the iterator and returns the row, column and value
func (s *sparseVectorIterator) Next() (int, int, float64) {
	s.next()

	s.matrix.RLock()
	defer s.matrix.RUnlock()

	return s.matrix.indices[s.old], 0, s.matrix.values[s.old]
}

// Map replace each element with the result of applying a function to its value
func (s *SparseVector) Map() Map {
	t := s.iterator()
	i := &sparseVectorMap{t}
	return i
}

type sparseVectorMap struct {
	*sparseVectorIterator
}

// HasNext checks the iterator has any more values
func (s *sparseVectorMap) HasNext() bool {
	return s.sparseVectorIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *sparseVectorMap) Map(f func(int, int, float64) float64) {
	s.next()

	s.matrix.Lock()
	defer s.matrix.Unlock()

	value := f(s.matrix.indices[s.old], 0, s.matrix.values[s.old])
	if value != 0 {
		s.matrix.values[s.old] = value
	} else {
		s.matrix.remove(s.old)
	}
}
