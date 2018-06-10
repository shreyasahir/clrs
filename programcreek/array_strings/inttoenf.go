package main

import (
	"fmt"
)

var (
	intDict = make(map[int]string, 10)
)

func intToEng(n int) {
	var str string
	for n > 0 {
		if n > 100 {
			str += intDict[n/100] + "hundred"
			n = n % 100
		} else if n > 10 {
			temp := n % 10
			str += intDict[n-temp]
			n = temp
		} else if n < 10 {
			str += intDict[n]
			n = 0
		}
	}
	fmt.Println("str", str)
}
func main() {
	n := 123
	intDict[20] = "twentty"
	intDict[30] = "thirty"
	intDict[40] = "forty"
	intDict[50] = "fifty"
	intDict[60] = "sixty"
	intDict[70] = "seventy"
	intDict[80] = "eighty"
	intDict[90] = "ninty"
	intDict[1] = "One"
	intDict[2] = "two"
	intDict[3] = "three"
	intDict[4] = "four"
	intDict[5] = "five"
	intDict[6] = "six"
	intDict[7] = "seven"
	intDict[8] = "eight"
	intDict[9] = "nine"
	intToEng(n)
}
