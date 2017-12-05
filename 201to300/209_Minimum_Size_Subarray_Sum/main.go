package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {

	N := len(nums)
	if N == 0 {
		return 0
	} else if N == 1 {
		if nums[0] >= s {
			return 1
		} else {
			return 0
		}
	}

	if nums[0] >= s {
		return 1
	}

	//local[i]表示包含了nums[i]的满足条件的最短length
	local := make([]int, N)
	//localSum是辅助local的字段，存放的是局部的和。
	localSum := make([]int, N)

	res := math.MaxInt64

	for i := 1; i < N; i++ {

		if local[i-1] == 0 {
			sum := 0
			var j int
			for j = i; j >= 0; j-- {
				sum += nums[j]
				if sum >= s {
					local[i] = i - j + 1
					localSum[i] = sum
					break
				}
			}
			if j == -1 {
				local[i] = 0
				localSum[i] = 0
			}
		} else {
			//看看第i个元素能够顶替几个。
			localSum[i] = localSum[i-1] + nums[i]
			for j := i - local[i-1]; j <= i; j++ {
				if localSum[i]-nums[j] >= s {
					localSum[i] -= nums[j]
				} else {
					local[i] = i - j + 1
					break
				}
			}
		}
		if local[i] != 0 && local[i] < res {
			res = local[i]
		}

	}
	if res == math.MaxInt64 {
		return 0
	} else {
		return res
	}
}

func main() {

	fmt.Printf("%v \n", minSubArrayLen(6, []int{10, 2, 3}))
	fmt.Printf("%v \n", minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}
