// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// A sparse linear algebra library that defines a set of matrix and vector operations
package f64

import "github.com/rossmerr/graphblas/binaryop/float64op"

const defaultFloat64 = float64(0)

var defaultMonoIDAddition = float64op.NewMonoIDFloat64(0, float64op.Addition)
var defaultMonoIDMaximum = float64op.NewMonoIDFloat64(0, float64op.Maximum)
