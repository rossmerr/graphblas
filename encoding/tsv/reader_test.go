// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package tsv_test

import (
	"reflect"
	"strings"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
	"github.com/RossMerr/Caudex.GraphBLAS/container/triples"
	"github.com/RossMerr/Caudex.GraphBLAS/encoding/tsv"
)

func TestTSV_ReadToMatrix(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want GraphBLAS.Matrix
	}{
		{
			name: "Tab-Separated Values Read Raw",
			in: `"1"	"1"	10
3	3	8
"2"	"2"	"3"`,
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
			r := tsv.NewReader(strings.NewReader(tt.in))

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

func TestTSV_ReadToTriples(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []*triples.Triple
	}{
		{
			name: "Tab-Separated Values Read Raw",
			in: `"1"	"1"	10
3	3	8
"2"	"2"	"3"`,
			want: func() []*triples.Triple {
				triples := []*triples.Triple{
					&triples.Triple{
						Row:    "1",
						Column: "1",
						Value:  float64(10),
					},
					&triples.Triple{
						Row:    "3",
						Column: "3",
						Value:  float64(8),
					},
					&triples.Triple{
						Row:    "2",
						Column: "2",
						Value:  float64(3),
					},
				}

				return triples
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tsv.NewReader(strings.NewReader(tt.in))
			if got, err := r.ReadToTriples(); err == nil {
				if len(tt.want) == len(got) {
					for i := 0; i < len(tt.want); i++ {
						if !reflect.DeepEqual(tt.want[i], got[i]) {
							t.Errorf("%+v ReadToTriples = got %+v, want %+v", tt.name, got[i], tt.want[i])
						}
					}
				} else {
					t.Errorf("%+v ReadToTriples length miss match = got %+v, want %+v", tt.name, len(got), len(tt.want))
				}
			} else {
				t.Errorf("%+v ReadToTriples error = %+v", tt.name, err)
			}
		})
	}
}
