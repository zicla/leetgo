package main

import "fmt"

func largestRectangleArea(heights []int) int {

	N := len(heights)

	return N
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

}
