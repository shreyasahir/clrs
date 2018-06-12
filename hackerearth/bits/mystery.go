package main

import (
	"fmt"
)

var (
	n int64
)

func main() {
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		//fmt.Println("value of n", n)
		count := 0
		for n > 0 {
			n = n & (n - 1)
			count++
		}
		fmt.Println(count)
	}
}
