package tsv

import (
	"encoding/csv"
	"io"
	"strconv"
)

// Reader Tab-Separated Values (TSV) file format
// (Row, Col, Value) tuple describing the adjacency matrix of the graph.
// Adjacency matrix is of size Num_vertices x Num_vertices
type Reader struct {
	csv *csv.Reader
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	reader := &Reader{
		csv: csv.NewReader(r),
	}

	reader.csv.Comma = '\t'
	return reader
}

// Read reads one record (a slice of fields) from r.
func (s *Reader) Read() (r, c int, value float64, err error) {
	record, err := s.csv.Read()

	if err != nil {
		return
	}

	r, err = strconv.Atoi(record[0])
	if err != nil {
		return
	}

	c, err = strconv.Atoi(record[1])
	if err != nil {
		return
	}

	value, err = strconv.ParseFloat(record[2], 64)
	return
}

// ReadAll reads all the remaining records from r.
func (s *Reader) ReadAll() (matrix [][]float64, err error) {
	for {
		r, c, value, err := s.Read()

		if err != nil {
			break
		}

		if len(matrix) < r {
			count := len(matrix) - r
			for i := 0; i < count; i++ {
				matrix = append(matrix, make([]float64, c))
			}
		}

		if len(matrix[r]) < c {
			count := len(matrix[r]) - c
			for i := 0; i < count; i++ {
				matrix[r] = append(matrix[r], 0)
			}
		}

		matrix[r][c] = value
	}

	return
}
