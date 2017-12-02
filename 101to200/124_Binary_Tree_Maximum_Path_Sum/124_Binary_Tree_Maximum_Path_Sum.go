package main

import (
	"fmt"

)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	res := root.Val
	maxPathSumDFS(root, &res);
	return res;
}

func maxPathSumMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//从顶点出发的单龙最大长度。
func maxPathSumDFS(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := maxPathSumDFS(root.Left, res)
	right := maxPathSumDFS(root.Right, res)
	total := root.Val + maxPathSumMax(left, 0) + maxPathSumMax(right, 0)
	*res = maxPathSumMax(*res, total)
	leftRightMax := maxPathSumMax(left, right)
	if leftRightMax > 0 {
		return leftRightMax + root.Val
	} else {
		return root.Val
	}
}

func main() {

	fmt.Printf("%v\n", maxPathSum(CreateBinaryTree([]int{1, 2, 3}, 0)))

}
