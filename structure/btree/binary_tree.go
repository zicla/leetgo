package btree

import (
	"math"
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
