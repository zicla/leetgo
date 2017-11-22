package btree

import (
	"math"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateBinaryTree(data []int, index int) *TreeNode {

	N := len(data)
	var pNode *TreeNode;
	//数组中MinInt64 表示该节点为空
	if index < N && data[index] != math.MinInt64 {

		left := CreateBinaryTree(data, 2*index+1)
		right := CreateBinaryTree(data, 2*index+2)
		pNode = &TreeNode{
			Val:   data[index],
			Left:  left,
			Right: right,
		}

	}

	return pNode;
}

func PrintLevelBinaryTree(root *TreeNode) {

	levelMap := make(map[int][]int)
	levelMax := []int{-1}
	PrintLevelBinaryTreeScan(root, 0, levelMax, levelMap)

	var res [][]int
	for i := 0; i <= levelMax[0]; i++ {
		res = append(res, levelMap[i])
	}
	fmt.Printf("%v\n", res)
}

func PrintLevelBinaryTreeScan(root *TreeNode, level int, levelMax []int, levelMap map[int][]int) {

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
	levelMap[level] = append(arr, root.Val)
	PrintLevelBinaryTreeScan(root.Left, level+1, levelMax, levelMap)
	PrintLevelBinaryTreeScan(root.Right, level+1, levelMax, levelMap)
}
