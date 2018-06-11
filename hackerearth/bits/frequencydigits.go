package main

import (
	"fmt"
	"strconv"
)

var (
	s    string
	sInt int64
)

func main() {
	freq := make([]int, 10)
	fmt.Scanf("%s", &s)
	for _, v := range s {
		number, _ := strconv.Atoi(string(v))
		freq[number]++
	}
	for k, v := range freq {
		fmt.Println(k, v)
	}
}
