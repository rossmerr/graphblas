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

// type Ordered interface {
// 	Integer | Float | ~string
// }

func init() {
	RegisterMatrix(reflect.TypeOf((*SparseVector[float64])(nil)).Elem())
}

// SparseVector compressed storage by indices
type SparseVector[T constraints.Number] struct {
	l       int // length of the sparse vector
	values  []T
	indices []int
}

// NewSparseVector returns a SparseVector
func NewSparseVector[T constraints.Number](l int) *SparseVector[T] {
	return newSparseVector[T](l, 0)
}

// NewSparseVectorFromArray returns a SparseVector
func NewSparseVectorFromArray[T constraints.Number](data []T) *SparseVector[T] {
	l := len(data)
	s := newSparseVector[T](l, 0)

	for i := 0; i < l; i++ {
		s.SetVec(i, data[i])
	}

	return s
}

func newSparseVector[T constraints.Number](l, s int) *SparseVector[T] {
	return &SparseVector[T]{l: l, values: make([]T, s), indices: make([]int, s)}
}

// Length of the vector
func (s *SparseVector[T]) Length() int {
	return s.l
}

// AtVec returns the value of a vector element at i-th
func (s *SparseVector[T]) AtVec(i int) T {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	pointer, length, _ := s.index(i)

	if pointer < length && s.indices[pointer] == i {
		return s.values[pointer]
	}

	return Zero[T]()
}

// SetVec sets the value at i-th of the vector
func (s *SparseVector[T]) SetVec(i int, value T) {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	pointer, length, _ := s.index(i)

	if pointer < length && s.indices[pointer] == i {
		if IsZero(value) {
			s.remove(pointer)
		} else {
			s.values[pointer] = value
		}
	} else {
		s.insert(pointer, i, Zero[T]())
	}
}

// Columns the number of columns of the vector
func (s *SparseVector[T]) Columns() int {
	return 1
}

// Rows the number of rows of the vector
func (s *SparseVector[T]) Rows() int {
	return s.l
}

// Update does a At and Set on the vector element at r-th, c-th
func (s *SparseVector[T]) Update(r, c int, f func(T) T) {
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
func (s *SparseVector[T]) At(r, c int) (value T) {
	s.Update(r, c, func(v T) T {
		value = v
		return v
	})

	return
}

// Set sets the value at r-th, c-th of the vector
func (s *SparseVector[T]) Set(r, c int, value T) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.SetVec(r, value)
}

// ColumnsAt return the columns at c-th
func (s *SparseVector[T]) ColumnsAt(c int) Vector[T] {
	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	return s.copy()
}

// RowsAt return the rows at r-th
func (s *SparseVector[T]) RowsAt(r int) Vector[T] {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector[T](1)

	v := s.AtVec(r)
	rows.SetVec(0, v)

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *SparseVector[T]) RowsAtToArray(r int) []T {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]T, 1)

	v := s.AtVec(r)
	rows[0] = v

	return rows
}

func (s *SparseVector[T]) insert(pointer, i int, value T) {
	if IsZero(value) {
		return
	}

	s.indices = append(s.indices[:pointer], append([]int{i}, s.indices[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]T{value}, s.values[pointer:]...)...)
}

func (s *SparseVector[T]) remove(pointer int) {
	s.indices = append(s.indices[:pointer], s.indices[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)
}

func (s *SparseVector[T]) index(i int) (int, int, error) {
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

func (s *SparseVector[T]) copy() *SparseVector[T] {
	vector := newSparseVector[T](s.l, len(s.indices))

	for i := range s.values {
		vector.values[i] = s.values[i]
		vector.indices[i] = s.indices[i]
	}

	return vector
}

// Copy copies the vector
func (s *SparseVector[T]) Copy() Matrix[T] {
	return s.copy()
}

// Scalar multiplication of a vector by alpha
func (s *SparseVector[T]) Scalar(alpha T) Matrix[T] {
	return Scalar[T](context.Background(), s, alpha)
}

// Multiply multiplies a vector by another vector
func (s *SparseVector[T]) Multiply(m Matrix[T]) Matrix[T] {
	matrix := newMatrix[T](m.Rows(), s.Columns(), nil)
	MatrixMatrixMultiply[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Add addition of a metrix by another metrix
func (s *SparseVector[T]) Add(m Matrix[T]) Matrix[T] {
	matrix := s.Copy()
	Add[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Subtract subtracts one metrix from another metrix
func (s *SparseVector[T]) Subtract(m Matrix[T]) Matrix[T] {
	matrix := m.Copy()
	Subtract[T](context.Background(), s, m, nil, matrix)
	return matrix
}

// Negative the negative of a metrix
func (s *SparseVector[T]) Negative() Matrix[T] {
	matrix := s.Copy()
	Negative[T](context.Background(), s, nil, matrix)
	return matrix
}

// Transpose swaps the rows and columns
func (s *SparseVector[T]) Transpose() Matrix[T] {
	matrix := newMatrix[T](s.Columns(), s.Rows(), nil)
	Transpose[T](context.Background(), s, nil, matrix)
	return matrix
}

// Equal the two metrics are equal
func (s *SparseVector[T]) Equal(m Matrix[T]) bool {
	return Equal[T](context.Background(), s, m)
}

// NotEqual the two metrix are not equal
func (s *SparseVector[T]) NotEqual(m Matrix[T]) bool {
	return NotEqual[T](context.Background(), s, m)
}

// Size of the vector
func (s *SparseVector[T]) Size() int {
	return s.l
}

// Values the number of non-zero elements in the Vector
func (s *SparseVector[T]) Values() int {
	return len(s.values)
}

// Clear removes all elements from a vector
func (s *SparseVector[T]) Clear() {
	s.values = make([]T, 0)
	s.indices = make([]int, 0)

}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *SparseVector[T]) Enumerate() Enumerate[T] {
	return s.iterator()
}

func (s *SparseVector[T]) iterator() *sparseVectorIterator[T] {
	i := &sparseVectorIterator[T]{
		matrix: s,
		size:   len(s.values),
		last:   0,
	}
	return i
}

type sparseVectorIterator[T constraints.Number] struct {
	matrix *SparseVector[T]
	size   int
	last   int
	old    int
}

// HasNext checks the iterator has any more values
func (s *sparseVectorIterator[T]) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *sparseVectorIterator[T]) next() {
	s.old = s.last
	s.last++
}

// Next moves the iterator and returns the row, column and value
func (s *sparseVectorIterator[T]) Next() (int, int, T) {
	s.next()

	return s.matrix.indices[s.old], 0, s.matrix.values[s.old]
}

// Map replace each element with the result of applying a function to its value
func (s *SparseVector[T]) Map() Map[T] {
	t := s.iterator()
	i := &sparseVectorMap[T]{t}
	return i
}

type sparseVectorMap[T constraints.Number] struct {
	*sparseVectorIterator[T]
}

// HasNext checks the iterator has any more values
func (s *sparseVectorMap[T]) HasNext() bool {
	return s.sparseVectorIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *sparseVectorMap[T]) Map(f func(int, int, T) T) {
	s.next()

	value := f(s.matrix.indices[s.old], 0, s.matrix.values[s.old])
	if !IsZero(value) {
		s.matrix.values[s.old] = value
	} else {
		s.matrix.remove(s.old)
	}
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *SparseVector[T]) Element(r, c int) bool {
	return s.AtVec(r) > Default[T]()
}
