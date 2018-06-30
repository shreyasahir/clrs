package main

import ()
import "fmt"

func checkParanthesis(s []string) bool {
	var stack []string

	for _, v := range s {
		if v == "[" || v == "(" || v == "{" {
			stack = append(stack, v)
		} else {
			var u string
			u, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if u == "(" && v != ")" || u == "[" && v != "]" || u == "{" && v != "}" {
				fmt.Println("u & v", u, v)
				return false
			}
		}
	}
	return true

}

func main() {
	s := []string{"[", "(", ")", "]", "{", "}", "{", "[", "(", ")", "(", ")", "]", "(", ")", "}"}
	fmt.Println("paranthesis is valid", checkParanthesis(s))
}
