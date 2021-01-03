// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64op

// MonoIDFloat64 is a set of float64's that closed under an associative binary operation
type MonoIDFloat64 interface {
	Zero() float64
	Reduce(done <-chan struct{}, slice <-chan float64) <-chan float64
}

type monoIDFloat64 struct {
	BinaryOpFloat64
	unit float64
}

// Zero the identity element
func (s *monoIDFloat64) Zero() float64 {
	return s.unit
}

// NewMonoIDFloat64 retun a MonoIDFloat64
func NewMonoIDFloat64(zero float64, operator BinaryOpFloat64) MonoIDFloat64 {
	return &monoIDFloat64{unit: zero, BinaryOpFloat64: operator}
}

// Reduce left folding over the monoID
func (s *monoIDFloat64) Reduce(done <-chan struct{}, slice <-chan float64) <-chan float64 {
	out := make(chan float64)
	go func() {
		result := s.unit
		for {
			select {
			case value := <-slice:
				result = s.BinaryOpFloat64.Apply(result, value)
			case <-done:
				out <- result
				close(out)
				return
			}
		}
	}()
	return out
}
