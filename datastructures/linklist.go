package main

import (
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

type linklist struct {
	head *node
}

func (l *linklist) add(x interface{}) {
	n := node{data: x, next: l.head}
	l.head = &n

}

func (l *linklist) printList() {

	n := l.head
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func main() {
	l1 := new(linklist)
	l1.add(1)
	l1.add(10)
	l1.add(40)
	l1.printList()
}
