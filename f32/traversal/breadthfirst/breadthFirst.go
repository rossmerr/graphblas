// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthfirst

import (
	"github.com/rossmerr/graphblas/f32"
	"context"
)

// Search a breadth-first search v is the source
func Search(ctx context.Context, a f32.Matrix, s int, c func(f32.Vector) bool) f32.Vector {
	n := a.Rows()
	// vertices visited in each level
	var frontier f32.Vector = f32.NewDenseVector(n)
	frontier.SetVec(s, 1)

	visited := frontier.Copy().(f32.Vector)

	// result
	result := f32.NewDenseVector(n)

	// level in BFS traversal
	d := 0

	// when some successor found
	for d < n {
		d++

		f32.MatrixVectorMultiply(ctx, a, frontier, visited, result)

		if c(result) {
			break
		}

		f32.ElementWiseVectorAdd(ctx, visited, result, nil, visited)
		frontier = result.Copy().(f32.Vector)
		result.Clear()
	}

	return result
}
