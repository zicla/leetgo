package main

import (
	"fmt"
	. "leetgo/structure/btree"
)

func inorderTraversal(root *TreeNode) []int {

	var res []int

	if root == nil {
		return res
	}
	var stack []*TreeNode
	cur := root
	for len(stack) != 0 || cur != nil {

		//先一通左边疯狂入栈。
		for cur != nil {
			stack = append([]*TreeNode{cur}, stack...)
			cur = cur.Left
		}

		if len(stack) != 0 {
			//出栈一个元素。
			cur = stack[0]
			stack = stack[1:]

			//打印出当前结点。
			res = append(res, cur.Val)

			//把这个元素指到右边
			cur = cur.Right
		}

	}
	return res

}

func main() {

	tree := &TreeNode{Val: 8, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7}}}, Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 11}}}

	fmt.Printf("%v", inorderTraversal(tree))

	tree = &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	fmt.Printf("%v", inorderTraversal(tree))

}
