package main

import (
	"fmt"
	. "leetgo/structure/btree"
	"math"
)

func sumNumbers(root *TreeNode) int {

	if root == nil {
		return 0
	}

	path := make([]int, math.MaxInt16)
	total := 0
	sumNumbersRecursive(root, &total, path, 0)
	return total
}

func sumNumbersRecursiveCal(path []int, pathLength int) int {
	s := 0
	for i := pathLength - 1; i >= 0; i-- {
		s += path[i] * int(math.Pow10(pathLength-1-i))

	}
	return s
}

func sumNumbersRecursive(root *TreeNode, total *int, path []int, pathLength int) {
	path[pathLength] = root.Val
	pathLength++
	if root.Left == nil && root.Right == nil {

		//结算
		for i := 0; i < pathLength; i++ {
			fmt.Printf("%v ", path[i])
		}
		fmt.Println()

		*total += sumNumbersRecursiveCal(path, pathLength)

	} else if root.Left == nil {
		sumNumbersRecursive(root.Right, total, path, pathLength)
	} else if root.Right == nil {
		sumNumbersRecursive(root.Left, total, path, pathLength)
	} else {
		sumNumbersRecursive(root.Left, total, path, pathLength)
		sumNumbersRecursive(root.Right, total, path, pathLength)
	}
}

func main() {

	fmt.Printf("%v\n", sumNumbers(CreateBinaryTree([]int{1, 2, 3, 4, 5}, 0)))
	fmt.Printf("%v\n", sumNumbers(CreateBinaryTree([]int{1, 2, 3}, 0)))

}
