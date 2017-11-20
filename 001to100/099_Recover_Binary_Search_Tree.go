package main

import (
	"math"
	"fmt"
)

type TreeNode099 struct {
	Val   int
	Left  *TreeNode099
	Right *TreeNode099
}

func recoverTree(root *TreeNode099) {

	//准备中序遍历。
	list := make([]*TreeNode099, math.MaxInt8)
	N := []int{0}
	recoverTreeScan(root, list, N)
	correctArray(list, N)
}

func recoverTreeScan(root *TreeNode099, list []*TreeNode099, N []int) {

	if root == nil {
		return
	}
	recoverTreeScan(root.Left, list, N)

	list[N[0]] = root
	N[0]++

	recoverTreeScan(root.Right, list, N)

}

func correctArray(list []*TreeNode099, N []int) {

	n := N[0]
	index1 := -1
	index2 := -1
	for i := 0; i < n-1; i++ {

		if list[i].Val > list[i+1].Val {
			if index1 == -1 {
				index1 = i

				index2 = i + 1
				for j := i + 2; j < n; j++ {
					if list[j].Val < list[index2].Val {
						index2 = j
					}
				}

			}

		}
	}
	if index1 == -1 {
		return
	}
	if index2 == -1 {
		index2 = index1 + 1
	}
	tmp := list[index1].Val
	list[index1].Val = list[index2].Val
	list[index2].Val = tmp
}

func printBTree(root *TreeNode099) {
	if root == nil {
		return
	}
	printBTree(root.Left)
	fmt.Printf("%v ", root.Val)
	printBTree(root.Right)

}

func main() {
	tree := &TreeNode099{
		Val: 8,
		Left: &TreeNode099{
			Val: 4,
			Left: &TreeNode099{
				Val: 2,
				Left: &TreeNode099{
					Val: 1},
				Right: &TreeNode099{
					Val: 3}},
			Right: &TreeNode099{
				Val: 6,
				Left: &TreeNode099{
					Val: 7},
				Right: &TreeNode099{
					Val: 5}}},
		Right: &TreeNode099{
			Val: 10,
			Left: &TreeNode099{
				Val: 9},
			Right: &TreeNode099{
				Val: 11}}}

	printBTree(tree)
	fmt.Println()

	recoverTree(tree)

	printBTree(tree)
	fmt.Println()
}
