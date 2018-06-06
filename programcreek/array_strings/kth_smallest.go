package main

import (
	"fmt"
)

func median(a, b []int) {
	n1 := len(a)
	n2 := len(b)
	if (n1+n2)%2 == 0 {
		fmt.Println("median even", (kth(a, n1, b, n2, (n1+n2)/2)+kth(a, n1, b, n2, ((len(a)+len(b))/2-1)))/2)
	} else {
		fmt.Println("median odd", kth(a, len(a), b, len(b), (n1+n2+1)/2))
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func kth(arr1 []int, n1 int, arr2 []int, n2 int, k int) int {
	if (k > n1+n2) || k < 1 {
		return -1
	}
	if n1 > n2 {
		return kth(arr2, n2, arr1, n1, k)
	}
	if n1 == 0 {
		return arr2[k-1]
	}
	if k == 1 {
		return min(arr1[0], arr2[0])
	}

	i := min(n1, k/2)
	j := min(n2, k/2)

	fmt.Println("arr1[i-1]", arr1[i-1])
	fmt.Println("arr2[j-1]", arr2[j-1])
	fmt.Println("arr2[j-1]", arr2[j:])

	if arr1[i-1] > arr2[j-1] {
		return kth(arr1, n1, arr2[j:], len(arr2[j:]), k-j)
	}
	return kth(arr1[i:], len(arr1[i:]), arr2, n2, k-i)
}

func main() {
	arr1 := []int{2, 3, 6, 7, 9}
	arr2 := []int{1, 4, 8, 10, 11}
	fmt.Println("kth", kth(arr1, len(arr1), arr2, len(arr2), 5))
	median(arr1, arr2)
}
