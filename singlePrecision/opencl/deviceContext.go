package opencl

import (
	"context"

	"github.com/RossMerr/Caudex.GraphBLAS/singlePrecision/opencl/kernels"
	"github.com/microo8/blackcl"
)

type deviceContext struct {
	context.Context
	*blackcl.Device
	globalWorkSize int
	localWorkSize  int
}

// NewDeviceContext returns a Context
func NewDeviceContext(device *blackcl.Device) (context.Context, error) {
	ctx := context.Background()
	deviceCtx := &deviceContext{ctx, device, 16, 1}
	err := kernels.Initialize(deviceCtx)
	return deviceCtx, err
}
