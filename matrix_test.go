package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestScalar(t *testing.T) {
	m := GraphBLAS.NewMatrix(2, 2)
	m[0][0] = 4
	m[0][1] = 0
	m[1][0] = 1
	m[1][1] = -9
	scale := m.Scalar(2)
	if scale[0][0] != 8 {
		t.Errorf("Expected 4 got %+v", scale[0][0])
	}
}

func TestMultiple(t *testing.T) {
	m := GraphBLAS.NewMatrix(2, 3)
	m[0][0] = 1
	m[0][1] = 2
	m[0][2] = 3
	m[1][0] = 4
	m[1][1] = 5
	m[1][2] = 6

	m2 := GraphBLAS.NewMatrix(3, 2)
	m2[0][0] = 7
	m2[0][1] = 8
	m2[1][0] = 9
	m2[1][1] = 10
	m2[2][0] = 11
	m2[2][1] = 12

	m3, ok := m.Multiply(m2)

	if ok == false {
		t.Error("Multiply failed")
	}

	if m3[0][0] != 58 {
		t.Errorf("Expected 58 got %+v", m3[0][0])
	}

	if m3[0][1] != 64 {
		t.Errorf("Expected 64 got %+v", m3[0][1])
	}

	if m3[1][0] != 139 {
		t.Errorf("Expected 139 got %+v", m3[1][0])
	}

	if m3[1][1] != 154 {
		t.Errorf("Expected 154 got %+v", m3[1][1])
	}
}

func TestAdd(t *testing.T) {
	m := GraphBLAS.NewMatrix(2, 2)
	m[0][0] = 3
	m[0][1] = 8
	m[1][0] = 4
	m[1][1] = 6

	m2 := GraphBLAS.NewMatrix(2, 2)
	m2[0][0] = 4
	m2[0][1] = 0
	m2[1][0] = 1
	m2[1][1] = -9

	m3, ok := m.Add(m2)

	if ok == false {
		t.Error("Add failed")
	}

	if m3[0][0] != 7 {
		t.Errorf("Expected 7 got %+v", m3[0][0])
	}

	if m3[0][1] != 8 {
		t.Errorf("Expected 8 got %+v", m3[0][1])
	}

	if m3[1][0] != 5 {
		t.Errorf("Expected 5 got %+v", m3[1][0])
	}

	if m3[1][1] != -3 {
		t.Errorf("Expected -3 got %+v", m3[1][1])
	}
}

func TestSubtract(t *testing.T) {
	m := GraphBLAS.NewMatrix(2, 2)
	m[0][0] = 3
	m[0][1] = 8
	m[1][0] = 4
	m[1][1] = 6

	m2 := GraphBLAS.NewMatrix(2, 2)
	m2[0][0] = 4
	m2[0][1] = 0
	m2[1][0] = 1
	m2[1][1] = -9

	m3, ok := m.Subtract(m2)

	if ok == false {
		t.Error("Add failed")
	}

	if m3[0][0] != -1 {
		t.Errorf("Expected -1 got %+v", m3[0][0])
	}

	if m3[0][1] != 8 {
		t.Errorf("Expected 8 got %+v", m3[0][1])
	}

	if m3[1][0] != 3 {
		t.Errorf("Expected 3 got %+v", m3[1][0])
	}

	if m3[1][1] != 15 {
		t.Errorf("Expected 15 got %+v", m3[1][1])
	}
}

func TestNegative(t *testing.T) {
	m := GraphBLAS.NewMatrix(2, 2)
	m[0][0] = 2
	m[0][1] = -4
	m[1][0] = 7
	m[1][1] = 10

	m3 := m.Negative()

	if m3[0][0] != -2 {
		t.Errorf("Expected -2 got %+v", m3[0][0])
	}

	if m3[0][1] != 4 {
		t.Errorf("Expected 4 got %+v", m3[0][1])
	}

	if m3[1][0] != -7 {
		t.Errorf("Expected -7 got %+v", m3[1][0])
	}

	if m3[1][1] != -10 {
		t.Errorf("Expected -10 got %+v", m3[1][1])
	}
}
