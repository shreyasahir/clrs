package main

import "fmt"

func main() {
	var array = []int{4, 8, 9, 1, 2, 5, 7, 6, 3, 0}
	var heap = new(Heap)

	fmt.Println(array)
	heap.heapSort(array)
	fmt.Println(array)
}

type Heap struct {
}

func (h *Heap) buildHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		h.heapify(array, i, len(array))
	}
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

func (h *Heap) heapify(array []int, root, length int) {
	var max = root
	var l, r = h.left(array, root), h.right(array, root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		h.heapify(array, max, length)
	}
}

func (*Heap) left(array []int, root int) int {
	return (root * 2) + 1
}

func (*Heap) right(array []int, root int) int {
	return (root * 2) + 2
}
