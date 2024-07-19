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
	root := &Node{data: 20}
	root.left = &Node{data: 5}
	root.right = &Node{data: 50}
	root.left.left = &Node{data: 120}
	root.left.right = &Node{data: 10}

	root.right.right = &Node{data: 55}
	root.right.left = &Node{data: 23}

	fmt.Println("")
	inorder(root)

	fmt.Println("")
	preorder(root)

	fmt.Println("")
	postorder(root)

	fmt.Println("Count of the tree", countLeaves(root))

	fmt.Println("Level of the tree", Level(root))
}

// Inorder traversal
// 1 First, visit all the nodes in the left subtree
// 2 Then the root node
// 3 Visit all the nodes in the right subtree
func inorder(n *Node) {
	if n != nil {
		inorder(n.left)
		fmt.Print(n.data, " ")
		inorder(n.right)
	}
}

// Preorder traversal
// 1 First print, the root node
// 2 Then, visit all the nodes in the left subtree
// 3 Visit all the nodes in the right subtree
func preorder(n *Node) {
	if n != nil {
		fmt.Print(n.data, " ")
		preorder(n.left)
		preorder(n.right)
	}
}

// Postorder traversal
// 1 First, visit all the nodes in the left subtree
// 2 Visit all the nodes in the right subtree
// 3 Print, the root node
func postorder(n *Node) {
	if n != nil {
		postorder(n.left)
		postorder(n.right)
		fmt.Print(n.data, " ")
	}
}

// Find the depth/level of the binary tree
func Level(n *Node) int {
	if n == nil {
		return -1 // Make it 0 for 1 indexing
	}
	leftLevel := Level(n.left)
	rightLevel := Level(n.right)
	if leftLevel > rightLevel {
		return leftLevel + 1
	}
	return rightLevel + 1
}

// Find the total number of leaf node
func countLeaves(n *Node) int {
	if n == nil {
		return 0
	}
	if n.left == nil && n.right == nil {
		return 1
	}
	return countLeaves(n.left) + countLeaves(n.right)
}
