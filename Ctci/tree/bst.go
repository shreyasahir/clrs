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

func (t *tree) insert(val int) {
	n := &node{value: val}
	if t.root == nil {
		t.root = n
	} else {
		insert(t.root, n)
	}
}

func insert(curr, n *node) error {
	switch {
	case curr.value <= n.value:
		if curr.right == nil {
			curr.right = n
			return nil
		}
		return insert(curr.right, n)
	case curr.value > n.value:
		if curr.left == nil {
			curr.left = n
			return nil
		}
		return insert(curr.left, n)
	}
	return nil
}

func (t *tree) inOrder(f func(*node)) {
	inOrder(t.root, f)
}
func inOrder(n *node, f func(*node)) {
	if n != nil {
		inOrder(n.left, f)
		f(n)
		inOrder(n.right, f)
	}
}

func (t *tree) remove(val int) {
	n := &node{value: val}
	remove(t.root, n)

}

func remove(curr, n *node) *node {

	if curr == nil {
		return nil
	}
	if curr.value > n.value {
		curr.left = remove(curr.left, n)
		return curr
	}
	if curr.value < n.value {
		curr.right = remove(curr.right, n)
		return curr
	}

	if curr.left == nil && curr.right == nil {
		curr = nil
		return curr
	}

	if curr.left == nil {
		curr = curr.right
		return curr
	}

	if curr.right == nil {
		curr = curr.left
		return curr
	}
	rightmostleft := curr.right

	for {
		if rightmostleft != nil && rightmostleft.left != nil {
			rightmostleft = rightmostleft.left
		} else {
			break
		}
	}
	curr.value = rightmostleft.value
	curr.right = remove(curr.right, curr)
	return curr
}

func invertTree(n *node) {
	queue := []*node{}
	if n != nil {
		queue = append(queue, n)
	}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.left != nil {
			queue = append(queue, p.left)
		}
		if p.right != nil {
			queue = append(queue, p.right)
		}
		temp := p.left
		p.left = p.right
		p.right = temp
	}
}
func main() {
	t := &tree{}
	t.insert(8)
	t.insert(6)
	t.insert(5)
	t.insert(15)
	t.insert(29)
	t.insert(44)
	t.insert(7)
	t.insert(45)
	t.insert(1)
	t.inOrder(func(n *node) { fmt.Println(n.value) })
	//t.remove(6)
	//fmt.Println("after delete")
	t.inOrder(func(n *node) { fmt.Println(n.value) })
	fmt.Println("after invert")
	invertTree(t.root)
	t.inOrder(func(n *node) { fmt.Println(n.value) })

}
