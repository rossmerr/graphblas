// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthfirst_test

import (
	"context"
	"testing"

	"github.com/rossmerr/graphblas/doubleprecision"
	"github.com/rossmerr/graphblas/doubleprecision/traversal/breadthfirst"
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
	g := doubleprecision.NewDenseMatrixFromArray(array)

	atx := breadthfirst.Search(context.Background(), g, 3, func(i doubleprecision.Vector) bool {
		return i.AtVec(5) == 1
	})

	if atx.AtVec(1) != 1 {
		t.Errorf("AtVec(%+v) wanted = %+v", 1, 1)
	}

	if atx.AtVec(5) != 1 {
		t.Errorf("AtVec(%+v) wanted = %+v", 5, 1)
	}
}
