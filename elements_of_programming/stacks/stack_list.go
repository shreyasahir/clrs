package main

import (
	"fmt"
)
type node struct {
	data interface{}
	next *node
}

type linklist struct {
	root *node
}

func (l *linklist) push(k interface{}) {
	n := &node{data:k,next:nil}
	if l.root == nil {
		l.root = n
	} else {
		v := l.root
		for v.next != nil {
			v = v.next
		}
		v.next = n
	}
}

func (l *linklist) pop() interface{}{
		v := l.root
		var prev *node
		var x interface{}
		if v == nil {
			panic("Pop from empty stack")
		} else {
			for v.next != nil {
				prev = v
				v = v.next
			}
			x = v.data
			prev.next = nil
		}
		return x
}

func (l * linklist) printList() {
	v := l.root
	if v == nil {
	} else {
		fmt.Println(v.data)
		v = v.next
		for v.next != nil {
			fmt.Println(v.data)
			v = v.next
		}
		fmt.Println(v.data)
	}
}
func main(){
	s := &linklist{}
	s.push(21)
	s.push(26)
	s.push(42)
	s.push(52)
	s.push(62)
	s.push(32)
	s.push(392)
	s.push("hello")
	fmt.Println("Stack contents")
	s.printList()
	fmt.Println("POP",s.pop())
	s.printList()
}