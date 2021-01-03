// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolop_test

import (
	"testing"

	"github.com/rossmerr/graphblas/binaryop/boolop"
)

func Test_Reduce(t *testing.T) {
	done := make(chan struct{})
	slice := make(chan bool)
	defer close(slice)
	defer close(done)

	monoID := boolop.NewMonoIDBool(true, boolop.LAND)

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
