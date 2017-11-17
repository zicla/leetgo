package main

import (
	"fmt"
	"math"
)

type TreeNode098 struct {
	Val   int
	Left  *TreeNode098
	Right *TreeNode098
}

func isValidBST(root *TreeNode098) bool {

	return isValidBSTRecursive(root, math.MinInt32, math.MaxInt32)

}

func isValidBSTRecursive(root *TreeNode098, min int, max int) bool {

	if root == nil {
		return true
	}
	return min < root.Val && root.Val < max && isValidBSTRecursive(root.Left, min, root.Val) && isValidBSTRecursive(root.Right, root.Val, max)

}

func main() {

	tree := &TreeNode098{Val: 2, Left: &TreeNode098{Val: 1}, Right: &TreeNode098{Val: 3}}
	fmt.Printf("%v\n", isValidBST(tree))

	tree = &TreeNode098{Val: 1, Left: &TreeNode098{Val: 2}, Right: &TreeNode098{Val: 3}}
	fmt.Printf("%v\n", isValidBST(tree))

	tree = &TreeNode098{Val: 1}
	fmt.Printf("%v\n", isValidBST(tree))

	tree = &TreeNode098{Val: 1, Left: &TreeNode098{Val: -1}}
	fmt.Printf("%v\n", isValidBST(tree))

}
