// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

type VectorSet struct {
	Value  float64
	Vector Vector
}

// Span of a set of Vectors
func Span(set ...VectorSet) Matrix {
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
