// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// Vector interface
type Vector interface {
	AtVec(i int) (float64, error)
	SetVec(i int, value float64) error
	Length() int

	Copy() Vector
	Scalar(Vector float64) Vector
	Multiply(m Vector) (Vector, error)
	Add(m Vector) (Vector, error)
	Subtract(m Vector) (Vector, error)
	Negative() Vector
	Equal(m Vector) bool
	NotEqual(m Vector) bool
}
