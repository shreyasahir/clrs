package main

type node struct {
	value int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func (t *tree) insert(value int) {
	n := &node{value: value}
	if t.root == nil {
		t.root = n
	} else {
		insert(t.root, n)
	}
}

func main() {
	t := &tree{}
	t.insert(1)
	t.insert(1)
	t.insert(1)
	t.insert(1)
	t.insert(1)

}
