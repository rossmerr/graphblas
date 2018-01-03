package GraphBLAS

type Graph struct {
	Vertices map[string]*Vertex
	Edges    map[string]*Edge
}

type Vertex struct {
	id       string
	position int // needs to be count up as vertices are added to the graph
	Edges    map[string]*Edge
}

func (v *Vertex) Index() int {
	return v.position
}

type Edge struct {
	id       string
	position int
}

func (e *Edge) Source() *Vertex {
	return nil
}

func (e *Edge) Sink() *Vertex {
	return nil
}
