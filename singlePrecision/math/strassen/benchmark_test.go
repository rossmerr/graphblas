package strassen_test

import (
	"math/rand"
	"testing"

	"golang.org/x/net/context"

	"github.com/RossMerr/Caudex.GraphBLAS/singlePrecision"
	"github.com/RossMerr/Caudex.GraphBLAS/singlePrecision/math/strassen"
)

var denseMatrix singlePrecision.Matrix
var csrMatrix singlePrecision.Matrix
var cscMatrix singlePrecision.Matrix

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

func dense(n int) singlePrecision.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := singlePrecision.NewDenseMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := singlePrecision.NewDenseMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csr(n int) singlePrecision.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := singlePrecision.NewCSRMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := singlePrecision.NewCSRMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csc(n int) singlePrecision.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := singlePrecision.NewCSCMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := singlePrecision.NewCSCMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}
