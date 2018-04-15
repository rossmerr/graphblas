package tsv_test

import (
	"strings"
	"testing"

	"github.com/RossMerr/Caudex.GraphBLAS/encoding/tsv"
)

func TestTSV_Reader(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want [][]float64
	}{
		{
			name: "TSV Read Raw",
			in: `"1"	"1"	10
3	3	8
"2"	"2"	"3"`,
			want: func() [][]float64 {
				matrix := make([][]float64, 3)
				matrix[0] = make([]float64, 3)
				matrix[1] = make([]float64, 3)
				matrix[2] = make([]float64, 3)
				matrix[0][0] = 10
				matrix[1][1] = 3
				matrix[2][2] = 8
				return matrix
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tsv.NewReader(strings.NewReader(tt.in))

			if got, err := r.ReadAll(); err == nil {
				for r, _ := range tt.want {
					for c, _ := range tt.want[r] {
						if got[r][c] != tt.want[r][c] {
							t.Errorf("%+v ReadAll = got %+v, want %+v", tt.name, got[r][c], tt.want[r][c])
						}
					}
				}
			} else {
				t.Errorf("%+v ReadAll error = %+v", tt.name, err)
			}
		})
	}
}
