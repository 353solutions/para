package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

var mu sync.Mutex
var counter int64

func BenchmarkMutex(b *testing.B) {
	for b.Loop() {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func BenchmarkAtomic(b *testing.B) {
	for b.Loop() {
		atomic.AddInt64(&counter, 1)
	}
}
