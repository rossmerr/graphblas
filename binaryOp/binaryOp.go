// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

// BinaryOp binaray operator
type BinaryOp interface {
	BinaryOp()
}

// Semigroup is a set together with an associative binary operation.
type Semigroup interface {
	BinaryOp
	Semigroup()
}
