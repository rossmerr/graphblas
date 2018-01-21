package GraphBLAS

type Matrix interface {
	At(r, c int) (float64, error)
	Set(r, c int, value float64) error
	ColumnsAt(c int) ([]float64, error)
	RowsAt(r int) ([]float64, error)
	Columns() int
	Rows() int

	Copy() Matrix
	Scalar(alpha float64) Matrix
	Multiply(m Matrix) (Matrix, error)
	Add(m Matrix) (Matrix, error)
	Subtract(m Matrix) (Matrix, error)
	Negative() Matrix
}
