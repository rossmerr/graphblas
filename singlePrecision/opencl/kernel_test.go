// Copyright (c) 2020 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package opencl

import (
	"fmt"
	"testing"

	"github.com/microo8/blackcl"
)

func TestVector_Update(t *testing.T) {
	ds, err := blackcl.GetDevices(blackcl.DeviceTypeAll)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range ds {
		t.Log(d.Name())
		t.Log(d.Profile())
		t.Log(d.OpenCLCVersion())
		t.Log(d.DriverVersion())
		t.Log(d.Extensions())
		t.Log(d.Vendor())
		err = d.Release()
		if err != nil {
			t.Fatal(err)
		}
	}

	d, err := blackcl.GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	if d == nil {
		t.Fatal("device is nil")
	}
	fmt.Println(d)
	err = d.Release()
	if err != nil {
		t.Fatal(err)
	}
}
