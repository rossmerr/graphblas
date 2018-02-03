// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import "fmt"

// DenseVector a vector
type DenseVector struct {
	l      int // length of the sparse vector
	values []float64
}

// NewDenseVector returns a GraphBLAS.DenseVector
func NewDenseVector(l int) *DenseVector {
	return &DenseVector{l: l, values: make([]float64, l)}
}

// Length of the vector
func (s *DenseVector) Length() int {
	return s.l
}

// AtVec returns the value of a vector element at i-th
func (s *DenseVector) AtVec(i int) (float64, error) {
	if i < 0 || i >= s.Length() {
		return 0, fmt.Errorf("Length '%+v' is invalid", i)
	}

	return s.values[i], nil
}

// SetVec sets the value at i-th of the vector
func (s *DenseVector) SetVec(i int, value float64) error {
	if i < 0 || i >= s.Length() {
		return fmt.Errorf("Length '%+v' is invalid", i)
	}

	s.values[i] = value

	return nil
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
func (s *DenseVector) Update(r, c int, f func(float64) float64) error {
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
func (s *DenseVector) At(r, c int) (float64, error) {
	value := 0.0
	err := s.Update(r, c, func(v float64) float64 {
		value = v
		return v
	})

	return value, err
}

// Set sets the value at r-th, c-th of the vector
func (s *DenseVector) Set(r, c int, value float64) error {
	if r < 0 || r >= s.Rows() {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	return s.SetVec(r, value)
}

// ColumnsAt return the columns at c-th
func (s *DenseVector) ColumnsAt(c int) (Vector, error) {
	if c < 0 || c >= s.Columns() {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	return s.copy(), nil
}

// RowsAt return the rows at r-th
func (s *DenseVector) RowsAt(r int) (Vector, error) {
	if r < 0 || r >= s.Rows() {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	v, _ := s.AtVec(1)
	rows := NewDenseVector(1)
	rows.SetVec(0, v)

	return rows, nil
}

// Iterator iterates through all non-zero elements, order is not guaranteed
func (s *DenseVector) Iterator(i func(r, c int, v float64) bool) bool {
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

func (s *DenseVector) copy() *DenseVector {
	vector := NewDenseVector(s.l)

	for i, v := range s.values {
		vector.SetVec(i, v)
	}

	return vector
}

// Copy copies the vector
func (s *DenseVector) Copy() Matrix {
	vector := NewDenseVector(s.l)

	for i, v := range s.values {
		vector.SetVec(i, v)
	}

	return vector
}

// Scalar multiplication of a vector by alpha
func (s *DenseVector) Scalar(alpha float64) Matrix {
	vector := NewDenseVector(s.l)

	for i, v := range s.values {
		vector.SetVec(i, alpha*v)
	}

	return vector
}

// Multiply multiplies a vector by another vector
func (s *DenseVector) Multiply(m Matrix) (Matrix, error) {
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

// Add addition of a vector by another vector
func (s *DenseVector) Add(m Matrix) (Matrix, error) {
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

// Subtract subtracts one vector from another vector
func (s *DenseVector) Subtract(m Matrix) (Matrix, error) {
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
func (s *DenseVector) Negative() Matrix {
	vector := NewDenseVector(s.l)

	for i := 0; i < s.l; i++ {
		v1, _ := s.AtVec(i)
		vector.SetVec(i, -v1)
	}

	return vector
}

// Transpose swaps the rows and columns
func (s *DenseVector) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), func(row []float64, c int) {
		for r := 0; r < s.Rows(); r++ {
			v, _ := s.At(r, c)
			row[r] = v
		}
	})

	return matrix
}

// Equal the two vectors are equal
func (s *DenseVector) Equal(m Matrix) bool {
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

// NotEqual the two vectors are not equal
func (s *DenseVector) NotEqual(m Matrix) bool {
	return !s.Equal(m)
}
