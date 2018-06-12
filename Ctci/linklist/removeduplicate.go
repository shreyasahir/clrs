package main

import (
	"fmt"
)

type node struct {
	value int
	next  *node
}

type linklist struct {
	head *node
}

func (l *linklist) removeDuplicate() {
	hash := make(map[int]int)
	curr := l.head
	var prev *node
	for curr != nil {
		if _, ok := hash[curr.value]; ok {
			prev.next = curr.next
		} else {
			hash[curr.value] = 1
			prev = curr
		}
		curr = curr.next
	}

}
func (l *linklist) insert(data int) {
	n := &node{value: data, next: nil}
	if l.head == nil {
		l.head = n
	} else {

		curr := l.head

		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
}

func (l *linklist) printList() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}
func main() {
	l := &linklist{}
	l.insert(50)
	l.insert(70)
	l.insert(30)
	l.insert(30)
	l.insert(40)
	l.printList()
	l.removeDuplicate()
	fmt.Println("after remove duplicate")
	l.printList()

}
