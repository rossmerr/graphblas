package GraphBLAS

type IncidenceMatrix struct {
	matrix [][]int
}

func NewIncidenceMatrix(vertices int, edges int) *IncidenceMatrix {
	m := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		m[i] = make([]int, edges)
	}
	return &IncidenceMatrix{matrix: m}

}

func NewIncidenceMatrixFromGraph(g Graph) *IncidenceMatrix {
	m := NewIncidenceMatrix(len(g.Vertices), len(g.Edges))
	mat := m.matrix
	i := 0
	for _, e := range g.Edges {
		from := e.Source()
		to := e.Sink()
		j := from.Index()
		k := to.Index()
		mat[j][i] = -1
		mat[k][i] = 1
	}

	return m
}
