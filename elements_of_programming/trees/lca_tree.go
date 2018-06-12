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

func lca(n *node, a, b int) *node {
	if n == nil {
		return nil
	}
	if n.value == a || n.value == b {
		return n
	}
	left := lca(n.left, a, b)
	right := lca(n.right, a, b)
	fmt.Println("value of n", n)
	fmt.Println("value of left", left)
	fmt.Println("value of right", right)

	if left != nil && right != nil {
		return n
	}

	if left != nil {
		return left
	}
	return right
}
func (t *tree) lca(n *node, a, b int) *node {
	return lca(t.root, a, b)
}

func main() {
	t := &tree{}
	t.insert(31)
	t.insert(41)
	t.insert(21)
	t.insert(51)
	t.insert(11)
	t.insert(8)
	t.insert(12)
	fmt.Println("Inorder Traversal")
	t.inorder(func(n *node) { fmt.Println("node value", n.value) })
	fmt.Println("LCA of nodes 8 and 12")
	fmt.Println(t.lca(t.root, 8, 41))
}
