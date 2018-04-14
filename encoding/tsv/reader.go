package tsv

import (
	"encoding/csv"
	"io"
)

// Reader Tab-Separated Values (TSV) file format
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
func (r *Reader) Read() (record []string, err error) {
	return r.csv.Read()
}

// ReadAll reads all the remaining records from r.
func (r *Reader) ReadAll() (records [][]string, err error) {

	return r.csv.ReadAll()
}
