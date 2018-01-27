package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestMatrix_Columns(t *testing.T) {
	s := GraphBLAS.NewDenseMatrix(3, 3)

	s.Set(0, 0, 31)
	s.Set(1, 0, 0)
	s.Set(2, 0, 41)
	s.Set(0, 1, 0)
	s.Set(1, 1, 59)
	s.Set(2, 1, 26)
	s.Set(0, 2, 53)
	s.Set(1, 2, 0)
	s.Set(2, 2, 0)

	col, _ := s.ColumnsAt(0)

	if col[0] != 31 {
		t.Errorf("Expected 31 got %+v", col[0])
	}

	if col[2] != 41 {
		t.Errorf("Expected 41 got %+v", col[2])
	}
}

func TestMatrix_Row(t *testing.T) {
	s := GraphBLAS.NewDenseMatrix(3, 3)

	s.Set(0, 0, 31)
	s.Set(1, 0, 0)
	s.Set(2, 0, 41)
	s.Set(0, 1, 0)
	s.Set(1, 1, 59)
	s.Set(2, 1, 26)
	s.Set(0, 2, 53)
	s.Set(1, 2, 0)
	s.Set(2, 2, 0)

	row, _ := s.RowsAt(0)

	if row[0] != 31 {
		t.Errorf("Expected 31 got %+v", row[0])
	}

	if row[2] != 53 {
		t.Errorf("Expected 53 got %+v", row[1])
	}
}

func TestScalar(t *testing.T) {
	m := GraphBLAS.NewDenseMatrix(2, 2)
	m.Set(0, 0, 4)
	m.Set(0, 1, 0)
	m.Set(1, 0, 1)
	m.Set(1, 1, -9)
	scale := m.Scalar(2)
	if v, _ := scale.At(0, 0); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}
}

func TestMultiple(t *testing.T) {
	m := GraphBLAS.NewDenseMatrix(2, 3)
	m.Set(0, 0, 1)
	m.Set(0, 1, 2)
	m.Set(0, 2, 3)
	m.Set(1, 0, 4)
	m.Set(1, 1, 5)
	m.Set(1, 2, 6)

	m2 := GraphBLAS.NewDenseMatrix(3, 2)
	m2.Set(0, 0, 7)
	m2.Set(0, 1, 8)
	m2.Set(1, 0, 9)
	m2.Set(1, 1, 10)
	m2.Set(2, 0, 11)
	m2.Set(2, 1, 12)

	m3, err := m.Multiply(m2)

	if err != nil {
		t.Error("Multiply failed")
	}

	if v, _ := m3.At(0, 0); v != 58 {
		t.Errorf("Expected 58 got %+v", v)
	}

	if v, _ := m3.At(0, 1); v != 64 {
		t.Errorf("Expected 64 got %+v", v)
	}

	if v, _ := m3.At(1, 0); v != 139 {
		t.Errorf("Expected 139 got %+v", v)
	}

	if v, _ := m3.At(1, 1); v != 154 {
		t.Errorf("Expected 154 got %+v", v)
	}
}

func TestAdd(t *testing.T) {
	m := GraphBLAS.NewDenseMatrix(2, 2)
	m.Set(0, 0, 3)
	m.Set(0, 1, 8)
	m.Set(1, 0, 4)
	m.Set(1, 1, 6)

	m2 := GraphBLAS.NewDenseMatrix(2, 2)
	m2.Set(0, 0, 4)
	m2.Set(0, 1, 0)
	m2.Set(1, 0, 1)
	m2.Set(1, 1, -9)

	m3, err := m.Add(m2)

	if err != nil {
		t.Error("Add failed")
	}

	if v, _ := m3.At(0, 0); v != 7 {
		t.Errorf("Expected 7 got %+v", v)
	}

	if v, _ := m3.At(0, 1); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}

	if v, _ := m3.At(1, 0); v != 5 {
		t.Errorf("Expected 5 got %+v", v)
	}

	if v, _ := m3.At(1, 1); v != -3 {
		t.Errorf("Expected -3 got %+v", v)
	}
}

func TestSubtract(t *testing.T) {
	m := GraphBLAS.NewDenseMatrix(2, 2)
	m.Set(0, 0, 3)
	m.Set(0, 1, 8)
	m.Set(1, 0, 4)
	m.Set(1, 1, 6)

	m2 := GraphBLAS.NewDenseMatrix(2, 2)
	m2.Set(0, 0, 4)
	m2.Set(0, 1, 0)
	m2.Set(1, 0, 1)
	m2.Set(1, 1, -9)

	m3, err := m.Subtract(m2)

	if err != nil {
		t.Error("Add failed")
	}

	if v, _ := m3.At(0, 0); v != -1 {
		t.Errorf("Expected -1 got %+v", v)
	}

	if v, _ := m3.At(0, 1); v != 8 {
		t.Errorf("Expected 8 got %+v", v)
	}

	if v, _ := m3.At(1, 0); v != 3 {
		t.Errorf("Expected 3 got %+v", v)
	}

	if v, _ := m3.At(1, 1); v != 15 {
		t.Errorf("Expected 15 got %+v", v)
	}
}

func TestNegative(t *testing.T) {
	m := GraphBLAS.NewDenseMatrix(2, 2)
	m.Set(0, 0, 2)
	m.Set(0, 1, -4)
	m.Set(1, 0, 7)
	m.Set(1, 1, 10)

	m3 := m.Negative()

	if v, _ := m3.At(0, 0); v != -2 {
		t.Errorf("Expected -2 got %+v", v)
	}

	if v, _ := m3.At(0, 1); v != 4 {
		t.Errorf("Expected 4 got %+v", v)
	}

	if v, _ := m3.At(1, 0); v != -7 {
		t.Errorf("Expected -7 got %+v", v)
	}

	if v, _ := m3.At(1, 1); v != -10 {
		t.Errorf("Expected -10 got %+v", v)
	}
}

func TestCopy(t *testing.T) {
	m := GraphBLAS.NewDenseMatrix(1, 1)
	m.Set(0, 0, 4)
	copy := m.Copy()
	if v, _ := copy.At(0, 0); v != 4 {
		t.Errorf("Expected 4 got %+v", v)
	}
}
