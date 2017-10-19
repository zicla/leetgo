package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func largestRectangleArea(heights []int) int {

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

func main() {

	matrix := []int{2, 1, 5, 6, 2, 3}

	fmt.Println(largestRectangleArea(matrix))

	matrix = []int{1, 1}

	fmt.Println(largestRectangleArea(matrix))

	matrix = []int{0, 9}

	fmt.Println(largestRectangleArea(matrix))

	matrix = []int{2, 1, 2}

	fmt.Println(largestRectangleArea(matrix))

	matrix = []int{5, 4, 1, 2}

	fmt.Println(largestRectangleArea(matrix))

	matrix = []int{4, 2, 0, 3, 2, 5}

	fmt.Println(largestRectangleArea(matrix))

	matrix = []int{3, 6, 5, 7, 4, 8, 1, 0}

	fmt.Println(largestRectangleArea(matrix))

}
