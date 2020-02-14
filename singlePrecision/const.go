// Copyright (c) 2020 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package singlePrecision

import "github.com/RossMerr/Caudex.GraphBLAS/binaryOp/float32Op"

const defaultFloat32 = float32(0)

var defaultMonoIDAddition = float32Op.NewMonoIDFloat32(0, float32Op.Addition)
var defaultMonoIDMaximum = float32Op.NewMonoIDFloat32(0, float32Op.Maximum)
