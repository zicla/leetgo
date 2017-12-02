package main

import (
	"fmt"
)

//重点参考：https://discuss.leetcode.com/topic/22821/an-general-way-to-handle-all-this-sort-of-questions?page=1
func singleNumber(nums []int) int {
	a := 0
	b := 0
	for _, c := range nums {
		ta := (^a & b & c) | (a & ^b & ^c)
		b = (^a & ^b & c) | (^a & b & ^c)
		a = ta
	}
	//we need find the number that is 01,10 => 1, 00 => 0.
	return b
}

func main() {

	fmt.Printf("%v\n", singleNumber([]int{1, 2, 2, 2}))
	fmt.Printf("%v\n", singleNumber([]int{2, 2, 3, 3, 2, 3, 5}))
	fmt.Printf("%v\n", singleNumber([]int{1, 2, 2, 1, 2, 10, 1}))
}
