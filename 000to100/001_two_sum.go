package main

import "fmt"

func twoSum(nums []int, target int) []int {

	N := len(nums)

	for start := 0; start <= N-2; start++ {

		for end := start + 1; end <= N-1; end++ {

			if nums[start]+nums[end] == target {
				return []int{start, end}
			}
		}
	}
	return []int{-1, -1}
}

func main() {

	nums := []int{2, 7, 11, 15}

	target := 18

	fmt.Printf("%v", twoSum(nums, target))
}
