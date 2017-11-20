package main

import (
	"fmt"
	"math"
)

func isValidBST(root *TreeNode) bool {

	return isValidBSTRecursive(root, math.MinInt64, math.MaxInt64)

}

func isValidBSTRecursive(root *TreeNode, min int, max int) bool {

	if root == nil {
		return true
	}
	return min < root.Val && root.Val < max && isValidBSTRecursive(root.Left, min, root.Val) && isValidBSTRecursive(root.Right, root.Val, max)

}

func main() {

	tree := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%v\n", isValidBST(tree))

	tree = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%v\n", isValidBST(tree))

	tree = &TreeNode{Val: 1}
	fmt.Printf("%v\n", isValidBST(tree))

	tree = &TreeNode{Val: 1, Left: &TreeNode{Val: -1}}
	fmt.Printf("%v\n", isValidBST(tree))

}
