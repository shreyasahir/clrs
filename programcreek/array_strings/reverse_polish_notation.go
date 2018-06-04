package main

import (
	"fmt"
	"strconv"
)

var (
	operators = []string{"+", "-", "*", "/"}
	stack     []string
)

func isOperator(s string) bool {
	for _, v := range operators {
		if s == v {
			return true
		}
	}
	return false
}
func polish(s []string) {
	var x, y string
	var result int64
	for _, v := range s {
		fmt.Println("stack is", stack)
		if isOperator(v) {
			x, y = stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch v {
			case "+":
				x1, _ := strconv.ParseInt(x, 10, 32)
				y1, _ := strconv.ParseInt(y, 10, 32)
				result = x1 + y1
			case "-":
				x1, _ := strconv.ParseInt(x, 10, 32)
				y1, _ := strconv.ParseInt(y, 10, 32)
				result = x1 - y1
			case "*":
				x1, _ := strconv.ParseInt(x, 10, 32)
				y1, _ := strconv.ParseInt(y, 10, 32)
				result = x1 * y1
			case "/":
				x1, _ := strconv.ParseInt(x, 10, 32)
				y1, _ := strconv.ParseInt(y, 10, 32)
				result = x1 / y1
				fmt.Println("result in division", result)
			}
			fmt.Println("result", result)
			stack = append(stack, strconv.FormatInt(result, 10))
			fmt.Println("stack after", stack)

		} else {
			stack = append(stack, v)
		}
	}

	fmt.Println("result is ", stack[0])
}
func main() {
	//s := []string{"2", "1", "+", "3", "*"}
	s := []string{"4", "13", "5", "/", "+"}
	polish(s)
}
