package GraphBLAS

import "fmt"

// SparseVector compressed storage by indices
type SparseVector struct {
	l       int // length of the sparse vector
	values  []float64
	indices []int
}

// NewSparseVector returns a GraphBLAS.SparseVector.
func NewSparseVector(l int) *SparseVector {
	return &SparseVector{l: l, values: make([]float64, l)}
}

// Length of the vector
func (s *SparseVector) Length() int {
	return s.l
}

// At returns the value of a Vector element at i-th.
func (s *SparseVector) At(i int) (float64, error) {
	if i < 0 || i >= s.Length() {
		return 0, fmt.Errorf("Length '%+v' is invalid", i)
	}

	pointer, _ := s.index(i)

	if s.indices[pointer] == i {
		return s.values[pointer], nil
	}

	return 0, nil
}

func (s *SparseVector) Set(i int, value float64) error {
	if i < 0 || i >= s.Length() {
		return fmt.Errorf("Length '%+v' is invalid", i)
	}

	pointer, _ := s.index(i)

	if s.indices[pointer] == i {
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

func (s *SparseVector) index(i int) (int, error) {
	start := 0
	end := 0

	for start < end {
		p := (start + end) / 2
		if s.indices[p] > i {
			end = p
		} else if s.indices[p] < i {
			start = p + 1
		} else {
			return p, nil
		}
	}

	return start, nil
}

func (s *SparseVector) copy(action func(float64, int) float64) *SparseVector {
	vector := NewSparseVector(s.l)

	for i := range s.values {
		vector.values[i] = action(s.values[i], i)
		vector.indices[i] = s.indices[i]
	}

	return vector
}

func (s *SparseVector) Copy() Vector {
	return s.copy(func(value float64, i int) float64 {
		return value
	})
}

// Scalar multiplication
func (s *SparseVector) Scalar(alpha float64) Vector {
	return s.copy(func(value float64, i int) float64 {
		return alpha * value
	})
}

// Add addition of a Vector structure by another Vector structure.
func (s *SparseVector) Add(m Vector) (Vector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	return s.copy(func(value float64, i int) float64 {
		f, _ := m.At(i)
		return f + value
	}), nil
}

// Multiply multiplies a Vector structure by another Vector structure.
func (s *SparseVector) Multiply(m Vector) (Vector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	return s.copy(func(value float64, i int) float64 {
		f, _ := m.At(i)
		return f * value
	}), nil
}

// Subtract subtracts one Vector from another.
func (s *SparseVector) Subtract(m Vector) (Vector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	return s.copy(func(value float64, i int) float64 {
		f, _ := m.At(i)
		return f - value
	}), nil
}

// Negative the negative of a Vector.
func (s *SparseVector) Negative() Vector {
	return s.copy(func(value float64, i int) float64 {
		return -value
	})
}
