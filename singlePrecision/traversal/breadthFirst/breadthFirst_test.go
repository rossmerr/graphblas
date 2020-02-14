// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthFirst_test

import (
	"context"
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/singlePrecision"
	"github.com/RossMerr/Caudex.GraphBLAS/singlePrecision/traversal/breadthFirst"
)

func TestBreadthFirstSearch(t *testing.T) {
	array := [][]float32{
		[]float32{0, 0, 0, 1, 0, 0, 0},
		[]float32{1, 0, 0, 0, 0, 0, 0},
		[]float32{0, 0, 0, 1, 0, 1, 1},
		[]float32{1, 0, 0, 0, 0, 0, 1},
		[]float32{0, 1, 0, 0, 0, 0, 1},
		[]float32{0, 0, 1, 0, 1, 0, 0},
		[]float32{0, 1, 0, 0, 0, 0, 0},
	}
	g := singlePrecision.NewDenseMatrixFromArray(array)

	atx := breadthFirst.Search(context.Background(), g, 3, func(i singlePrecision.Vector) bool {
		return i.AtVec(5) == 1
	})

	if atx.AtVec(1) != 1 {
		t.Errorf("AtVec(%+v) wanted = %+v", 1, 1)
	}

	if atx.AtVec(5) != 1 {
		t.Errorf("AtVec(%+v) wanted = %+v", 5, 1)
	}
}
