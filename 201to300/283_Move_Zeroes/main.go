package main

import "fmt"

func moveZeroes(nums []int) {

	N := len(nums)
	if N == 0 {
		return
	}

	for i := 1; i < N; i++ {
		//不是0的数往前浮。
		if nums[i] != 0 {

			k := i - 1
			for k >= 0 && nums[k] == 0 {
				k--
			}
			k++
			nums[i], nums[k] = nums[k], nums[i]

		}
	}
}
func main() {

	arr := []int{0, 1, 0, 15, 12}
	moveZeroes(arr)
	fmt.Printf("%v\n", arr)
}
