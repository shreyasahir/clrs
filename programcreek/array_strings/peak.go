package main

import (
	"fmt"
)

func findPeak(arr []int, start, end int) int {
	mid := (start + end) / 2
	fmt.Println("mid", mid)
	if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
		return mid
	}
	if arr[mid] < arr[mid-1] {
		return findPeak(arr, start, mid-1)
	}
	if arr[mid] < arr[mid+1] {
		return findPeak(arr, mid+1, end)
	}
	return -1
}

func main() {
	arr := []int{10, 20, 15, 2, 23, 90, 67}
	fmt.Println(findPeak(arr, 0, len(arr)))
}
