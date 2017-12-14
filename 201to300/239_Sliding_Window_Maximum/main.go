package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {

	N := len(nums)
	var res []int
	if N == 0 {
		return res
	}

	// write your code here
	var myDeq []int
	for i := 0; i < N; i++ {
		if i-k >= 0 {
			// 出队列踢头节点判断
			if len(myDeq) != 0 && nums[i-k] == myDeq[0] {
				myDeq = myDeq[1:]
			}
		}
		// 入队列踢值判断
		for len(myDeq) != 0 && nums[i] > myDeq[len(myDeq)-1] {
			myDeq = myDeq[:len(myDeq)-1]
		}
		myDeq = append(myDeq, nums[i])
		if i+1-k >= 0 {
			// 窗口大小满足判断
			res = append(res, myDeq[0])
		}
	}
	if k > len(nums) {
		res = append(res, myDeq[0])
	}

	return res
}
func main() {

	fmt.Printf("%v\n", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Printf("%v\n", maxSlidingWindow([]int{}, 0))
}
