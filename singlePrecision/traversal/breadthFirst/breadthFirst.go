// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirst

import (
	"github.com/RossMerr/Caudex.GraphBLAS/singlePrecision"
	"golang.org/x/net/context"
)

// Search a breadth-first search v is the source
func Search(ctx context.Context, a singlePrecision.Matrix, s int, c func(singlePrecision.Vector) bool) singlePrecision.Vector {
	n := a.Rows()
	// vertices visited in each level
	var frontier singlePrecision.Vector = singlePrecision.NewDenseVector(n)
	frontier.SetVec(s, 1)

	visited := frontier.Copy().(singlePrecision.Vector)

	// result
	result := singlePrecision.NewDenseVector(n)

	// level in BFS traversal
	d := 0

	// when some successor found
	for d < n {
		d++

		singlePrecision.MatrixVectorMultiply(ctx, a, frontier, visited, result)

		if c(result) {
			break
		}

		singlePrecision.ElementWiseVectorAdd(ctx, visited, result, nil, visited)
		frontier = result.Copy().(singlePrecision.Vector)
		result.Clear()
	}

	return result
}
