package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	N   int
	Arr []*TreeNode
}

func (this *Stack) push(node *TreeNode) {
	if this.Arr == nil {
		this.Arr = make([]*TreeNode, 1<<7-1)
	}
	this.Arr[this.N] = node
	this.N++
}

func (this *Stack) pop() *TreeNode {
	if this.N == 0 {
		return nil
	} else {
		node := this.Arr[this.N-1]
		this.N--
		return node
	}
}

func (this *Stack) empty() bool {
	return this.N == 0
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 1<<15-1)
	N := 0
	stack := &Stack{}
	cur := root
	for !stack.empty() || cur != nil {
		for cur != nil {
			stack.push(cur)
			cur = cur.Left
		}
		//此时到达了最左下角，需要出栈了。
		if !stack.empty() {
			cur = stack.pop()
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
