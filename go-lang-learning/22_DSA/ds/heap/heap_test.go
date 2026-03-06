package main

import (
	"math/rand"
	"testing"
)

// This benchmarks the insertion of 1,000 random elements
func BenchmarkHeap_Insert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Stop timer to ignore setup (creating the node)
		b.StopTimer()
		h := NewMinHeap(1000)
		b.StartTimer()

		for j := 0; j < 1000; j++ {
			h.CreateHeap(rand.Intn(1000000))
		}
	}
}

// This benchmarks extracting all 1,000 elements (DeleteRoot)
func BenchmarkHeap_DeleteRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Stop timer to setup a full heap first
		b.StopTimer()
		h := &Node{}
		for j := 0; j < 1000; j++ {
			h.CreateHeap(rand.Intn(1000000))
		}
		b.StartTimer()

		// Benchmark how long it takes to empty the heap
		for len(h.arr) > 0 {
			h.DeleteRoot()
		}
	}
}
