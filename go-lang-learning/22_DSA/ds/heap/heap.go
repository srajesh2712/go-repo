package main

import (
	"fmt"
)

type Node struct {
	arr          []int
	currentIndex int
}

func NewMinHeap(capacity int) *Node {
	return &Node{
		// Pre-allocate the slice capacity!
		arr: make([]int, 0, capacity),
	}
}
func (node *Node) CreateHeap(val int) *Node {
	node.arr = append(node.arr, val)
	node.BubbleUp(len(node.arr) - 1)
	return node
}
func (node *Node) BubbleUp(currentIndex int) {
	if currentIndex <= 0 {
		return
	}
	parentIndex := (currentIndex - 1) / 2

	//fmt.Println(currentIndex, parentIndex)
	if node.arr[parentIndex] >= node.arr[currentIndex] {
		node.arr[parentIndex], node.arr[currentIndex] = node.arr[currentIndex], node.arr[parentIndex]

		node.BubbleUp(parentIndex)
	}
}
func (node *Node) BubbleDown(currentIndex int) {

	smallest := currentIndex
	left := 2*currentIndex + 1
	right := 2*currentIndex + 2

	if left < len(node.arr) && node.arr[left] < node.arr[smallest] {
		smallest = left
	}
	if right < len(node.arr) && node.arr[right] < node.arr[smallest] {
		smallest = right
	}
	if smallest == currentIndex {
		return // The parent is already the smallest, we are done!
	}
	if smallest != currentIndex {
		node.arr[smallest], node.arr[currentIndex] = node.arr[currentIndex], node.arr[smallest]
		node.BubbleDown(smallest)

	}

}

func (node *Node) DeleteRoot() int {
	if len(node.arr) == 0 {
		return -1
	}
	root := node.arr[0]
	node.arr[0] = node.arr[len(node.arr)-1]
	node.arr = node.arr[:len(node.arr)-1]
	node.BubbleDown(0)
	return root

}
func main() {
	node := &Node{}
	unsorted := []int{40, 10, 30, 5, 20, 15, 1}
	for _, val := range unsorted {
		node.CreateHeap(val)
	}
	fmt.Println("Heap Structure", node.arr)

	fmt.Print("Sorted output : ")
	for len(node.arr) > 0 {
		fmt.Printf("%d ", node.DeleteRoot())
	}
	fmt.Println()
}
