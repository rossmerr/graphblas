// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthfirst

import (
	"github.com/rossmerr/graphblas/f64"
	"golang.org/x/net/context"
)

// Search a breadth-first search v is the source
func Search(ctx context.Context, a f64.Matrix, s int, c func(f64.Vector) bool) f64.Vector {
	n := a.Rows()
	// vertices visited in each level
	var frontier f64.Vector = f64.NewDenseVector(n)
	frontier.SetVec(s, 1)

	visited := frontier.Copy().(f64.Vector)

	// result
	result := f64.NewDenseVector(n)

	// level in BFS traversal
	d := 0

	// when some successor found
	for d < n {
		d++

		f64.MatrixVectorMultiply(ctx, a, frontier, visited, result)

		if c(result) {
			break
		}

		f64.ElementWiseVectorAdd(ctx, visited, result, nil, visited)
		frontier = result.Copy().(f64.Vector)
		result.Clear()
	}

	return result
}
