package main

import (
	"fmt"
)

func singleNumber137(nums []int) int {

	N := len(nums)
	var box = make(map[int]int)
	for i := 0; i < N; i++ {
		box[nums[i]]++
	}
	for k, v := range box {
		if v == 1 {
			return k
		}
	}

	if N == 0 {
		return 0
	} else {
		return nums[0]
	}

}

func main() {

	a := 2 ^ 2
	fmt.Println(a)

	fmt.Printf("%v\n", singleNumber137([]int{1, 2, 2}))
	fmt.Printf("%v\n", singleNumber137([]int{2, 2, 3, 3, 5}))
	fmt.Printf("%v\n", singleNumber137([]int{1, 2, 2, 10, 1}))
}
