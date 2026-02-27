package main

import (
	"testing"
)

// LinearSearch simulates looking through a standard list/array
func LinearSearch(data []int8, target int8) bool {
	for _, v := range data {
		if v == target {
			return true
		}
	}
	return false
}

func BenchmarkTreeSearch(b *testing.B) {
	// This happens ONCE
	root := NewNode(50)
	for i := int8(0); i < 100; i++ {
		InsertNode(root, int8((int(i)*31)%100))
	}

	b.ResetTimer() // This tells Go: "Ignore the time spent building the tree above"

	// This happens b.N times
	for i := 0; i < b.N; i++ {
		Search(root, 99)
	}
}

// BenchmarkLinearSearch tests a standard array
func BenchmarkLinearSearch(b *testing.B) {
	data := make([]int8, 100)
	for i := int8(0); i < 100; i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(data, 99)
	}
}
