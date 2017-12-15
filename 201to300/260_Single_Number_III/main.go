package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {

	//全部异或起来就是 a^b 重要结论，0和任何数异或均为自己。
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor = xor ^ nums[i]
	}

	//找出a,b不相同的最低位。负数的表示是各位取反加1.
	xor = xor & (-xor)
	//此时借这一位分为两组。
	var res = []int{0, 0}
	for i := 0; i < len(nums); i++ {
		if xor&nums[i] > 0 {
			res[0] = res[0] ^ nums[i]
		} else {
			res[1] = res[1] ^ nums[i]
		}
	}
	return res
}
func main() {

	fmt.Printf("%v\n", singleNumber([]int{-1, 0}))
}
