// Copyright (c) 2020 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package f32

import "github.com/rossmerr/graphblas/binaryop/float32op"

const defaultFloat32 = float32(0)

var defaultMonoIDAddition = float32op.NewMonoIDFloat32(0, float32op.Addition)
var defaultMonoIDMaximum = float32op.NewMonoIDFloat32(0, float32op.Maximum)
