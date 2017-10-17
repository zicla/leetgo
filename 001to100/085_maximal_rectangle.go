package main

import "fmt"

func maximalRectangle(matrix [][]byte) int {

	R := len(matrix)
	if R == 0 {
		return 0
	}
	C := len(matrix[0])

	area := 0

	//TODO: here.
	return area
}

func main() {

	matrix := [][]byte{{'0'}}

	fmt.Println(maximalRectangle(matrix))

	matrix = [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}

	fmt.Println(maximalRectangle(matrix))

}
