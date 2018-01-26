package GraphBLAS

type Vector interface {
	At(i int) (float64, error)
	Set(i int, value float64) error
	Length() int

	Copy() Vector
	Scalar(Vector float64) Vector
	Multiply(m Vector) (Vector, error)
	Add(m Vector) (Vector, error)
	Subtract(m Vector) (Vector, error)
	Negative() Vector
}
