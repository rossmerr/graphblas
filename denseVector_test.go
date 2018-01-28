// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestVectorScalar(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.Set(0, 4)
	m.Set(1, 0)
	scale := m.Scalar(2)
	if v, _ := scale.At(0); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}
}

func TestVectorMultiple(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.Set(0, 2)
	m.Set(1, 2)

	m2 := GraphBLAS.NewDenseVector(2)
	m2.Set(0, 7)
	m2.Set(1, 8)

	m3, err := m.Multiply(m2)

	if err != nil {
		t.Error("Multiply failed")
	}

	if v, _ := m3.At(0); v != 14 {
		t.Errorf("Expected 14 got %+v", v)
	}

	if v, _ := m3.At(1); v != 16 {
		t.Errorf("Expected 16 got %+v", v)
	}

}

func TestVectorAdd(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.Set(0, 3)
	m.Set(1, 8)

	m2 := GraphBLAS.NewDenseVector(2)
	m2.Set(0, 4)
	m2.Set(1, 0)

	m3, err := m.Add(m2)

	if err != nil {
		t.Error("Add failed")
	}

	if v, _ := m3.At(0); v != 7 {
		t.Errorf("Expected 7 got %+v", v)
	}

	if v, _ := m3.At(1); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}
}

func TestVectorSubtract(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.Set(0, 3)
	m.Set(1, 8)

	m2 := GraphBLAS.NewDenseVector(2)
	m2.Set(0, 4)
	m2.Set(1, 0)

	m3, err := m.Subtract(m2)

	if err != nil {
		t.Error("Add failed")
	}

	if v, _ := m3.At(0); v != -1 {
		t.Errorf("Expected -1 got %+v", v)
	}

	if v, _ := m3.At(1); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}
}

func TestVectorNegative(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.Set(0, 2)
	m.Set(1, -4)

	m3 := m.Negative()

	if v, _ := m3.At(0); v != -2 {
		t.Errorf("Expected -2 got %+v", v)
	}

	if v, _ := m3.At(1); v != 4 {
		t.Errorf("Expected 4 got %+v", v)
	}
}

func TestVectorCopy(t *testing.T) {
	m := GraphBLAS.NewDenseVector(1)
	m.Set(0, 4)
	copy := m.Copy()
	if v, _ := copy.At(0); v != 4 {
		t.Errorf("Expected 4 got %+v", v)
	}
}
