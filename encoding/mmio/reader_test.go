// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package mmio_test

import (
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/container/triple"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
	"github.com/RossMerr/Caudex.GraphBLAS/encoding/mmio"
)

func TestMMIO_ReadToMatrix(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want GraphBLAS.Matrix
	}{
		{
			name: "Matrix Market sparse ",
			in: `%%MatrixMarket matrix coordinate real general
			3 3 3
			1 1 10
3 3 8
2 2 3`,
			want: func() GraphBLAS.Matrix {
				matrix := make([][]float64, 3)
				matrix[0] = make([]float64, 3)
				matrix[1] = make([]float64, 3)
				matrix[2] = make([]float64, 3)
				matrix[0][0] = 10
				matrix[1][1] = 3
				matrix[2][2] = 8
				return GraphBLAS.NewDenseMatrixFromArray(matrix)
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mmio.NewReader(strings.NewReader(tt.in))

			if got, err := r.ReadToMatrix(); err == nil {

				for r := 0; r < tt.want.Rows(); r++ {
					for c := 0; c < tt.want.Columns(); c++ {
						if got.At(r, c) != tt.want.At(r, c) {
							t.Errorf("%+v ReadToMatrix = got %+v, want %+v", tt.name, got.At(r, c), tt.want.At(r, c))
						}
					}
				}

			} else {
				t.Errorf("%+v ReadToMatrix error = %+v", tt.name, err)
			}
		})
	}
}

func TestMMIO_ReadToTripleStore(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want *triple.Store
	}{
		{
			name: "Matrix Market sparse",
			in: `%%MatrixMarket matrix coordinate real general
			3 3 3
			1 1 10
3 3 8
2 2 3`,
			want: func() *triple.Store {
				store := &triple.Store{}
				return store
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := mmio.NewReader(strings.NewReader(tt.in))
			if got, err := r.ReadToTripleStore(); err == nil {
				for i := 0; i < len(tt.want.Triples); i++ {
					if tt.want.Triples[i] != got.Triples[i] {
						t.Errorf("%+v ReadToTripleStore = got %+v, want %+v", tt.name, got.Triples[i], tt.want.Triples[i])
					}
				}
			} else {
				t.Errorf("%+v ReadToTripleStore error = %+v", tt.name, err)
			}
		})
	}
}
