package GraphBLAS_test

import (
	"math/rand"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func BenchmarkMatrix(b *testing.B) {
	for _, fn := range benchmarks {
		fn.fn(b)
	}
}

var benchmarks = []struct {
	name string
	fn   func(*testing.B)
}{
	{
		name: "matrix_multiply_dense",
		fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := dense(1)
				v := c.At(0, 0)
				if !(v >= 0) {
					b.Fatal("assert failed")
				}
			}
		},
	},
}

func dense(n int) GraphBLAS.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := GraphBLAS.NewDenseMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := GraphBLAS.NewDenseMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

// func csr(n int) GraphBLAS.Matrix {
// 	aData := make([][]float64, n)
// 	for r := range aData {
// 		aData[r] = make([]float64, n)
// 		for c := range aData {
// 			aData[r][c] = rand.Float64()

// 		}
// 	}
// 	a := GraphBLAS.NewCSRMatrix(aData)

// 	bData := make([][]float64, n)
// 	for r := range bData {
// 		bData[r] = make([]float64, n)
// 		for c := range bData {
// 			bData[r][c] = rand.Float64()

// 		}
// 	}
// 	b := GraphBLAS.NewCSRMatrix(bData)

// 	c := a.Multiply(b)
// 	return c
// }
