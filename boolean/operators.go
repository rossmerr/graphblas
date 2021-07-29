// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package boolean

import (
	"log"

	"context"

	GraphBLAS "github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/binaryop/boolop"
	boolUnaryOp "github.com/rossmerr/graphblas/unaryop/boolop"
)

// Apply modifies edge weights by the UnaryOperator
//  C ⊕= f(A)
func Apply(ctx context.Context, in Matrix, mask GraphBLAS.Mask, u boolUnaryOp.UnaryOpBool, matrix Matrix) {
	if mask == nil {
		mask = GraphBLAS.NewEmptyMask(matrix.Rows(), matrix.Columns())
	}

	if mask.Rows() != matrix.Rows() {
		log.Panicf("Can not apply mask found rows mismatch %+v, %+v", mask.Rows(), matrix.Rows())
	}

	if mask.Columns() != matrix.Columns() {
		log.Panicf("Can not apply mask found columns mismatch %+v, %+v", mask.Columns(), matrix.Columns())
	}

	if in == matrix {
		for iterator := in.Map(); iterator.HasNext(); {
			select {
			case <-ctx.Done():
				return
			default:
				iterator.Map(func(r, c int, value bool) bool {
					if mask.Element(r, c) {
						return u.Apply(value)
					}

					return value
				})
			}
		}

		return
	}

	for iterator := in.Enumerate(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			return
		default:
			r, c, value := iterator.Next()
			if !mask.Element(r, c) {
				matrix.Set(c, r, u.Apply(value))
			}
		}
	}
}

// Transpose swaps the rows and columns
//  C ⊕= Aᵀ
func Transpose(ctx context.Context, s Matrix, mask GraphBLAS.Mask, matrix Matrix) {
	if mask == nil {
		mask = GraphBLAS.NewEmptyMask(matrix.Rows(), matrix.Columns())
	}

	if mask.Rows() != matrix.Rows() {
		log.Panicf("Can not apply mask found rows mismatch %+v, %+v", mask.Rows(), matrix.Rows())
	}

	if mask.Columns() != matrix.Columns() {
		log.Panicf("Can not apply mask found columns mismatch %+v, %+v", mask.Columns(), matrix.Columns())
	}

	for iterator := s.Enumerate(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			return
		default:
			r, c, value := iterator.Next()

			if !mask.Element(r, c) {
				matrix.Set(c, r, value)
			}
		}
	}
}

// TransposeToCSR swaps the rows and columns and returns a compressed storage by rows (CSR) matrix
func TransposeToCSR(ctx context.Context, s Matrix) Matrix {
	matrix := NewCSRMatrix(s.Columns(), s.Rows())

	Transpose(ctx, s, nil, matrix)
	return matrix
}

// TransposeToCSC swaps the rows and columns and returns a compressed storage by columns (CSC) matrix
func TransposeToCSC(ctx context.Context, s Matrix) Matrix {
	matrix := NewCSCMatrix(s.Columns(), s.Rows())

	Transpose(ctx, s, nil, matrix)
	return matrix
}

// Equal the two matrices are equal
func Equal(ctx context.Context, s, m Matrix) bool {
	if s == nil && m == nil {
		return true
	}

	if s != nil && m == nil {
		return false
	}

	if m != nil && s == nil {
		return false
	}

	if s.Columns() != m.Columns() {
		return false
	}

	if s.Rows() != m.Rows() {
		return false
	}

	isSparseMatrixS := IsSparseMatrix(s)
	isSparseMatrixM := IsSparseMatrix(m)

	if (isSparseMatrixS && isSparseMatrixM) || (!isSparseMatrixS && !isSparseMatrixM) {
		if s.Values() != m.Values() {
			return false
		}
	}

	var iterator Enumerate
	var matrix Matrix

	// Check for a sparse matrix as we want to use its Enumerate operation
	// Because the use of the At operation on a sparse matrix is expensive
	if isSparseMatrixS {
		iterator = s.Enumerate()
		matrix = m
	} else {
		iterator = m.Enumerate()
		matrix = s
	}

	for iterator.HasNext() {
		select {
		case <-ctx.Done():
			return false
		default:
			sR, sC, sV := iterator.Next()
			mV := matrix.At(sR, sC)
			if sV != mV {
				return false
			}
		}
	}

	return true
}

// NotEqual the two matrices are not equal
func NotEqual(ctx context.Context, s, m Matrix) bool {
	return !Equal(ctx, s, m)
}

// ReduceMatrixToScalarWithMonoID perform's a reduction on the Matrix
// monoid used in the element-wise reduction operation
func ReduceMatrixToScalarWithMonoID(ctx context.Context, s Matrix, monoID boolop.MonoIDBool, mask GraphBLAS.Mask) bool {
	done := make(chan struct{})
	slice := make(chan bool)
	defer close(slice)
	defer close(done)

	out := monoID.Reduce(done, slice)

	if mask == nil {
		mask = GraphBLAS.NewEmptyMask(s.Rows(), s.Columns())
	}

	if mask.Rows() != s.Rows() {
		log.Panicf("Can not apply mask found rows mismatch %+v, %+v", mask.Rows(), s.Rows())
	}

	if mask.Columns() != s.Columns() {
		log.Panicf("Can not apply mask found columns mismatch %+v, %+v", mask.Columns(), s.Columns())
	}

	go func() {
		for iterator := s.Enumerate(); iterator.HasNext(); {
			select {
			case <-ctx.Done():
				return
			default:
				r, c, value := iterator.Next()
				if !mask.Element(r, c) {
					slice <- value
				}
			}
		}
		done <- struct{}{}
	}()

	return <-out

}

// AssignConstantVector the contents of a subset of a vector
// func AssignConstantVector(w, mask Vector, val float64, nindices int) {

// 	// if u.Length() != nindices {
// 	// 	log.Panicf("The number of values in indices array. Must be equal to Length of u %+v", u.Length())
// 	// }

// 	for iterator := mask.Enumerate(); iterator.HasNext(); {
// 		r, _, _ := iterator.Next()
// 		if r < nindices {
// 			w.SetVec(r, val)
// 		} else {
// 			break
// 		}
// 	}
// }
