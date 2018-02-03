// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// Vector interface
type Vector interface {
	Matrix
	AtVec(i int) (float64, error)
	SetVec(i int, value float64) error
	Length() int
}
