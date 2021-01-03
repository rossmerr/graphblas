// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthfirst

import (
	"github.com/rossmerr/graphblas/doubleprecision"
	"golang.org/x/net/context"
)

// Search a breadth-first search v is the source
func Search(ctx context.Context, a doubleprecision.Matrix, s int, c func(doubleprecision.Vector) bool) doubleprecision.Vector {
	n := a.Rows()
	// vertices visited in each level
	var frontier doubleprecision.Vector = doubleprecision.NewDenseVector(n)
	frontier.SetVec(s, 1)

	visited := frontier.Copy().(doubleprecision.Vector)

	// result
	result := doubleprecision.NewDenseVector(n)

	// level in BFS traversal
	d := 0

	// when some successor found
	for d < n {
		d++

		doubleprecision.MatrixVectorMultiply(ctx, a, frontier, visited, result)

		if c(result) {
			break
		}

		doubleprecision.ElementWiseVectorAdd(ctx, visited, result, nil, visited)
		frontier = result.Copy().(doubleprecision.Vector)
		result.Clear()
	}

	return result
}
