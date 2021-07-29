// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolean

import (
	"context"
	"log"
)

// DenseMatrix a dense matrix
type DenseMatrix struct {
	c    int // number of rows in the sparse matrix
	r    int // number of columns in the sparse matrix
	data [][]bool
}

// NewDenseMatrix returns a DenseMatrix
func NewDenseMatrix(r, c int) *DenseMatrix {
	return newMatrix(r, c, nil)
}

// NewDenseMatrixFromArray returns a DenseMatrix
func NewDenseMatrixFromArray(data [][]bool) *DenseMatrix {
	r := len(data)
	c := len(data[0])
	s := &DenseMatrix{data: data, r: r, c: c}

	return s
}

func newMatrix(r, c int, initialise func([]bool, int)) *DenseMatrix {
	s := &DenseMatrix{data: make([][]bool, r), r: r, c: c}

	for i := 0; i < r; i++ {
		s.data[i] = make([]bool, c)

		if initialise != nil {
			initialise(s.data[i], i)
		}
	}

	return s
}

// Columns the number of columns of the matrix
func (s *DenseMatrix) Columns() int {
	return s.c
}

// Rows the number of rows of the matrix
func (s *DenseMatrix) Rows() int {
	return s.r
}

// Update does a At and Set on the matrix element at r-th, c-th
func (s *DenseMatrix) Update(r, c int, f func(bool) bool) {
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
func (s *DenseMatrix) At(r, c int) bool {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	return s.data[r][c]
}

// Set sets the value at r-th, c-th of the matrix
func (s *DenseMatrix) Set(r, c int, value bool) {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	s.data[r][c] = value
}

// ColumnsAt return the columns at c-th
func (s *DenseMatrix) ColumnsAt(c int) Vector {
	if c < 0 || c >= s.Columns() {
		log.Panicf("Column '%+v' is invalid", c)
	}

	columns := NewDenseVector(s.r)

	for r := 0; r < s.r; r++ {
		columns.SetVec(r, s.data[r][c])
	}

	return columns
}

// RowsAt return the rows at r-th
func (s *DenseMatrix) RowsAt(r int) Vector {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := NewDenseVector(s.c)

	for i := 0; i < s.c; i++ {
		rows.SetVec(i, s.data[r][i])
	}

	return rows
}

// RowsAtToArray return the rows at r-th
func (s *DenseMatrix) RowsAtToArray(r int) []bool {
	if r < 0 || r >= s.Rows() {
		log.Panicf("Row '%+v' is invalid", r)
	}

	rows := make([]bool, s.c)

	for i := 0; i < s.c; i++ {
		rows[i] = s.data[r][i]
	}

	return rows
}

// Copy copies the matrix
func (s *DenseMatrix) Copy() Matrix {
	matrix := newMatrix(s.Rows(), s.Columns(), func(row []bool, r int) {
		for c := 0; c < s.Columns(); c++ {
			row[c] = s.data[r][c]
		}
	})

	return matrix
}

// Transpose swaps the rows and columns
func (s *DenseMatrix) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), nil)
	Transpose(context.Background(), s, nil, matrix)
	return matrix
}

// Equal the two matrices are equal
func (s *DenseMatrix) Equal(m Matrix) bool {
	return Equal(context.Background(), s, m)
}

// NotEqual the two matrices are not equal
func (s *DenseMatrix) NotEqual(m Matrix) bool {
	return NotEqual(context.Background(), s, m)
}

// Size of the matrix
func (s *DenseMatrix) Size() int {
	return s.r * s.c
}

// Values the number of elements in the matrix
func (s *DenseMatrix) Values() int {
	return s.r * s.c
}

// Clear removes all elements from a matrix
func (s *DenseMatrix) Clear() {
	s.data = make([][]bool, s.r)
	for i := 0; i < s.r; i++ {
		s.data[i] = make([]bool, s.c)
	}
}

// RawMatrix returns the raw matrix
func (s *DenseMatrix) RawMatrix() [][]bool {
	return s.data
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *DenseMatrix) Enumerate() Enumerate {
	return s.iterator()
}

func (s *DenseMatrix) iterator() *denseMatrixIterator {
	i := &denseMatrixIterator{
		matrix: s,
		size:   s.Values(),
		last:   0,
		c:      0,
		r:      0,
	}
	return i
}

type denseMatrixIterator struct {
	matrix *DenseMatrix
	size   int
	last   int
	c      int
	r      int
	cOld   int
	rOld   int
}

// HasNext checks the iterator has any more values
func (s *denseMatrixIterator) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *denseMatrixIterator) next() {
	if s.c == s.matrix.Columns() {
		s.c = 0
		s.r++
	}
	s.cOld = s.c
	s.c++
	s.last++
}

// Next moves the iterator and returns the row, column and value
func (s *denseMatrixIterator) Next() (int, int, bool) {
	s.next()

	return s.r, s.cOld, s.matrix.At(s.r, s.cOld)
}

// Map replace each element with the result of applying a function to its value
func (s *DenseMatrix) Map() Map {
	t := s.iterator()
	i := &denseMatrixMap{t}
	return i
}

type denseMatrixMap struct {
	*denseMatrixIterator
}

// HasNext checks the iterator has any more values
func (s *denseMatrixMap) HasNext() bool {
	return s.denseMatrixIterator.HasNext()
}

// Map move the iterator and uses a higher order function to changes the elements current value
func (s *denseMatrixMap) Map(f func(int, int, bool) bool) {
	s.next()

	s.matrix.Set(s.r, s.cOld, f(s.r, s.cOld, s.matrix.At(s.r, s.cOld)))
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *DenseMatrix) Element(r, c int) bool {
	return s.element(r, c)
}

func (s *DenseMatrix) element(r, c int) bool {
	return s.At(r, c)
}
