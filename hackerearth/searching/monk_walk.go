package main

import (
	"fmt"
)

func main() {
	var noOfCases, i int8
	var count int32
	var inputString string
	var vowelList = map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1}
	fmt.Scanf("%d", &noOfCases)
	for i = 0; i < noOfCases; i++ {
		count = 0
		fmt.Scanf("%s", &inputString)
		inputByte := []byte(inputString)
		//fmt.Println("inputString", inputString)
		//fmt.Println("inputByte", inputByte)
		for _, v := range inputByte {
			//fmt.Println("value of v", v)
			_, ok := vowelList[v]
			if ok {
				count++
			}
		}

		fmt.Println(count)
	}
}
