package graphblas_test

import (
	"math/rand"
	"testing"

	"github.com/rossmerr/graphblas"
)

var denseMatrix graphblas.Matrix[float64]
var csrMatrix graphblas.Matrix[float64]
var cscMatrix graphblas.Matrix[float64]

func init() {
	denseMatrix = dense(100)
	csrMatrix = csr(100)
	cscMatrix = csc(100)
}

func BenchmarkMatrixDenseAt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := denseMatrix.At(50, 50)
		if !(v >= 0) {
			b.Fatal("assert failed")
		}
	}
}

func BenchmarkMatrixCSRAt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := csrMatrix.At(50, 50)
		if !(v >= 0) {
			b.Fatal("assert failed")
		}
	}
}

func BenchmarkMatrixCSCAt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		v := cscMatrix.At(50, 50)
		if !(v >= 0) {
			b.Fatal("assert failed")
		}
	}
}

func BenchmarkMatrixDenseMultiply(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		denseMatrix.Multiply(denseMatrix)
	}
}

func BenchmarkMatrixCSRMultiply(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		csrMatrix.Multiply(denseMatrix)
	}
}

func BenchmarkMatrixCSCMultiply(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cscMatrix.Multiply(denseMatrix)
	}
}

func BenchmarkMatrixDenseAdd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		denseMatrix.Add(denseMatrix)
	}
}

func BenchmarkMatrixCSRAdd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		csrMatrix.Add(denseMatrix)
	}
}

func BenchmarkMatrixCSCAdd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		cscMatrix.Add(denseMatrix)
	}
}

func dense(n int) graphblas.Matrix[float64] {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := graphblas.NewDenseMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := graphblas.NewDenseMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csr(n int) graphblas.Matrix[float64] {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := graphblas.NewCSRMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := graphblas.NewCSRMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csc(n int) graphblas.Matrix[float64] {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := graphblas.NewCSCMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := graphblas.NewCSCMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}
