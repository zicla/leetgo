package main

import "fmt"

func majorityElement(nums []int) []int {
	N := len(nums)
	var res []int
	if N == 0 {
		return res
	}

	dict := make(map[int]int)
	for _, v := range nums {
		dict[v]++
	}
	for k, v := range dict {
		if v > N/3 {
			res = append(res, k)
		}
	}

	return res
}

func main() {

	fmt.Printf("%v\n", majorityElement([]int{1, 1, 2, 1, 2, 2}))
}
