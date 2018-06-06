// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolOp_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/binaryOp/boolOp"
)

func Test_Reduce(t *testing.T) {
	done := make(chan struct{})
	slice := make(chan bool)
	defer close(slice)
	defer close(done)

	monoID := boolOp.NewMonoIDBool(true, boolOp.LAND)

	result := monoID.Reduce(done, slice)

	zero := monoID.Zero()

	if zero != true {
		t.Errorf("Zero = %+v want %+v", zero, true)
	}

	slice <- true
	done <- struct{}{}
	for out := range result {
		if !out {
			t.Errorf("MonoIDBool = %+v, want %+v", false, true)
		}
	}
}
