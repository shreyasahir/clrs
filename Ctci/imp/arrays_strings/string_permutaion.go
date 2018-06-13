package main

import (
	"fmt"
)

func perm(str []rune, i int) {
	if i == len(str) {
		fmt.Println(string(str))
	} else {
		for j := i; j < len(str); j++ {
			str[i], str[j] = str[j], str[i]
			perm(str, i+1)
			str[i], str[j] = str[j], str[i]
		}
	}
}

func main() {
	s := "ABC"
	perm([]rune(s), 0)
}
