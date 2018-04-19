// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package tsv

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/RossMerr/Caudex.GraphBLAS"
)

// Reader Tab-Separated Values (TSV) file format
// (Row, Col, Value) tuple describing the adjacency matrix of the graph.
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
func (s *Reader) read() (r, c int, value float64, err error) {
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
func (s *Reader) ReadAll() (GraphBLAS.Matrix, error) {
	columnMax := 0
	matrix := [][]float64{}
	for {
		r, c, value, err := s.read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if columnMax < c {
			columnMax = c
		}

		if len(matrix) < r {
			count := r - len(matrix)
			for i := 0; i < count; i++ {
				matrix = append(matrix, make([]float64, columnMax))
			}
		}

		if len(matrix[r-1]) < c {
			count := columnMax - len(matrix[r-1])
			for i := 0; i < count; i++ {
				matrix[r-1] = append(matrix[r-1], 0)
			}
		}

		matrix[r-1][c-1] = value
	}

	// Set all zero elements in the matrix
	for r := range matrix {
		if len(matrix[r]) < columnMax {
			count := columnMax - len(matrix[r])
			for i := 0; i < count; i++ {
				matrix[r] = append(matrix[r], 0)
			}
		}
	}

	graph := GraphBLAS.NewDenseMatrixFromArray(matrix)

	return graph, nil
}
