// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package triples

import (
	"github.com/RossMerr/Caudex.GraphBLAS/container/table"
)

// Triple (Row, Col, Value) tuple describing the adjacency matrix of the graph
type Triple struct {
	Row    string
	Column string
	Value  interface{}
}

// NewTriplesFromTable returns a []*Triple
func NewTriplesFromTable(t table.Table) []*Triple {
	triples := make([]*Triple, 0)

	t.ReadAll()

	t.Iterator(func(r, c string, v interface{}) {
		triple := &Triple{Row: r, Column: c, Value: v}
		triples = append(triples, triple)
	})

	return triples
}

// NewTripleTransposeFromTable returns a []*Triple transposed
func NewTripleTransposeFromTable(t table.Table) []*Triple {
	triples := make([]*Triple, 0)

	t.ReadAll()

	t.Iterator(func(r, c string, v interface{}) {
		triple := &Triple{Row: c, Column: r, Value: v}
		triples = append(triples, triple)
	})

	return triples
}

// Transpose swap the row's and column's
func Transpose(tt []*Triple) []*Triple {
	triples := make([]*Triple, 0)

	for _, t := range tt {
		triple := &Triple{Row: t.Column, Column: t.Row, Value: t.Value}
		triples = append(triples, triple)
	}

	return triples
}
