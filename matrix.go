package GraphBLAS

type Matrix [][]int

func NewMatrix(x, y int, initialise func([]int, int)) Matrix {
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
	matrix := NewMatrix(len(m), len(m[0]), func(n []int, x int) {
		for v := 0; v < len(m[0]); v++ {
			n[v] = s * m[x][v]
		}
	})

	return matrix
}

// Multiply Multiplies a Matrix structure by another Matrix structure.
func (m Matrix) Multiply(m2 Matrix) (Matrix, bool) {
	if len(m[0]) != len(m2) {
		return nil, false
	}

	matrix := NewMatrix(len(m), len(m2[0]), func(row []int, x int) {
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
