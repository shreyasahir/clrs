package main

import (
	"fmt"
)

func isParanthesis(s string) bool {
	if s == "(" {
		return true
	} else if s == ")" {
		return true
	} else {
		return false
	}
}

func isValidString(s string) bool {
	var count int
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			count++
		}
		if string(s[i]) == ")" {
			count--
		}
		if count < 0 {
			return false
		}
	}

	if count == 0 {
		return true
	}
	return false
}
func removeInvalidParanthesis(s string) {
	visited := make([]string, len(s))
	var level bool
	queue := []string{}
	queue = append(queue, s)
	visited = append(visited, s)
	for len(queue) > 0 {
		str := queue[0]
		queue = queue[:len(queue)-1]

		if isValidString(str) {
			fmt.Println(str)
			level = true
		}

		if level {
			continue
		}

		for i := 0; i < len(s); i++ {
			if !isParanthesis(string(s[i])) {
				continue
			}
		}
	}
}

func main() {
	s := "()())()"
	removeInvalidParanthesis(s)
}
