// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirst

import (
	"github.com/RossMerr/Caudex.GraphBLAS/doublePrecision"
	"golang.org/x/net/context"
)

// Search a breadth-first search v is the source
func Search(ctx context.Context, a doublePrecision.Matrix, s int, c func(doublePrecision.Vector) bool) doublePrecision.Vector {
	n := a.Rows()
	// vertices visited in each level
	var frontier doublePrecision.Vector = doublePrecision.NewDenseVector(n)
	frontier.SetVec(s, 1)

	visited := frontier.Copy().(doublePrecision.Vector)

	// result
	result := doublePrecision.NewDenseVector(n)

	// level in BFS traversal
	d := 0

	// when some successor found
	for d < n {
		d++

		doublePrecision.MatrixVectorMultiply(ctx, a, frontier, visited, result)

		if c(result) {
			break
		}

		doublePrecision.ElementWiseVectorAdd(ctx, visited, result, nil, visited)
		frontier = result.Copy().(doublePrecision.Vector)
		result.Clear()
	}

	return result
}
