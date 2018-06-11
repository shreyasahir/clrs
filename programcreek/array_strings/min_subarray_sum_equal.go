package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// int minSubArrayLen(int s, vector<int>& nums)
// {
//     int n = nums.size();
//     int ans = INT_MAX;
//     int left = 0;
//     int sum = 0;
//     for (int i = 0; i < n; i++) {
//         sum += nums[i];
//         while (sum >= s) {
//             ans = min(ans, i + 1 - left);
//             sum -= nums[left++];
//         }
//     }
//     return (ans != INT_MAX) ? ans : 0;
// }

func minSubArray(arr []int, n int) {
	fmt.Println("arr is ", arr)
	j := 0
	sum := 0
	ans := 1000
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		fmt.Println("summ is ", sum)
		for sum >= n {
			fmt.Println("sum is", sum, "ans is", ans)
			ans = min(ans, i+1-j)
			sum -= arr[j]
			j++
		}
	}
	fmt.Println("ans", ans)

}
func main() {

	arr := []int{2, 3, 1, 2, 4, 3}
	n := 7
	minSubArray(arr, n)
}
