package GraphBLAS

import "fmt"

// CSRMatrix compressed storage by rows (CSR)
type CSRMatrix struct {
	r        int // number of rows in the sparse matrix
	c        int // number of columns in the sparse matrix
	values   []float64
	cols     []int
	rowStart []int
}

// NewCSRMatrix returns an GraphBLAS.CSRMatrix.
func NewCSRMatrix(r, c int) *CSRMatrix {
	return newCSRMatrix(r, c, 0)
}

func newCSRMatrix(r, c, l int) *CSRMatrix {
	s := &CSRMatrix{
		r:        r,
		c:        c,
		values:   make([]float64, l),
		cols:     make([]int, l),
		rowStart: make([]int, r+1),
	}

	return s
}

func (s *CSRMatrix) Columns() int {
	return s.c
}

func (s *CSRMatrix) Rows() int {
	return s.r
}

// At returns the value of a matrix element at r-th, c-th.
func (s *CSRMatrix) At(r, c int) (float64, error) {
	if r < 0 || r >= s.r {
		return 0, fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return 0, fmt.Errorf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.columnIndex(r, c)

	if pointerStart < pointerEnd && s.cols[pointerStart] == c {
		return s.values[pointerStart], nil
	}

	return 0, nil
}

func (s *CSRMatrix) Set(r, c int, value float64) error {
	if r < 0 || r >= s.r {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.c {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	pointerStart, pointerEnd := s.columnIndex(r, c)

	if pointerStart < pointerEnd && s.cols[pointerStart] == c {
		if value == 0 {
			s.remove(pointerStart, r)
		} else {
			s.values[pointerStart] = value
		}
	} else {
		s.insert(pointerStart, r, c, value)
	}

	return nil
}

func (s *CSRMatrix) ColumnsAt(c int) ([]float64, error) {
	if c < 0 || c >= s.c {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	columns := make([]float64, s.c)

	for r := range s.rowStart[:s.r] {
		pointerStart, _ := s.columnIndex(r, c)
		columns[r] = s.values[pointerStart]
	}

	return columns, nil

}

func (s *CSRMatrix) RowsAt(r int) ([]float64, error) {
	if r < 0 || r >= s.r {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	start := s.rowStart[r]
	end := start

	if r+1 != s.r {
		end = s.rowStart[r+1]
	}

	rows := make([]float64, s.r)
	for i := start; i < end; i++ {
		rows[s.cols[i]] = s.values[i]
	}

	return rows, nil
}

func (s *CSRMatrix) insert(pointer, r, c int, value float64) {
	if value == 0 {
		return
	}

	s.cols = append(s.cols[:pointer], append([]int{c}, s.cols[pointer:]...)...)
	s.values = append(s.values[:pointer], append([]float64{value}, s.values[pointer:]...)...)

	for i := r + 1; i <= s.r; i++ {
		s.rowStart[i]++
	}
}

func (s *CSRMatrix) remove(pointer, r int) {
	s.cols = append(s.cols[:pointer], s.cols[pointer+1:]...)
	s.values = append(s.values[:pointer], s.values[pointer+1:]...)

	for i := r + 1; i <= s.r; i++ {
		s.rowStart[i]--
	}
}

func (s *CSRMatrix) columnIndex(r, c int) (int, int) {

	start := s.rowStart[r]
	end := s.rowStart[r+1]

	if start-end == 0 {
		return start, end
	}

	if c > s.cols[end-1] {
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

func (s *CSRMatrix) Copy() SparseMatrix {
	return s.copy(func(value float64) float64 {
		return value
	})
}

func (s *CSRMatrix) copy(action func(float64) float64) *CSRMatrix {
	matrix := newCSRMatrix(s.r, s.c, len(s.values))

	for i := range s.values {
		matrix.values[i] = action(s.values[i])
		matrix.cols[i] = s.cols[i]
	}

	for i := range s.rowStart {
		matrix.rowStart[i] = s.rowStart[i]
	}

	return matrix
}

// Scalar multiplication
func (s *CSRMatrix) Scalar(alpha float64) SparseMatrix {
	return s.copy(func(value float64) float64 {
		return alpha * value
	})
}

// Multiply multiplies a Matrix structure by another Matrix structure.
func (s *CSRMatrix) Multiply(m SparseMatrix) (SparseMatrix, error) {
	if s.Rows() != m.Columns() {
		return nil, fmt.Errorf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	matrix := newCSRMatrix(s.Rows(), m.Columns(), 0)

	for r := 0; r < s.Rows(); r++ {
		rows, err := s.RowsAt(r)
		if err != nil {
			return nil, err
		}

		for c := 0; c < m.Columns(); c++ {
			column, err := m.ColumnsAt(c)
			if err != nil {
				return nil, err
			}

			sum := 0.0
			for l := 0; l < len(rows); l++ {
				sum += rows[l] * column[l]
			}

			matrix.Set(r, c, sum)
		}

	}

	return matrix, nil
}

// Add addition of a Matrix structure by another Matrix structure.
func (s *CSRMatrix) Add(m SparseMatrix) (SparseMatrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := newCSRMatrix(s.Rows(), m.Columns(), 0)

	for r := 0; r < s.Rows(); r++ {
		sRows, err := s.RowsAt(r)
		if err != nil {
			return nil, err
		}

		mRows, err := m.RowsAt(r)
		if err != nil {
			return nil, err
		}

		for c := 0; c < s.Columns(); c++ {
			matrix.Set(r, c, sRows[c]+mRows[c])
		}
	}

	return matrix, nil
}

// Subtract subtracts one matrix from another.
func (s *CSRMatrix) Subtract(m SparseMatrix) (SparseMatrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := newCSRMatrix(s.Rows(), m.Columns(), 0)

	for r := 0; r < s.Rows(); r++ {
		sRows, err := s.RowsAt(r)
		if err != nil {
			return nil, err
		}

		mRows, err := m.RowsAt(r)
		if err != nil {
			return nil, err
		}

		for c := 0; c < s.Columns(); c++ {
			matrix.Set(r, c, sRows[c]-mRows[c])
		}
	}

	return matrix, nil
}

// Negative the negative of a matrix.
func (s *CSRMatrix) Negative() SparseMatrix {
	return s.copy(func(value float64) float64 {
		return -value
	})
}
