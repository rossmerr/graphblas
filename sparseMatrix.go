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

	size := r * c

	s := &SparseMatrix{
		r:        r,
		c:        c,
		values:   make([]int, size),
		rows:     make([]int, size),
		colStart: make([]int, c+1),
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

	pointer := s.searchForRowIndex(r, s.colStart[c], s.colStart[c+1])

	if pointer < s.colStart[c+1] && s.rows[pointer] == r {
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

	// 	s.values = append(s.values[:pointer], append([]int{value}, s.values[pointer+1:]...)...)
	// 	s.rows = append(s.rows[:pointer], append([]int{r}, s.rows[pointer+1:]...)...)

	s.values[pointer] = value
	s.rows[pointer] = r

	for cc := c + 1; cc < s.c+1; cc++ {
		s.colStart[cc]++
	}

}

func (s *SparseMatrix) remove(pointer, c int) {
	// 	s.values = append(s.values[:pointer+1], s.values[pointer:]...)
	// 	s.rows = append(s.rows[:pointer+1], s.rows[pointer:]...)

	for cc := c + 1; cc < s.c+1; cc++ {
		s.colStart[cc]--
	}
}

func (s *SparseMatrix) searchForRowIndex(r, left, right int) int {
	if right-left == 0 || r > s.rows[right-1] {
		return right
	}

	for left < right {
		p := (left + right) / 2
		if s.rows[p] > r {
			right = p
		} else if s.rows[p] < r {
			left = p + 1
		} else {
			return p
		}
	}

	return left
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
