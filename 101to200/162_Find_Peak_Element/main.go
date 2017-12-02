package main

import "fmt"

func findPeakElement(nums []int) int {

	N := len(nums)
	if N == 1 {
		return 0
	}

	left := 0
	right := N - 1

	for left <= right {

		if left == right {
			return left
		}

		mid := (left + right) / 2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}

	}

	return -1
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))

}
