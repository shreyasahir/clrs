package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	t = &tree{}
)

type node struct {
	left  *node
	right *node
	key   int
	value interface{}
}

type tree struct {
	root *node
}

func (t *tree) insert(k int, v interface{}) error {
	n := &node{left: nil, right: nil, key: k, value: v}
	if t.root == nil {
		t.root = n
	} else {
		insert(t.root, n)
	}
	return nil
}

func insert(n1, n2 *node) error {

	switch {
	case n1.key == n2.key:
		return errors.New("duplicate value")
	case n1.key < n2.key:
		if n1.right == nil {
			n1.right = n2
			return nil
		}
		return insert(n1.right, n2)
	case n1.key > n2.key:
		if n1.left == nil {
			n1.left = n2
			return nil
		}
		return insert(n1.left, n2)
	}
	return nil
}

func (t *tree) inorder(f func(*node)) error {
	inorder(t.root, f)
	return nil
}

func inorder(n *node, f func(*node)) {
	if n != nil {
		inorder(n.left, f)
		f(n)
		inorder(n.right, f)
	}
}

func (t *tree) preorder(f func(*node)) error {
	preorder(t.root, f)
	return nil
}

func preorder(n *node, f func(*node)) {
	if n != nil {
		f(n)
		preorder(n.left, f)
		preorder(n.right, f)
	}
}

func (t *tree) findMin() *node {
	n := t.root
	if n == nil {
		return nil
	}

	for {
		if n.left == nil {
			return n
		}
		n = n.left
	}
}

func (t *tree) findMax() *node {
	n := t.root
	if n == nil {
		return nil
	}

	for {
		if n.right == nil {
			return n
		}
		n = n.right
	}
}
func addnodes() {
	values := []int{8, 4, 10, 2, 6, 9, 12, 11, 13}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha", "tango", "mango", "ff", "gg"}
	//Create a tree and fill it from the values.

	for i := 0; i < len(values); i++ {
		err := t.insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}
}

func find(n *node, k int) (string, bool) {

	if n == nil {
		return "", false
	}
	switch {
	case n.key == k:
		return n.value.(string), true
	case n.key < k:
		return find(n.right, k)
	case n.key > k:
		return find(n.left, k)
	}
	return "", false
}

func (t *tree) remove(k int) {
	remove(t.root, k)
}

func remove(n *node, k int) *node {
	if n == nil {
		return nil
	}
	if n.key < k {
		n.right = remove(n.right, k)
		return n
	}
	if n.key > k {
		n.left = remove(n.left, k)
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

	leftmostrightside := n.right

	for {
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}

	n.key, n.value = leftmostrightside.key, leftmostrightside.value
	n.right = remove(n.right, n.key)
	return n

}
func main() {

	addnodes()
	fmt.Println("inorder")
	t.inorder(func(n *node) { fmt.Print(n.value, ": ", n.key, " | \n") })
	fmt.Println("preorder")
	t.preorder(func(n *node) { fmt.Println(n.value, ": ", n.key, "|") })
	fmt.Println()
	fmt.Println("Min value", t.findMin().value)
	fmt.Println()
	fmt.Println("Max value", t.findMax().value)
	fmt.Println()
	val, err := find(t.root, 9)
	fmt.Println(err)
	fmt.Println("find", val)

	fmt.Println("remove element")
	t.remove(10)
	t.inorder(func(n *node) { fmt.Print(n.value, ": ", n.key, " | \n") })

}
