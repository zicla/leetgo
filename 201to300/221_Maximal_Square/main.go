package main

import (
	"math"
	"fmt"
)

func maximalSquare(matrix [][]byte) int {

	R := len(matrix)
	if R == 0 {
		return 0
	}
	C := len(matrix[0])
	if C == 0 {
		return 0
	}

	//squre[i][j]表示以i,j为正方形右下角点，最大的正方形边长。
	var square = make([][]int, R)
	for i := 0; i < R; i++ {
		square[i] = make([]int, C)
	}
	res := math.MinInt64
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {

			if i == 0 {
				if matrix[i][j] == '1' {
					square[i][j] = 1
				} else {
					square[i][j] = 0
				}
			} else if j == 0 {
				if matrix[i][j] == '1' {
					square[i][j] = 1
				} else {
					square[i][j] = 0
				}
			} else {
				if matrix[i][j] == '1' {
					left := square[i][j-1]
					leftTop := square[i-1][j-1]
					top := square[i-1][j]

					min := left
					if top < min {
						min = top
					}

					if min >= leftTop {
						square[i][j] = leftTop + 1
					} else {
						square[i][j] = min + 1
					}

				} else {
					square[i][j] = 0
				}
			}

			if square[i][j] > res {
				res = square[i][j]
			}

		}
	}
	return res * res
}

func main() {

	fmt.Printf("%v \n", maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}))
		fmt.Printf("%v \n", maximalSquare([][]byte{
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'}}))

}
