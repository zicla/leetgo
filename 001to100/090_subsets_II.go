package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {

	//分类
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}



	fmt.Printf("%v\n", numMap)

	return nil
}

func main() {

	a := []int{1, 2, 2}
	fmt.Printf("%v\n", subsetsWithDup(a))

	b := []int{1, 2, 3}
	fmt.Printf("%v\n", subsetsWithDup(b))

	var result [][]int
	result = append(result, []int{1, 2, 3})
	fmt.Printf("%v\n", result)
}
