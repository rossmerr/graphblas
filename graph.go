package GraphBLAS

type Graph struct {
	Vertices map[string]*Vertex
	Edges    map[string]*Edge
}

func NewGraph() *Graph {
	g := &Graph{Vertices: make(map[string]*Vertex), Edges: make(map[string]*Edge)}
	return g
}

func NewGraphNamedNodes(s []string, t []string) *Graph {
	g := NewGraph()

	getVertex := func(ID string) *Vertex {
		v, ok := g.Vertices[ID]
		if !ok {
			v := NewVertex(ID, len(g.Vertices))
			g.Vertices[ID] = v
			return v
		}
		return v
	}

	for i := 0; i < len(s); i++ {
		source := getVertex(s[i])
		sink := getVertex(t[i])
		key := s[i] + "-" + t[i]
		g.Edges[key] = NewEdge(key, source, sink)

	}

	return g
}

type Vertex struct {
	id    string
	index int // needs to be count up as vertices are added to the graph
	Edges map[string]*Edge
}

func NewVertex(ID string, index int) *Vertex {
	v := &Vertex{id: ID, index: index}
	return v
}

func (v *Vertex) Index() int {
	return v.index
}

type Edge struct {
	id       string
	position int
	source   *Vertex
	sink     *Vertex
}

func NewEdge(ID string, source, sink *Vertex) *Edge {
	e := &Edge{id: ID, source: source, sink: sink}
	return e
}

func (e *Edge) Source() *Vertex {
	return e.source
}

func (e *Edge) Sink() *Vertex {
	return e.sink
}
