// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirst

import (
	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
	"golang.org/x/net/context"
)

// Search a breadth-first search v is the source
func Search(ctx context.Context, a GraphBLAS.Matrix, s int, c func(GraphBLAS.Vector) bool) GraphBLAS.Vector {
	n := a.Rows()
	// vertices visited in each level
	var frontier GraphBLAS.Vector = GraphBLAS.NewDenseVector(n)
	frontier.SetVec(s, 1)

	visited := frontier.Copy().(GraphBLAS.Vector)

	// result
	result := GraphBLAS.NewDenseVector(n)

	// level in BFS traversal
	d := 0

	// when some successor found
	for d < n {
		d++

		GraphBLAS.MatrixVectorMultiply(ctx, a, frontier, visited, result)

		if c(result) {
			break
		}

		GraphBLAS.ElementWiseVectorAdd(ctx, visited, result, nil, visited)
		frontier = result.Copy().(GraphBLAS.Vector)
		result.Clear()
	}

	return result
}
