package GraphBLAS

// SparseMatrix is a matrix in which most of the elements are zero
type SparseMatrix interface {
	Get(r, c int) (int, error)
	Set(r, c, value int) error
	sparse()
}
