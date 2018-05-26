// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

type BinaryOp interface {
	Operator()
	BinaryOp()
}

type Semigroup interface {
	BinaryOp
	Semigroup()
}
