package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestCSRMatrix_Set(t *testing.T) {
	s := GraphBLAS.NewCSRMatrix(3, 3)

	s.Set(0, 0, 31)
	s.Set(0, 1, 0)
	s.Set(0, 2, 53)
	s.Set(1, 0, 0)
	s.Set(1, 1, 59)
	s.Set(1, 2, 0)
	s.Set(2, 0, 41)
	s.Set(2, 1, 26)
	s.Set(2, 2, 0)

	i, _ := s.At(0, 0)
	if i != 31 {
		t.Errorf("Expected 31 got %+v", i)
	}

	i, _ = s.At(0, 1)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	i, _ = s.At(0, 2)
	if i != 53 {
		t.Errorf("Expected 53 got %+v", i)
	}

	i, _ = s.At(1, 0)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	i, _ = s.At(1, 1)
	if i != 59 {
		t.Errorf("Expected 59 got %+v", i)
	}

	i, _ = s.At(1, 2)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	i, _ = s.At(2, 0)
	if i != 41 {
		t.Errorf("Expected 41 got %+v", i)
	}

	i, _ = s.At(2, 1)
	if i != 26 {
		t.Errorf("Expected 26 got %+v", i)
	}

	i, _ = s.At(2, 2)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	s.Set(2, 1, 0)
	i, _ = s.At(2, 1)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	s.Set(2, 1, 62)
	i, _ = s.At(2, 1)
	if i != 62 {
		t.Errorf("Expected 62 got %+v", i)
	}
}

func TestCSRMatrix_Columns(t *testing.T) {
	s := GraphBLAS.NewCSRMatrix(3, 3)

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

func TestCSRMatrix_Row(t *testing.T) {
	s := GraphBLAS.NewCSRMatrix(3, 3)

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

func TestCSRMultiple(t *testing.T) {
	m := GraphBLAS.NewCSRMatrix(2, 3)
	m.Set(0, 0, 1)
	m.Set(0, 1, 2)
	m.Set(0, 2, 3)
	m.Set(1, 0, 4)
	m.Set(1, 1, 5)
	m.Set(1, 2, 6)

	m2 := GraphBLAS.NewCSRMatrix(3, 2)
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

func TestCSRAdd(t *testing.T) {
	m := GraphBLAS.NewCSRMatrix(2, 2)
	m.Set(0, 0, 3)
	m.Set(0, 1, 8)
	m.Set(1, 0, 4)
	m.Set(1, 1, 6)

	m2 := GraphBLAS.NewCSRMatrix(2, 2)
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

func TestCSRSubtract(t *testing.T) {
	m := GraphBLAS.NewCSRMatrix(2, 2)
	m.Set(0, 0, 3)
	m.Set(0, 1, 8)
	m.Set(1, 0, 4)
	m.Set(1, 1, 6)

	m2 := GraphBLAS.NewCSRMatrix(2, 2)
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

func TestCSRNegative(t *testing.T) {
	m := GraphBLAS.NewCSRMatrix(2, 2)
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
