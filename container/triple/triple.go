package triple

import "github.com/RossMerr/Caudex.GraphBLAS/container/table"

type Store struct {
	Triples []*Tuple
}

// Tuple (Row, Col, Value) tuple describing the adjacency matrix of the graph
type Tuple struct {
	Row    string
	Column string
	Value  float64
}

type tripleStart map[string]int

func NewTripleStoreFromTable(t table.Table) *Store {
	ts := tripleStart{}
	store := &Store{Triples: make([]*Tuple, 0)}

	t.Iterator(func(r, c string, v float64) {
		store.newTriple(ts, r, c, v)
	})

	return store
}

func NewTripleStoreTransposeFromTable(t table.Table) *Store {
	ts := tripleStart{}
	store := &Store{Triples: make([]*Tuple, 0)}

	t.Iterator(func(r, c string, v float64) {
		store.newTriple(ts, c, r, v)
	})

	return store
}

func (s *Store) newTriple(ts tripleStart, r, c string, v float64) {
	triple := &Tuple{Row: r, Column: c, Value: v}

	if start, ok := ts[r]; ok {
		s.Triples = append(s.Triples[:start], append([]*Tuple{triple}, s.Triples[start:]...)...)

	} else {
		length := len(ts)
		ts[r] = length
		s.Triples = append(s.Triples[:length], append([]*Tuple{triple}, s.Triples[length:]...)...)

	}
}

// Transpose swap the row's and column's
func (s *Store) Transpose() *Store {
	store := &Store{Triples: make([]*Tuple, 0)}

	for _, t := range s.Triples {
		triple := &Tuple{Row: t.Column, Column: t.Row, Value: t.Value}

		store.Triples = append(store.Triples, triple)
	}

	return store
}
