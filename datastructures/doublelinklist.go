package main

import (
	"fmt"
)

//Node is for each node in a link list
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

//LinkList act as collection of Nodes
type LinkList struct {
	head *Node
}

func (l *LinkList) add(x interface{}) {
	head := l.head
	n := Node{data: x, prev: nil, next: nil}
	if head == nil {
		l.head = &n
	} else {
		for head.next != nil {
			head = head.next
		}
		head.next = &n
		n.prev = head
	}
}

func (l *LinkList) printList() {
	n := l.head
	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}
}

func (l *LinkList) addAtFront(x interface{}) {

	n := Node{data: x, prev: nil, next: nil}
	head := l.head
	l.head = &n
	n.next = head
}
func (l *LinkList) delete(x interface{}) {
	head := l.head
	if head.data == x {
		l.head = head.next
		l.head.prev = nil
	} else {
		for head.next != nil {
			if head.data == x {
				break
			}
			head = head.next
		}

		if head.data == x {
			if head.next != nil {
				head.next.prev = head.prev
				head.prev.next = head.next
			} else if head.next == nil {
				head.prev.next = nil
			}
		} else {
			fmt.Println("Node doesn't exist ")
		}
	}

}

func (l *LinkList) reverse() {
	head := l.head
	var temp *Node
	for head != nil {
		temp = head.prev
		head.prev = head.next
		head.next = temp
		head = head.prev
	}
	if temp != nil {
		l.head = temp.prev
	}
}

func main() {
	l1 := new(LinkList)
	l1.add(1)
	l1.add(2)
	l1.add(3)
	l1.add(4)
	l1.printList()
	fmt.Println("#####")
	// l1.delete(5)
	// l1.printList()
	// fmt.Println("#####")
	// l1.delete(3)
	// l1.printList()
	// fmt.Println("#####")
	// l1.delete(4)
	// l1.printList()
	// fmt.Println("#####")
	// l1.delete(7)
	// l1.printList()
	fmt.Println("#####")
	l1.reverse()
	l1.printList()
	fmt.Println("#####")

}
