package main

// Min,max,delete remaining
import (
	"fmt"
	"log"
)

type Node struct {
	left  *Node
	right *Node
	key   int
	value interface{}
}

type Tree struct {
	root *Node
}

func (t *Tree) find(k int) (string, bool) {
	if t.root == nil {
		return "", false
	} else {
		return find(t.root, k)
	}
}

func find(n *Node, k int) (string, bool) {
	if n == nil {
		return "", false
	}

	switch {
	case n.key == k:
		return n.value.(string), true
	case n.key < k:
		return find(n.left, k)
	case n.key > k:
		return find(n.right, k)

	}
	return "", false
}

func (t *Tree) insert(k int, x interface{}) error {
	node := &Node{left: nil, right: nil, key: k, value: x}

	if t.root == nil {
		t.root = node
	} else {
		insert(t.root, node)
	}
	return nil
}

func insert(node *Node, n *Node) error {

	switch {
	case n.key == node.key:
		return nil
	case node.key < n.key:
		if node.left == nil {
			node.left = n
			return nil
		}
		return insert(node.left, n)
	case node.key > n.key:
		if node.right == nil {
			node.right = n
			return nil
		}
		return insert(node.right, n)
	}
	return nil
}

func (t *Tree) traverse(f func(*Node)) {
	traverse(t.root, f)
}

func traverse(n *Node, f func(*Node)) {
	if n != nil {
		traverse(n.left, f)
		f(n)
		traverse(n.right, f)
	}
}

func (t *Tree) preorder(f func(*Node)) {
	preorder(t.root, f)
}

func preorder(n *Node, f func(*Node)) {
	if n != nil {
		f(n)
		preorder(n.left, f)
		preorder(n.right, f)
	}
}

func (t *Tree) postorder(f func(*Node)) {
	postorder(t.root, f)
}

func postorder(n *Node, f func(*Node)) {
	if n != nil {
		postorder(n.left, f)
		postorder(n.right, f)
		f(n)
	}
}

func main() {
	values := []int{45, 23, 56, 12, 10}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}
	//Create a tree and fill it from the values.
	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}
	tree.traverse(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	tree.preorder(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	tree.postorder(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	value, ok := tree.find(4)
	if ok {
		fmt.Println(value)
	}
}
