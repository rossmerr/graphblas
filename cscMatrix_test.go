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

	i, _ := s.Get(0, 0)
	if i != 31 {
		t.Errorf("Expected 31 got %+v", i)
	}

	i, _ = s.Get(1, 0)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	i, _ = s.Get(2, 0)
	if i != 41 {
		t.Errorf("Expected 41 got %+v", i)
	}

	i, _ = s.Get(0, 1)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	i, _ = s.Get(1, 1)
	if i != 59 {
		t.Errorf("Expected 59 got %+v", i)
	}

	i, _ = s.Get(2, 1)
	if i != 26 {
		t.Errorf("Expected 26 got %+v", i)
	}

	i, _ = s.Get(0, 2)
	if i != 53 {
		t.Errorf("Expected 53 got %+v", i)
	}

	i, _ = s.Get(1, 2)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	i, _ = s.Get(2, 2)
	if i != 0 {
		t.Errorf("Expected 0 got %+v", i)
	}

	fmt.Printf("%+v", s)

	// s.Set(2, 0, 0) //was 41

	//fmt.Printf("%+v", s)

	// s.Set(2, 0, 41) //was 41

	//fmt.Printf("%+v", s)
}
