package main

import (
	"math"
	"fmt"
)

func levelOrderBottom(root *TreeNode) [][]int {

	var res [][]int

	levelOrderBottomRecursive(root, 0, &res)

	return res
}

func levelOrderBottomRecursive(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}

	if level+1 > len(*res) {
		*res = append([][]int{{}}, *res...)
	}
	N := len(*res)

	(*res)[N-1-level] = append((*res)[N-1-level], root.Val)
	levelOrderBottomRecursive(root.Left, level+1, res)
	levelOrderBottomRecursive(root.Right, level+1, res)

}

func main() {
	tree1 := CreateBinaryTree([]int{3, 9, 20, math.MinInt64, math.MinInt64, 15, 7}, 0)
	tree2 := CreateBinaryTree([]int{8, 4, 10, 3, 6, 9, 11, 1, 3, 7, 5}, 0)
	tree3 := CreateBinaryTree([]int{}, 0)

	fmt.Printf("%v\n", levelOrderBottom(tree1))
	fmt.Printf("%v\n", levelOrderBottom(tree2))
	fmt.Printf("%v\n", levelOrderBottom(tree3))

}
