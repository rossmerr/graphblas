// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package table

import (
	"io"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
	"github.com/RossMerr/Caudex.GraphBLAS/container"
)

const (
	emptyFloat64 = 0.0
)

// Table is a set of data elements using a model of columns and rows
type Table interface {
	Iterator(i func(string, string, interface{})) bool
	Columns() int
	Rows() int
	Get(r, c string) interface{}
	GetFloat64(r, c string) float64
}

type table struct {
	matrix        GraphBLAS.Matrix
	rowIndices    []string
	columnIndices []string
	columns       map[string]int
}

func newTable(r, c int) *table {
	return &table{
		matrix:        GraphBLAS.NewCSCMatrix(r, c),
		rowIndices:    make([]string, r),
		columnIndices: make([]string, c),
		columns:       make(map[string]int, c),
	}
}

// NewTableFromReader returns a table.Table
func NewTableFromReader(r, c int, reader container.Reader) (Table, error) {
	table := newTable(r, c)

	// Read the header
	line, err := reader.Read()
	if err != nil {
		return nil, err
	}

	header := line

	// Read the body
	count := 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		table.read(header, count, line)
		count++
	}

	return table, nil
}

func (s *table) read(header []string, r int, row []string) {
	indice := header[0]
	s.rowIndices[r] = indice + "|" + row[0]

	for i := 1; i < len(row); i++ {
		// Column header name
		uniqueTypeValuePair := header[i] + "|" + row[i]
		v := 1.0

		if c, ok := s.columns[uniqueTypeValuePair]; ok {
			v += s.matrix.At(r, c)
			s.matrix.Set(r, c, v)
		} else {
			c = len(s.columns)
			s.columns[uniqueTypeValuePair] = c
			s.columnIndices[c] = uniqueTypeValuePair
			s.matrix.Set(r, c, v)
		}
	}
}

// Columns the number of columns of the matrix
func (s *table) Columns() int {
	return s.matrix.Columns()
}

// Rows the number of rows of the matrix
func (s *table) Rows() int {
	return s.matrix.Rows()
}

// Get (unoptimized) returns the value of a table element at r-th, c-th
func (s *table) Get(r, c string) interface{} {
	cPointer := s.columns[c]
	rPointer := -1
	for i, value := range s.rowIndices {
		if value == r {
			rPointer = i
			break
		}
	}

	return s.matrix.At(rPointer, cPointer)
}

func (s *table) GetFloat64(r, c string) float64 {
	v := s.Get(r, c)
	if value, ok := v.(float64); ok {
		return value
	}
	return emptyFloat64
}

// Iterator iterates through all non-zero elements, order is not guaranteed
func (s *table) Iterator(i func(string, string, interface{})) bool {
	enumerator := s.matrix.Enumerate()
	if enumerator.HasNext() {
		r, c, v := enumerator.Next()
		i(s.rowIndices[r], s.columnIndices[c], v)
		return true
	}

	return false
}
