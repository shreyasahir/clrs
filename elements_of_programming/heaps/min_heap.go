package main

import (
	"fmt"
)

type heap struct {
}

func (h *heap) left(root int) int {
	return (root * 2) + 1
}

func (h *heap) right(root int) int {
	return (root * 2) + 2
}

func (h *heap) heapify(arr []int, i, length int) {
	var max = i
	var left, right = h.left(i), h.right(i)
	if left < length && arr[left] > arr[max] {
		max = left
	}
	if right < length && arr[right] > arr[max] {
		max = right
	}

	if max != i {
		arr[max], arr[i] = arr[i], arr[max]
		h.heapify(arr, max, length)
	}
}
func (h *heap) buildHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		h.heapify(arr, i, len(arr))
	}
}

func (h *heap) removeTop(array []int, length int) {
	lastIndex := length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	h.heapify(array, 0, lastIndex)
}

func (h *heap) heapSort(arr []int) {
	h.buildHeap(arr)
	for i := len(arr); i > 1; i-- {
		h.removeTop(arr, i)
	}
}

func main() {
	var array = []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
	var heap = new(heap)
	heap.buildHeap(array)
	fmt.Println("heapify arr", array)
	heap.heapSort(array)
	fmt.Println(array)
}
