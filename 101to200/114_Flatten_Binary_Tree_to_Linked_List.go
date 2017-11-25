package main

import (
	. "leetgo/structure/btree"
	"math"
)

func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	flattenRecursive(root)
}

//处理为一条龙，并且返回一条龙的最后一个结点。
func flattenRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	} else if root.Left == nil {
		return flattenRecursive(root.Right)
	} else if root.Right == nil {
		tail := flattenRecursive(root.Left)
		root.Right = root.Left
		root.Left = nil
		return tail
	} else {
		tail1 := flattenRecursive(root.Left)
		tail2 := flattenRecursive(root.Right)
		right := root.Right
		root.Right = root.Left
		root.Left = nil
		tail1.Right = right
		return tail2
	}

}

func main() {

	tree := CreateBinaryTree([]int{1, 2, 5, 3, 4, math.MinInt64, 6}, 0)
	flatten(tree)
	PrintLevelBinaryTree(tree)

}
