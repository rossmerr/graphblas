// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package graphblas

import (
	"github.com/rossmerr/graphblas/constraints"
)

type vector[T constraints.Type] interface {
	// AtVec returns the value of a vector element at i-th
	AtVec(i int) T

	// SetVec sets the value at i-th of the vector
	SetVec(i int, value T)

	// Length of the vector
	Length() int
}

type VectorLogial[T constraints.Type] interface {
	MatrixLogical[T]

	vector[T]
}

type Vector[T constraints.Number] interface {
	Matrix[T]

	vector[T]
}

type VectorRune interface {
	MatrixLogical[rune]

	vector[rune]

	Compare(v VectorRune) int
}
