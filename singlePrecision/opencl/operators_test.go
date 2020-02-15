package opencl

import (
	"context"
	"reflect"
	"testing"

	"github.com/microo8/blackcl"
)

func TestAddOne(t *testing.T) {
	type args struct {
		ctx context.Context
		f1  []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "AddOne",
			args: args{
				ctx: func() context.Context {
					d, _ := blackcl.GetDefaultDevice()
					ctx, _ := NewDeviceContext(d)
					return ctx
				}(),
				f1: []float32{1, 2, 3, 4},
			},
			want: []float32{2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		v1, _ := NewDenseVector(tt.args.ctx, 4)
		v1.Copy(tt.args.f1)
		AddOne(tt.args.ctx, v1, v1)
		got, _ := v1.Data()
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
