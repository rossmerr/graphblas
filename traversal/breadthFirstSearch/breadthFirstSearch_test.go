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
	array := [][]float64{
		[]float64{1, 0, 0, 1, 0, 0, 0},
		[]float64{1, 1, 0, 0, 0, 0, 0},
		[]float64{0, 0, 1, 1, 0, 1, 1},
		[]float64{1, 0, 0, 1, 0, 0, 1},
		[]float64{0, 1, 0, 0, 1, 0, 1},
		[]float64{0, 0, 1, 0, 1, 1, 0},
		[]float64{0, 1, 0, 0, 0, 0, 1},
	}
	g := GraphBLAS.NewDenseMatrixFromArray(array)

	v := GraphBLAS.NewDenseVector(7)
	v.SetVec(4, float64(1))

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
