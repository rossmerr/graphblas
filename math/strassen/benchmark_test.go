package strassen_test

import (
	"math/rand"
	"testing"

	"golang.org/x/net/context"

	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/math/strassen"
)

var denseMatrix graphblas.Matrix[float64]
var csrMatrix graphblas.Matrix[float64]
var cscMatrix graphblas.Matrix[float64]

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

func dense(n int) graphblas.Matrix[float64] {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := graphblas.NewDenseMatrixFromArrayN(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := graphblas.NewDenseMatrixFromArrayN(bData)

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
