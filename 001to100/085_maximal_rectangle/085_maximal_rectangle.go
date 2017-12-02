package main

import "fmt"

func largestRectangleArea1(heights []int) int {

	arr := append(heights, 0)
	N := len(arr)
	stack := make([]int, N)
	topIndex := -1
	area := 0
	for i := 0; i < N; i++ {
		item := arr[i]
		//入栈
		if topIndex == -1 || arr[stack[topIndex]] <= item {
			topIndex++
			stack[topIndex] = i
		} else {

			//出栈。
			top := stack[topIndex]
			topIndex--

			//最后一个元素出栈时，宽度为全部。
			if topIndex == -1 {
				if area < arr[top]*i {
					area = arr[top] * i
				}
			} else {
				if area < arr[top]*(i-stack[topIndex]-1) {
					area = arr[top] * (i - stack[topIndex] - 1)
				}
			}
			i--
		}
	}
	return area
}

func maximalRectangle(matrix [][]byte) int {

	R := len(matrix)
	if R == 0 {
		return 0
	}
	C := len(matrix[0])

	area := 0

	for r := 0; r < R; r++ {
		arr := make([]int, C)
		for n := 0; n < C; n++ {
			for m := r; m >= 0; m-- {
				if matrix[m][n] == '0' {
					break
				} else {
					arr[n]++
				}
			}
		}

		tmp := largestRectangleArea1(arr)
		if tmp > area {
			area = tmp
		}
	}
	return area
}

func main() {

	matrix := [][]byte{{'0'}}

	fmt.Println(maximalRectangle(matrix))

	matrix = [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}

	fmt.Println(maximalRectangle(matrix))

	matrix = [][]byte{{'0', '1'}, {'1', '0'}}

	fmt.Println(maximalRectangle(matrix))

}
