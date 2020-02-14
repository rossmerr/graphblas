// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// A sparse linear algebra library that defines a set of matrix and vector operations
package doublePrecision

import "github.com/RossMerr/Caudex.GraphBLAS/binaryOp/float64Op"

const defaultFloat64 = float64(0)

var defaultMonoIDAddition = float64Op.NewMonoIDFloat64(0, float64Op.Addition)
var defaultMonoIDMaximum = float64Op.NewMonoIDFloat64(0, float64Op.Maximum)
