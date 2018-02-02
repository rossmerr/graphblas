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
func (s *SparseVector) Copy() Vector {
	return s.copy(func(value float64, i int) float64 {
		return value
	})
}

// Scalar multiplication of a vector by alpha
func (s *SparseVector) Scalar(alpha float64) Vector {
	return s.copy(func(value float64, i int) float64 {
		return alpha * value
	})
}

// Multiply multiplies a vector by another vector
func (s *SparseVector) Multiply(m Vector) (Vector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	return s.copy(func(value float64, i int) float64 {
		f, _ := m.AtVec(i)
		return value * f
	}), nil
}

// Add addition of a vector by another vector
func (s *SparseVector) Add(m Vector) (Vector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	return s.copy(func(value float64, i int) float64 {
		f, _ := m.AtVec(i)
		return value + f
	}), nil
}

// Subtract subtracts one vector from another vector
func (s *SparseVector) Subtract(m Vector) (Vector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	return s.copy(func(value float64, i int) float64 {
		f, _ := m.AtVec(i)
		return value - f
	}), nil
}

// Negative the negative of a vector
func (s *SparseVector) Negative() Vector {
	return s.copy(func(value float64, i int) float64 {
		return -value
	})
}

// Equal the two vectors are equal
func (s *SparseVector) Equal(m Vector) bool {
	if m.Length() != s.Length() {
		return false
	}

	for i := 0; i < s.Length(); i++ {
		v1, _ := s.AtVec(i)
		v2, _ := m.AtVec(i)
		if v1 != v2 {
			return false
		}
	}

	return true
}

// NotEqual the two vectors are not equal
func (s *SparseVector) NotEqual(m Vector) bool {
	return !s.Equal(m)
}
