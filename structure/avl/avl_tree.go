package avl

import "fmt"

//左右子树高度一样
const EH = 0

//左子树高
const LH = 1

//右子树高
const RH = -1

//数据类型
type ElemType int

type AVLTree struct {
	//数据元素
	Val ElemType
	//平衡因子
	Factor int
	//左子树指针
	Left *AVLTree
	//右子树指针
	Right *AVLTree
}

//向平衡树T中插入一个元素e，taller反应T长高与否。
//返回是否成功插入元素。
func InsertAVL(root **AVLTree, e ElemType, taller *bool) bool {

	//本身是空树，那么插入新节点，树长高。
	if *root == nil {
		*root = &AVLTree{
			Val:    e,
			Factor: EH,
		}
		*taller = true
		return true
	} else {

		//存在相同大小的节点，不插入
		if e == (*root).Val {
			*taller = false
			return false;
		} else if e < (*root).Val {

			//往左子树中插入。
			if !InsertAVL(&((*root).Left), e, taller) {
				//没有插入成功（和左子树某个元素值相同导致的）
				return false
			} else {
				//插入成功，已经顺利安放在左子树的某个位置上了。

				//如果变高了，需要看看是否需要进行平衡性调整。
				if *taller {

					switch (*root).Factor {
					case EH:
						//原本左右等高，现在因为左子树增高而增高了。
						(*root).Factor = LH
						*taller = true
					case LH:
						//原本就是左子树高，现在需要做平衡调整了
						LeftBalance(root)
						*taller = false
					case RH:
						//原本是右边高，现在等高了
						(*root).Factor = EH
						*taller = false
					}
				}
			}

		} else {

			//往右子树中插入。
			if !InsertAVL(&((*root).Right), e, taller) {
				//没有插入成功。
				return false
			} else {
				//插入成功

				//如果变高了
				if *taller {

					switch (*root).Factor {
					case EH:
						//原本左右等高，现在因为右子树增高而增高了。
						(*root).Factor = RH
						*taller = true
					case LH:
						//原本是左边高，现在等高了
						(*root).Factor = EH
						*taller = false
					case RH:
						//原本就是右子树高，现在需要做平衡调整了
						RightBalance(root)
						*taller = false
					}
				}
			}
		}
	}

	return true;
}

//从AVL中删除掉一个元素。返回值为是否删除了。
func DeleteAVL(root **AVLTree, e ElemType, shorter *bool) bool {
	if (*root) == nil {
		return false
	}

	if e < (*root).Val {
		//去左子树中删除。
		if DeleteAVL(&(*root).Left, e, shorter) {
			//如果真的成功删除了一个元素。

			if *shorter {
				//如果左子树变矮了，那么看看是否需要平衡性调整。
				switch (*root).Factor {
				case EH:
					//原本两边一样高，左边变矮了，那就右边自然变高。
					(*root).Factor = RH
					*shorter = false
				case LH:
					//原本左边高，现在左边变矮了，那么左右就一样高了。
					(*root).Factor = EH
					*shorter = true
				case RH:
					//原本右边高，现在右边更高了，直接对右边进行平衡处理即可。平衡处理内部调整平衡因子。
					RightBalance(root)
					*shorter = true
				}
			}

		} else {
			return false
		}
	} else if e > (*root).Val {
		//去右子树中删除。
		if DeleteAVL(&(*root).Right, e, shorter) {
			//如果真的成功删除了一个元素。

			if *shorter {
				//如果右子树变矮了，那么看看是否需要平衡性调整。
				switch (*root).Factor {
				case EH:
					//原本两边一样高，右边变矮了，那就左边自然变高。
					(*root).Factor = LH
					*shorter = false
				case LH:
					//原本左边高，现在左边更高了，直接对左边进行平衡处理即可。平衡处理内部调整平衡因子。
					//这里不能使用LeftBalance方法，比如 6 4 3 5这个结点就会无法调整。此处直接对root右旋即可。
					LeftBalance(root)
					*shorter = true
				case RH:
					//原本右边高，现在右边变矮了，那么左右就一样高了。
					(*root).Factor = EH
					*shorter = true
				}
			}

		} else {
			return false
		}
	} else {
		//那就是当前这个节点了。
		if (*root).Left == nil && (*root).Right == nil {
			//情况1.该结点是叶子节点，即没有左右树。直接删除即可。
			*root = nil
			*shorter = true
			return true
		} else if (*root).Left == nil {
			//情况2.该节点没有左子树，但是有右子树。把右子树接上即可。
			*root = (*root).Right
			*shorter = true
			return true
		} else if (*root).Right == nil {
			//情况3.该结点没有右子树，但是有左子树。把左子树街上即可。
			*root = (*root).Left
			*shorter = true
			return true
		} else {

			//情况4.该结点包含了左右子树。
			if (*root).Factor == RH {
				//右边更深，那么就从右边寻找替死鬼。
				maxNode := FindMaxNode(*root)
				//和替死鬼交换身份。
				(*root).Val, maxNode.Val = maxNode.Val, (*root).Val
				//这个时候从root结点对替死鬼进行重新删除。
				return DeleteAVL(root, e, shorter)

			} else {
				//左边更深，那么从左边寻找替死鬼。
				minNode := FindMinNode(*root)
				//和替死鬼交换身份。
				(*root).Val, minNode.Val = minNode.Val, (*root).Val
				//这个时候从root结点对替死鬼进行重新删除。
				return DeleteAVL(root, e, shorter)
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

func LeftBalance(root **AVLTree) {

	left := (*root).Left
	//检查左子树的平衡状态。
	switch left.Factor {
	case LH:
		//此时的情况是LL型，做右旋转即可。右旋之前需要先调整平衡因子。
		(*root).Factor = EH
		left.Factor = EH
		RightRotate(root)
	case RH:
		//此时是LR型，做先右旋，再左旋即可。
		leftRight := left.Right
		switch leftRight.Factor {
		case LH:
			(*root).Factor = RH
			left.Factor = EH
		case RH:
			(*root).Factor = EH
			left.Factor = LH
		}

		leftRight.Factor = EH
		LeftRotate(&left)
		(*root).Left = left;
		RightRotate(root)
	case EH:
		//左右结点相等，这种情况
	}
}
func RightBalance(root **AVLTree) {
	right := (*root).Right
	//检查左子树的平衡状态。
	switch right.Factor {
	case RH:
		//此时的情况是RR型，做左旋转即可。左旋之前需要先调整平衡因子。
		(*root).Factor = EH
		right.Factor = EH
		LeftRotate(root)
	case LH:
		//此时是RL型，做先左旋，再右旋即可。
		rightLeft := right.Left
		switch rightLeft.Factor {
		case LH:
			(*root).Factor = EH
			right.Factor = RH
		case RH:
			(*root).Factor = LH
			right.Factor = EH
		}
		rightLeft.Factor = EH
		RightRotate(&right)
		(*root).Right = right;
		LeftRotate(root)
	}
}

//右旋转，对于LL型可以一步到位。
func RightRotate(root **AVLTree) {
	left := (*root).Left
	(*root).Left = left.Right
	left.Right = *root
	*root = left
}

//左旋转，对于RR型可以一步到位。
func LeftRotate(root **AVLTree) {
	right := (*root).Right
	(*root).Right = right.Left
	right.Left = *root
	*root = right
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
