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
	RegisterMatrix(reflect.TypeOf((*SparseVector)(nil)).Elem())
}

// SparseVector compressed storage by indices
type SparseVector struct {
	l       int // length of the sparse vector
	values  []bool
	indices []int
}

// NewSparseVector returns a SparseVector
func NewSparseVector(l int) *SparseVector {
	return newSparseVector(l, 0)
}

// NewSparseVectorFromArray returns a SparseVector
func NewSparseVectorFromArray(data []bool) *SparseVector {
	l := len(data)
	s := newSparseVector(l, 0)

	for i := 0; i < l; i++ {
		s.SetVec(i, data[i])
	}

	return s
}

func newSparseVector(l, s int) *SparseVector {
	return &SparseVector{l: l, values: make([]bool, s), indices: make([]int, s)}
}

// Length of the vector
func (s *SparseVector) Length() int {
	return s.l
}

// AtVec returns the value of a vector element at i-th
func (s *SparseVector) AtVec(i int) bool {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	pointer, length, _ := s.index(i)

	if pointer < length && s.indices[pointer] == i {
		return s.values[pointer]
	}

	return false
}

// SetVec sets the value at i-th of the vector
func (s *SparseVector) SetVec(i int, value bool) {
	if i < 0 || i >= s.Length() {
		log.Panicf("Length '%+v' is invalid", i)
	}

	pointer, length, _ := s.index(i)

	if pointer < length && s.indices[pointer] == i {
		if value == false {
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
func (s *SparseVector) Update(r, c int, f func(bool) bool) {
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
func (s *SparseVector) At(r, c int) (value bool) {
	s.Update(r, c, func(v bool) bool {
		value = v
		return v
	})

	return
}

// Set sets the value at r-th, c-th of the vector
func (s *SparseVector) Set(r, c int, value bool) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.SetVec(r, value)
}

// ColumnsAt return the columns at c-th
func (s *SparseVector) ColumnsAt(c int) Vector {
	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	return s.copy()
}

// RowsAt return the rows at r-th
func (s *SparseVector) RowsAt(r int) Vector {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewSparseVector(1)

	v := s.AtVec(r)
	rows.SetVec(0, v)

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *SparseVector) RowsAtToArray(r int) []bool {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]bool, 1)

	v := s.AtVec(r)
	rows[0] = v

	return rows
}

func (s *SparseVector) insert(pointer, i int, value bool) {
	if value == false {
		return
	}

	s.indices = append(s.indices[:pointer], append([]int{i}, s.indices[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]bool{value}, s.values[pointer:]...)...)
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
	return s.copy()
}

// Transpose swaps the rows and columns
func (s *SparseVector) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), nil)
	Transpose(context.Background(), s, nil, matrix)
	return matrix
}

// Equal the two metrics are equal
func (s *SparseVector) Equal(m Matrix) bool {
	return Equal(context.Background(), s, m)
}

// NotEqual the two metrix are not equal
func (s *SparseVector) NotEqual(m Matrix) bool {
	return NotEqual(context.Background(), s, m)
}

// Size of the vector
func (s *SparseVector) Size() int {
	return s.l
}

// Values the number of non-zero elements in the Vector
func (s *SparseVector) Values() int {
	return len(s.values)
}

// Clear removes all elements from a vector
func (s *SparseVector) Clear() {
	s.values = make([]bool, 0)
	s.indices = make([]int, 0)

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
func (s *sparseVectorIterator) Next() (int, int, bool) {
	s.next()

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
func (s *sparseVectorMap) Map(f func(int, int, bool) bool) {
	s.next()

	value := f(s.matrix.indices[s.old], 0, s.matrix.values[s.old])
	if value != false {
		s.matrix.values[s.old] = value
	} else {
		s.matrix.remove(s.old)
	}
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *SparseVector) Element(r, c int) bool {
	return s.AtVec(r)
}
