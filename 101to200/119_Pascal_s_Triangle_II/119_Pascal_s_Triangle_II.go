package main

import "fmt"

func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	var res = make([][]int, numRows)
	if numRows == 0 {
	}else{
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
	}
	return res[rowIndex]
}

func main() {

	fmt.Printf("%v\n", getRow(0))
	fmt.Printf("%v\n", getRow(1))
	fmt.Printf("%v\n", getRow(2))
	fmt.Printf("%v\n", getRow(3))
	fmt.Printf("%v\n", getRow(10))
}
