package strassen_test

import (
	"math/rand"
	"testing"

	"context"

	"github.com/rossmerr/graphblas/f32"
	"github.com/rossmerr/graphblas/f32/math/strassen"
)

var denseMatrix f32.Matrix
var csrMatrix f32.Matrix
var cscMatrix f32.Matrix

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

func dense(n int) f32.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := f32.NewDenseMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := f32.NewDenseMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csr(n int) f32.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := f32.NewCSRMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := f32.NewCSRMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}

func csc(n int) f32.Matrix {
	aData := make([][]float32, n)
	for r := range aData {
		aData[r] = make([]float32, n)
		for c := range aData {
			aData[r][c] = rand.Float32()

		}
	}
	a := f32.NewCSCMatrixFromArray(aData)

	bData := make([][]float32, n)
	for r := range bData {
		bData[r] = make([]float32, n)
		for c := range bData {
			bData[r][c] = rand.Float32()

		}
	}
	b := f32.NewCSCMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}
