// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package strassen

import (
	"context"
	"log"

	"github.com/rossmerr/graphblas/f64"
)

// Multiply multiplies a matrix by another matrix using the Strassen algorithm
func Multiply(ctx context.Context, a, b f64.Matrix) f64.Matrix {
	return MultiplyCrossoverPoint(ctx, a, b, 64)
}

// MultiplyCrossoverPoint multiplies a matrix by another matrix using the Strassen algorithm
// the crossover point is when to switch standard methods of matrix multiplication for more efficiency
func MultiplyCrossoverPoint(ctx context.Context, a, b f64.Matrix, crossover int) f64.Matrix {
	if a.Columns() != b.Rows() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", a.Columns(), b.Rows())
	}

	n := b.Rows()
	if n <= crossover {
		matrix := f64.NewDenseMatrix(a.Rows(), b.Columns())
		f64.MatrixMatrixMultiply(ctx, a, b, nil, matrix)
		return matrix
	}

	size := n / 2

	a11 := f64.NewDenseMatrix(size, size)
	a12 := f64.NewDenseMatrix(size, size)
	a21 := f64.NewDenseMatrix(size, size)
	a22 := f64.NewDenseMatrix(size, size)

	b11 := f64.NewDenseMatrix(size, size)
	b12 := f64.NewDenseMatrix(size, size)
	b21 := f64.NewDenseMatrix(size, size)
	b22 := f64.NewDenseMatrix(size, size)

	// dividing the matrices in 4 sub-matrices:
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			select {
			case <-ctx.Done():
				break
			default:
				a11.Set(r, c, a.At(r, c))           // top left
				a12.Set(r, c, a.At(r, c+size))      // top right
				a21.Set(r, c, a.At(r+size, c))      // bottom left
				a22.Set(r, c, a.At(r+size, c+size)) // bottom right

				b11.Set(r, c, b.At(r, c))           // top left
				b12.Set(r, c, b.At(r, c+size))      // top right
				b21.Set(r, c, b.At(r+size, c))      // bottom left
				b22.Set(r, c, b.At(r+size, c+size)) // bottom right
			}
		}
	}

	out := make(chan *mPlace)

	go subMatrixM(ctx, out, 1, a11.Add(a22), b11.Add(b22), crossover)
	go subMatrixM(ctx, out, 2, a21.Add(a22), b11, crossover)
	go subMatrixM(ctx, out, 3, a11, b12.Subtract(b22), crossover)
	go subMatrixM(ctx, out, 4, a22, b21.Subtract(b11), crossover)
	go subMatrixM(ctx, out, 5, a11.Add(a12), b22, crossover)
	go subMatrixM(ctx, out, 6, a21.Subtract(a11), b11.Add(b12), crossover)
	go subMatrixM(ctx, out, 7, a12.Subtract(a22), b21.Add(b22), crossover)

	m := [8]f64.Matrix{}
	for i := 0; i < 7; i++ {
		mtx := <-out
		m[mtx.m] = mtx.matrix
	}

	c11 := m[1].Add(m[4]).Subtract(m[5]).Add(m[7])
	c12 := m[3].Add(m[5])
	c21 := m[2].Add(m[4])
	c22 := m[1].Subtract(m[2]).Add(m[3]).Add(m[6])

	matrix := f64.NewDenseMatrix(c11.Rows()*2, c11.Rows()*2)
	shift := c11.Rows()

	// Combine the results
	for r := 0; r < c11.Rows(); r++ {
		for c := 0; c < c11.Columns(); c++ {
			select {
			case <-ctx.Done():
				break
			default:
				matrix.Set(r, c, c11.At(r, c))
				matrix.Set(r, c+shift, c12.At(r, c))
				matrix.Set(r+shift, c, c21.At(r, c))
				matrix.Set(r+shift, c+shift, c22.At(r, c))
			}
		}
	}

	return matrix
}

func subMatrixM(ctx context.Context, out chan *mPlace, m int, a, b f64.Matrix, crossover int) {
	out <- &mPlace{
		m:      m,
		matrix: MultiplyCrossoverPoint(ctx, a, b, crossover),
	}
}

type mPlace struct {
	m      int
	matrix f64.Matrix
}
