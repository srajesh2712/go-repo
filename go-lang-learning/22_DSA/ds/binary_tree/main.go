package main

type Node struct {
	value int8
	left  *Node
	right *Node
}

func NewNode(value int8) *Node {
	return &Node{value: value, left: nil, right: nil}
}
func InsertNode(n *Node, value int8) *Node {
	if n == nil {
		return nil
	}
	if n.value > value {
		if n.left == nil {
			n.left = NewNode(value)
		} else {
			InsertNode(n.left, value)
		}
	}
	if n.value < value {
		if n.right == nil {
			n.right = NewNode(value)
		} else {
			InsertNode(n.right, value)
		}

	}
	return n
}
func PrintTree(n *Node, depth int) {
	if n == nil {
		return
	}
	PrintTree(n.right, depth+1)
	padding := ""
	for i := 0; i < depth; i++ {
		padding += "    "
	}
	//fmt.Printf("%s└── %d\n", padding, n.value)
	PrintTree(n.left, depth+1)
}
func Search(n *Node, target int8) bool {
	if n == nil {
		return false
	}
	if n.value == target {
		return true
	}
	if target < n.value {
		return Search(n.left, target)
	}
	return Search(n.right, target)
}

func main() {
	root := NewNode(10)
	InsertNode(root, 5)
	InsertNode(root, 15)
	InsertNode(root, 2)
	InsertNode(root, 7)
	InsertNode(root, 12)
	InsertNode(root, 20)

	PrintTree(root, 0)
}
