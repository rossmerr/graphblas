// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthfirst_test

import (
	"context"
	"testing"

	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/breadthfirst"
)

func TestBreadthFirstSearch(t *testing.T) {
	array := [][]float64{
		{0, 0, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 0},
	}
	g := graphblas.NewDenseMatrixFromArrayN(array)

	atx := breadthfirst.Search[float64](context.Background(), g, 3, func(i graphblas.Vector[float64]) bool {
		return i.AtVec(5) == 1
	})

	if atx.AtVec(1) != 1 {
		t.Errorf("AtVec(%+v) wanted = %+v got %v", 1, 1, atx.AtVec(5))
	}

	if atx.AtVec(5) != 1 {
		t.Errorf("AtVec(%+v) wanted = %+v got %v", 5, 1, atx.AtVec(5))
	}
}
