package main

import (
	"fmt"
	"math"
)

func Race(v1, v2, g int) [3]int {
	if v1 > v2 {
		return [3]int{-1, -1, -1}
	}
	totalSeconds := float64(g) * 3600 / float64(v2-v1)

	// Use a tiny epsilon and Floor to handle precision errors like 19.99999
	totalSecondsInt := int(math.Floor(totalSeconds + 0.00001))

	h := totalSecondsInt / 3600
	m := (totalSecondsInt % 3600) / 60
	s := totalSecondsInt % 60
	return [3]int{h, m, int(s)}
}

func main() {
	fmt.Println(Race(720, 850, 70))
}
