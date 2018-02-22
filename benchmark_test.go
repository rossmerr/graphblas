package GraphBLAS_test

import (
	"math/rand"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func BenchmarkMatrixDense(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := dense(1)
		v := c.At(0, 0)
		if !(v >= 0) {
			b.Fatal("assert failed")
		}
	}

}

func BenchmarkMatrixCSR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := csr(1)
		v := c.At(0, 0)
		if !(v >= 0) {
			b.Fatal("assert failed")
		}
	}
}

func BenchmarkMatrixCSC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := csc(1)
		v := c.At(0, 0)
		if !(v >= 0) {
			b.Fatal("assert failed")
		}
	}
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

func csr(n int) GraphBLAS.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := GraphBLAS.NewCSRMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := GraphBLAS.NewCSRMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csc(n int) GraphBLAS.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := GraphBLAS.NewCSCMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := GraphBLAS.NewCSCMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}
