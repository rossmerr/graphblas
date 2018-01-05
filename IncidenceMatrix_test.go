package GraphBLAS_test

import (
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS"
)

func TestIncidenceMatrix(t *testing.T) {
	//target = [1 1 1 1 1];
	//target = [2 3 4 5 6];
	source := []string{}
	source = append(source, "1", "1", "1", "1", "1")
	target := []string{}
	target = append(target, "2", "3", "4", "5", "6")
	g := GraphBLAS.NewGraphNamedNodes(source, target)
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
