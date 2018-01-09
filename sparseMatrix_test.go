package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestSparseMatrix_Set(t *testing.T) {
	s := GraphBLAS.NewSparseMatrix(3, 3)
	s.Set(0, 0, 31)
	s.Set(0, 1, 0)
	s.Set(0, 2, 53)
	s.Set(1, 0, 0)
	s.Set(1, 1, 59)
	s.Set(1, 2, 0)
	s.Set(2, 0, 41)
}
