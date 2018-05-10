package main

import (
	"fmt"
)

func binToDec(s string) int64 {
	var decValue int64
	var base int64 = 1
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		if int(s[i]-'0') == 1 {
			decValue += base
		}
		base *= 2
	}
	return decValue
}

func main() {
	noOfCases := 0
	fmt.Scanf("%d", &noOfCases)
	for noOfCases > 0 {
		count := 0
		var n int64
		var s string
		fmt.Scanf("%d", &n)
		fmt.Scanf("%s", &s)
		var i int64
		for i < n {
			// s = s[1:] + string(s[0])
			// y := binToDec(s)
			// if y&1 == 0 {
			// 	count++
			// }
			if int(s[i]-'0') == 0 {
				count++
			}
			i++
		}
		fmt.Println(count)
		noOfCases--
	}

}
