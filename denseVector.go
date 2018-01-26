package GraphBLAS

import "fmt"

// DenseVector a vector
type DenseVector struct {
	l    int
	data []float64
}

// NewVector returns a GraphBLAS.DenseVector.
func NewDenseVector(l int) *DenseVector {
	return &DenseVector{l: l, data: make([]float64, l)}
}

// Length of the vector
func (s *DenseVector) Length() int {
	return s.l
}

func (s *DenseVector) At(i int) (float64, error) {
	if i < 0 || i >= s.Length() {
		return 0, fmt.Errorf("Length '%+v' is invalid", i)
	}

	return s.data[i], nil
}

func (s *DenseVector) Set(i int, value float64) error {
	if i < 0 || i >= s.Length() {
		return fmt.Errorf("Length '%+v' is invalid", i)
	}

	s.data[i] = value

	return nil
}

// Scalar multiplication
func (s *DenseVector) Scalar(alpha float64) *DenseVector {
	vector := NewDenseVector(s.l)

	for i, v := range s.data {
		vector.Set(i, alpha*v)
	}

	return vector
}

// Multiply multiplies a Vector structure by another Vector structure.
func (s *DenseVector) Multiply(m *DenseVector) (*DenseVector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	vector := NewDenseVector(s.l)

	for i := 0; i < s.l; i++ {
		v1, _ := s.At(i)
		v2, _ := m.At(i)
		vector.Set(i, v1*v2)
	}

	return vector, nil
}

// Add addition of a Vector structure by another Vector structure.
func (s *DenseVector) Add(m *DenseVector) (*DenseVector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	vector := NewDenseVector(s.l)

	for i := 0; i < s.l; i++ {
		v1, _ := s.At(i)
		v2, _ := m.At(i)
		vector.Set(i, v1+v2)
	}

	return vector, nil
}

// Subtract subtracts one Vector from another.
func (s *DenseVector) Subtract(m *DenseVector) (*DenseVector, error) {
	if s.Length() != m.Length() {
		return nil, fmt.Errorf("Length miss match %+v %+v", s.Length(), m.Length())
	}

	vector := NewDenseVector(s.l)

	for i := 0; i < s.l; i++ {
		v1, _ := s.At(i)
		v2, _ := m.At(i)
		vector.Set(i, v1-v2)
	}

	return vector, nil
}

// Negative the negative of a Vector.
func (s *DenseVector) Negative() *DenseVector {
	vector := NewDenseVector(s.l)

	for i := 0; i < s.l; i++ {
		v1, _ := s.At(i)
		vector.Set(i, -v1)
	}

	return vector
}

func (s *DenseVector) Copy() *DenseVector {
	vector := NewDenseVector(s.l)

	for i, v := range s.data {
		vector.Set(i, v)
	}

	return vector
}
