package main

import (
	"fmt"
	"math"
)

type TreeNode099 struct {
	Val   int
	Left  *TreeNode099
	Right *TreeNode099
}

func recoverTree(root *TreeNode099) {

	//准备中序遍历。
	list := make([]int, math.MaxInt8)
	N := []int{0}
	recoverTreeScan(root, list, N)

	fmt.Printf("%v\n", list)
}

func recoverTreeScan(root *TreeNode099, list []int, N []int) {

	if root == nil {
		return
	}
	recoverTreeScan(root.Left, list, N)

	list[N[0]] = root.Val
	N[0]++

	recoverTreeScan(root.Right, list, N)

}

func correctArray(list []int, N []int) {


}

func main() {
	tree := &TreeNode099{Val: 8, Left: &TreeNode099{Val: 4, Left: &TreeNode099{Val: 2, Left: &TreeNode099{Val: 1}, Right: &TreeNode099{Val: 3}}, Right: &TreeNode099{Val: 6, Left: &TreeNode099{Val: 5}, Right: &TreeNode099{Val: 7}}}, Right: &TreeNode099{Val: 10, Left: &TreeNode099{Val: 9}, Right: &TreeNode099{Val: 11}}}
	recoverTree(tree)

}
