package main

import (
	"fmt"
	"math"
)

func maxDepth(root *TreeNode) int {
	var a int = 0
	maxDepthRecursive(root, 1, &a)
	return a
}

func maxDepthRecursive(root *TreeNode, level int, maxLevel *int) {
	if root == nil {
		return
	}

	if level > *maxLevel {
		*maxLevel = level
	}
	maxDepthRecursive(root.Left, level+1, maxLevel)
	maxDepthRecursive(root.Right, level+1, maxLevel)

}
func main() {

	tree1 := CreateBinaryTree([]int{3, 9, 20, math.MinInt64, math.MinInt64, 15, 7}, 0)
	tree2 := CreateBinaryTree([]int{8, 4, 10, 3, 6, 9, 11, 1, 3, 7, 5}, 0)
	tree3 := CreateBinaryTree([]int{}, 0)

	fmt.Printf("%v\n", maxDepth(tree1))
	fmt.Printf("%v\n", maxDepth(tree2))
	fmt.Printf("%v\n", maxDepth(tree3))

}
