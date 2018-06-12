package main

import (
	"fmt"
)

type heap struct {
}

func main() {
	var array = []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
	h := new(heap)
	fmt.Println("array before sort", array)
	h.heapSort(array)
	fmt.Println("array after sort", array)

}

func (h *heap) heapSort(arr []int) {
	h.buidlHeap(arr)
	for i := len(arr); i > 1; i-- {
		h.removeTop(arr, i, len(arr))
	}
}

func (h *heap) removeTop(arr []int, i, n int) {
	lastIndex := i - 1
	arr[0], arr[lastIndex] = arr[lastIndex], arr[0]
	h.heapify(arr, 0, lastIndex)
}
func (h *heap) buidlHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		h.heapify(arr, i, len(arr))
	}
}

func (h *heap) heapify(arr []int, root, n int) {
	var max = root
	var l, r = h.left(arr, max), h.right(arr, max)
	if l < n && arr[l] > arr[max] {
		max = l
	}

	if r < n && arr[r] > arr[max] {
		max = r
	}

	if max != root {
		arr[root], arr[max] = arr[max], arr[root]
		h.heapify(arr, max, n)
	}
}

func (h *heap) left(arr []int, root int) int {
	return (root * 2) + 1
}

func (h *heap) right(arr []int, root int) int {
	return (root * 2) + 2
}
