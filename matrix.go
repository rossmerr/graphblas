package GraphBLAS

type Matrix [][]int

func NewMatrix(x, y int) Matrix {
	return newMatrix(x, y, nil)
}

func newMatrix(x, y int, initialise func([]int, int)) Matrix {
	m := make(Matrix, x)

	for i := 0; i < x; i++ {
		m[i] = make([]int, y)
		if initialise != nil {
			initialise(m[i], i)
		}
	}

	return m
}

// Scalar multiplication
func (m Matrix) Scalar(s int) Matrix {
	matrix := newMatrix(len(m), len(m[0]), func(n []int, x int) {
		for v := 0; v < len(m[0]); v++ {
			n[v] = s * m[x][v]
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
