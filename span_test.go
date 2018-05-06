// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

import (
	"reflect"
	"testing"
)

func TestSpan(t *testing.T) {
	type args struct {
		set []VectorSet
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "Span",
			want: func() Matrix {
				matrix := NewDenseVector(2)
				matrix.SetVec(0, 9)
				matrix.SetVec(1, 23)
				return matrix
			}(),
			args: args{
				set: func() []VectorSet {
					results := make([]VectorSet, 0)
					results = append(results, VectorSet{
						Value: 1,
						Vector: func() Vector {
							vector := NewDenseVector(2)
							vector.SetVec(0, 1)
							vector.SetVec(1, 3)
							return vector
						}(),
					})

					results = append(results, VectorSet{
						Value: 4,
						Vector: func() Vector {
							vector := NewDenseVector(2)
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
			if got := Span(tt.args.set...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Span() = %v, want %v", got, tt.want)
			}
		})
	}
}
