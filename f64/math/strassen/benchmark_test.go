package strassen_test

import (
	"math/rand"
	"testing"

	"golang.org/x/net/context"

	"github.com/rossmerr/graphblas/f64"
	"github.com/rossmerr/graphblas/f64/math/strassen"
)

var denseMatrix f64.Matrix
var csrMatrix f64.Matrix
var cscMatrix f64.Matrix

func init() {
	denseMatrix = dense(100)
	csrMatrix = csr(100)
	cscMatrix = csc(100)
}

func BenchmarkMatrixDenseMultiplyStrassen(b *testing.B) {
	b.ReportAllocs()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		strassen.Multiply(ctx, denseMatrix, denseMatrix)
	}
}

func BenchmarkMatrixCSRMultiplyStrassen(b *testing.B) {
	b.ReportAllocs()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		strassen.Multiply(ctx, csrMatrix, denseMatrix)
	}
}

func BenchmarkMatrixCSCMultiplyStrassen(b *testing.B) {
	b.ReportAllocs()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		strassen.Multiply(ctx, cscMatrix, denseMatrix)
	}
}

func dense(n int) f64.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := f64.NewDenseMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := f64.NewDenseMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csr(n int) f64.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := f64.NewCSRMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := f64.NewCSRMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csc(n int) f64.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := f64.NewCSCMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := f64.NewCSCMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}
