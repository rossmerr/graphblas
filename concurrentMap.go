package GraphBLAS

import (
	"sync"
)

// Map type that can be safely shared between
type ConcurrentMap struct {
	sync.RWMutex
	items map[int]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	s := ConcurrentMap{items: make(map[int]interface{})}
	return &s
}

// Concurrent map item
type ConcurrentMapItem struct {
	Key   int
	Value interface{}
}

// Sets a key in a concurrent map
func (cm *ConcurrentMap) Set(key int, value interface{}) {
	cm.Lock()
	defer cm.Unlock()

	cm.items[key] = value
}

// Gets a key from a concurrent map
func (cm *ConcurrentMap) Get(key int) (interface{}, bool) {
	cm.Lock()
	defer cm.Unlock()

	value, ok := cm.items[key]

	return value, ok
}

// Iterates over the items in a concurrent map
// Each item is sent over a channel, so that
// we can iterate over the map using the builtin range keyword
func (cm *ConcurrentMap) Iter() <-chan ConcurrentMapItem {
	c := make(chan ConcurrentMapItem)

	f := func() {
		cm.Lock()
		defer cm.Unlock()

		for k, v := range cm.items {
			c <- ConcurrentMapItem{k, v}
		}
		close(c)
	}
	go f()

	return c
}
