package main

import (
	"fmt"
)

var (
	n         int
	noOfCases int
	count     int
)

func main() {
	fmt.Scanf("%d", &noOfCases)
	for noOfCases > 0 {
		count = 0
		fmt.Scanf("%d", &n)
		for n > 0 {
			n = n & (n - 1)
			//fmt.Println(n)
			count++
		}
		fmt.Println(count)
		noOfCases--
	}

}
