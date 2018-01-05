package GraphBLAS_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS"
)

func TestIncidenceMatrix(t *testing.T) {
	//s = [1 1 1 1 1];
	//t = [2 3 4 5 6];
	s := []string{}
	s = append(s, "1", "1", "1", "1", "1")
	tt := []string{}
	tt = append(tt, "2", "3", "4", "5", "6")
	g := GraphBLAS.NewGraphNamedNodes(s, tt)
	i := GraphBLAS.NewIncidenceMatrixFromGraph(g)

	if i.Element(2, 1) != -1 {
		t.Errorf("Expect -1 got %+v", i.Element(2, 1))
	}
	// 	I =
	//    (1,1)       -1
	//    (2,1)        1
	//    (1,2)       -1
	//    (3,2)        1
	//    (1,3)       -1
	//    (4,3)        1
	//    (1,4)       -1
	//    (5,4)        1
	//    (1,5)       -1
	//    (6,5)        1

}
