package GraphBLAS

import (
	"fmt"
)

// CSCMatrix compressed storage by columns (CSC)
type CSCMatrix struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []int
	rows     []int
	colStart []int
}

// NewCSCMatrix returns an GraphBLAS.CSCMatrix.
func NewCSCMatrix(r, c int) *CSCMatrix {
	s := &CSCMatrix{
		r:        r,
		c:        c,
		values:   make([]int, 0),
		rows:     make([]int, 0),
		colStart: make([]int, c+1),
	}

	return s
}

// At returns the value of a matrix element at r-th, c-th.
func (s *CSCMatrix) At(r, c int) (int, error) {
	if r < 0 || r >= s.r {
		return 0, fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return 0, fmt.Errorf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.rowIndex(r, c)

	if pointerStart < pointerEnd && s.rows[pointerStart] == r {
		return s.values[pointerStart], nil
	}

	return 0, nil
}

func (s *CSCMatrix) Set(r, c, value int) error {
	if r < 0 || r >= s.r {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.rowIndex(r, c)

	if pointerStart < pointerEnd && s.rows[pointerStart] == r {
		if value == 0 {
			s.remove(pointerStart, c)
		} else {
			s.values[pointerStart] = value
		}
	} else {
		s.insert(pointerStart, r, c, value)
	}

	return nil
}

func (s *CSCMatrix) Columns(c int) ([]int, error) {
	if c < 0 || c >= s.c {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	start := s.colStart[c]
	end := start

	if c+1 != s.c {
		end = s.colStart[c+1]
	}

	columns := make([]int, s.c)
	for i := start; i < end; i++ {
		columns[s.rows[i]] = s.values[i]
	}

	return columns, nil
}

func (s *CSCMatrix) Rows(r int) ([]int, error) {
	if r < 0 || r >= s.r {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	rows := make([]int, s.r)

	for c := range s.colStart[:s.c] {
		pointerStart, _ := s.rowIndex(r, c)
		rows[c] = s.values[pointerStart]
	}

	return rows, nil
}

func (s *CSCMatrix) insert(pointer, r, c, value int) {
	if value == 0 {
		return
	}

	s.rows = append(s.rows[:pointer], append([]int{r}, s.rows[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]int{value}, s.values[pointer:]...)...)

	for i := c + 1; i <= s.c; i++ {
		s.colStart[i]++
	}
}

func (s *CSCMatrix) remove(pointer, c int) {
	s.rows = append(s.rows[:pointer], s.rows[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)

	for i := c + 1; i <= s.c; i++ {
		s.colStart[i]--
	}
}

func (s *CSCMatrix) rowIndex(r, c int) (int, int) {

	start := s.colStart[c]
	end := s.colStart[c+1]

	if start-end == 0 {
		return start, end
	}

	if r > s.rows[end-1] {
		return end, end
	}

	for start < end {
		p := (start + end) / 2
		if s.rows[p] > r {
			end = p
		} else if s.rows[p] < r {
			start = p + 1
		} else {
			return p, end
		}
	}

	return start, end
}

func (s *CSCMatrix) sparse() {
}
