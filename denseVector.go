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
	return s.CopyArithmetic(func(value float64) float64 {
		return value
	})
}

// CopyArithmetic copies the matrix and applies a arithmetic function through all non-zero elements, order is not guaranteed
func (s *DenseVector) CopyArithmetic(action func(float64) float64) Matrix {
	vector := NewDenseVector(s.l)

	for i, v := range s.values {
		if v != 0.0 {
			vector.SetVec(i, action(v))
		} else {
			vector.SetVec(i, v)
		}
	}

	return vector
}

// Scalar multiplication of a vector by alpha
func (s *DenseVector) Scalar(alpha float64) Matrix {
	return scalar(s, alpha)
}

// Multiply multiplies a vector by another vector
func (s *DenseVector) Multiply(m Matrix) (Matrix, error) {
	return multiplyVector(s, m)
}

// Add addition of a vector by another vector
func (s *DenseVector) Add(m Matrix) (Matrix, error) {
	return add(s, m)
}

// Subtract subtracts one vector from another vector
func (s *DenseVector) Subtract(m Matrix) (Matrix, error) {
	return subtract(s, m)
}

// Negative the negative of a metrix
func (s *DenseVector) Negative() Matrix {
	return negative(s)
}

// Transpose swaps the rows and columns
func (s *DenseVector) Transpose() Matrix {
	matrix := newMatrix(s.Columns(), s.Rows(), nil, nil)

	return transpose(s, matrix)
}

// Equal the two vectors are equal
func (s *DenseVector) Equal(m Matrix) bool {
	return equal(s, m)
}

// NotEqual the two vectors are not equal
func (s *DenseVector) NotEqual(m Matrix) bool {
	return notEqual(s, m)
}
