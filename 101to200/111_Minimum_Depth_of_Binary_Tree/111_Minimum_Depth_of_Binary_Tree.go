package main

import (
	"fmt"
	"math"
)

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	ml := minDepth(root.Left)
	mr := minDepth(root.Right)

	if ml < mr {
		return ml + 1
	} else {
		return mr + 1
	}

}

func main() {

	tree := CreateBinaryTree([]int{1, 2, math.MinInt64, 3}, 0)
	fmt.Println(minDepth(tree))
}
