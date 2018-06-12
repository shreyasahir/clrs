package main

import (
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func (t *tree) insert(data int) error {
	n := &node{value: data, left: nil, right: nil}
	if t.root == nil {
		t.root = n
	} else {
		insert(t.root, n)
	}
	return nil
}

func insert(n *node, data *node) error {
	if n.value <= data.value {
		if n.right == nil {
			n.right = data
			return nil
		}
		return insert(n.right, data)
	} else {
		if n.left == nil {
			n.left = data
			return nil
		}
		return insert(n.left, data)
	}
	return nil
}

func (t *tree) inorder(f func(n *node)) {
	inorder(t.root, f)
}

func inorder(n *node, f func(n *node)) {
	if n != nil {
		inorder(n.left, f)
		f(n)
		inorder(n.right, f)
	}
}

func getHeight(n *node) int {
	if n == nil {
		return -1
	}

	leftHeight := getHeight(n.left)
	if leftHeight == -2 {
		return -2
	}

	rightHeight := getHeight(n.right)
	if rightHeight == -2 {
		return -2
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -2
	}

	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a-b > 0 {
		return a
	}
	return b
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * (a)

}
func main() {
	t := &tree{}
	t.insert(31)
	t.insert(41)
	t.insert(21)
	t.insert(51)
	t.insert(11)
	//t.insert(8)
	//t.insert(1)
	fmt.Println("Inorder Traversal")
	t.inorder(func(n *node) { fmt.Println("node value", n.value) })
	fmt.Println("Is tree balanced")
	fmt.Println(getHeight(t.root))
}
