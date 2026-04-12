package main

import (
	"testing"
)

// Benchmark für Ihre ArrayDiff-Funktion
func BenchmarkArrayDiff(b *testing.B) {
	// Vorbereitung der Testdaten
	sliceA := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		sliceA[i] = i
	}
	sliceB := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// Der eigentliche Benchmark-Lauf
	b.ResetTimer() // Ignoriert die Zeit für die Vorbereitung oben
	for i := 0; i < b.N; i++ {
		ArrayDiff(sliceA, sliceB)
	}
}

// Benchmark für Ihre ArrayDiff-Funktion
func BenchmarkArrayDiffInplaceReplacement(b *testing.B) {
	// Vorbereitung der Testdaten
	sliceA := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		sliceA[i] = i
	}
	sliceB := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// Der eigentliche Benchmark-Lauf
	b.ResetTimer() // Ignoriert die Zeit für die Vorbereitung oben
	for i := 0; i < b.N; i++ {
		ArrayDiffTwoPointerApproach(sliceA, sliceB)
	}
}
