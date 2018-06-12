package main

import (
	"fmt"
)

type stack struct {
	data []int
}

func (s *stack) pop() int {

	if len(s.data) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	var x int
	x, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return x
}

func (s *stack) push(x int) {
	s.data = append(s.data, x)
}

func main() {
	s1 := new(stack)
	s1.push(1)
	s1.push(45)
	s1.push(34)
	s1.push(90)
	fmt.Println(s1)
	fmt.Println(s1.pop())
	fmt.Println(s1)
	fmt.Println(s1.pop())
	fmt.Println(s1)
	fmt.Println(s1.pop())

}
