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
	rA := len(m)
	cA := len(m[0])
	rB := len(m2)
	cB := len(m2[0])
	if cA != rB {
		return nil, false
	}

	matrix := NewMatrix(rA, cB, nil)

	for i := 0; i < rA; i++ {
		for j := 0; j < cB; j++ {
			temp := 0
			for k := 0; k < cA; k++ {
				temp += m[i][k] * m2[k][j]
			}
			matrix[i][j] = temp
		}
	}

	return matrix, true
}
