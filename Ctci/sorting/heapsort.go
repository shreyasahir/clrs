package main

import (
	"fmt"
)

//Heap datastructure
type Heap struct {
}

func main() {
	var arr = []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
	h := new(Heap)
	fmt.Println("arr before sorting", arr)
	h.heapSort(arr)
	fmt.Println("arr after sorting", arr)
}

func (h *Heap) heapSort(arr []int) {
	h.buildHeap(arr)
	for i := len(arr); i > 1; i-- {
		h.removeTop(arr, i)
	}
}

func (h *Heap) removeTop(array []int, length int) {
	lastIndex := length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	h.heapify(array, 0, lastIndex)
}

func (h *Heap) buildHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		h.heapify(array, i, len(array))
	}
}

func (h *Heap) heapify(arr []int, root int, n int) {
	var max = root
	var l, r = h.left(arr, root), h.right(arr, root)
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

func (h *Heap) left(arr []int, root int) int {
	return (root * 2) + 1
}

func (h *Heap) right(arr []int, root int) int {
	return (root * 2) + 2
}
