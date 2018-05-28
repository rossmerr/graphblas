// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64Op_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/binaryOp/float64Op"
)

func Test_Reduce(t *testing.T) {
	done := make(chan interface{})
	slice := make(chan float64)
	defer close(slice)
	defer close(done)

	monoID := float64Op.NewMonoIDFloat64(1, float64Op.Addition)

	result := monoID.Reduce(done, slice)

	slice <- 1
	done <- nil
	for out := range result {
		if 2 != out {
			t.Errorf("MonoIDBool = %+v, want %+v", out, 2)
		}
	}

}
