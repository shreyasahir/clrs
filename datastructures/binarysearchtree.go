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
		if node.right == nil {
			node.right = n
			return nil
		}
		return insert(node.right, n)
	case node.key > n.key:
		if node.left == nil {
			node.left = n
			return nil
		}
		return insert(node.left, n)
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

func (t *Tree) findMax() *Node {
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

func (t *Tree) findMin() *Node {
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

func (t *Tree) remove(k int) {
	remove(t.root, k)
}

func remove(node *Node, k int) *Node {
	if node == nil {
		return nil
	}

	if k < node.key {
		node.left = remove(node.left, k)
		return node
	}

	if k > node.key {
		node.right = remove(node.right, k)
		return node
	}

	//case where key = node.key

	if node.left == nil && node.right == nil {
		node = nil
		return node
	}

	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}

	leftMostRightSide := node.right

	for {
		if leftMostRightSide != nil && leftMostRightSide.left != nil {
			leftMostRightSide = leftMostRightSide.left
		} else {
			break
		}
	}
	node.key, node.value = leftMostRightSide.key, leftMostRightSide.value
	node.right = remove(node.right, node.key)
	return node
}

func (t *Tree) maxDepth() int {
	return maxDepth(t.root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(n *Node) int {
	//var depth, int
	if n == nil {
		return 0
	} else {
		return 1 + max(maxDepth(n.left), maxDepth(n.right))
	}
}
func main() {
	//values := []int{45, 23, 56, 12, 10,}
	values := []int{8, 4, 10, 2, 6, 9, 12, 11, 13}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha", "tango", "mango", "ff", "gg"}
	//Create a tree and fill it from the values.
	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		err := tree.insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}

	fmt.Println("inorder")
	tree.traverse(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	fmt.Println("preorder")
	tree.preorder(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	fmt.Println("postorder")
	tree.postorder(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	value, ok := tree.find(4)
	if ok {
		fmt.Println(value)
	}

	fmt.Println(tree.findMax())
	fmt.Println(tree.findMin())
	tree.remove(10)
	fmt.Println(tree.root)
	fmt.Println(tree.root.left)
	fmt.Println(tree.root.right)
	fmt.Println("inorder")
	tree.traverse(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	fmt.Println("preorder")
	tree.preorder(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	fmt.Println("postorder")
	tree.postorder(func(n *Node) { fmt.Print(n.value, ": ", n.key, " | ") })
	fmt.Println("\n")
	depth := tree.maxDepth()
	fmt.Println("height of tree", depth)
}
