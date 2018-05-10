package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arraysize, searchValue, n, i, key int64

	var arr []int64
	var input1, input2 string

	fmt.Scanln(&input1, &input2)
	arraysize, _ = strconv.ParseInt(input1, 10, 64)
	searchValue, _ = strconv.ParseInt(input2, 10, 64)
	for i = 0; i < arraysize; i++ {
		fmt.Scanf("%d", &n)
		arr = append(arr, n)
	}
	key = -1
	for i := arraysize - 1; i > 0; i-- {
		if arr[i] == searchValue {
			key = i
			break
		}
	}
	fmt.Println(key + 1)
}
