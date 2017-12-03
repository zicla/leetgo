package main

import "fmt"

func twoSum(numbers []int, target int) []int {

	var dict = make(map[int]int)
	for k, v := range numbers {

		if index, ok := dict[v]; ok {
			if v*2 == target {
				return []int{index, k + 1}
			}
		}

		dict[v] = k + 1

	}

	for _, v := range numbers {
		if index1, ok := dict[v]; ok {
			if index2, ok := dict[target-v]; ok {
				return []int{index1, index2}
			}
		}
	}
	return []int{}
}

func main() {

	fmt.Printf("%v \n", twoSum([]int{0, 0, 3, 4}, 0))
	fmt.Printf("%v \n", twoSum([]int{2, 7, 11, 15}, 9))

}
