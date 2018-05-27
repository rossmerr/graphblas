// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64Op

type MonoIDFloat64ToBool interface {
	BinaryOpFloat64ToBool
	Zero() float64
	Reduce(done <-chan interface{}, slice <-chan float64) <-chan bool
}

type monoIDFloat64ToBool struct {
	BinaryOpFloat64ToBool
	unit float64
}

func (s *monoIDFloat64ToBool) Zero() float64 {
	return s.unit
}

// NewMonoIDFloat64ToBool retuns a MonoIDFloat64ToBool
func NewMonoIDFloat64ToBool(zero float64, operator BinaryOpFloat64ToBool) MonoIDFloat64ToBool {
	return &monoIDFloat64ToBool{unit: zero, BinaryOpFloat64ToBool: operator}
}

func (s *monoIDFloat64ToBool) Reduce(done <-chan interface{}, slice <-chan float64) <-chan bool {
	out := make(chan bool)
	go func() {
		for {
			select {
			case value := <-slice:
				out <- s.BinaryOpFloat64ToBool.Apply(s.unit, value)
			case <-done:
				close(out)
				return
			}
		}
	}()
	return out
}
