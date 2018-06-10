package main

func arrIntersection(arr1, arr2 []int) {
	arrHash1 := make(map[int]bool, len(arr1))
	arrHash2 := make(map[int]bool, len(arr2))

	for i := 0; i < len(arr1); i++ {
		arrHash1[arr1[i]] = 1
	}

	for i := 0; i < len(arr2); i++ {
		arrHash2[arr2[i]] = 1
	}
}
func main() {
	arr1 := []int{1, 4, 7, 9, 12}
	arr2 := []int{7, 9, 12}
	arrIntersection(arr1, arr2)
}
