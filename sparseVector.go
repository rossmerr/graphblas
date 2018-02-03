// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import "fmt"

// SparseVector compressed storage by indices
type SparseVector struct {
	l       int // length of the sparse vector
	values  []float64
	indices []int
}

// NewSparseVector returns a GraphBLAS.SparseVector
func NewSparseVector(l int) *SparseVector {
	return newSparseVector(l, 0)
}

func newSparseVector(l, s int) *SparseVector {
	return &SparseVector{l: l, values: make([]float64, s), indices: make([]int, s)}
}

// Length of the vector
func (s *SparseVector) Length() int {
	return s.l
}

// AtVec returns the value of a vector element at i-th
func (s *SparseVector) AtVec(i int) (float64, error) {
	if i < 0 || i >= s.Length() {
		return 0, fmt.Errorf("Length '%+v' is invalid", i)
	}

	pointer, length, _ := s.index(i)

	if pointer < length && s.indices[pointer] == i {
		return s.values[pointer], nil
	}

	return 0, nil
}

// SetVec sets the value at i-th of the vector
func (s *SparseVector) SetVec(i int, value float64) error {
	if i < 0 || i >= s.Length() {
		return fmt.Errorf("Length '%+v' is invalid", i)
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

	return nil
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
func (s *SparseVector) Update(r, c int, f func(float64) float64) error {
	if r < 0 || r >= s.Rows() {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	v, _ := s.AtVec(r)
	return s.SetVec(r, f(v))
}

// At returns the value of a vector element at r-th, c-th
func (s *SparseVector) At(r, c int) (float64, error) {
	value := 0.0
	err := s.Update(r, c, func(v float64) float64 {
		value = v
		return v
	})

	return value, err
}

// Set sets the value at r-th, c-th of the vector
func (s *SparseVector) Set(r, c int, value float64) error {
	if r < 0 || r >= s.Rows() {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	return s.SetVec(r, value)
}

// ColumnsAt return the columns at c-th
func (s *SparseVector) ColumnsAt(c int) (Vector, error) {
	if c < 0 || c >= s.Columns() {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	return s.copy(func(value float64, i int) float64 {
		return value
	}), nil
}

// RowsAt return the rows at r-th
func (s *SparseVector) RowsAt(r int) (Vector, error) {
	if r < 0 || r >= s.Rows() {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	v, _ := s.AtVec(1)
	rows := NewSparseVector(1)
	rows.SetVec(0, v)

	return rows, nil
}

// Iterator iterates through all non-zero elements, order is not guaranteed
func (s *SparseVector) Iterator(i func(r, c int, v float64) bool) bool {
	for c := 0; c < s.Columns(); c++ {
		for r := 0; r < s.Rows(); r++ {
			v, _ := s.At(r, c)
			if v != 0.0 {
				if i(r, c, v) == false {
					return false
				}
			}
		}
	}

	return true
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

func (s *SparseVector) copy(action func(float64, int) float64) *SparseVector {
	vector := newSparseVector(s.l, len(s.indices))

	for i := range s.values {
		vector.values[i] = action(s.values[i], i)
		vector.indices[i] = s.indices[i]
	}

	return vector
}

// Copy copies the vector
func (s *SparseVector) Copy() Matrix {
	return s.copy(func(value float64, i int) float64 {
		return value
	})
}

// Scalar multiplication of a vector by alpha
func (s *SparseVector) Scalar(alpha float64) Matrix {
	return s.copy(func(value float64, i int) float64 {
		return alpha * value
	})
}

// Multiply multiplies a vector by another vector
func (s *SparseVector) Multiply(m Matrix) (Matrix, error) {
	if s.Rows() != m.Columns() {
		return nil, fmt.Errorf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	matrix := newCSRMatrix(m.Rows(), s.Columns(), 0)

	for r := 0; r < m.Rows(); r++ {
		rows, _ := m.RowsAt(r)
		for c := 0; c < s.Columns(); c++ {
			column, _ := s.ColumnsAt(c)
			sum := 0.0
			for l := 0; l < rows.Length(); l++ {
				vC, _ := column.AtVec(l)
				vR, _ := rows.AtVec(l)
				sum += vR * vC
			}

			matrix.Set(r, c, sum)
		}

	}
	return matrix, nil
}

// Add addition of a metrix by another metrix
func (s *SparseVector) Add(m Matrix) (Matrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, r int) {
		for c := 0; c < m.Columns(); c++ {
			v, _ := m.At(r, c)
			v2, _ := s.At(r, c)
			row[c] = v2 + v
		}
	})

	return matrix, nil
}

// Subtract subtracts one metrix from another metrix
func (s *SparseVector) Subtract(m Matrix) (Matrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, r int) {
		for c := 0; c < m.Columns(); c++ {
			v, _ := m.At(r, c)
			v2, _ := s.At(r, c)
			row[c] = v2 - v
		}
	})

	return matrix, nil
}

// Negative the negative of a metrix
func (s *SparseVector) Negative() Matrix {
	vector := NewSparseVector(s.l)

	for i := 0; i < s.l; i++ {
		v1, _ := s.AtVec(i)
		vector.SetVec(i, -v1)
	}

	return vector
}

// Transpose swaps the rows and columns
func (s *SparseVector) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), func(row []float64, c int) {
		for r := 0; r < s.Rows(); r++ {
			v, _ := s.At(r, c)
			row[r] = v
		}
	})

	return matrix
}

// Equal the two metrics are equal
func (s *SparseVector) Equal(m Matrix) bool {
	if s.Columns() != m.Columns() {
		return false
	}

	if s.Rows() != m.Rows() {
		return false
	}

	return s.Iterator(func(r, c int, v float64) bool {
		value, _ := m.At(r, c)
		if v != value {

			return false
		}
		return true
	})
}

// NotEqual the two metrix are not equal
func (s *SparseVector) NotEqual(m Matrix) bool {
	return !s.Equal(m)
}
