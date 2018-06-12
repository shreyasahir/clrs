package main

import (
	"fmt"
)

type queue struct {
	data []int
}

func (q *queue) pop() int {

	if len(q.data) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	var x int
	x, q.data = q.data[0], q.data[1:]
	return x
}

func (q *queue) push(x int) {
	q.data = append(q.data, x)
}

func main() {
	s1 := new(queue)
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
