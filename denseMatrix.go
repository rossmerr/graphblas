package GraphBLAS

import "fmt"

// DenseMatrix a dense matrix
type DenseMatrix struct {
	data [][]float64
}

func NewDenseMatrix(r, c int) *DenseMatrix {
	return newMatrix(r, c, nil)
}

func newMatrix(r, c int, initialise func([]float64, int)) *DenseMatrix {
	s := &DenseMatrix{data: make([][]float64, r)}

	for i := 0; i < r; i++ {
		s.data[i] = make([]float64, c)
		if initialise != nil {
			initialise(s.data[i], i)
		}
	}

	return s
}

func (s *DenseMatrix) Columns() int {
	return len(s.data[0])
}

func (s *DenseMatrix) Rows() int {
	return len(s.data)
}

func (s *DenseMatrix) At(r, c int) (float64, error) {
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
	matrix := newMatrix(s.Rows(), s.Columns(), func(n []float64, x int) {
		for v := 0; v < s.Columns(); v++ {
			n[v] = alpha * s.data[x][v]
		}
	})

	return matrix
}

// Multiply multiplies a Matrix structure by another Matrix structure.
func (s *DenseMatrix) Multiply(m Matrix) (Matrix, error) {
	if s.Rows() != m.Columns() {
		return nil, fmt.Errorf("Can not multiply matrices found length miss match %+v, %+v", s.Rows(), m.Columns())
	}

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, x int) {
		for j := 0; j < m.Columns(); j++ {
			temp := 0.0
			for k := 0; k < s.Columns(); k++ {
				f, _ := m.At(k, j)
				temp += s.data[x][k] * f
			}
			row[j] = temp
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

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, x int) {
		for j := 0; j < m.Columns(); j++ {
			f, _ := m.At(x, j)
			row[j] = s.data[x][j] + f
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

	matrix := newMatrix(s.Rows(), m.Columns(), func(row []float64, x int) {
		for j := 0; j < m.Columns(); j++ {
			f, _ := m.At(x, j)
			row[j] = s.data[x][j] - f
		}
	})

	return matrix, nil
}

// Negative the negative of a matrix.
func (s *DenseMatrix) Negative() Matrix {
	matrix := newMatrix(s.Rows(), s.Columns(), func(row []float64, x int) {
		for j := 0; j < s.Columns(); j++ {
			row[j] = -s.data[x][j]
		}
	})

	return matrix
}

func (s *DenseMatrix) ColumnsAt(c int) ([]float64, error) {
	if c < 0 || c >= s.Columns() {
		return nil, fmt.Errorf("Column '%+v' is invalid", c)
	}

	return nil, nil
}

func (s *DenseMatrix) RowsAt(r int) ([]float64, error) {
	if r < 0 || r >= s.Rows() {
		return nil, fmt.Errorf("Row '%+v' is invalid", r)
	}

	return nil, nil
}

func (s *DenseMatrix) Copy() Matrix {
	return nil
}
