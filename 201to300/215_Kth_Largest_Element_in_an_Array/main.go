package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	N := len(nums)
	sort.Ints(nums)
	return nums[N-k]
}
func main() {

	fmt.Printf("%v \n", findKthLargest([]int{99, 99}, 1))
	fmt.Printf("%v \n", findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
