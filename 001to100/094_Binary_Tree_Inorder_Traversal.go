package main

import (
	"fmt"
	. "leetgo/structure/btree"
)

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 1<<15-1)
	N := 0
	stack := &Stack{}
	cur := root
	for !stack.Empty() || cur != nil {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		//此时到达了最左下角，需要出栈了。
		if !stack.Empty() {
			cur = stack.Pop()
			result[N] = cur.Val
			N++
			cur = cur.Right
		}
	}
	return result[0:N]
}

func main() {

	tree := &TreeNode{Val: 8, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7}}}, Right: &TreeNode{Val: 10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 11}}}

	fmt.Printf("%v", inorderTraversal(tree))

	tree = &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}

	fmt.Printf("%v", inorderTraversal(tree))

}
