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

func add(l1, l2 *linklist) {
	curr := l1.head
	curr1 := l2.head
	var carry int
	for curr != nil {
		sum := curr.value + curr1.value + carry
		if sum > 9 {
			carry = sum % 9
		}
		fmt.Println(curr.value, curr1.value)

		if sum >= 10 {
			sum -= 10
		}
		fmt.Println("sum", sum)
		fmt.Println("carry", carry)
		curr = curr.next
		curr1 = curr1.next
	}
}

func main() {
	l1 := &linklist{}
	l1.insert(3)
	l1.insert(1)
	l1.insert(5)
	l1.printList()
	fmt.Println("linklist 2 below")
	l2 := &linklist{}
	l2.insert(5)
	l2.insert(9)
	l2.insert(2)
	l2.printList()
	add(l1, l2)
}
