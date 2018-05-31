// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"log"

	"github.com/RossMerr/Caudex.GraphBLAS/binaryOp/float64Op"
	float64UnaryOp "github.com/RossMerr/Caudex.GraphBLAS/unaryOp/float64Op"
	"golang.org/x/net/context"
)

func multiply(ctx context.Context, s, m, matrix Matrix) {
	if m.Rows() != s.Columns() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", m.Rows(), s.Columns())
	}

	for r := 0; r < s.Rows(); r++ {
		rows := s.RowsAt(r)

		for c := 0; c < m.Columns(); c++ {
			column := m.ColumnsAt(c)

			sum := 0.0
			for l := 0; l < rows.Length(); l++ {
				select {
				case <-ctx.Done():
					break
				default:
					vC := column.AtVec(l)
					vR := rows.AtVec(l)
					sum += vR * vC
				}
			}

			matrix.Set(r, c, sum)
		}

	}
}

// MatrixMatrixMultiply multiplies a matrix by another matrix
// mxm
func MatrixMatrixMultiply(ctx context.Context, s, m, matrix Matrix) {
	multiply(ctx, s, m, matrix)
}

// VectorMatrixMultiply multiplies a vector by a matrix
// vxm
func VectorMatrixMultiply(ctx context.Context, s Vector, m Matrix, vector Vector) {
	multiply(ctx, m, s, vector)
}

// MatrixVectorMultiply multiplies a matrix by a vector
// mxv
func MatrixVectorMultiply(ctx context.Context, s Matrix, m Vector, vector Vector) {
	multiply(ctx, s, m, vector)
}

func elementWiseMultiply(ctx context.Context, s, m, matrix Matrix) {
	var iterator Enumerate
	var source Matrix

	target := func(r, c int, value float64) {
		v := source.At(r, c)
		if value == v {
			matrix.Set(r, c, value)
		}
	}

	// when the matrix is the same object of s or m we use the setSource func to self update
	setSource := func(r, c int, value float64) {
		source.Update(r, c, func(v float64) float64 {
			if value != v {
				return defaultFloat64
			}
			return v
		})
	}

	if m == matrix {
		target = setSource
		iterator = s.Enumerate()
		source = m
	} else if s == matrix {
		target = setSource
		iterator = m.Enumerate()
		source = s
	} else if IsSparseMatrix(s) {
		iterator = s.Enumerate()
		source = m
	} else {
		iterator = m.Enumerate()
		source = s
	}

	for iterator.HasNext() {
		select {
		case <-ctx.Done():
			break
		default:
			r, c, value := iterator.Next()
			target(r, c, value)
		}
	}
}

// ElementWiseMatrixMultiply Element-wise multiplication on a matrix
// eWiseMult
func ElementWiseMatrixMultiply(ctx context.Context, s, m, matrix Matrix) {
	if m.Rows() != s.Columns() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", m.Rows(), s.Columns())
	}

	elementWiseMultiply(ctx, s, m, matrix)
}

// ElementWiseVectorMultiply Element-wise multiplication on a vector
// eWiseMult
func ElementWiseVectorMultiply(ctx context.Context, s, m, vector Vector) {
	if m.Rows() != s.Rows() {
		log.Panicf("Can not multiply vectors found length miss match %+v, %+v", m.Rows(), s.Rows())
	}

	elementWiseMultiply(ctx, s, m, vector)
}

// Add addition of a matrix by another matrix
func Add(s, m, matrix Matrix) {
	if s.Columns() != m.Columns() {
		log.Panicf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	var iterator Enumerate
	var source Matrix
	if IsSparseMatrix(s) {
		iterator = s.Enumerate()
		source = m
	} else {
		iterator = m.Enumerate()
		source = s
	}

	for iterator.HasNext() {
		r, c, v := iterator.Next()
		value := source.At(r, c)
		matrix.Set(r, c, value+v)
	}
}

func elementWiseAdd(ctx context.Context, s, m, matrix Matrix) {

	if s != matrix {
		for iterator := s.Enumerate(); iterator.HasNext(); {
			select {
			case <-ctx.Done():
				break
			default:
				r, c, value := iterator.Next()
				if value != defaultFloat64 {
					matrix.Set(r, c, value)
				}
			}
		}
	}

	if m != matrix {
		for iterator := m.Enumerate(); iterator.HasNext(); {
			select {
			case <-ctx.Done():
				break
			default:
				r, c, value := iterator.Next()
				if value != defaultFloat64 {
					matrix.Set(r, c, value)
				}
			}
		}
	}
}

// ElementWiseMatrixAdd Element-wise addition on a matrix
// eWiseMult
func ElementWiseMatrixAdd(ctx context.Context, s, m, matrix Matrix) {
	if m.Rows() != s.Columns() {
		log.Panicf("Can not multiply matrices found length miss match %+v, %+v", m.Rows(), s.Columns())
	}

	elementWiseAdd(ctx, s, m, matrix)
}

// ElementWiseVectorAdd Element-wise addition on a vector
// eWiseMult
func ElementWiseVectorAdd(ctx context.Context, s, m, vector Vector) {
	if m.Rows() != s.Rows() {
		log.Panicf("Can not multiply vectors found length miss match %+v, %+v", m.Rows(), s.Rows())
	}

	elementWiseAdd(ctx, s, m, vector)
}

// Subtract subtracts one matrix from another matrix
func Subtract(ctx context.Context, s, m, matrix Matrix) {
	if s.Columns() != m.Columns() {
		log.Panicf("Column miss match %+v, %+v", s.Columns(), m.Columns())
	}

	if s.Rows() != m.Rows() {
		log.Panicf("Row miss match %+v, %+v", s.Rows(), m.Rows())
	}

	for iterator := s.Enumerate(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			break
		default:
			r, c, value := iterator.Next()
			matrix.Update(r, c, func(v float64) float64 {
				return value - v
			})
		}
	}
}

