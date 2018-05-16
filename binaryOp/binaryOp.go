// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

type BinaryOp interface {
	operator()
	binaryOp()
}

type Semigroup interface {
	BinaryOp
	semigroup()
}
