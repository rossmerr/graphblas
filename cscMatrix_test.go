package GraphBLAS_test

import (
	"fmt"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestCSCMatrix_Set(t *testing.T) {
	s := GraphBLAS.NewCSCMatrix(3, 3)

	s.Set(0, 0, 31)
	s.Set(1, 0, 0)
	s.Set(2, 0, 41)
	s.Set(0, 1, 0)
	s.Set(1, 1, 59)
	s.Set(2, 1, 26)
	s.Set(0, 2, 53)
	s.Set(1, 2, 0)
	s.Set(2, 2, 0)

	fmt.Printf("%+v", s)

	i, _ := s.At(0, 0)
	if i != 31 {
		t.Errorf("Expected 31 got %+v", i)
	}

	i, _ = s.At(1, 0)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	i, _ = s.At(2, 0)
	if i != 41 {
		t.Errorf("Expected 41 got %+v", i)
	}

	i, _ = s.At(0, 1)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	i, _ = s.At(1, 1)
	if i != 59 {
		t.Errorf("Expected 59 got %+v", i)
	}

	i, _ = s.At(2, 1)
	if i != 26 {
		t.Errorf("Expected 26 got %+v", i)
	}

	i, _ = s.At(0, 2)
	if i != 53 {
		t.Errorf("Expected 53 got %+v", i)
	}

	i, _ = s.At(1, 2)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
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

func TestCSCMatrix_Columns(t *testing.T) {
	s := GraphBLAS.NewCSCMatrix(3, 3)

	s.Set(0, 0, 31)
	s.Set(1, 0, 0)
	s.Set(2, 0, 41)
	s.Set(0, 1, 0)
	s.Set(1, 1, 59)
	s.Set(2, 1, 26)
	s.Set(0, 2, 53)
	s.Set(1, 2, 0)
	s.Set(2, 2, 0)

	col, _ := s.Columns(0)

	if col[0] != 31 {
		t.Errorf("Expected 31 got %+v", col[0])
	}

	if col[2] != 41 {
		t.Errorf("Expected 41 got %+v", col[2])
	}
}

func TestCSCMatrix_Row(t *testing.T) {
	s := GraphBLAS.NewCSCMatrix(3, 3)

	s.Set(0, 0, 31)
	s.Set(1, 0, 0)
	s.Set(2, 0, 41)
	s.Set(0, 1, 0)
	s.Set(1, 1, 59)
	s.Set(2, 1, 26)
	s.Set(0, 2, 53)
	s.Set(1, 2, 0)
	s.Set(2, 2, 0)

	row, _ := s.Rows(0)

	if row[0] != 31 {
		t.Errorf("Expected 31 got %+v", row[0])
	}

	if row[2] != 53 {
		t.Errorf("Expected 53 got %+v", row[1])
	}
}
