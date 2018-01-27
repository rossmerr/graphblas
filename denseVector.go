package GraphBLAS

import "fmt"

// DenseVector a vector
type DenseVector struct {
	l      int // length of the sparse vector
	values []float64
}

// NewDenseVector returns a GraphBLAS.DenseVector.
func NewDenseVector(l int) *DenseVector {
	return &DenseVector{l: l, values: make([]float64, 0)}
}

// Length of the vector
func (s *DenseVector) Length() int {
	return s.l
}

func (s *DenseVector) At(i int) (float64, error) {
	if i < 0 || i >= s.Length() {
		return 0, fmt.Errorf("Length '%+v' is invalid", i)
	}

	return s.values[i], nil
}

func (s *DenseVector) Set(i int, value float64) error {
	if i < 0 || i >= s.Length() {
		return fmt.Errorf("Length '%+v' is invalid", i)
	}

	s.values[i] = value

	return nil
}

// Scalar multiplication
func (s *DenseVector) Scalar(alpha float64) Vector {
	vector := NewDenseVector(s.l)

	for i, v := range s.values {
		vector.Set(i, alpha*v)
	}

	return vector
}

// Multiply multiplies a Vector structure by another Vector structure.
func (s *DenseVector) Multiply(m Vector) (Vector, error) {
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
func (s *DenseVector) Add(m Vector) (Vector, error) {
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
func (s *DenseVector) Subtract(m Vector) (Vector, error) {
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
func (s *DenseVector) Negative() Vector {
	vector := NewDenseVector(s.l)

	for i := 0; i < s.l; i++ {
		v1, _ := s.At(i)
		vector.Set(i, -v1)
	}

	return vector
}

func (s *DenseVector) Copy() Vector {
	vector := NewDenseVector(s.l)

	for i, v := range s.values {
		vector.Set(i, v)
	}

	return vector
}
