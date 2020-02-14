// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float32Op

// MonoIDFloat32 is a set of float64's that closed under an associative binary operation
type MonoIDFloat32 interface {
	Zero() float32
	Reduce(done <-chan struct{}, slice <-chan float32) <-chan float32
}

type monoIDFloat32 struct {
	BinaryOpFloat32
	unit float32
}

// Zero the identity element
func (s *monoIDFloat32) Zero() float32 {
	return s.unit
}

// NewMonoIDFloat32 retun a MonoIDFloat32
func NewMonoIDFloat32(zero float32, operator BinaryOpFloat32) MonoIDFloat32 {
	return &monoIDFloat32{unit: zero, BinaryOpFloat32: operator}
}

// Reduce left folding over the monoID
func (s *monoIDFloat32) Reduce(done <-chan struct{}, slice <-chan float32) <-chan float32 {
	out := make(chan float32)
	go func() {
		result := s.unit
		for {
			select {
			case value := <-slice:
				result = s.BinaryOpFloat32.Apply(result, value)
			case <-done:
				out <- result
				close(out)
				return
			}
		}
	}()
	return out
}
