package main

import "fmt"

func findPeakElement(nums []int) int {

	N := len(nums)
	if N == 1 {
		return 0
	}
	for i := 0; i < N; i++ {

		if i == 0 {
			if nums[0] > nums[1] {
				return 0
			}
		} else if i == N-1 {
			if nums[N-1] > nums[N-2] {
				return N - 1
			}
		} else {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))

}
