package main

import (
	"fmt"
	"math"
)

func zigzagLevelOrder(root *TreeNode) [][]int {

	levelMap := make(map[int][]int)
	levelMax := []int{-1}
	zigzagLevelOrderScan(root, 0, levelMax, levelMap)

	var res [][]int
	for i := 0; i <= levelMax[0]; i++ {
		res = append(res, levelMap[i])
	}
	return res
}

func zigzagLevelOrderScan(root *TreeNode, level int, levelMax []int, levelMap map[int][]int) {

	if root == nil {
		return
	}
	if level > levelMax[0] {
		levelMax[0] = level
	}
	var arr []int
	if levelMap[level] != nil {
		arr = levelMap[level]
	}
	if level%2 == 0 {
		levelMap[level] = append(arr, root.Val)
	} else {
		levelMap[level] = append([]int{root.Val}, arr...)
	}

	zigzagLevelOrderScan(root.Left, level+1, levelMax, levelMap)
	zigzagLevelOrderScan(root.Right, level+1, levelMax, levelMap)
}

func main() {

	tree1 := CreateBinaryTree([]int{3, 9, 20, math.MinInt64, math.MinInt64, 15, 7}, 0)
	tree2 := CreateBinaryTree([]int{8, 4, 10, 3, 6, 9, 11, 1, 3, 7, 5}, 0)
	tree3 := CreateBinaryTree([]int{}, 0)

	fmt.Printf("%v\n", zigzagLevelOrder(tree1))
	fmt.Printf("%v\n", zigzagLevelOrder(tree2))
	fmt.Printf("%v\n", zigzagLevelOrder(tree3))

}
