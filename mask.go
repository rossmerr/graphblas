// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// Mask is used to control how computed values are stored in the output from a method
type Mask interface {
	// Columns the number of columns of the mask
	Columns() int

	// Rows the number of rows of the mask
	Rows() int

	// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
	Element(r, c int) bool
}

// EmptyMask is a mask with no elements but will always returns false
type EmptyMask struct {
	r int
	c int
}

// NewEmptyMask returns a EmptyMask
func NewEmptyMask(r, c int) *EmptyMask {
	return &EmptyMask{r: r, c: c}
}

// Columns the number of columns of the mask
func (s *EmptyMask) Columns() int {
	return s.c
}

// Rows the number of rows of the mask
func (s *EmptyMask) Rows() int {
	return s.r
}

// Element of the mask for each tuple that exists in the matrix for which the value of the tuple cast to Boolean is true
func (s *EmptyMask) Element(r, c int) bool {
	return false
}
