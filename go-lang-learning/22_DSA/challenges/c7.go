package main

// go test -bench=. -benchmem c7.go c7_test.go
import "fmt"

func ArrayDiff(a, b []int) []int {
	bmap := make(map[int]struct{}, len(b))
	for _, v := range b {
		bmap[v] = struct{}{}
	}
	result := make([]int, 0, len(a))

	for _, v := range a {

		if _, found := bmap[v]; !found {
			result = append(result, v)
		}

	}
	return result
}

func ArrayDiffTwoPointerApproach(a, b []int) []int {
	bmap := make(map[int]struct{}, len(b))
	for _, v := range b {
		bmap[v] = struct{}{}
	}

	k := 0
	for i := 0; i < len(a); i++ {

		if _, found := bmap[a[i]]; !found {
			a[k] = a[i]
			k++
		}

	}
	return a[:k]
}

func main() {
	a := []int{1, 2}
	b := []int{1}
	c := ArrayDiff(a, b)
	fmt.Println(c)
	fmt.Println(ArrayDiffTwoPointerApproach(a, b))
}
