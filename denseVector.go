package GraphBLAS

type DenseVector struct {
	l    int
	data []float64
}

// NewDenseVector returns an GraphBLAS.DenseVector.
func NewDenseVector(l int) *DenseVector {
	return &DenseVector{l: l, data: make([]float64, l)}
}
