// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package breadthfirst_test

import (
	"context"
	"testing"

	"github.com/rossmerr/graphblas/f32"
	"github.com/rossmerr/graphblas/f32/traversal/breadthfirst"
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
	g := f32.NewDenseMatrixFromArray(array)

	atx := breadthfirst.Search(context.Background(), g, 3, func(i f32.Vector) bool {
		return i.AtVec(5) == 1
	})

	if atx.AtVec(1) != 1 {
		t.Errorf("AtVec(%+v) wanted = %+v", 1, 1)
	}

	if atx.AtVec(5) != 1 {
		t.Errorf("AtVec(%+v) wanted = %+v", 5, 1)
	}
}
