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
	m.SetVec(0, 4)
	m.SetVec(1, 0)
	scale := m.Scalar(2)

	if v, _ := scale.At(0, 0); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}
}

func TestVectorMultiple(t *testing.T) {
	m := GraphBLAS.NewDenseVector(3)
	m.SetVec(0, 3)
	m.SetVec(1, 4)
	m.SetVec(2, 3)

	m2 := GraphBLAS.NewDenseMatrix(2, 3)
	m2.Set(0, 0, 0)
	m2.Set(0, 1, 3)
	m2.Set(0, 2, 5)
	m2.Set(1, 0, 5)
	m2.Set(1, 1, 5)
	m2.Set(1, 2, 2)

	m3, err := m.Multiply(m2)

	if err != nil {
		t.Error(err)
	}

	if v, _ := m3.At(0, 0); v != 27 {
		t.Errorf("Expected 27 got %+v", v)
	}

	if v, _ := m3.At(1, 0); v != 41 {
		t.Errorf("Expected 41 got %+v", v)
	}

}

func TestVectorAdd(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.SetVec(0, 3)
	m.SetVec(1, 8)

	m2 := GraphBLAS.NewDenseVector(2)
	m2.SetVec(0, 4)
	m2.SetVec(1, 0)

	m3, err := m.Add(m2)

	if err != nil {
		t.Error("Add failed")
	}

	if v, _ := m3.At(0, 0); v != 7 {
		t.Errorf("Expected 7 got %+v", v)
	}

	if v, _ := m3.At(1, 0); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}
}

func TestVectorSubtract(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.SetVec(0, 3)
	m.SetVec(1, 8)

	m2 := GraphBLAS.NewDenseVector(2)
	m2.SetVec(0, 4)
	m2.SetVec(1, 0)

	m3, err := m.Subtract(m2)

	if err != nil {
		t.Error("Add failed")
	}

	if v, _ := m3.At(0, 0); v != -1 {
		t.Errorf("Expected -1 got %+v", v)
	}

	if v, _ := m3.At(1, 0); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}
}

func TestVectorNegative(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.SetVec(0, 2)
	m.SetVec(1, -4)

	m3 := m.Negative()

	if v, _ := m3.At(0, 0); v != -2 {
		t.Errorf("Expected -2 got %+v", v)
	}

	if v, _ := m3.At(1, 0); v != 4 {
		t.Errorf("Expected 4 got %+v", v)
	}
}

func TestVectorCopy(t *testing.T) {
	m := GraphBLAS.NewDenseVector(1)
	m.SetVec(0, 4)
	copy := m.Copy()
	if v, _ := copy.At(0, 0); v != 4 {
		t.Errorf("Expected 4 got %+v", v)
	}
}

func TestVectorEqual(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.SetVec(0, 3)
	m.SetVec(1, 8)

	m2 := GraphBLAS.NewDenseVector(2)
	m2.SetVec(0, 3)
	m2.SetVec(1, 8)

	if m.Equal(m2) == false {
		t.Error("Equal failed")
	}
}

func TestVectorNotEqual(t *testing.T) {
	m := GraphBLAS.NewDenseVector(2)
	m.SetVec(0, 3)
	m.SetVec(1, 8)

	m2 := GraphBLAS.NewDenseVector(2)
	m2.SetVec(0, 8)
	m2.SetVec(1, 3)

	if m.NotEqual(m2) == false {
		t.Error("NotEqual failed")
	}
}
