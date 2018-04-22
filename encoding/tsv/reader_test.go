// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package tsv_test

import (
	"strings"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
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
