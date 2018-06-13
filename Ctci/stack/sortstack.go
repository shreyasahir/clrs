package main

import (
	"fmt"
)

type stack struct {
	data []int
}

func (s *stack) push(k int) {
	s.data = append(s.data, k)
}

func (s *stack) pop() int {
	var x int
	x, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return x
}

func (s *stack) peek() int {
	return s.data[len(s.data)-1]
}
func (s *stack) sort() *stack {
	s2 := new(stack)
	for len(s.data) > 0 {
		tmp := s.pop()
		for len(s2.data) > 0 && s2.peek() > tmp {
			s.push(s2.pop())
		}
		s2.push(tmp)
	}
	return s2
}

func main() {
	s := new(stack)
	s.push(1)
	s.push(93)
	s.push(83)
	s.push(53)
	s.push(63)
	s.push(73)
	fmt.Printf("%v", s)
	fmt.Printf("%v", s.sort())
}
