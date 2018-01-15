package GraphBLAS

import (
	"fmt"
)

// SparseMatrix compressed storage by columns (CSC)
type SparseMatrix struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []int
	rows     []int
	colStart []int
}

// NewSparseMatrix returns an GraphBLAS.SparseMatrix.
func NewSparseMatrix(r, c int) *SparseMatrix {
	s := &SparseMatrix{
		r:        r,
		c:        c,
		values:   make([]int, 0),
		rows:     make([]int, 0),
		colStart: make([]int, c),
	}

	return s
}

func (s *SparseMatrix) Set(r, c, value int) error {
	if r < 0 || r >= s.r {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	c1 := c + 1
	if c+1 == s.c {
		c1 = c
	}

	pointer := s.rowIndex(r, c)

	if pointer < s.colStart[c1] && s.rows[pointer] == r {
		if value == 0 {
			s.remove(pointer, c)
		} else {
			s.values[pointer] = value
		}
	} else {
		s.insert(pointer, r, c, value)
	}

	return nil
}

func (s *SparseMatrix) insert(pointer, r, c, value int) {
	if value == 0 {
		return
	}

	s.values[pointer] = value
	s.rows[pointer] = r

	for cc := c + 1; cc < s.c; cc++ {
		s.colStart[cc]++
	}
}

func (s *SparseMatrix) remove(pointer, c int) {
	for cc := c + 1; cc < s.c; cc++ {
		s.colStart[cc]--
	}
}

func (s *SparseMatrix) rowIndex(r, c int) int {

	start := s.colStart[c]
	end := start

	if c+1 != s.c {
		end = s.colStart[c+1]
	}

	if len(s.rows) <= end {
		s.rows = append(s.rows, -1)
		s.values = append(s.values, -1)
	}

	if start-end == 0 {
		return start
	}

	if len(s.rows) <= end-1 {
		s.rows = append(s.rows, -1)
		s.values = append(s.values, -1)
	}

	if r > s.rows[end-1] {
		return end
	}

	for start < end {
		p := (start + end) / 2
		if s.rows[p] > r {
			end = p
		} else if s.rows[p] < r {
			start = p + 1
		} else {
			return p
		}
	}

	return start
}

func (s *SparseMatrix) Output() {
	fmt.Print("\ncolStart \n")

	for k, v := range s.colStart {
		fmt.Printf("%+v: %+v\n", k, v)
	}

	fmt.Print("\nrows \n")
	for k, v := range s.rows {
		fmt.Printf("%+v: %+v\n", k, v)
	}
	fmt.Print("\nvalues \n")
	for k, v := range s.values {
		fmt.Printf("%+v: %+v\n", k, v)
	}
}
