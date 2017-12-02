package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type AVLTree struct {
	//数据元素
	Val int
	//树的高度。如果采用平衡因子的方式，维护起来会更加困难一些
	Height int
	//左子树指针
	Left *AVLTree
	//右子树指针
	Right *AVLTree
}

func TreeHeight(root *AVLTree) int {
	if root == nil {
		return 0
	} else {
		return root.Height
	}
}

func getAVL(root *TreeNode, balanced *bool) *AVLTree {
	if root == nil {
		return nil
	} else {

		left := getAVL(root.Left, balanced)
		if !*balanced {
			return nil
		}

		right := getAVL(root.Right, balanced)
		if !*balanced {
			return nil
		}

		l := TreeHeight(left)
		r := TreeHeight(right)

		if l-r > 1 || l-r < -1 {
			*balanced = false
			return nil
		}

		max := l
		if r > max {
			max = r
		}
		return &AVLTree{
			Val:    root.Val,
			Height: max + 1,
			Left:   left,
			Right:  right,
		}
	}
}

func isBalanced(root *TreeNode) bool {
	var balanced = true
	getAVL(root, &balanced)
	return balanced
}

func main() {

	tree := CreateBinaryTree([]int{1, 2, math.MinInt64, 3}, 0)
	fmt.Println(isBalanced(tree))
}
