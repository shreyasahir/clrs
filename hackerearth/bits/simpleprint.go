package main

import (
	"fmt"
)

func main() {

	var i int
	var mystring string
	fmt.Scanf("%d", &i) // Reading input from STDIN
	fmt.Scanf("%s", &mystring)
	fmt.Println(i * 2)
	fmt.Println(mystring)
}
