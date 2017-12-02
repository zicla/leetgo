package main

import "fmt"

func subsetsWithDup(nums []int) [][]int {

	//分类
	numMap := make(map[int]int)
	var keys []int
	for _, v := range nums {
		numMap[v]++
		if numMap[v] == 1 {
			keys = append(keys, v)
		}
	}

	var arr [][]int
	arr = append(arr, []int{})
	for _, v := range keys {

		arr = subsetsWithDupMethod(arr, v, numMap[v])

	}

	return arr
}

func subsetsWithDupMethod(arr [][]int, key int, num int) [][]int {

	N := len(arr)
	for i := 0; i < N; i++ {
		row := arr[i]
		rowLen := len(row)
		for j := 1; j <= num; j++ {
			row1 := make([]int, rowLen+j, rowLen+j)
			for k := 0; k < rowLen; k++ {
				row1[k] = row[k]
			}
			for n := 0; n < j; n++ {
				row1[rowLen+n] = key
			}
			arr = append(arr, row1)
		}
	}
	return arr
}

func main() {

	a := []int{1, 2, 2}
	fmt.Printf("%v\n", subsetsWithDup(a))

	b := []int{1, 0, 3, 4, 5}
	fmt.Printf("%v\n", subsetsWithDup(b))

	c := []int{9, 0, 3, 5, 7}
	fmt.Printf("%v\n", subsetsWithDup(c))

}
