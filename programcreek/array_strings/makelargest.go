package main

import (
	"fmt"
	"sort"
	"strconv"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	//return len(s[i]) < len(s[j])
	num1, _ := strconv.ParseInt((s[i] + s[j]), 10, 32)
	num2, _ := strconv.ParseInt((s[j] + s[i]), 10, 32)
	return num1 > num2
}

// type maxNum string

// func (a maxNum) Len() int { return len(a) }
// func (a maxNum) Swap(i, j int) {
// 	fmt.Println("i&j", i, j)
// 	a[i], a[j] = a[j], a[i]
// }
// func (a maxNum) Less(i, j int) bool {
// 	return strconv.ParseInt(a[i], 10, 32) < strconv.ParseInt(a[j], 10, 32)
// }

func longestNumber(s []string) {
	str := ""
	sort.Sort(ByLength(s))
	for _, v := range s {
		str += v
	}
	fmt.Println("str", str)
}
func main() {
	//s := []string{"54", "546", "548", "60"}
	s := []string{"1", "34", "3", "98", "9", "76", "45", "4"}
	longestNumber(s)
}
