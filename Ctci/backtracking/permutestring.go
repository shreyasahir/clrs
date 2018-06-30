package main

import (
	"fmt"
)

func permute(a []string, l, r int) {

	if l == r {
		fmt.Println(a)
	} else {
		for i := l; i < r+1; i++ {
			a[i], a[l] = a[l], a[i]
			permute(a, l+1, r)
			a[i], a[l] = a[l], a[i]
		}
	}

}
func main() {
	//s := "ABC"
	a := []string{"A", "B", "C"}
	n := len(a)
	permute(a, 0, n-1)
}
