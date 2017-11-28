package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	var exist = make(map[int]bool)
	var left = make(map[int]int)
	var right = make(map[int]int)
	N := len(nums)
	if N == 0 {
		return 0
	}

	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = 1
	}

	exist[nums[0]] = true
	for i := 1; i < N; i++ {
		item := nums[i]

		length := 1
		if exist[item-1] && exist[item+1] {

			length = left[item-1] + right[item+1] + 3

			left[item] = left[item-1] + 1
			right[item] = right[item+1] + 1

			for k := item - 1; k >= (item - left[item]); k-- {
				right[k] = (item - 1) - k + right[item] + 1
			}

			for k := item + 1; k <= (item + right[item]); k++ {
				left[k] = k - (item + 1) + left[item] + 1
			}

		} else if exist[item-1] {

			length = left[item-1] + 2

			left[item] = left[item-1] + 1
			right[item] = 0

			for k := item - 1; k >= (item - left[item]); k-- {
				right[k] = (item - 1) - k + right[item] + 1
			}

		} else if exist[item+1] {

			length = right[item+1] + 2

			left[item] = 0
			right[item] = right[item+1] + 1

			for k := item + 1; k <= (item + right[item]); k++ {
				left[k] = k - (item + 1) + left[item] + 1
			}

		} else {
			length = 1
			left[item] = 0
			right[item] = 0
			right[item-1] = 1
			left[item+1] = 1
		}

		if length > dp[i-1] {
			dp[i] = length
		} else {
			dp[i] = dp[i-1]
		}

		exist[item] = true
	}

	return dp[N-1]
}

func main() {

	fmt.Printf("%v\n", longestConsecutive([]int{1, 4, 8, 3, 6, 9, 5, 7, 1, 2}))
	fmt.Printf("%v\n", longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Printf("%v\n", longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Printf("%v\n", longestConsecutive([]int{100, 4, 5, 1, 3, 2}))
	fmt.Printf("%v\n", longestConsecutive([]int{}))

}
