// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirstSearch

import (
	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

// BreadthFirstSearch a breadth-first search v is the source
func BreadthFirstSearch(a GraphBLAS.Matrix, s int) map[int]GraphBLAS.Matrix {
	v := make(map[int]GraphBLAS.Matrix)
	// vertices visited in each level
	visited := GraphBLAS.NewDenseVector(a.Rows())
	visited.SetVec(s, float64(1))
	var q GraphBLAS.Matrix = visited

	// level in BFS traversal
	d := 0
	// true when some successor found
	succ := false

	for succ {
		d++
		v[d] = q.Copy()
		GraphBLAS.ElementWiseMatrixMultiply(q, a, q)
		if GraphBLAS.ReduceMatrixToScalar(q) == 1 {
			succ = true
		}
	}

	return v
}
