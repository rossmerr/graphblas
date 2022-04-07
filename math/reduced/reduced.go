// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package reduced

import (
	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/constraints"
)

// Reduced row echelon form of matrix (Gauss-Jordan elimination)
// rref
func Reduced[T constraints.Number](s graphblas.Matrix[T]) graphblas.Matrix[T] {
	m := s.Copy()
	lead := 0
	rowCount := m.Rows()
	columnCount := m.Columns()

	for r := 0; r < rowCount; r++ {
		if lead >= columnCount {
			return m
		}
		i := r
		for m.At(i, lead) == 0 {
			i++
			if rowCount == i {
				i = r
				lead++
				if columnCount == lead {
					return m
				}
			}
		}

		if i != r {
			v1 := m.RowsAtToArray(i)
			v2 := m.RowsAtToArray(r)

			for c := 0; c < len(v1); c++ {
				m.Set(r, c, v1[c])
			}

			for c := 0; c < len(v2); c++ {
				m.Set(i, c, v2[c])
			}
		}

		f := 1 / m.At(r, lead)

		vector := m.RowsAtToArray(r)
		for c := 0; c < len(vector); c++ {
			value := vector[c]
			value *= f
			m.Set(r, c, value)
		}

		for i = 0; i < rowCount; i++ {
			if i != r {
				f = m.At(i, lead)
				vector := m.RowsAtToArray(r)
				for c := 0; c < len(vector); c++ {
					value := vector[c]
					m.Update(i, c, func(v T) T {
						v -= value * f
						return v
					})

				}
			}
		}
		lead++
	}

	return m
}
