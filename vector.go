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

func (s *DenseVector) Copy() *DenseVector {
	vector := NewDenseVector(s.l)

	for i, v := range s.data {
		vector.Set(i, v)
	}

	return vector
}
