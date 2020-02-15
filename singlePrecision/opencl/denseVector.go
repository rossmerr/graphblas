// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package opencl

import (
	"context"
	"fmt"

	"github.com/microo8/blackcl"
)

type Vector interface {
	Length() int
	//RowsAtToArray(r int, data VectorFloat32) <-chan error
	Copy(data VectorFloat32) <-chan error
}

// DenseVector a vector
type DenseVector struct {
	v *blackcl.Vector
}

// NewDenseVector returns a DenseVector
func NewDenseVector(ctx context.Context, length int) (*DenseVector, error) {
	deviceCtx, ok := ctx.(*deviceContext)
	if !ok {
		return nil, fmt.Errorf("No opencl device found in context")
	}
	v, err := deviceCtx.Device.NewVector(length)
	if err != nil {
		return nil, fmt.Errorf("could not allocate buffer")
	}

	return &DenseVector{v: v}, nil
}

// Columns the number of columns of the vector
func (s *DenseVector) Columns() int {
	return 1
}

// Rows the number of rows of the vector
func (s *DenseVector) Rows() int {
	return s.Length()
}

// Length of the vector
func (s *DenseVector) Length() int {
	return s.v.Length()
}

// // RowsAtToArray return the rows at r-th
// func (s *DenseVector) RowsAtToArray(r int, data VectorFloat32) <-chan error {
// 	if r < 0 || r >= s.Length() {
// 		log.Panicf("Row '%+v' is invalid", r)
// 	}

// 	return s.v.Copy(data)
// }

// Copy copies the vector
func (s *DenseVector) Copy(data VectorFloat32) <-chan error {
	return s.v.Copy(data)
}

// Date out of the vector
func (s *DenseVector) Data() (VectorFloat32, error) {
	return s.v.Data()
}
