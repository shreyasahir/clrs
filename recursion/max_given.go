// Given an array of integers, a number and a maximum value, task is to compute the maximum
// value that can be obtained from the array elements. Every value on the array traversing from the
// beginning can be either added to or subtracted from the result obtained from previous index such that
// at any point the result is not less than 0 and not greater than the given maximum value. For index 0 take previous
// result equal to given number.
//  In case of no possible answer print -1.
// Input : arr[] = {2, 1, 7}
//         Number = 3
//         Maximum value = 7
// Output : 7
// The order of addition and subtraction
// is: 3(given number) - 2(arr[0]) -
// 1(arr[1]) + 7(arr[2]).

// Input : arr[] = {3, 10, 6, 4, 5}
//         Number = 1
//         Maximum value = 15
// Output : 9
// The order of addition and subtraction
// is: 1 + 3 + 10 - 6 - 4 + 5
package main

func findMax(arr []int, n, m int) {
	var result, currSum int
	if 
}
func main() {
	arr := []int{2, 1, 7}
	n := 3
	m := 7
}
