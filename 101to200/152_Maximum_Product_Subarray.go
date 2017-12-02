package main

import (
	"fmt"
)

func maxProduct(nums []int) int {

	N := len(nums)

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	//tempMin 表示带上 i 的最小。 tempMax表示带上 i 的最大
	tempMin := nums[0]
	tempMax := nums[0]
	finalMax := tempMax
	for i := 1; i < N; i++ {

		preTempMin := tempMin
		preTempMax := tempMax
		//因为 i-1一定参与了战斗，要么是贡献了
		tempMin = min(min(preTempMin*nums[i], preTempMax*nums[i]), nums[i])
		tempMax = max(max(preTempMin*nums[i], preTempMax*nums[i]), nums[i])
		finalMax = max(finalMax, tempMax)
	}

	return finalMax
}
func main() {

	fmt.Printf("%v \n", maxProduct([]int{2, 3, -2, 4}))

}
