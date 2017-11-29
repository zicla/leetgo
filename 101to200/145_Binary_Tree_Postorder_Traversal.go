package main

import (
	. "leetgo/structure/btree"
)

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	arr1 := postorderTraversal(root.Left)
	arr2 := postorderTraversal(root.Right)

	return append(append(arr1, arr2...), root.Val)
}

func main() {

}
