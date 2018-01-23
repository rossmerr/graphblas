package GraphBLAS

// DenseVector a dense vector
type DenseVector struct {
	l    int
	data []float64
}

// NewDenseVector returns a GraphBLAS.DenseVector.
func NewDenseVector(l int) *DenseVector {
	return &DenseVector{l: l, data: make([]float64, l)}
}

// Length of the vector
func (s *DenseVector) Length() int {
	return s.l
}
