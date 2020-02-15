package opencl

import (
	"context"
	"fmt"

	"github.com/RossMerr/Caudex.GraphBLAS/singlePrecision/opencl/kernels"
)

func AddOne(ctx context.Context, v1, v2 *DenseVector) <-chan error {
	ch := make(chan error, 1)
	deviceCtx, ok := ctx.(*deviceContext)
	if !ok {
		ch <- fmt.Errorf("No opencl device found in context")
		return ch
	}

	kernel := deviceCtx.Kernel(kernels.AddOne)

	//run kernel (global work size 16 and local work size 1)
	return kernel.Global(deviceCtx.globalWorkSize).Local(deviceCtx.localWorkSize).Run(v1.v)
}
