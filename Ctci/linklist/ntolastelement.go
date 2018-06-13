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

func (l *linklist) nthtolast(k int) {
	//var prev,curr *node
	prev := l.head
	curr := l.head
	for i := 1; i < k; i++ {
		curr = curr.next
		fmt.Println("i", i)
	}
	fmt.Println("curr", curr.value)
	for curr.next != nil {
		curr = curr.next
		prev = prev.next
	}
	fmt.Println("nth is", prev.value)
}

func (l *linklist) deletenode(n *node) {
	curr := l.head
	for curr != nil {
		if curr.value == 30 {
			break
		}
		curr = curr.next
	}
	fmt.Println("curr", curr.value)
	curr.value = curr.next.value
	curr.next = curr.next.next
}
func main() {
	l := &linklist{}
	l.insert(50)
	l.insert(70)
	l.insert(30)
	l.insert(40)
	l.insert(20)
	l.insert(10)
	l.insert(80)
	l.insert(100)
	l.printList()
	l.nthtolast(4)

}
