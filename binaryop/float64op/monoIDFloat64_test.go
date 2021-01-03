// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64op_test

import (
	"testing"

	"github.com/rossmerr/graphblas/binaryop/float64op"
)

func Test_Reduce(t *testing.T) {
	done := make(chan struct{})
	slice := make(chan float64)
	defer close(slice)
	defer close(done)

	monoID := float64op.NewMonoIDFloat64(1, float64op.Addition)

	result := monoID.Reduce(done, slice)

	zero := monoID.Zero()

	if zero != 1 {
		t.Errorf("Zero = %+v want %+v", zero, 1)
	}

	slice <- 1
	done <- struct{}{}
	for out := range result {
		if 2 != out {
			t.Errorf("MonoIDBool = %+v, want %+v", out, 2)
		}
	}
}
