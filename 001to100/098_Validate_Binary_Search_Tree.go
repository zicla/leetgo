package main

import "fmt"

type TreeNode098 struct {
	Val   int
	Left  *TreeNode098
	Right *TreeNode098
}

func minBST(root *TreeNode098) int {
	cur := root
	min := root.Val
	for cur != nil {
		if cur.Val < min {
			min = cur.Val
		}
		cur = cur.Left
	}
	return min
}

func maxBST(root *TreeNode098) int {

	cur := root
	max := root.Val
	for cur != nil {
		if cur.Val > max {
			max = cur.Val
		}
		cur = cur.Right
	}
	return max
}

func isValidBST(root *TreeNode098) bool {

	if root == nil {
		return true
	}

	left := root.Left
	right := root.Right
	if left == nil && right == nil {
		return true
	} else if left == nil {
		if root.Val >= right.Val {
			return false
		} else {
			b1 := isValidBST(right)

			return b1 && root.Val < minBST(right)
		}

	} else if right == nil {
		if root.Val < left.Val {
			return false
		} else {
			b1 := isValidBST(left)
			return b1 && root.Val > maxBST(left)
		}

	} else {
		if root.Val >= right.Val || root.Val < left.Val {
			return false
		} else {

			//b1 := isValidBST(right)
			//b2 := root.Val < minBST(right)
			//b3 := isValidBST(left)
			//b4 := root.Val > maxBST(left)
			//return b1 && b2 && b3 && b4

			return isValidBST(right) && root.Val < minBST(right) && isValidBST(left) && root.Val > maxBST(left)
		}

	}
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
