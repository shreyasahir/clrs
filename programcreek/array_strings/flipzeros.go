// You are given with an array of 1s and 0s. And you are given with an integer m, which signifies number of flips allowed.

// find the position of zeros which when flipped will produce maximum continuous series of 1s.

// e.g.
// input:
// arr={1 1 0 1 1 0 0 1 1 1 } m=1
// output={1 1 1 1 1 0 0 1 1 1} position=2

// arr={1 1 0 1 1 0 0 1 1 1 } m=2
// output={1 1 0 1 1 1 1 1 1 1} position=5,6

// wL = 0; wR = 0;
// nZero = 0;
// bestWindowWidth = -1;

// while(wR < A.length()){
// 	//expand to the right, update '0' count:
// 	if (nZero <= m){
// 		wR++;
// 		if (A[wR] == '0') nZero++;
// 	};

// 	//shrink from left, update '0' count:
// 	if (nZero > m){
// 		if (A[wL] == '0') nZero--;
// 		wL++;
// 	};

// 	//update best window:
// 	if (wR - wL > bestWindowWidth){
// 		bestWindowWidth = wR - wL;
// 		bestWR = wR;
// 		bestWL = wL;
// 	};
// };

package main

import (
	"fmt"
)

func flipZeros(arr []int, m int) {
	var left, right, currZero, bestRight, bestLeft int
	bestWindow := -1

	for right < len(arr) {
		if currZero <= m {
			if arr[right] == 0 {
				currZero++
			}
			right++
		}
		if currZero > m {
			if arr[left] == 0 {
				currZero--
			}
			left++
		}
		if right-left > bestWindow {
			bestWindow = right - left
			bestRight = right
			bestLeft = left
		}
	}
	fmt.Println(bestWindow, bestRight, bestLeft)
}
func main() {
	m := 2
	arr := []int{1, 1, 0, 1, 1, 0, 0, 1, 1, 1}
	flipZeros(arr, m)
}
