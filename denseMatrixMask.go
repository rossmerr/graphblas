// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

func (s *DenseMatrix) Element(r, c int) bool {
	s.RLock()
	defer s.RUnlock()

	return s.element(r, c)
}

func (s *DenseMatrix) element(r, c int) bool {
	return s.at(r, c) > 0
}

// EnumerateMask iterates through all non-zero elements, order is not guaranteed
func (s *DenseMatrix) EnumerateMask() EnumerateMask {
	return s.enumerateMask()
}

func (s *DenseMatrix) enumerateMask() *denseMatrixMask {
	i := &denseMatrixMask{
		matrix: s,
		size:   s.Values(),
		last:   0,
		c:      0,
		r:      0,
	}
	return i
}

type denseMatrixMask struct {
	matrix *DenseMatrix
	size   int
	last   int
	c      int
	r      int
	cOld   int
	rOld   int
}

// HasNext checks the iterator has any more values
func (s *denseMatrixMask) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

func (s *denseMatrixMask) next() {
	if s.c == s.matrix.Columns() {
		s.c = 0
		s.r++
	}
	s.cOld = s.c
	s.c++
	s.last++
}

// Next moves the iterator and returns the row, column and value
func (s *denseMatrixMask) Next() (int, int, bool) {
	s.next()

	s.matrix.RLock()
	defer s.matrix.RUnlock()

	return s.r, s.cOld, s.matrix.element(s.r, s.cOld)
}
