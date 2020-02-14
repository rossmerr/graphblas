package doublePrecision_test

import (
	"math/rand"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS/doublePrecision"
)

var denseMatrix GraphBLAS.Matrix
var csrMatrix GraphBLAS.Matrix
var cscMatrix GraphBLAS.Matrix

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
