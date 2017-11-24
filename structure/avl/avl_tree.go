package avl

import "fmt"

//数据类型
type ElemType int

type AVLTree struct {
	//数据元素
	Val ElemType
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
func InsertAVL(root **AVLTree, e ElemType) bool {

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

//从AVL中删除掉一个元素。返回值为是否删除了。
func DeleteAVL(root **AVLTree, e ElemType) bool {
	if root == nil {
		return false
	}

	if e < (*root).Val {
		//去左子树中删除。
		if DeleteAVL(&(*root).Left, e) {
			//如果真的成功删除了一个元素。
			(*root).Height = Max(TreeHeight((*root).Left), TreeHeight((*root).Right)) + 1
			if TreeHeight((*root).Right)-TreeHeight((*root).Left) == 2 {
				*root = LeftRotate(*root)
			}
		} else {
			return false
		}
	} else if e > (*root).Val {
		//去右子树中删除。
		if DeleteAVL(&(*root).Right, e) {

			//如果真的成功删除了一个元素。
			(*root).Height = Max(TreeHeight((*root).Left), TreeHeight((*root).Right)) + 1
			if TreeHeight((*root).Left)-TreeHeight((*root).Right) == 2 {
				*root = RightRotate(*root)
			}

		} else {
			return false
		}
	} else {
		//那就是当前这个节点了。
		if (*root).Left == nil && (*root).Right == nil {
			//情况1.该结点是叶子节点，即没有左右树。直接删除即可。
			*root = nil
			return true
		} else if (*root).Left == nil {
			//情况2.该节点没有左子树，但是有右子树。把右子树接上即可。
			*root = (*root).Right
			return true
		} else if (*root).Right == nil {
			//情况3.该结点没有右子树，但是有左子树。把左子树街上即可。
			*root = (*root).Left
			return true
		} else {

			//情况4.该结点包含了左右子树。
			if (*root).Left.Height <= (*root).Right.Height {
				//右边更深，那么就从右边寻找一个最小的替死鬼。
				maxNode := FindMinNode((*root).Right)
				//和替死鬼交换身份。
				(*root).Val, maxNode.Val = maxNode.Val, (*root).Val
				//这个时候从root结点对替死鬼进行重新删除。
				return DeleteAVL(&((*root).Right), e)

			} else {
				//左边更深，那么从左边寻找一个最大的替死鬼。
				minNode := FindMaxNode((*root).Left)
				//和替死鬼交换身份。
				(*root).Val, minNode.Val = minNode.Val, (*root).Val
				//这个时候从root结点对替死鬼进行重新删除。
				return DeleteAVL(&((*root).Left), e)
			}

		}
	}

	return true
}

//找寻一颗avl树最小的结点。
func FindMinNode(root *AVLTree) *AVLTree {

	if root == nil {
		return nil
	}

	if root.Left == nil {
		return root
	} else {
		return FindMinNode(root.Left)
	}
}

//找寻一颗avl树最大的结点。
func FindMaxNode(root *AVLTree) *AVLTree {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	} else {
		return FindMaxNode(root.Right)
	}
}

//按层打印AVL
func PrintLevelAVLTree(root *AVLTree) {

	levelMap := make(map[int][]ElemType)
	levelMax := []int{-1}
	PrintLevelAVLTreeScan(root, 0, levelMax, levelMap)

	var res [][]ElemType
	for i := 0; i <= levelMax[0]; i++ {
		res = append(res, levelMap[i])
	}
	fmt.Printf("%v\n", res)
}

func PrintLevelAVLTreeScan(root *AVLTree, level int, levelMax []int, levelMap map[int][]ElemType) {

	if root == nil {
		return
	}
	if level > levelMax[0] {
		levelMax[0] = level
	}
	var arr []ElemType
	if levelMap[level] != nil {
		arr = levelMap[level]
	}
	levelMap[level] = append(arr, root.Val)
	PrintLevelAVLTreeScan(root.Left, level+1, levelMax, levelMap)
	PrintLevelAVLTreeScan(root.Right, level+1, levelMax, levelMap)
}
