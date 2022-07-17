// Copyright (c) 2022 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sort_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/rossmerr/graphblas"
	"github.com/rossmerr/graphblas/sort"
)

func TestBubbleColumns(t *testing.T) {

	setup := graphblas.NewDenseMatrixFromArray([][]rune{
		{'c', 'a', 'b'},
		{'b', 'a', 'c'},
		{'a', 'b', 'c'},
	})

	want := graphblas.NewDenseMatrixFromArray([][]rune{
		{'a', 'b', 'c'},
		{'a', 'c', 'b'},
		{'b', 'c', 'a'},
	})

	type args struct {
		ctx context.Context
		a   graphblas.MatrixRune
	}
	tests := []struct {
		name string
		args args
		want graphblas.MatrixRune
	}{

		{
			name: "column",
			args: args{
				ctx: context.Background(),
				a:   setup,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort.BubbleColumns(tt.args.ctx, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleRow(t *testing.T) {

	setup := graphblas.NewDenseMatrixFromArray([][]rune{
		{'c', 'a', 'b'},
		{'b', 'a', 'c'},
		{'a', 'b', 'c'},
	})

	want := graphblas.NewDenseMatrixFromArray([][]rune{
		{'a', 'b', 'c'},
		{'b', 'a', 'c'},
		{'c', 'a', 'b'},
	})

	type args struct {
		ctx context.Context
		a   graphblas.MatrixRune
	}
	tests := []struct {
		name string
		args args
		want graphblas.MatrixRune
	}{

		{
			name: "column",
			args: args{
				ctx: context.Background(),
				a:   setup,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort.BubbleRow(tt.args.ctx, tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
