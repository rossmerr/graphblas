// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas

import (
	"log"
	"strings"

	"context"

	"github.com/rossmerr/graphblas/binaryop"
	"github.com/rossmerr/graphblas/constraints"
	"github.com/rossmerr/graphblas/unaryop"
)

func multiply[T constraints.Number](ctx context.Context, s, m Matrix[T], mask Mask, matrix Matrix[T]) {
	if m.Rows() != s.Columns() {
		log.Panicf("Can not multiply matrices found length mismatch %+v, %+v", m.Rows(), s.Columns())
	}

	if mask == nil {
		mask = NewEmptyMask(matrix.Rows(), matrix.Columns())
	}

	if mask.Rows() != matrix.Rows() {
		log.Panicf("Can not apply mask found rows mismatch %+v, %+v", mask.Rows(), matrix.Rows())
	}

	if mask.Columns() != matrix.Columns() {
		log.Panicf("Can not apply mask found columns mismatch %+v, %+v", mask.Columns(), matrix.Columns())
	}

	for r := 0; r < s.Rows(); r++ {
		rows := s.RowsAt(r)

		for c := 0; c < m.Columns(); c++ {
			column := m.ColumnsAt(c)

			sum := Default[T]()
			for l := 0; l < rows.Length(); l++ {
				select {
				case <-ctx.Done():
					return
				default:
					vC := column.AtVec(l)
					vR := rows.AtVec(l)
					sum += vR * vC
				}
			}

			if !mask.Element(r, c) {
				matrix.Set(r, c, sum)
			}
		}

	}
}

// MatrixMatrixMultiply multiplies a matrix by another matrix
//
// mxm
func MatrixMatrixMultiply[T constraints.Number](ctx context.Context, s, m Matrix[T], mask Mask, matrix Matrix[T]) {
	multiply(ctx, s, m, mask, matrix)
}

// VectorMatrixMultiply multiplies a vector by a matrix
//
// vxm
func VectorMatrixMultiply[T constraints.Number](ctx context.Context, s Vector[T], m Matrix[T], mask Mask, vector Vector[T]) {
	multiply[T](ctx, m, s, mask, vector)
}

// MatrixVectorMultiply multiplies a matrix by a vector
//
// mxv
func MatrixVectorMultiply[T constraints.Number](ctx context.Context, s Matrix[T], m Vector[T], mask Mask, vector Vector[T]) {
	multiply[T](ctx, s, m, mask, vector)
}

func elementWiseMultiply[T constraints.Number](ctx context.Context, s, m Matrix[T], mask Mask, matrix Matrix[T]) {
	var iterator Enumerate[T]
	var source Matrix[T]

	if mask == nil {
		mask = NewEmptyMask(matrix.Rows(), matrix.Columns())
	}

	if mask.Rows() != matrix.Rows() {
		log.Panicf("Can not apply mask found rows mismatch %+v, %+v", mask.Rows(), matrix.Rows())
	}

	if mask.Columns() != matrix.Columns() {
		log.Panicf("Can not apply mask found columns mismatch %+v, %+v", mask.Columns(), matrix.Columns())
	}

	target := func(r, c int, value T) {
		v := source.At(r, c)
		if value == v {
			if !mask.Element(r, c) {
				matrix.Set(r, c, value)
			}
		}
	}

	// when the matrix is the same object of s or m we use the setSource func to self update
	setSource := func(r, c int, value T) {
		if !mask.Element(r, c) {
			source.Update(r, c, func(v T) T {
				if value != v {
					return Default[T]()
				}
				return v
			})
		}
	}

	if m == matrix {
		target = setSource
		iterator = s.Enumerate()
		source = m
	} else if s == matrix {
		target = setSource
		iterator = m.Enumerate()
		source = s
	} else if IsSparseMatrix[T](s) {
		iterator = s.Enumerate()
		source = m
	} else {
		iterator = m.Enumerate()
		source = s
	}

	for iterator.HasNext() {
		select {
		case <-ctx.Done():
			return
		default:
			r, c, value := iterator.Next()
			target(r, c, value)
		}
	}
}

// ElementWiseMatrixMultiply Element-wise multiplication on a matrix
//
// eWiseMult
func ElementWiseMatrixMultiply[T constraints.Number](ctx context.Context, s, m Matrix[T], mask Mask, matrix Matrix[T]) {
	if m.Rows() != s.Columns() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", m.Rows(), s.Columns())
	}

	elementWiseMultiply(ctx, s, m, mask, matrix)
}

// ElementWiseVectorMultiply Element-wise multiplication on a vector
//
// eWiseMult
func ElementWiseVectorMultiply[T constraints.Number](ctx context.Context, s, m Vector[T], mask Mask, vector Vector[T]) {
	if m.Rows() != s.Rows() {
		log.Panicf("Can not multiply vectors found length mismatch %+v, %+v", m.Rows(), s.Rows())
	}

	elementWiseMultiply[T](ctx, s, m, mask, vector)
}

