package GraphBLAS

// AdjacencyMatrix rows and columns represent vertices
type AdjacencyMatrix struct {
	matrix Matrix
}

func NewAdjacencyMatrix(vertices int) *AdjacencyMatrix {
	m := make(Matrix, vertices)
	for i := 0; i < vertices; i++ {
		m[i] = make([]int, vertices)
	}
	return &AdjacencyMatrix{matrix: m}

}
