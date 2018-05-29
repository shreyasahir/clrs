package main

import(
	"fmt"
)
type stack struct {
	value []interface{}
	max []interface{}
}
func (s * stack) push(k interface{}) {
	s.value = append(s.value,k)
	s.stackMax(k,"add")

}

func (s *stack) stackMax(k interface{}, action string) interface{} {
	if action == "add" {
		if len(s.max) == 0 {
			s.max = append(s.max,k)
		} else if s.max[len(s.max)-1].(int) < k.(int) {
			s.max = append(s.max,k)
		} else {
			s.max = append(s.max,s.max[len(s.max)-1])
		}
	} else {
		s.max = s.max[:len(s.max)-1]
	}
	return s.max[len(s.max)-1]
}
func (s * stack) pop() interface{}{
	if len(s.value) == 0 {
		panic("Pop from empty stack")
	} 
	var x interface{}
	x, s.value = s.value[len(s.value)-1], s.value[:len(s.value)-1]
	s.stackMax(x,"remove")
	return x
}

func main() {
	s := &stack{}
	s.push(21)
	fmt.Println("max",s.max)
	s.push(405)
	s.push(42)
	s.push(52)
	s.push(62)
	s.push(32)
	fmt.Println("POP",s.pop())
	fmt.Println("S after POP",s.value)
	s.push(392)
	fmt.Println("S after POP",s.value)
	//s.push("hello")
	fmt.Println("S after POP",s.value)
	fmt.Println("max",s.max)
	fmt.Println("POP",s.pop())
	fmt.Println("S after POP",s.value)
		fmt.Println("POP",s.pop())
	fmt.Println("S after POP",s.value)
		fmt.Println("POP",s.pop())
	fmt.Println("S after POP",s.value)
		fmt.Println("POP",s.pop())
	fmt.Println("S after POP",s.value)
		fmt.Println("max",s.max)

}