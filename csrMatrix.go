package GraphBLAS

import "fmt"

// CSRMatrix compressed storage by rows (CSR)
type CSRMatrix struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []int
	cols     []int
	rowStart []int
}

// NewCSRMatrix returns an GraphBLAS.CSRMatrix.
func NewCSRMatrix(r, c int) *CSRMatrix {
	s := &CSRMatrix{
		r:        r,
		c:        c,
		values:   make([]int, 0),
		cols:     make([]int, 0),
		rowStart: make([]int, r),
	}

	return s
}

// At returns the value of a matrix element at r-th, c-th.
func (s *CSRMatrix) At(r, c int) (int, error) {
	if r < 0 || r >= s.r {
		return 0, fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return 0, fmt.Errorf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.columnIndex(r, c)

	if pointerStart <= pointerEnd && s.cols[pointerStart] == r {
		return s.values[pointerStart], nil
	}

	return 0, nil
}

func (s *CSRMatrix) Set(r, c, value int) error {
	if r < 0 || r >= s.r {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.columnIndex(r, c)

	if pointerStart < pointerEnd && s.cols[pointerStart] == r {
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

func (s *CSRMatrix) Columns(c int) ([]int, error) {
	if c < 0 || c >= s.c {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	start := s.rowStart[c]
	end := start

	if c+1 != s.c {
		end = s.rowStart[c+1]
	}

	columns := make([]int, s.c)
	for i := start; i < end; i++ {
		columns[s.cols[i]] = s.values[i]
	}

	return columns, nil
}

func (s *CSRMatrix) Rows(r int) ([]int, error) {
	if r < 0 || r >= s.r {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	rows := make([]int, s.r)

	for c := range s.rowStart {
		pointerStart, _ := s.columnIndex(r, c)
		rows[c] = s.values[pointerStart]
	}

	return rows, nil
}

func (s *CSRMatrix) insert(pointer, r, c, value int) {
	if value == 0 {
		return
	}

	s.cols = append(s.cols[:pointer], append([]int{c}, s.cols[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]int{value}, s.values[pointer:]...)...)

	for i := r + 1; i < s.r; i++ {
		s.rowStart[i]++
	}
}

func (s *CSRMatrix) remove(pointer, r int) {
	s.cols = append(s.cols[:pointer], s.cols[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)

	for i := r + 1; i < s.c; i++ {
		s.rowStart[i]--
	}
}

func (s *CSRMatrix) columnIndex(r, c int) (int, int) {

	start := s.rowStart[r]
	end := start

	if r+1 != s.r {
		end = s.rowStart[r+1]
	}

	if start-end == 0 {
		return start, end
	}

	if r > s.cols[end-1] {
		return end, end
	}

	for start < end {
		p := (start + end) / 2
		if s.cols[p] > c {
			end = p
		} else if s.cols[p] < c {
			start = p + 1
		} else {
			return p, end
		}
	}

	return start, end
}

func (s *CSRMatrix) sparse() {
}
