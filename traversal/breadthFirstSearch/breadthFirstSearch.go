// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirstSearch

import (
	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
	"golang.org/x/net/context"
)

// BreadthFirstSearch a breadth-first search v is the source
func BreadthFirstSearch(ctx context.Context, a GraphBLAS.Matrix, s int) map[int]GraphBLAS.Matrix {
	v := make(map[int]GraphBLAS.Matrix)
	// vertices visited in each level
	visited := GraphBLAS.NewDenseVector(a.Rows())
	visited.SetVec(s, 1)
	var q GraphBLAS.Matrix = visited

	// level in BFS traversal
	d := 0
	// true when some successor found
	succ := false

	for succ {
		d++
		v[d] = q.Copy()
		GraphBLAS.ElementWiseMatrixMultiply(ctx, q, a, q)
		if GraphBLAS.ReduceMatrixToScalar(ctx, q) == 1 {
			succ = true
		} else {
			succ = false
		}
	}

	return v
}
