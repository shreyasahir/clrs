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

// func (t *tree) inorder(f func(n *node)) {
// 	inorder(t.root, f)
// }

// func inorder(n *node, f func(n *node)) {
// 	if n != nil {
// 		inorder(n.left, f)
// 		f(n)
// 		inorder(n.right, f)
// 	}
// }

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

// func (t *tree) morrisInorder() {
// 	var curr, pre *node
// 	curr = t.root

// 	for curr != nil {
// 		if curr.left == nil {
// 			fmt.Println(curr.value)
// 			curr = curr.right
// 		} else {
// 			pre = curr.left
// 			for pre.right != nil && pre.right != curr {
// 				pre = pre.right
// 			}
// 			if pre.right == nil {
// 				pre.right = curr
// 				curr = curr.left
// 			} else {
// 				pre.right = nil
// 				fmt.Println(curr.value)
// 				curr = curr.right
// 			}
// 		}
// 	}
// }

func morrisPostorder(n *node) {
	var curr, pre *node
	curr = n

	for curr != nil {
		if curr.left == nil {
			fmt.Println(curr.value)
			curr = curr.right
		} else {
			pre = curr.left
			for pre.right != nil && pre.right != curr {
				pre = pre.right
			}
			if pre.right == curr {
				pre.right = nil
				curr = curr.right
			} else {
				fmt.Println(curr.value)
				pre.right = curr
				curr = curr.left
			}
		}
	}
}

func (t *tree) postorder() {
	if t.root != nil {
		fmt.Println(t.root.value)
		morrisPostorder(t.root.right)
		morrisPostorder(t.root.left)

	}
}

func (t *tree) inorder() {
	if t.root != nil {
		morrisPostorder(t.root.left)
		fmt.Println(t.root.value)
		morrisPostorder(t.root.right)

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
	t.insert(12)
	// fmt.Println("Inorder Traversal")
	// t.inorder(func(n *node) { fmt.Println("node value", n.value) })
	fmt.Println("Morris inorder")
	t.inorder()
	fmt.Println("Morris postorder")
	t.postorder()

}
