// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirstSearch_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
	"github.com/RossMerr/Caudex.GraphBLAS/traversal/breadthFirstSearch"
)

func TestBreadthFirstSearch(t *testing.T) {
	set := float64(1)
	g := GraphBLAS.NewDenseMatrix(7, 7)
	g.Set(0, 0, set)
	g.Set(0, 3, set)
	g.Set(1, 0, set)
	g.Set(1, 1, set)
	g.Set(2, 2, set)
	g.Set(2, 3, set)
	g.Set(2, 5, set)
	g.Set(2, 6, set)
	g.Set(3, 0, set)
	g.Set(3, 3, set)
	g.Set(3, 6, set)
	g.Set(4, 1, set)
	g.Set(4, 4, set)
	g.Set(4, 6, set)
	g.Set(5, 2, set)
	g.Set(5, 4, set)
	g.Set(5, 5, set)
	g.Set(6, 1, set)
	g.Set(6, 6, set)

	v := GraphBLAS.NewDenseVector(7)
	v.SetVec(4, set)

	breadthFirstSearch.BreadthFirstSearch(g, 4)

	//atx := breadthFirstSearch.BreadthFirstSearch(g, v)

	// if atx.At(0, 0) != set {
	// 	t.Errorf("At(%+v, %+v) wanted = %+v", 0, 0, set)
	// }

	// if atx.At(0, 2) != set {
	// 	t.Errorf("At(%+v, %+v) wanted = %+v", 0, 2, set)
	// }

	// if atx.At(0, 3) != set {
	// 	t.Errorf("At(%+v, %+v) wanted = %+v", 0, 3, set)
	// }
}
