// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthfirst

import (
	"github.com/rossmerr/graphblas/singleprecision"
	"context"
)

// Search a breadth-first search v is the source
func Search(ctx context.Context, a singleprecision.Matrix, s int, c func(singleprecision.Vector) bool) singleprecision.Vector {
	n := a.Rows()
	// vertices visited in each level
	var frontier singleprecision.Vector = singleprecision.NewDenseVector(n)
	frontier.SetVec(s, 1)

	visited := frontier.Copy().(singleprecision.Vector)

	// result
	result := singleprecision.NewDenseVector(n)

	// level in BFS traversal
	d := 0

	// when some successor found
	for d < n {
		d++

		singleprecision.MatrixVectorMultiply(ctx, a, frontier, visited, result)

		if c(result) {
			break
		}

		singleprecision.ElementWiseVectorAdd(ctx, visited, result, nil, visited)
		frontier = result.Copy().(singleprecision.Vector)
		result.Clear()
	}

	return result
}
