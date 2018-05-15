// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

// MonoID is algebra that is closed under an associative binary operation and has an identity element
type MonoID interface {
	Type() Type
	Unit() interface{}
	Op() Operator
}

type monoID struct {
	typ       Type
	unit      interface{}
	operation Operator
}

// NewBoolMonoID returns a bool type MonoID
func NewBoolMonoID(b bool, o Operator) MonoID {
	return &monoID{
		typ:       Bool,
		unit:      b,
		operation: o,
	}
}

// NewFloat64MonoID returns a float64 type MonoID
func NewFloat64MonoID(b float64, o Operator) MonoID {
	return &monoID{
		typ:       Float64,
		unit:      b,
		operation: o,
	}
}

func (s *monoID) Type() Type {
	return s.typ
}

func (s *monoID) Unit() interface{} {
	return s.unit
}

func (s *monoID) Op() Operator {
	return s.operation
}
