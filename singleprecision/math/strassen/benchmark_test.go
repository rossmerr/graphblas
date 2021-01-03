package strassen_test

import (
	"math/rand"
	"testing"

	"context"

	"github.com/rossmerr/graphblas/singleprecision"
	"github.com/rossmerr/graphblas/singleprecision/math/strassen"
)

var denseMatrix singleprecision.Matrix
var csrMatrix singleprecision.Matrix
var cscMatrix singleprecision.Matrix

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

func dense(n int) singleprecision.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := singleprecision.NewDenseMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := singleprecision.NewDenseMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csr(n int) singleprecision.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := singleprecision.NewCSRMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := singleprecision.NewCSRMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csc(n int) singleprecision.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := singleprecision.NewCSCMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := singleprecision.NewCSCMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}
