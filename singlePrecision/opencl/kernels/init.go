package kernels

import (
	"context"
	"fmt"
	"io/ioutil"
)

func Initialize(ctx context.Context) error {
	deviceCtx, ok := ctx.(deviceContext)
	if !ok {
		return fmt.Errorf("No opencl device found in context")
	}

	byt, err := ioutil.ReadFile("addOne.cl")
	if err != nil {
		return err
	}
	deviceCtx.AddProgram(string(byt))

	return nil
}
