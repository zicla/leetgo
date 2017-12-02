package main

import (
	"fmt"
)

func hasPathSum(root *TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	} else if root.Left == nil {
		return hasPathSum(root.Right, sum-root.Val)
	} else if root.Right == nil {
		return hasPathSum(root.Left, sum-root.Val)
	} else {
		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	}
}

func main() {

	tree := CreateBinaryTree([]int{1, 2}, 0)
	fmt.Println(hasPathSum(tree, 1))
}
