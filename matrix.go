package GraphBLAS

// Matrix a dense matrix
type Matrix [][]int

func NewMatrix(m, n int) Matrix {
	return newMatrix(m, n, nil)
}

func newMatrix(m, n int, initialise func([]int, int)) Matrix {
	s := make(Matrix, m)

	for i := 0; i < m; i++ {
		s[i] = make([]int, n)
		if initialise != nil {
			initialise(s[i], i)
		}
	}

	return s
}

// Scalar multiplication
func (s Matrix) Scalar(alpha int) Matrix {
	matrix := newMatrix(len(s), len(s[0]), func(n []int, x int) {
		for v := 0; v < len(s[0]); v++ {
			n[v] = alpha * s[x][v]
		}
	})

	return matrix
}

// Multiply multiplies a Matrix structure by another Matrix structure.
func (m Matrix) Multiply(m2 Matrix) (Matrix, bool) {
	if len(m[0]) != len(m2) {
		return nil, false
	}

	matrix := newMatrix(len(m), len(m2[0]), func(row []int, x int) {
		for j := 0; j < len(m2[0]); j++ {
			temp := 0
			for k := 0; k < len(m[0]); k++ {
				temp += m[x][k] * m2[k][j]
			}
			row[j] = temp
		}
	})

	return matrix, true
}

// Add addition of a Matrix structure by another Matrix structure.
func (m Matrix) Add(m2 Matrix) (Matrix, bool) {
	if len(m) != len(m2) {
		return nil, false
	}

	if len(m[0]) != len(m2[0]) {
		return nil, false
	}

	matrix := newMatrix(len(m), len(m2[0]), func(row []int, x int) {
		for j := 0; j < len(m2[0]); j++ {
			row[j] = m[x][j] + m2[x][j]
		}
	})

	return matrix, true
}

// Subtract subtracts one matrix from another.
func (m Matrix) Subtract(m2 Matrix) (Matrix, bool) {
	if len(m) != len(m2) {
		return nil, false
	}

	if len(m[0]) != len(m2[0]) {
		return nil, false
	}

	matrix := newMatrix(len(m), len(m2[0]), func(row []int, x int) {
		for j := 0; j < len(m2[0]); j++ {
			row[j] = m[x][j] - m2[x][j]
		}
	})

	return matrix, true
}

// Negative the negative of a matrix.
func (m Matrix) Negative() Matrix {

	matrix := newMatrix(len(m), len(m[0]), func(row []int, x int) {
		for j := 0; j < len(m[0]); j++ {
			row[j] = -m[x][j]
		}
	})

	return matrix
}
