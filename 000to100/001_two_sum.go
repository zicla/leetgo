package main

import "fmt"

func twoSum(nums []int, target int) []int {

	N := len(nums)

	dict := make(map[int]int)
	for i := 0; i < N; i++ {

		//解决两个相同数字组成的。
		old, ok := dict[nums[i]]
		if ok && target%2 == 0 && target/2 == nums[i] {
			return []int{old, i}
		}

		dict[nums[i]] = i
	}

	for start := 0; start < N; start++ {
		positive := target - nums[start]
		if positive != nums[start] {
			i, ok := dict[positive]
			if ok {
				return []int{start, i}
			}
		}
	}
	return []int{-1, -1}
}

func main() {

	nums := []int{2, 7, 11, 15}

	target := 22

	fmt.Printf("%v", twoSum(nums, target))
}