// Add addition of a matrix by another matrix
func Add[T constraints.Number](ctx context.Context, s, m Matrix[T], mask Mask, matrix Matrix[T]) {
	if s.Columns() != m.Columns() {
		log.Panicf("Column mismatch %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row mismatch %+v, %+v", s.Rows(), m.Rows())
	}

	if mask == nil {
		mask = NewEmptyMask(matrix.Rows(), matrix.Columns())
	}

	if mask.Rows() != matrix.Rows() {
		log.Panicf("Can not apply mask found rows mismatch %+v, %+v", mask.Rows(), matrix.Rows())
	}

	if mask.Columns() != matrix.Columns() {
		log.Panicf("Can not apply mask found columns mismatch %+v, %+v", mask.Columns(), matrix.Columns())
	}

	var iterator Enumerate[T]
	var source Matrix[T]
	if IsSparseMatrix[T](s) {
		iterator = s.Enumerate()
		source = m
	} else {
		iterator = m.Enumerate()
		source = s
	}

	for iterator.HasNext() {
		select {
		case <-ctx.Done():
			return
		default:
			r, c, v := iterator.Next()
			value := source.At(r, c)
			if !mask.Element(r, c) {
				matrix.Set(r, c, value+v)
			}
		}
	}
}

func elementWiseAdd[T constraints.Number](ctx context.Context, s, m Matrix[T], mask Mask, matrix Matrix[T]) {
	if mask == nil {
		mask = NewEmptyMask(matrix.Rows(), matrix.Columns())
	}

	if mask.Rows() != matrix.Rows() {
		log.Panicf("Can not apply mask found rows mismatch %+v, %+v", mask.Rows(), matrix.Rows())
	}

	if mask.Columns() != matrix.Columns() {
		log.Panicf("Can not apply mask found columns mismatch %+v, %+v", mask.Columns(), matrix.Columns())
	}

	if s != matrix {
		for iterator := s.Enumerate(); iterator.HasNext(); {
			select {
			case <-ctx.Done():
				return
			default:
				r, c, value := iterator.Next()
				if value != Default[T]() {
					if !mask.Element(r, c) {
						matrix.Set(r, c, value)
					}
				}
			}
		}
	}

	if m != matrix {
		for iterator := m.Enumerate(); iterator.HasNext(); {
			select {
			case <-ctx.Done():
				return
			default:
				r, c, value := iterator.Next()
				if value != Default[T]() {
					if !mask.Element(r, c) {
						matrix.Set(r, c, value)
					}
				}
			}
		}
	}
}

// ElementWiseMatrixAdd Element-wise addition on a matrix
//
// eWiseMult
func ElementWiseMatrixAdd[T constraints.Number](ctx context.Context, s, m Matrix[T], mask Mask, matrix Matrix[T]) {
	if m.Rows() != s.Columns() {
		log.Panicf("Can not multiply matrices found length mismatch %+v, %+v", m.Rows(), s.Columns())
	}

	elementWiseAdd(ctx, s, m, mask, matrix)
}

// ElementWiseVectorAdd Element-wise addition on a vector
//
// eWiseMult
func ElementWiseVectorAdd[T constraints.Number](ctx context.Context, s, m Vector[T], mask Mask, vector Vector[T]) {
	if m.Rows() != s.Rows() {
		log.Panicf("Can not multiply vectors found length mismatch %+v, %+v", m.Rows(), s.Rows())
	}

	elementWiseAdd[T](ctx, s, m, mask, vector)
}

// Subtract subtracts one matrix from another matrix
func Subtract[T constraints.Number](ctx context.Context, s, m Matrix[T], mask Mask, matrix Matrix[T]) {
	if s.Columns() != m.Columns() {
		log.Panicf("Column mismatch %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row mismatch %+v, %+v", s.Rows(), m.Rows())
	}

	if mask == nil {
		mask = NewEmptyMask(matrix.Rows(), matrix.Columns())
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
				matrix.Update(r, c, func(v T) T {
					return value - v
				})
			}
		}
	}
}

