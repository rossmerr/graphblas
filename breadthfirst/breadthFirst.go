// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthfirst

import (
	"context"

	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/constraints"
)

// Search a breadth-first search s is the source
func Search[T constraints.Number](ctx context.Context, a graphblas.Matrix[T], s int, c func(graphblas.Vector[T]) bool) graphblas.Vector[T] {
	n := a.Rows()
	// vertices visited in each level
	var frontier graphblas.Vector[T] = graphblas.NewDenseVectorN[T](n)
	frontier.SetVec(s, 1)

	visited := frontier.Copy().(graphblas.Vector[T])

	// result
	result := graphblas.NewDenseVectorN[T](n)

	// level in BFS traversal
	d := 0

	// when some successor found
	for d < n {
		d++

		graphblas.MatrixVectorMultiply[T](ctx, a, frontier, visited, result)

		if c(result) {
			break
		}

		graphblas.ElementWiseVectorAdd[T](ctx, visited, result, nil, visited)
		frontier = result.Copy().(graphblas.Vector[T])
		result.Clear()
	}

	return result
}
