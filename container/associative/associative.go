package associative

// Array associative array
type Array interface {
	// Enumerate iterates through all elements
	Enumerate() Enumerate

	// The number of elements in the array
	Size() int

	// At returns the value of a array element at n-th
	At(n string) interface{}
	Append(n string, value interface{})
	Update(n string, value interface{})
	Delete(n string) interface{}
}

// Enumerate iterates over the array
type Enumerate interface {
	// HasNext checks for the next element in the matrix
	HasNext() bool

	// Next move the iterator over the array
	Next() (n string, v interface{})
}

type array struct {
	elements map[string]interface{}
	index    []string
}

// Enumerate iterates through all non-zero elements, order is not guaranteed
func (s *array) Enumerate() Enumerate {
	return s.iterator()
}

// The number of elements in the array
func (s *array) Size() int {
	return len(s.elements)
}

// At returns the value of a array element at n-th
func (s *array) At(n string) interface{} {
	return s.elements[n]
}

func (s *array) Append(n string, value interface{}) {
	s.Update(n, value)
	s.index = append(s.index, n)
}

func (s *array) Update(n string, value interface{}) {
	s.elements[n] = value
}

func (s *array) Delete(n string) interface{} {
	value := s.elements[n]
	delete(s.elements, n)
	for i, v := range s.index {
		if v == n {
			s.index = append(s.index[:i], s.index[i+1:]...)
			break
		}
	}

	return value
}

func (s *array) iterator() *arrayIterator {
	i := &arrayIterator{
		array: s,
		size:  s.Size(),
		last:  0,
	}
	return i
}

type arrayIterator struct {
	array *array
	size  int
	last  int
}

// HasNext checks the iterator has any more values
func (s *arrayIterator) HasNext() bool {
	if s.last >= s.size {
		return false
	}
	return true
}

// Next moves the iterator and returns the index and value
func (s *arrayIterator) Next() (n string, v interface{}) {
	p := s.last
	s.last++
	i := s.array.index[p]
	return i, s.array.elements[i]
}
