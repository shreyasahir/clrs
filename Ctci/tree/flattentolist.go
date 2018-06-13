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

func flattenTree(n *node) {
	stack := []*node{}
	p := n
	for p != nil || len(stack) > 0 {
		if p.right != nil {
			stack = append(stack, p.right)
		}
		if p.left != nil {
			p.right = p.left
			p.left = nil
		} else if len(stack) > 0 {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p.right = temp
		}
		p = p.right
	}
}

// Stack<TreeNode> stack = new Stack<TreeNode>();
// TreeNode p = root;

// while(p != null || !stack.empty()){

// 	if(p.right != null){
// 		stack.push(p.right);
// 	}

// 	if(p.left != null){
// 		p.right = p.left;
// 		p.left = null;
// 	}else if(!stack.empty()){
// 		TreeNode temp = stack.pop();
// 		p.right=temp;
// 	}

// 	p = p.right;
// }
// }

func main() {
	t := &tree{}
	t.root = &node{value: 1}
	t.root.left = &node{value: 2}
	t.root.right = &node{value: 5}
	t.root.left.left = &node{value: 3}
	t.root.left.right = &node{value: 4}
	t.root.right.right = &node{value: 6}
	t.inorder(func(n *node) { fmt.Println(n.value) })
	fmt.Println("after flatten")
	flattenTree(t.root)
	t.inorder(func(n *node) { fmt.Println(n.value) })

}
