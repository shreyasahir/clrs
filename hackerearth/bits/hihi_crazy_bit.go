package main

import (
	"fmt"
	"strconv"
)

var (
	n         int64
	noOfCases int64
	count     int64
)

func main() {
	fmt.Scanf("%d", &noOfCases)
	var flipInt int
	for noOfCases > 0 {
		count = 0
		fmt.Scanf("%d", &n)
		nbin := fmt.Sprintf("%064b", n)
		fmt.Println("value of nbin", nbin)
		nInt, _ := strconv.ParseInt(nbin, 10, 64)
		fmt.Println(nInt)
		iter := len(nbin)
		for {
			y := nInt % 10
			y ~= y
			iter--
			nInt = nInt / 10
			fmt.Println("iter", iter)
			fmt.Println("nbin[:iter]", nbin[:iter])
			fmt.Println("nbin[iter:]", nbin[iter:])
			fmt.Println("strconv.FormatInt(y, 10)", strconv.FormatInt(y, 10))
			nIntString := ""
			if iter != 63 {
				nIntString = nbin[:iter] + strconv.FormatInt(y, 10) + nbin[iter+1:len(nbin)]
			}
			//fmt.Println("string(y)", string(y))
			fmt.Println("nIntString", nIntString)
			// if iter >= 0 {
			// 	nIntString := nbin[:iter] + string(y)
			// } else {
			// 	nIntString := nbin[:iter] + string(y)
			// }
			flipInt, _ := strconv.ParseInt(nIntString, 2, 64)
			fmt.Println("flipInt", flipInt)

			if flipInt > n {
				break
			}
		}
		// for i := len(nbin) - 1; i >= 0; i-- {
		// 	y := ^nbin[i]
		// 	fmt.Println("value of nbin[i]", nbin[i])
		// 	fmt.Println("value of v", int(y-'0'))
		// 	flipBit := nbin[:i] + string(int(y-'0')) + nbin[i+1:]
		// 	fmt.Println("flipbit", flipBit)
		// 	fmt.Println("nbin[:k]", nbin[:i])
		// 	fmt.Println("nbin[k+1:]", nbin[i+1:])
		// 	flipInt, _ := strconv.ParseInt(flipBit, 2, 64)
		// 	fmt.Println("flipint", flipInt)

		// 	if flipInt > n {
		// 		break
		// 	}
		// }
		fmt.Println(flipInt)
		noOfCases--
	}

}
