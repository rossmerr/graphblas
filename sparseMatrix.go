package GraphBLAS

import (
	"fmt"
)

// SparseMatrix compressed storage by columns (CSC)
type SparseMatrix struct {
	m        int // number of rows in the sparse matrix
	n        int // number of columns in the sparse matrix
	values   []int
	rows     []int
	colStart map[int]int
}

// NewSparseMatrix returns an GraphBLAS.SparseMatrix.
func NewSparseMatrix(m, n int) *SparseMatrix {
	s := &SparseMatrix{m: m, n: n, values: make([]int, 0), rows: make([]int, 0), colStart: make(map[int]int)}
	return s
}

// Set the matrix element value at the r-th row and c-th column.
func (s *SparseMatrix) Set(r, c, value int) {
	if value == 0 {
		return
	}

	if pointer, ok := s.colStart[c]; ok {
		indexStart := s.rows[pointer]

		if indexStart == r {
			s.values[indexStart] = value
			return
		}
		pointerEnd := s.colStart[c+1]

		indexEnd := s.rows[pointerEnd]

		for i := indexStart + 1; i < indexEnd; i++ {

			if i == r {
				s.values[i] = value
				return
			} else if i > r {
				s.rows = append(s.rows[:i], append([]int{i}, s.rows[i:]...)...)
				s.values = append(s.values[:i], append([]int{value}, s.values[i:]...)...)
				return
			}

		}

		// need to insert here

	} else {
		s.rows = append(s.rows, r)
		s.values = append(s.values, value)
		s.colStart[c] = len(s.rows) - 1
		fmt.Print("hit \n")
	}

}

// rows
// 0: 0
// 1: 0
// 2: 1

// colStart
// 0: 0
// 2: 1
// 1: 2
