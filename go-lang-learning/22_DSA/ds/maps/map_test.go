package main

//go test -bench=. -benchmem
import (
	"fmt"
	"testing"
)

// Benchmark for the Set operation
func BenchmarkGoMap_Set(b *testing.B) {
	gomap := &GoMap{}

	// b.N is a number Go automatically adjusts to get a stable average
	for i := 0; i < b.N; i++ {
		// We use a different key each time to see real performance
		key := fmt.Sprintf("key%d", i)
		gomap.Set(key, "value")
	}
}

// Benchmark for the Get operation
func BenchmarkGoMap_Get(b *testing.B) {
	gomap := &GoMap{}
	// Pre-fill the map so we have something to Get
	for i := 0; i < 1000; i++ {
		gomap.Set(fmt.Sprintf("key%d", i), "value")
	}

	// Reset the timer so the "Set" time doesn't ruin our "Get" stats
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// We look for a key that definitely exists
		gomap.Get("key500")
	}
}
