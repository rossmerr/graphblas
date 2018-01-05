package GraphBLAS

type Graph struct {
	Vertices map[int]*Vertex
	Edges    map[int]*Edge
}

func NewGraph() *Graph {
	g := &Graph{Vertices: make(map[int]*Vertex), Edges: make(map[int]*Edge)}
	return g
}

func NewGraphNamedNodes(s []string, t []string) *Graph {
	g := NewGraph()

	for i := 0; i < len(s); i++ {
		key := s[i]
		g.Vertices[i] = NewVertex(key)
	}

	for i := 0; i < len(t); i++ {
		key := t[i]
		g.Edges[i] = NewEdge(key)
	}

	return g
}

type Vertex struct {
	id       string
	position int // needs to be count up as vertices are added to the graph
	Edges    map[string]*Edge
}

func NewVertex(ID string) *Vertex {
	v := &Vertex{id: ID}
	return v
}

func (v *Vertex) Index() int {
	return v.position
}

type Edge struct {
	id       string
	position int
}

func NewEdge(ID string) *Edge {
	e := &Edge{id: ID}
	return e
}

func (e *Edge) Source() *Vertex {
	return nil
}

func (e *Edge) Sink() *Vertex {
	return nil
}
