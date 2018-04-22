// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package triple

import "github.com/RossMerr/Caudex.GraphBLAS/container/table"

// Store is a triplestore for the storage and retrieval of triples
type Store struct {
	Triples []*Triple
}

// Triple (Row, Col, Value) tuple describing the adjacency matrix of the graph
type Triple struct {
	Row    string
	Column string
	Value  float64
}

type tripleStart map[string]int

// NewTripleStoreFromTable returns a triple.Store
func NewTripleStoreFromTable(t table.Table) *Store {
	ts := tripleStart{}
	store := &Store{Triples: make([]*Triple, 0)}

	t.Iterator(func(r, c string, v float64) {
		store.newTriple(ts, r, c, v)
	})

	return store
}

// NewTripleStoreTransposeFromTable returns a triple.Store transposed
func NewTripleStoreTransposeFromTable(t table.Table) *Store {
	ts := tripleStart{}
	store := &Store{Triples: make([]*Triple, 0)}

	t.Iterator(func(r, c string, v float64) {
		store.newTriple(ts, c, r, v)
	})

	return store
}

func (s *Store) newTriple(ts tripleStart, r, c string, v float64) {
	triple := &Triple{Row: r, Column: c, Value: v}

	if start, ok := ts[r]; ok {
		s.Triples = append(s.Triples[:start], append([]*Triple{triple}, s.Triples[start:]...)...)

	} else {
		length := len(ts)
		ts[r] = length
		s.Triples = append(s.Triples[:length], append([]*Triple{triple}, s.Triples[length:]...)...)

	}
}

// Transpose swap the row's and column's
func (s *Store) Transpose() *Store {
	store := &Store{Triples: make([]*Triple, 0)}

	for _, t := range s.Triples {
		triple := &Triple{Row: t.Column, Column: t.Row, Value: t.Value}

		store.Triples = append(store.Triples, triple)
	}

	return store
}
