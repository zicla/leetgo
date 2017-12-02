package main

import "fmt"

func generate(numRows int) [][]int {
	var res = make([][]int, numRows)
	if numRows == 0 {
		return res
	}

	res[0] = []int{1}
	for i := 2; i <= numRows; i++ {

		preRow := res[i-2]
		row := make([]int, i)
		row[0] = 1
		row[i-1] = 1
		for j := 1; j < i-1; j++ {
			row[j] = preRow[j-1] + preRow[j]
		}
		res[i-1] = row
	}

	return res
}


func main() {

	fmt.Printf("%v\n", generate(0))
	fmt.Printf("%v\n", generate(1))
	fmt.Printf("%v\n", generate(2))
	fmt.Printf("%v\n", generate(3))
	fmt.Printf("%v\n", generate(10))
}
