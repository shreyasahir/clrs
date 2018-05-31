package main

import (
	"fmt"
)

func reverse(s string) string {
	//var result string
	length := len(s)
	if length == 1 {
		return s
	}
	return string(s[length-1]) + reverse(s[:length-1])
}

func main() {
	input := "geeksforgeeks"
	fmt.Println(reverse(input))
}
