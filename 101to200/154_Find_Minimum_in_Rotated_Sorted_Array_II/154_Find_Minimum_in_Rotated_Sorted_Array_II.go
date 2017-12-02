package main

import (
	"fmt"
)

func findMin(nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	//找到突然变小的那个元素。
	curElement := nums[0]

	for i := 1; i < N; i++ {
		if nums[i] < curElement {
			return nums[i]
		}
	}
	return nums[0]
}

func main() {

	fmt.Printf("%v \n", findMin([]int{4, 5, 6, 7, 0, 1, 2}))

}
