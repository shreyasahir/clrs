package main

import (
	"fmt"
)

func subSeq(s string) {
	//var arr []string
	//fmt.Println("value of s", s)
	for i := 0; i < len(s); i++ {
		res := ""
		for j := i; j < len(s); j++ {
			res += string(s[j])
			fmt.Println(res)
		}
	}
}
func main() {
	s := "abcd"
	subSeq(s)

}

// def substring(str):
//     for i in range(len(str)):
//         print(str[i:])
//         substring(str[i:i-1])
