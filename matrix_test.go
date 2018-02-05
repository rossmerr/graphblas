// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS_test

import (
	"math"
	"math/rand"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestMatrix_ColumnsAt(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := GraphBLAS.NewDenseVector(2)
	want.SetVec(0, 4)
	want.SetVec(1, 1)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got, _ := tt.s.ColumnsAt(0); !got.Equal(want) {
				t.Errorf("%+v ColumnsAt = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_RowAt(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := GraphBLAS.NewDenseVector(2)
	want.SetVec(0, 4)
	want.SetVec(1, 0)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got, _ := tt.s.RowsAt(0); !got.Equal(want) {
				t.Errorf("%+v RowsAt = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Scalar(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 4)
		m.Set(0, 1, 0)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, 8)
	want.Set(0, 1, 0)
	want.Set(1, 0, 2)
	want.Set(1, 1, -18)

	tests := []struct {
		name  string
		s     GraphBLAS.Matrix
		alpha float64
	}{
		{
			name:  "DenseMatrix",
			s:     GraphBLAS.NewDenseMatrix(2, 2),
			alpha: 2,
		},
		{
			name:  "CSCMatrix",
			s:     GraphBLAS.NewCSCMatrix(2, 2),
			alpha: 2,
		},
		{
			name:  "CSRMatrix",
			s:     GraphBLAS.NewCSRMatrix(2, 2),
			alpha: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Scalar(tt.alpha); !got.Equal(want) {
				t.Errorf("%+v Scalar = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Negative(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 2)
		m.Set(0, 1, -4)
		m.Set(1, 0, 7)
		m.Set(1, 1, 10)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, -2)
	want.Set(0, 1, 4)
	want.Set(1, 0, -7)
	want.Set(1, 1, -10)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Negative(); !got.Equal(want) {
				t.Errorf("%+v Negative = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Transpose(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 6)
		m.Set(0, 1, 4)
		m.Set(0, 2, 24)
		m.Set(1, 0, 1)
		m.Set(1, 1, -9)
		m.Set(1, 2, 8)
	}

	want := GraphBLAS.NewDenseMatrix(3, 2)
	want.Set(0, 0, 6)
	want.Set(0, 1, 1)
	want.Set(1, 0, 4)
	want.Set(1, 1, -9)
	want.Set(2, 0, 24)
	want.Set(2, 1, 8)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Transpose(); !got.Equal(want) {
				t.Errorf("%+v Transpose = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Equal(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 3)
	want.Set(0, 0, 1)
	want.Set(0, 1, 2)
	want.Set(0, 2, 3)
	want.Set(1, 0, 4)
	want.Set(1, 1, 5)
	want.Set(1, 2, 6)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if !tt.s.Equal(want) {
				t.Errorf("%+v Equal = %+v, want %+v", tt.name, tt.s, want)
			}
		})
	}
}

func TestMatrix_NotEqual(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 3)
	want.Set(0, 0, 2)
	want.Set(0, 1, 3)
	want.Set(0, 2, 4)
	want.Set(1, 0, 5)
	want.Set(1, 1, 6)
	want.Set(1, 2, 7)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if !tt.s.NotEqual(want) {
				t.Errorf("%+v NotEqual = %+v, want %+v", tt.name, tt.s, want)
			}
		})
	}
}

func TestMatrix_Copy(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 3)
	want.Set(0, 0, 1)
	want.Set(0, 1, 2)
	want.Set(0, 2, 3)
	want.Set(1, 0, 4)
	want.Set(1, 1, 5)
	want.Set(1, 2, 6)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got := tt.s.Copy(); !got.Equal(want) {
				t.Errorf("%+v Copy = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}
func TestMatrix_Multiply(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 1)
		m.Set(0, 1, 2)
		m.Set(0, 2, 3)
		m.Set(1, 0, 4)
		m.Set(1, 1, 5)
		m.Set(1, 2, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, 58)
	want.Set(0, 1, 64)
	want.Set(1, 0, 139)
	want.Set(1, 1, 154)

	matrix := GraphBLAS.NewDenseMatrix(3, 2)
	matrix.Set(0, 0, 7)
	matrix.Set(0, 1, 8)
	matrix.Set(1, 0, 9)
	matrix.Set(1, 1, 10)
	matrix.Set(2, 0, 11)
	matrix.Set(2, 1, 12)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 3),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 3),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got, _ := tt.s.Multiply(matrix); !got.Equal(want) {
				t.Errorf("%+v Multiply = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Add(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 3)
		m.Set(0, 1, 8)
		m.Set(1, 0, 4)
		m.Set(1, 1, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, 7)
	want.Set(0, 1, 8)
	want.Set(1, 0, 5)
	want.Set(1, 1, -3)

	matrix := GraphBLAS.NewDenseMatrix(2, 2)
	matrix.Set(0, 0, 4)
	matrix.Set(0, 1, 0)
	matrix.Set(1, 0, 1)
	matrix.Set(1, 1, -9)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got, _ := tt.s.Add(matrix); !got.Equal(want) {
				t.Errorf("%+v Add = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

func TestMatrix_Subtract(t *testing.T) {

	setup := func(m GraphBLAS.Matrix) {
		m.Set(0, 0, 3)
		m.Set(0, 1, 8)
		m.Set(1, 0, 4)
		m.Set(1, 1, 6)
	}

	want := GraphBLAS.NewDenseMatrix(2, 2)
	want.Set(0, 0, -1)
	want.Set(0, 1, 8)
	want.Set(1, 0, 3)
	want.Set(1, 1, 15)

	matrix := GraphBLAS.NewDenseMatrix(2, 2)
	matrix.Set(0, 0, 4)
	matrix.Set(0, 1, 0)
	matrix.Set(1, 0, 1)
	matrix.Set(1, 1, -9)

	tests := []struct {
		name string
		s    GraphBLAS.Matrix
	}{
		{
			name: "DenseMatrix",
			s:    GraphBLAS.NewDenseMatrix(2, 2),
		},
		{
			name: "CSCMatrix",
			s:    GraphBLAS.NewCSCMatrix(2, 2),
		},
		{
			name: "CSRMatrix",
			s:    GraphBLAS.NewCSRMatrix(2, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(tt.s)
			if got, _ := tt.s.Subtract(matrix); !got.Equal(want) {
				t.Errorf("%+v Subtract = %+v, want %+v", tt.name, got, want)
			}
		})
	}
}

// ################################################################################################

func BenchmarkMatrix(b *testing.B) {
	for _, fn := range benchmarks {
		fn.fn(b)
	}
}

var benchmarks = []struct {
	name string
	fn   func(*testing.B)
}{

	{
		name: "iteration_pi_sum",
		fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if math.Abs(pisum()-1.644834071848065) >= 1e-6 {
					b.Fatal("pi_sum out of range")
				}
			}
		},
	},
	// {
	// 	name: "matrix_statistics",
	// 	fn: func(b *testing.B) {
	// 		for i := 0; i < b.N; i++ {
	// 			c1, c2 := randmatstat(1000)
	// 			assert(b, 0.5 < c1)
	// 			assert(b, c1 < 1.0)
	// 			assert(b, 0.5 < c2)
	// 			assert(b, c2 < 1.0)
	// 		}
	// 	},
	// },

	{
		name: "matrix_multiply",
		fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := randmatmul(1000)
				v, _ := c.At(0, 0)
				if !(v >= 0) {
					b.Fatal("assert failed")
				}
			}
		},
	},
}

func pisum() float64 {
	var sum float64
	for i := 0; i < 500; i++ {
		sum = 0.0
		for k := 1.0; k <= 10000; k += 1 {
			sum += 1.0 / (k * k)
		}
	}
	return sum
}

// func randmatstat(t int) (float64, float64) {
// 	n := 5
// 	v := make([]float64, t)
// 	w := make([]float64, t)
// 	ad := make([][]float64, n)
// 	bd := make([][]float64, n)
// 	cd := make([][]float64, n)
// 	dd := make([][]float64, n)
// 	P := GraphBLAS.NewDenseMatrix(n, 4*n)
// 	Q := GraphBLAS.NewDenseMatrix(2*n, 2*n)
// 	pTmp := GraphBLAS.NewDenseMatrix(4*n, 4*n)
// 	qTmp := GraphBLAS.NewDenseMatrix(2*n, 2*n)
// 	for i := 0; i < t; i++ {
// 		for r := range ad {
// 			ad[r] = make([]float64, n)
// 			bd[r] = make([]float64, n)
// 			cd[r] = make([]float64, n)
// 			dd[r] = make([]float64, n)
// 			for c := range ad[0] {
// 				ad[r][c] = rand.NormFloat64()
// 				bd[r][c] = rand.NormFloat64()
// 				cd[r][c] = rand.NormFloat64()
// 				dd[r][c] = rand.NormFloat64()
// 			}

// 		}
// 		a := GraphBLAS.NewCSCMatrixFromArray(n, n, ad)
// 		b := GraphBLAS.NewCSCMatrixFromArray(n, n, bd)
// 		c := GraphBLAS.NewCSCMatrixFromArray(n, n, cd)
// 		d := GraphBLAS.NewCSCMatrixFromArray(n, n, dd)
// 		// P.Copy(a)
// 		// P.View(0, n, n, n).(*mat64.Dense).Copy(b)
// 		// P.View(0, 2*n, n, n).(*mat64.Dense).Copy(c)
// 		// P.View(0, 3*n, n, n).(*mat64.Dense).Copy(d)

// 		// Q.Copy(a)
// 		// Q.View(0, n, n, n).(*mat64.Dense).Copy(b)
// 		// Q.View(n, 0, n, n).(*mat64.Dense).Copy(c)
// 		// Q.View(n, n, n, n).(*mat64.Dense).Copy(d)

// 		// pTmp.Mul(P.T(), P)
// 		// pTmp.Pow(pTmp, 4)

// 		// qTmp.Mul(Q.T(), Q)
// 		// qTmp.Pow(qTmp, 4)

// 		// v[i] = mat64.Trace(pTmp)
// 		// w[i] = mat64.Trace(qTmp)
// 	}
// 	mv, stdv := stat.MeanStdDev(v, nil)
// 	mw, stdw := stat.MeanStdDev(v, nil)
// 	return stdv / mv, stdw / mw
// }

func randmatmul(n int) GraphBLAS.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := GraphBLAS.NewDenseMatrixFromArray(n, n, aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := GraphBLAS.NewDenseMatrixFromArray(n, n, bData)

	c, _ := a.Multiply(b)
	return c
}