// Apply modifies edge weights by the UnaryOperator
//
//	C ⊕= f(A)
func Apply[T constraints.Number](ctx context.Context, in Matrix[T], mask Mask, u unaryop.UnaryOp[T], matrix Matrix[T]) {
	if mask == nil {
		mask = NewEmptyMask(matrix.Rows(), matrix.Columns())
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
				iterator.Map(func(r, c int, value T) T {
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

// Negative the negative of a matrix
func Negative[T constraints.Number](ctx context.Context, s Matrix[T], mask Mask, matrix Matrix[T]) {
	if mask == nil {
		mask = NewEmptyMask(matrix.Rows(), matrix.Columns())
	}

	if mask.Rows() != matrix.Rows() {
		log.Panicf("Can not apply mask found rows mismatch %+v, %+v", mask.Rows(), matrix.Rows())
	}

	if mask.Columns() != matrix.Columns() {
		log.Panicf("Can not apply mask found columns mismatch %+v, %+v", mask.Columns(), matrix.Columns())
	}

	for iterator := matrix.Map(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			return
		default:
			iterator.Map(func(r, c int, value T) T {
				if !mask.Element(r, c) {
					return -value
				}

				return value
			})
		}
	}
}

// Transpose swaps the rows and columns
//
//	C ⊕= Aᵀ
func Transpose[T constraints.Type](ctx context.Context, s MatrixLogical[T], mask Mask, matrix MatrixLogical[T]) {
	if mask == nil {
		mask = NewEmptyMask(matrix.Rows(), matrix.Columns())
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
func TransposeToCSR[T constraints.Number](ctx context.Context, s Matrix[T]) Matrix[T] {
	matrix := NewCSRMatrix[T](s.Columns(), s.Rows())

	Transpose[T](ctx, s, nil, matrix)
	return matrix
}

// TransposeToCSC swaps the rows and columns and returns a compressed storage by columns (CSC) matrix
func TransposeToCSC[T constraints.Number](ctx context.Context, s Matrix[T]) Matrix[T] {
	matrix := NewCSCMatrix[T](s.Columns(), s.Rows())

	Transpose[T](ctx, s, nil, matrix)
	return matrix
}

// Compare returns an integer comparing two matrices lexicographically.
func Compare(ctx context.Context, s, m MatrixLogical[rune]) int {
	select {
	case <-ctx.Done():
		return 0
	default:
		if s.Equal(m) {
			return 0
		}
	}
	select {
	case <-ctx.Done():
		return 0
	default:
		if Less(ctx, s, m) {
			return -1
		}
	}
	return +1
}

func Less(ctx context.Context, s, m MatrixLogical[rune]) bool {
	return String(ctx, s) < String(ctx, m)
}

func Greater(ctx context.Context, s, m MatrixLogical[rune]) bool {
	return String(ctx, s) > String(ctx, m)
}

func String(ctx context.Context, s MatrixLogical[rune]) string {
	var b strings.Builder
	b.Grow(s.Size())
	enumator := s.Enumerate()

	for enumator.HasNext() {
		select {
		case <-ctx.Done():
			return b.String()
		default:
			_, _, r := enumator.Next()
			b.WriteRune(r)
		}
	}

	return b.String()
}

// Equal the two matrices are equal
func Equal[T constraints.Type](ctx context.Context, s, m MatrixLogical[T]) bool {
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

	var iterator Enumerate[T]
	var matrix MatrixLogical[T]

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
func NotEqual[T constraints.Type](ctx context.Context, s, m MatrixLogical[T]) bool {
	return !Equal(ctx, s, m)
}

// Scalar multiplication of a matrix by alpha
func Scalar[T constraints.Number](ctx context.Context, s Matrix[T], alpha T) Matrix[T] {
	matrix := s.Copy()
	for iterator := matrix.Map(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			break
		default:
			iterator.Map(func(r, c int, v T) T {
				return alpha * v
			})
		}
	}
	return matrix
}

// ReduceMatrixToVector perform's a reduction on the Matrix
func ReduceMatrixToVector[T constraints.Number](ctx context.Context, s Matrix[T]) Vector[T] {
	return ReduceMatrixToVectorWithMonoID(ctx, s, DefaultMonoIDMaximum[T](), nil)
}

// ReduceMatrixToVectorWithMonoID perform's a reduction on the Matrix
// monoid used in the element-wise reduction operation
func ReduceMatrixToVectorWithMonoID[T constraints.Number](ctx context.Context, s Matrix[T], monoID binaryop.MonoID[T], mask Mask) Vector[T] {

	vector := NewDenseVectorN[T](s.Columns())
	for c := 0; c < s.Columns(); c++ {
		v := s.ColumnsAt(c)
		scaler := ReduceVectorToScalarWithMonoID(ctx, v, monoID, mask)
		vector.SetVec(c, scaler)
	}

	return vector
}

// ReduceVectorToScalar perform's a reduction on the Matrix
func ReduceVectorToScalar[T constraints.Number](ctx context.Context, s VectorLogial[T], mask Mask) T {
	return ReduceMatrixToScalar[T](ctx, s, mask)
}

// ReduceVectorToScalarWithMonoID perform's a reduction on the Matrix
// monoid used in the element-wise reduction operation
func ReduceVectorToScalarWithMonoID[T constraints.Number](ctx context.Context, s VectorLogial[T], monoID binaryop.MonoID[T], mask Mask) T {
	return ReduceMatrixToScalarWithMonoID[T](ctx, s, monoID, mask)
}

// ReduceMatrixToScalar perform's a reduction on the Matrix
func ReduceMatrixToScalar[T constraints.Number](ctx context.Context, s MatrixLogical[T], mask Mask) T {
	return ReduceMatrixToScalarWithMonoID(ctx, s, DefaultMonoIDAddition[T](), mask)
}

// ReduceMatrixToScalarWithMonoID perform's a reduction on the Matrix
// monoid used in the element-wise reduction operation
func ReduceMatrixToScalarWithMonoID[T constraints.Number](ctx context.Context, s MatrixLogical[T], monoID binaryop.MonoID[T], mask Mask) T {
	done := make(chan struct{})
	slice := make(chan T)
	defer close(slice)
	defer close(done)

	out := monoID.Reduce(done, slice)

	if mask == nil {
		mask = NewEmptyMask(s.Rows(), s.Columns())
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
