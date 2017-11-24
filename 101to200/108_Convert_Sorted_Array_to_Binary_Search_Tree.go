package main

import (
	. "leetgo/structure/btree"
)

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

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//向平衡树T中插入一个元素e，taller反应T长高与否。
//返回是否成功插入元素。
func InsertAVL(root **AVLTree, e int) bool {

	//本身是空树，那么插入新节点
	if *root == nil {
		*root = &AVLTree{
			Val:    e,
			Height: 1,
		}
		return true
	} else {

		//存在相同大小的节点，不插入
		if e == (*root).Val {
			return false;
		} else if e < (*root).Val {

			//往左子树中插入。
			if !InsertAVL(&((*root).Left), e) {
				//没有插入成功（和左子树某个元素值相同导致的）
				return false
			} else {
				//插入成功，已经顺利安放在左子树的某个位置上了。
				(*root).Height = Max(TreeHeight((*root).Left), TreeHeight((*root).Right)) + 1

				//看看当前的树是否已经失衡。
				if TreeHeight((*root).Left)-TreeHeight((*root).Right) == 2 {

					if e < (*root).Left.Val {
						//LL型
						*root = RightRotate(*root)
					} else {
						//LR型
						*root = LeftRightRotate(*root)
					}

				}
			}

		} else {

			//往右子树中插入。
			if !InsertAVL(&((*root).Right), e) {
				//没有插入成功。
				return false
			} else {
				//插入成功
				(*root).Height = Max(TreeHeight((*root).Left), TreeHeight((*root).Right)) + 1

				//看看当前的树是否已经失衡。
				if TreeHeight((*root).Right)-TreeHeight((*root).Left) == 2 {

					if e < (*root).Right.Val {
						//RL型
						*root = RightLeftRotate(*root)
					} else {
						//RR型
						*root = LeftRotate(*root)
					}

				}
			}
		}
	}
	return true;
}

//获取一颗树的高度，为了方便处理nil而已。
func TreeHeight(root *AVLTree) int {
	if root == nil {
		return 0
	} else {
		return root.Height
	}
}

/*
	                              100                              85
                                 /  \               右旋         /    \
                               85   120         ------ ->       60    100
                              /  \                                \   /   \
                            60    90                              80 90   120
                              \
                               80
*/
//右旋转，对于LL型可以一步到位。
func RightRotate(root *AVLTree) *AVLTree {
	left := root.Left
	root.Left = left.Right
	left.Right = root

	root.Height = Max(TreeHeight(root.Left), TreeHeight(root.Right)) + 1
	left.Height = Max(TreeHeight(left.Left), root.Height) + 1

	return left
}

/*
                                  80                                   90
                                 /  \             左旋               /    \
                               60    90          ---- ->            80    120
                                    /  \                           /  \   /
                                   85  120                        60  85 100
                                      /
                                    100
*/
//左旋转，对于RR型可以一步到位。
func LeftRotate(root *AVLTree) *AVLTree {
	right := root.Right
	root.Right = right.Left
	right.Left = root

	root.Height = Max(TreeHeight(root.Left), TreeHeight(root.Right)) + 1
	right.Height = Max(TreeHeight(right.Right), root.Height) + 1

	return right
}

/*
	                          100                          100                        90
                             /  \         左旋            /  \       右旋           /    \
                            80  120     ------>          90  120    ------>        80   100
                           / \                          /                         /  \     \
                          60 90                        80                        60  85    120
                            /                         / \
                           85                        60 85
*/
//先左旋再右旋转，对于LR型可以一步到位。
func LeftRightRotate(root *AVLTree) *AVLTree {

	root.Left = LeftRotate(root.Left)
	return RightRotate(root)
}

/*
                         80                               80                                   85
                        /   \             右 旋          /  \               左 旋             /  \
                       60   100          ------>        60   85            ------->          80 100
                            /  \                               \                            /   /  \
                           85  120                             100                         60  90  120
                            \                                  /  \
                            90                                90  120
*/
//先右旋转再左旋，对于RL型可以一步到位。
func RightLeftRotate(root *AVLTree) *AVLTree {

	root.Right = RightRotate(root.Right)
	return LeftRotate(root)
}

//复制一个二叉树
func constructTreeNode(avlTree *AVLTree) *TreeNode {

	if avlTree == nil {
		return nil
	} else {
		return &TreeNode{
			Val:   avlTree.Val,
			Left:  constructTreeNode(avlTree.Left),
			Right: constructTreeNode(avlTree.Right),
		}
	}

}

func sortedArrayToBST(nums []int) *TreeNode {

	var root *AVLTree
	for _, v := range nums {
		InsertAVL(&root, v)
	}
	return constructTreeNode(root)
}

func main() {

	PrintLevelBinaryTree(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
