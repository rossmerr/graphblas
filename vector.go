package GraphBLAS

import "fmt"

// Vector a vector
type Vector struct {
	l    int
	data []float64
}

// NewVector returns a GraphBLAS.Vector.
func NewVector(l int) *Vector {
	return &Vector{l: l, data: make([]float64, l)}
}

// Length of the vector
func (s *Vector) Length() int {
	return s.l
}

func (s *Vector) At(i int) (float64, error) {
	if i < 0 || i >= s.Length() {
		return 0, fmt.Errorf("Length '%+v' is invalid", i)
	}

	return s.data[i], nil
}

func (s *Vector) Set(i int, value float64) error {
	if i < 0 || i >= s.Length() {
		return fmt.Errorf("Length '%+v' is invalid", i)
	}

	s.data[i] = value

	return nil
}

func (s *Vector) Copy() *Vector {
	vector := NewVector(s.l)

	for i, v := range s.data {
		vector.Set(i, v)
	}

	return vector
}
