// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS_test

import (
	"reflect"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func TestSpan(t *testing.T) {
	type args struct {
		set GraphBLAS.VectorSet
	}
	tests := []struct {
		name string
		args args
		want GraphBLAS.Matrix
	}{
		{
			name: "Span",
			want: func() GraphBLAS.Matrix {
				matrix := GraphBLAS.NewDenseVector(2)
				matrix.SetVec(0, 9)
				matrix.SetVec(1, 23)
				return matrix
			}(),
			args: args{
				set: func() GraphBLAS.VectorSet {
					results := GraphBLAS.VectorSet{}
					results = append(results, &GraphBLAS.VectorConstant{
						Value: 1,
						Vector: func() GraphBLAS.Vector {
							vector := GraphBLAS.NewDenseVector(2)
							vector.SetVec(0, 1)
							vector.SetVec(1, 3)
							return vector
						}(),
					})

					results = append(results, &GraphBLAS.VectorConstant{
						Value: 4,
						Vector: func() GraphBLAS.Vector {
							vector := GraphBLAS.NewDenseVector(2)
							vector.SetVec(0, 2)
							vector.SetVec(1, 5)
							return vector
						}(),
					})
					return results
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GraphBLAS.Span(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Span() = %v, want %v", got, tt.want)
			}
		})
	}
}
