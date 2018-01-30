package GraphBLAS

// Iterator over the matrix returning all non-zero elements and false when complete
type Iterator func() (int, int, float64, bool)
