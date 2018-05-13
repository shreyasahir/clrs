package main

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) pop() interface{} {
	n := len(s.data) - 1
	var x interface{}
	x, s.data = s.data[n], s.data[:n]
	return x
}

func (s *Stack) push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *Stack) isEmpty() bool {
	return len(s.data) > 0
}

func (s *Stack) getMin() interface{} {

}

func main() {

	s1 := new(Stack)
	s1.push(1)
	s1.push(45)
	s1.push(55)
	s1.push(65)
	fmt.Println("Stack:", s1)
	//fmt.Printf("%v", s1)
	s1.pop()
}
