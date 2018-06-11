package main

import (
	"fmt"
)

func reverse(s string) string {
	var result string
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}

	return result
}
func reverseStr(s string) string {
	var result, subset string
	s = reverse(s)
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			result += subset
			result += " "
			subset = ""
		} else {
			subset = string(s[i]) + subset
		}
	}
	result += subset
	return result
}
func main() {
	s := "the sky is blue"
	fmt.Println(reverseStr(s))
}
