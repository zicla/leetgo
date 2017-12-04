package main

import "fmt"

func rob(nums []int) int {

	N := len(nums)
	if N == 0 {
		return 0
	} else if N == 1 {
		return nums[0]
	} else if N == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	//dp[i]表示，第i间被盗最大收益。
	dp := make([]int, N)
	dp[0] = nums[0]
	dp[1] = nums[1]
	//f[i]表示，第i间不抢的最大收益
	f := make([]int, N)
	f[0] = 0
	f[1] = nums[0]
	for i := 2; i < N; i++ {
		dp[i] = max(dp[i-2]+nums[i], f[i-2]+nums[i])
		f[i] = max(dp[i-1], f[i-1])
	}
	return max(f[N-1], dp[N-1])
}
func main() {

	fmt.Printf("%v \n", rob([]int{2, 5, 8, 1}))
}
