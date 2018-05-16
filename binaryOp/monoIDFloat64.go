// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package binaryOp

type MonoIDFloat64 interface {
	BinaryOpFloat64
	Zero() float64
	Reduce([]float64) []float64
}

type monoIDFloat64 struct {
	BinaryOpFloat64
	unit float64
}

func (s *monoIDFloat64) Zero() float64 {
	return s.unit
}

// NewMonoIDFloat64 retuns a MonoIDFloat64
func NewMonoIDFloat64(zero float64, operator BinaryOpFloat64) MonoIDFloat64 {
	return &monoIDFloat64{unit: zero, BinaryOpFloat64: operator}
}

func (s *monoIDFloat64) Reduce(slice []float64) []float64 {
	results := make([]float64, len(slice))
	for i, value := range slice {
		results[i] = s.BinaryOpFloat64.Apply(s.unit, value)
	}

	return results
}
