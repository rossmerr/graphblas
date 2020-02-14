package strassen_test

import (
	"math/rand"
	"testing"

	"golang.org/x/net/context"

	"github.com/RossMerr/Caudex.GraphBLAS/doublePrecision"
	"github.com/RossMerr/Caudex.GraphBLAS/doublePrecision/math/strassen"
)

var denseMatrix doublePrecision.Matrix
var csrMatrix doublePrecision.Matrix
var cscMatrix doublePrecision.Matrix

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

func dense(n int) doublePrecision.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := doublePrecision.NewDenseMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := doublePrecision.NewDenseMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csr(n int) doublePrecision.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := doublePrecision.NewCSRMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := doublePrecision.NewCSRMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csc(n int) doublePrecision.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := doublePrecision.NewCSCMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := doublePrecision.NewCSCMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}
