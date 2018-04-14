package triple

import "github.com/RossMerr/Caudex.GraphBLAS/container/table"

type Store struct {
	Triples []*Triple
}

type Triple struct {
	Row    string
	Column string
	Value  float64
}

type tripleStart map[string]int

func NewTripleStoreFromTable(t table.Table) *Store {
	ts := tripleStart{}
	store := &Store{Triples: make([]*Triple, 0)}

	t.Iterator(func(r, c string, v float64) {
		store.newTriple(ts, r, c, v)
	})

	return store
}

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
