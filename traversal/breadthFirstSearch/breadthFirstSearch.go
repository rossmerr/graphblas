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
func BreadthFirstSearch(ctx context.Context, a GraphBLAS.Matrix, s int) GraphBLAS.Matrix {
	n := a.Rows()
	// vertices visited in each level
	var q GraphBLAS.Vector = GraphBLAS.NewDenseVector(n)
	q.SetVec(s, 1)

	// result
	v := GraphBLAS.NewDenseVector(n)

	// // level in BFS traversal
	// d := float64(0)

	GraphBLAS.ElementWiseMatrixMultiply(ctx, a, q, v)

	// // when some successor found
	// for {
	// 	d++
	// 	GraphBLAS.ElementWiseMatrixMultiply(ctx, a, q, v)
	// 	GraphBLAS.AssignConstantVector(v, q, d, n)
	// 	GraphBLAS.VectorMatrixMultiply(ctx, q, a, q)
	// 	if GraphBLAS.ReduceVectorToScalar(ctx, q) != 1 {
	// 		break
	// 	}
	// }

	return v
}
