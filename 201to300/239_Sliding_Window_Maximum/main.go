package main

import (
	"math"
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {

	N := len(nums)
	var res []int
	if N == 0 {
		return res
	}
	for i := 0; i <= N-k; i++ {
		max := math.MinInt64
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}

	return res
}
func main() {

	fmt.Printf("%v\n", maxSlidingWindow([]int{}, 0))
	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
