package GraphBLAS

// SparseMatrix is a matrix in which most of the elements are zero
type SparseMatrix interface {
	At(r, c int) (float64, error)
	Set(r, c int, value float64) error
	ColumnsAt(c int) ([]float64, error)
	RowsAt(r int) ([]float64, error)
	Columns() int
	Rows() int

	Copy() SparseMatrix
	Scalar(alpha float64) SparseMatrix
	Multiply(m SparseMatrix) (SparseMatrix, error)
	Add(m SparseMatrix) (SparseMatrix, error)
	Subtract(m SparseMatrix) (SparseMatrix, error)
	Negative() SparseMatrix
}
