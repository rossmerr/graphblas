// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package bellmanFord

import GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"

// Bellman–Ford(A, s)
// d = ∞
// d(s) = 0
// for k = 1 to N − 1
//		do d = d min.+ A
// if d != d min.+ A
//		then return “A negative-weight cycle exists.”
// return d

// BellmanFord an algebraic implementation of the Bellman–Ford algorithm. s is the source
func BellmanFord(a GraphBLAS.Matrix, s GraphBLAS.Vector) []GraphBLAS.Matrix {
	d := []GraphBLAS.Matrix{}
	for i := 0; i < s.Length(); i++ {
		d = append(d, GraphBLAS.NewDenseVector(a.Rows()))
	}

	// need to build out
	n := a.Rows() - 1
	for k := 1; k < n; k++ {
		d[k] = d[k-1].Add(a)
	}

	return d
}
