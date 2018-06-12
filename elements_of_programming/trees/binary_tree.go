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

func (t *tree) remove(data int) {
	//n := &node{value: data}
	remove(t.root, data)
}

func remove(n *node, data int) *node {
	if n == nil {
		return nil
	}
	if n.value > data {
		n.left = remove(n.left, data)
		return n
	}
	if n.value < data {
		n.right = remove(n.right, data)
		return n
	}

	if n.left == nil && n.right == nil {
		n = nil
		return n
	}

	if n.left == nil {
		n = n.right
		return n
	}

	if n.right == nil {
		n = n.left
		return n
	}

	leftmostright := n.right
	for {
		if leftmostright != nil && leftmostright.left != nil {
			leftmostright = leftmostright.left
		} else {
			break
		}
	}
	n.value = leftmostright.value
	n.right = remove(n.right, n.value)
	return n
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

func (t *tree) preorder(f func(n *node)) {
	preorder(t.root, f)
}

func preorder(n *node, f func(n *node)) {
	if n != nil {
		f(n)
		preorder(n.left, f)
		preorder(n.right, f)
	}
}

func main() {
	t := &tree{}
	t.insert(31)
	t.insert(41)
	t.insert(21)
	t.insert(51)
	t.insert(11)
	t.insert(8)
	t.insert(1)
	t.inorder(func(n *node) { fmt.Println("node value", n.value) })
	fmt.Println("Preorder below")
	t.preorder(func(n *node) { fmt.Println("node value", n.value) })
	t.remove(21)
	fmt.Println("print tree below")
	t.inorder(func(n *node) { fmt.Println("node value", n.value) })

}
