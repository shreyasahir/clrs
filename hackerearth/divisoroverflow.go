package main

import (
	"fmt"
)

var (
	n         int64
	noOfCases int
	count     int
	//binaryReprentation string
)

func main() {
	fmt.Scanf("%d", &noOfCases)
	for noOfCases > 0 {
		count = 0
		fmt.Scanf("%d", &n)
		//binaryReprentation := strconv.FormatInt(n, 2)
		//fmt.Sscanf(string(n), "%b", &binaryReprentation)
		binaryReprentation := fmt.Sprintf("%b", n)
		//fmt.Println("binary", binaryReprentation)
		noOfBits := len(string(binaryReprentation))
		//fmt.Println("string", noOfBits)
		for i := int64(1); i <= n; i++ {
			quotient := n / i
			QuotBinaryReprentation := fmt.Sprintf("%b", quotient)
			qnoOfBits := len(string(QuotBinaryReprentation))
			if noOfBits > qnoOfBits {
				count++
			}
		}
		fmt.Println(count)
		noOfCases--
	}

}
