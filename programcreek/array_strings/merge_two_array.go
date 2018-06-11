package main

import (
	"fmt"
)

// public class Solution {
//     public void merge(int A[], int m, int B[], int n) {

//         while(m > 0 && n > 0){
//             if(A[m-1] > B[n-1]){
//                 A[m+n-1] = A[m-1];
//                 m--;
//             }else{
//                 A[m+n-1] = B[n-1];
//                 n--;
//             }
//         }

//         while(n > 0){
//             A[m+n-1] = B[n-1];
//             n--;
//         }
//     }
// }

func mergeArray(arr1 []int, m int, arr2 []int, n int) {
	fmt.Println("arr1", arr1)
	fmt.Println("m", m)
	fmt.Println("n", n)

	for m > 0 && n > 0 {
		if arr1[m-1] > arr2[n-1] {
			arr1[m+n-1] = arr2[m-1]
			m--
		} else {
			fmt.Println("m+n-1", m+n-1)
			arr1[m+n-1] = arr2[n-1]
			n--
		}
	}

	for n > 0 {
		arr1[m+n-1] = arr2[n-1]
		n--
	}
	fmt.Println("merged array", arr1)

}
func main() {
	var arr1 = make([]int, 20)
	arr2 := []int{2, 3}
	for k, v := range arr2 {
		arr1[k] = v
	}
	arr3 := []int{1, 4}
	mergeArray(arr1, len(arr1), arr3, len(arr3))
	fmt.Println("merged array", arr1)
}
