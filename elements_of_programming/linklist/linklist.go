package main

import (
	"fmt"
)
type node struct {
	data int
	next *node
}

type linklist struct {
	head * node
}

func (l *linklist) add (data int) {
	n := &node{data: data,next:nil}

	if l.head == nil {
		l.head = n
	} else {
		n1 := l.head
		for n1.next != nil {
			n1 = n1.next
		}
		n1.next = n
	}
}

func (l * linklist) printList() {
	n := l.head
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

//Not accurate just for testing my cycle code currently assumed it points to head.
func (l * linklist) printListWithCycle() {
	n := l.head
	fmt.Println(n.data)
	n = n.next
	for n != l.head {
		fmt.Println(n.data)
		n = n.next
	}
	fmt.Println(n.data)
}

func (l *linklist) reverseIterative() {
	var prev *node
	curr := l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next 
	}
	l.head = prev
}

func mergeList(l1,l2 *node) *node {
	if l1 == nil {
		return l2
	} 
	if l2 == nil {
		return l1
	}

	if l1.data <= l2.data {
		l1.next = mergeList(l1.next,l2)
		return l1
	} else {
		l2.next = mergeList(l1,l2.next)
		return l2
	}
}

func (l *linklist) addCycle() {

	n := l.head
	//curr := n.next
	for n.next != nil {
		//fmt.Println("addcycle",n.data)
		n = n.next
	}
	//fmt.Println("addcycle outer",n.data)
	n.next = l.head
}

func (l * linklist) hasCycle() *node{
	if l.head == nil {
		fmt.Println("No")
	}
	fast:= l.head
	slow := l.head
	for slow != nil && slow.next != nil  && fast != nil && fast.next != nil  && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			slow = l.head
			for slow != fast {
				slow = slow.next
				fast = fast.next
			}
			return slow
		}
	}

	return nil
}

func main() {
l := linklist{}
l.add(4)
l.add(5)
l.add(7)
l.add(9)
l.printList()
fmt.Println("###########")
//l.reverseIterative()
//l.printList()
l.addCycle()
l.printListWithCycle()
fmt.Println(l.hasCycle())
// l1:= linklist{}
// l1.add(2)
// l1.add(6)
// l1.add(10)
// l1.add(11)
// l1.printList()
// //l2 :=  linklist{}
// fmt.Println("###########")
// head := mergeList(l.head,l1.head)
// l2 := linklist{head:head }
// //l2.add()
// l2.printList()

}