package GraphBLAS_test

import (
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestSparseMatrix_Set(t *testing.T) {
	s := GraphBLAS.NewSparseMatrix(3, 3)
	s.Set(0, 0, 31)

}
