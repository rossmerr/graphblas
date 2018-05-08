// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package GraphBLAS

type Operator func(in1, in2, out *Value)

// MonoID is algebra that is closed under an associative binary operation and has an identity element
type MonoID interface {
	Type() Type
}

// GraphBLASMonoID a Value MonoID
type GraphBLASMonoID struct {
	unit      *Value
	operation Operator
}

// NewBoolMonoID returns a BoolMonoID
func NewBoolMonoID(b bool, o Operator) *GraphBLASMonoID {
	return &GraphBLASMonoID{
		unit:      ToBool(b),
		operation: o,
	}
}

// Type returns the GraphBLAS type
func (s *GraphBLASMonoID) Type() Type {
	return s.unit.typ
}

func (s *GraphBLASMonoID) Bool(in bool) bool {
	var v *Value
	s.operation(s.unit, ToBool(in), v)
	return v.Bool()
}

type Value struct {
	typ   Type
	value interface{}
}

func ToBool(b bool) *Value {
	return &Value{
		typ:   Bool,
		value: b,
	}
}

func (s *Value) Bool() bool {
	if s.typ == Bool {
		if v, ok := s.value.(bool); ok {
			return v
		}
	}

	return false
}
