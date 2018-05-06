// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// VectorSet holds a array of VectorConstant's
type VectorSet []*VectorConstant

// VectorConstant holds a vector and a constant
type VectorConstant struct {
	Value  float64
	Vector Vector
}

// NewVectorSet returns a VectorSet
func NewVectorSet(value float64, v Vector) *VectorConstant {
	return &VectorConstant{
		Value:  value,
		Vector: v,
	}
}

// Span of a set of Vectors
func Span(set VectorSet) Matrix {
	results := make([]Matrix, len(set))

	for i, s := range set {
		results[i] = s.Vector.Scalar(s.Value)
	}

	result := results[0]

	for i := 1; i < len(results); i++ {
		result = result.Add(results[i])
	}

	return result
}
