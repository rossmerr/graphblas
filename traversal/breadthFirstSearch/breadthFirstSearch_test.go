// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirstSearch_test

import (
	"context"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
	"github.com/RossMerr/Caudex.GraphBLAS/traversal/breadthFirstSearch"
)

func TestBreadthFirstSearch(t *testing.T) {
	array := [][]float64{
		[]float64{0, 0, 0, 1, 0, 0, 0},
		[]float64{1, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 1, 0, 1, 1},
		[]float64{1, 0, 0, 0, 0, 0, 1},
		[]float64{0, 1, 0, 0, 0, 0, 1},
		[]float64{0, 0, 1, 0, 1, 0, 0},
		[]float64{0, 1, 0, 0, 0, 0, 0},
	}
	g := GraphBLAS.NewDenseMatrixFromArray(array)

	atx := breadthFirstSearch.BreadthFirstSearch(context.Background(), g, 3, func(i GraphBLAS.Vector) bool {
		return i.AtVec(5) == 1
	})

	if atx.At(0, 0) != 1 {
		t.Errorf("At(%+v, %+v) wanted = %+v", 0, 0, 1)
	}

	if atx.At(0, 2) != 1 {
		t.Errorf("At(%+v, %+v) wanted = %+v", 0, 2, 1)
	}
}
