package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(v1, v2 string) int {
	arr1 := strings.Split(v1, ".")
	arr2 := strings.Split(v2, ".")
	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)

	i := 0
	for i < len(arr1) || len(arr2) > i {
		if i < len(arr1) && len(arr2) > i {
			a, _ := strconv.ParseInt(arr1[i], 10, 32)
			b, _ := strconv.ParseInt(arr2[i], 10, 32)
			fmt.Println("a", a)
			fmt.Println("b", b)
			if a > b {
				return 1
			}
			if a < b {
				return -1
			}
		} else {
			if len(arr1) > i {
				a, _ := strconv.ParseInt(arr1[i], 10, 32)
				if a != 0 {
					return 1
				}
			}
			if len(arr2) > i {
				b, _ := strconv.ParseInt(arr2[i], 10, 32)
				if b != 0 {
					return -1
				}
			}
		}

		i++
	}
	return 0
}

func main() {
	version1 := "1"
	version2 := "1.0"
	fmt.Println("compare", compareVersion(version1, version2))
}
