package GraphBLAS_test

import (
	"sync"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

func BenchmarkMap(b *testing.B) {
	s := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		s[i] = i
	}
}

func BenchmarkConcurrentMap(b *testing.B) {
	s := GraphBLAS.NewConcurrentMap()
	for i := 0; i < 1000000; i++ {
		s.Set(i, i)
	}
}

func BenchmarkSyncMap(b *testing.B) {
	s := sync.Map{}
	for i := 0; i < 1000000; i++ {
		s.Store(i, i)
	}
}
