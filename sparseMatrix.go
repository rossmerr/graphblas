package GraphBLAS

// SparseMatrix compressed Storage by Columns (CSC)
type SparseMatrix struct {
	values   []int
	rows     []int
	colStart map[int]int
}

func NewSparseMatrix(x, y int) *SparseMatrix {
	m := &SparseMatrix{values: make([]int, 0), rows: make([]int, 0), colStart: make(map[int]int)}
	return m
}

func (s *SparseMatrix) Set(x, y, value int) {
	if value == 0 {
		return
	}

	if pointer, ok := s.colStart[y]; ok {
		indexStart := s.rows[pointer]

		if indexStart == x {
			s.values[indexStart] = value
			return
		}

		pointerEnd := s.colStart[y+1]

		indexEnd := s.rows[pointerEnd]

		for i := indexStart + 1; i < indexEnd; i++ {

			if i == x {
				s.values[i] = value
			} else if i > x {
				s.rows = append(s.rows[:i], append([]int{i}, s.rows[i:]...)...)
				s.values = append(s.values[:i], append([]int{value}, s.values[i:]...)...)
				break
			}

		}

	} else {
		s.rows = append(s.rows, x)
		s.values = append(s.values, value)
		s.colStart[y] = len(s.rows)
	}

}
