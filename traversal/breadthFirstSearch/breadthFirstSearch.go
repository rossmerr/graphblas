// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirstSearch

import GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"

// BreadthFirstSearch a breadth-first search v is the source
func BreadthFirstSearch(a GraphBLAS.Matrix, s int) map[int]GraphBLAS.Matrix {
	v := make(map[int]GraphBLAS.Matrix)
	// vertices visited in each level
	q := GraphBLAS.NewDenseVector(a.Rows())
	q.SetVec(s, float64(1))

	qm := q.Copy()

	// level in BFS traversal
	d := 0
	// true when some successor found
	succ := false

	for succ {
		d++
		v[d] = qm.Copy()
		qm = qm.Multiply(a)
		if qm.ReduceScalar() == 1 {
			succ = true
		}
	}

	return v
}