// Apply modifies edge weights by the UnaryOperator
// C ⊕= f(A)
func Apply(ctx context.Context, in, out Matrix, u float64UnaryOp.UnaryOpFloat64) {
	if in == out {
		for iterator := in.Map(); iterator.HasNext(); {
			select {
			case <-ctx.Done():
				break
			default:
				iterator.Map(func(r, c int, value float64) float64 {
					return u.Apply(value)
				})
			}
		}

		return
	}

	for iterator := in.Enumerate(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			break
		default:
			r, c, value := iterator.Next()
			out.Set(c, r, u.Apply(value))
		}
	}
}

// Negative the negative of a matrix
func Negative(ctx context.Context, s, matrix Matrix) {
	for iterator := matrix.Map(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			break
		default:
			iterator.Map(func(r, c int, v float64) float64 {
				return -v
			})
		}
	}
}

// Transpose swaps the rows and columns
// C ⊕= Aᵀ
func Transpose(ctx context.Context, s, m Matrix) {
	for iterator := s.Enumerate(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			break
		default:
			r, c, value := iterator.Next()
			m.Set(c, r, value)
		}
	}
}

// TransposeToCSR swaps the rows and columns and returns a compressed storage by rows (CSR) matrix
func TransposeToCSR(ctx context.Context, s Matrix) Matrix {
	matrix := NewCSRMatrix(s.Columns(), s.Rows())

	Transpose(ctx, s, matrix)
	return matrix
}

// TransposeToCSC swaps the rows and columns and returns a compressed storage by columns (CSC) matrix
func TransposeToCSC(ctx context.Context, s Matrix) Matrix {
	matrix := NewCSCMatrix(s.Columns(), s.Rows())

	Transpose(ctx, s, matrix)
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
			break
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

// Scalar multiplication of a matrix by alpha
func Scalar(ctx context.Context, s Matrix, alpha float64) Matrix {
	matrix := s.Copy()
	for iterator := matrix.Map(); iterator.HasNext(); {
		select {
		case <-ctx.Done():
			break
		default:
			iterator.Map(func(r, c int, v float64) float64 {
				return alpha * v
			})
		}
	}
	return matrix
}

// ReduceMatrixToVector perform's a reduction on the Matrix
func ReduceMatrixToVector(ctx context.Context, s Matrix, properties ...interface{}) Vector {
	return ReduceMatrixToVectorWithMonoID(ctx, s, defaultMonoIDMaximum)
}

// ReduceMatrixToVectorWithMonoID perform's a reduction on the Matrix
// monoid used in the element-wise reduction operation
func ReduceMatrixToVectorWithMonoID(ctx context.Context, s Matrix, monoID float64Op.MonoIDFloat64) Vector {

	vector := NewDenseVector(s.Columns())
	for c := 0; c < s.Columns(); c++ {
		v := s.ColumnsAt(c)
		scaler := ReduceVectorToScalarWithMonoID(ctx, v, monoID)
		vector.SetVec(c, scaler)
	}

	return vector
}

// ReduceVectorToScalar perform's a reduction on the Matrix
func ReduceVectorToScalar(ctx context.Context, s Vector) float64 {
	return ReduceMatrixToScalar(ctx, s)
}

// ReduceVectorToScalarWithMonoID perform's a reduction on the Matrix
// monoid used in the element-wise reduction operation
func ReduceVectorToScalarWithMonoID(ctx context.Context, s Vector, monoID float64Op.MonoIDFloat64) float64 {
	return ReduceMatrixToScalarWithMonoID(ctx, s, monoID)
}

// ReduceMatrixToScalar perform's a reduction on the Matrix
func ReduceMatrixToScalar(ctx context.Context, s Matrix) float64 {
	return ReduceMatrixToScalarWithMonoID(ctx, s, defaultMonoIDAddition)
}

// ReduceMatrixToScalarWithMonoID perform's a reduction on the Matrix
// monoid used in the element-wise reduction operation
func ReduceMatrixToScalarWithMonoID(ctx context.Context, s Matrix, monoID float64Op.MonoIDFloat64) float64 {
	done := make(chan struct{})
	slice := make(chan float64)
	defer close(slice)
	defer close(done)

	out := monoID.Reduce(done, slice)

	go func() {
		for iterator := s.Enumerate(); iterator.HasNext(); {
			select {
			case <-ctx.Done():
				break
			default:
				_, _, value := iterator.Next()
				slice <- value
			}
		}
		done <- struct{}{}
	}()

	return <-out

}
