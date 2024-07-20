// Binary Search Tree
package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	var root *Node
	root = Insert(root, 5)
	root = Insert(root, 6)
	root = Insert(root, 60)

	fmt.Println()
	Inorder(root)

	fmt.Println()
	Preorder(root)

	fmt.Println()
	Postorder(root)

	fmt.Print("\n", Search(root, 61))
	fmt.Print("\n", Level(root))
}
func Insert(n *Node, data int) *Node {
	if n == nil {
		return &Node{data: data}
	}
	if data <= n.data {
		n.left = Insert(n.left, data)
	} else {
		n.right = Insert(n.right, data)
	}
	return n
}
func Inorder(n *Node) {
	if n != nil {
		Inorder(n.left)
		fmt.Print(n.data, " ")
		Inorder(n.right)
	}
}
func Preorder(n *Node) {
	if n != nil {
		fmt.Print(n.data, " ")
		Preorder(n.left)
		Preorder(n.right)
	}
}
func Postorder(n *Node) {
	if n != nil {
		Preorder(n.left)
		Preorder(n.right)
		fmt.Print(n.data, " ")
	}
}
func Search(n *Node, item int) bool {
	if n == nil {
		return false
	}
	if n.data == item {
		return true
	}
	if n.data > item {
		return Search(n.left, item)
	}
	return Search(n.right, item)
}

func Level(n *Node) int {
	if n == nil {
		return 0
	}
	leftLevel := Level(n.left)
	rightLevel := Level(n.right)
	if leftLevel > rightLevel {
		return leftLevel + 1
	}
	return rightLevel + 1
}
