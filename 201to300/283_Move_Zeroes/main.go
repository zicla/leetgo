package main

import "fmt"

func moveZeroes(nums []int) {

	N := len(nums)
	if N == 0 {
		return
	}

	lastNonZero := 0
	for i := 0; i < N; i++ {
		//不是0的数往前浮。
		if nums[i] != 0 {
			nums[i], nums[lastNonZero] = nums[lastNonZero], nums[i]
			lastNonZero++
		}
	}
}
func main() {

	arr := []int{0, 1, 0, 15, 12}
	moveZeroes(arr)
	fmt.Printf("%v\n", arr)
}
