package GraphBLAS

import "fmt"

// DenseMatrix a dense matrix
type DenseMatrix struct {
	c    int
	r    int
	data [][]float64
}

func NewDenseMatrix(r, c int) *DenseMatrix {
	return newMatrix(r, c, nil)
}

func newMatrix(r, c int, initialise func([]float64, int)) *DenseMatrix {
	s := &DenseMatrix{data: make([][]float64, r), r: r, c: c}

	for i := 0; i < r; i++ {
		s.data[i] = make([]float64, c)
		if initialise != nil {
			initialise(s.data[i], i)
		}
	}

	return s
}

func (s *DenseMatrix) Columns() int {
	return s.c
}

func (s *DenseMatrix) Rows() int {
	return s.r
}

func (s *DenseMatrix) At(r, c int) (float64, error) {
	if r < 0 || r >= s.Rows() {
		return 0, fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		return 0, fmt.Errorf("Column '%+v' is invalid", c)
	}

	return s.data[r][c], nil
}

func (s *DenseMatrix) Set(r, c int, value float64) error {
	if r < 0 || r >= s.Rows() {
		return fmt.Errorf("Row '%+v' is invalid", r)
	}

	if c < 0 || c >= s.Columns() {
		return fmt.Errorf("Column '%+v' is invalid", c)
	}

	s.data[r][c] = value

	return nil
}

// Scalar multiplication
func (s *DenseMatrix) Scalar(alpha float64) Matrix {
	matrix := newMatrix(s.Rows(), s.Columns(), func(rows []float64, r int) {
		for c := 0; c < s.Columns(); c++ {
			rows[c] = alpha * s.data[r][c]
		}
	})

	return matrix
}

// Multiply multiplies a Matrix structure by another Matrix structure.
func (s *DenseMatrix) Multiply(m Matrix) (Matrix, error) {
	if s.Rows() != m.Columns() {
		return nil, fmt.Errorf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, r int) {
		for c := 0; c < m.Columns(); c++ {
			total := 0.0
			for k := 0; k < s.Columns(); k++ {
				v, _ := m.At(k, c)
				total += v * s.data[r][k]
			}
			row[c] = total
		}
	})

	return matrix, nil
}

// Add addition of a Matrix structure by another Matrix structure.
func (s *DenseMatrix) Add(m Matrix) (Matrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, r int) {
		for c := 0; c < m.Columns(); c++ {
			v, _ := m.At(r, c)
			row[c] = s.data[r][c] + v
		}
	})

	return matrix, nil
}

// Subtract subtracts one matrix from another.
func (s *DenseMatrix) Subtract(m Matrix) (Matrix, error) {
	if s.Columns() != m.Columns() {
		return nil, fmt.Errorf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		return nil, fmt.Errorf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, r int) {
		for c := 0; c < m.Columns(); c++ {
			v, _ := m.At(r, c)
			row[c] = s.data[r][c] - v
		}
	})

	return matrix, nil
}

// Negative the negative of a matrix.
func (s *DenseMatrix) Negative() Matrix {
	matrix := newMatrix(s.Rows(), s.Columns(), func(row []float64, r int) {
		for c := 0; c < s.Columns(); c++ {
			row[c] = -s.data[r][c]
		}
	})

	return matrix
}

func (s *DenseMatrix) ColumnsAt(c int) ([]float64, error) {
	if c < 0 || c >= s.Columns() {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	columns := make([]float64, s.c)

	for r := 0; r < s.r; r++ {
		columns[r] = s.data[r][c]
	}

	return columns, nil
}

func (s *DenseMatrix) RowsAt(r int) ([]float64, error) {
	if r < 0 || r >= s.Rows() {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	return s.data[r], nil
}

func (s *DenseMatrix) Copy() Matrix {
	matrix := newMatrix(s.Rows(), s.Columns(), func(row []float64, r int) {
		for c := 0; c < s.Columns(); c++ {
			row[c] = s.data[r][c]
		}
	})

	return matrix
}
