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

//通过一个数组创建一颗二叉树。
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

//获取一颗二叉树的总层数。
func BinaryTreeHeight(root *TreeNode) int {
	if root == nil {
		return 0;
	}
	left := BinaryTreeHeight(root.Left)
	right := BinaryTreeHeight(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func PrintLevelBinaryTree(root *TreeNode) {

	if root == nil {
		return
	}

	//树的层数
	height := BinaryTreeHeight(root)
	//最后一层的结点个数
	lastLevelNum := 1 << uint(height-1)
	//最后一层的位置数。
	seatNum := 2*lastLevelNum - 1

	//每个位置初始值都为无穷大
	matrix := make([][]int, height)
	for level := 0; level < height; level++ {
		tmp := make([]int, seatNum)
		for i := 0; i < seatNum; i++ {
			tmp[i] = math.MaxInt64
		}
		matrix[level] = tmp
	}

	PrintLevelBinaryTreeScan(root, height, 0, lastLevelNum-1, matrix)

	for i := 0; i < height; i++ {
		for j := 0; j < seatNum; j++ {
			t := matrix[i][j]
			if t == math.MaxInt64 {
				fmt.Printf("%3s", ".")
			} else {
				fmt.Printf("%3d", t)
			}
		}
		fmt.Println()
	}

}

func PrintLevelBinaryTreeScan(root *TreeNode, height int, level int, seatIndex int, matrix [][]int) {

	if root == nil {
		return
	}

	matrix[level][seatIndex] = root.Val

	PrintLevelBinaryTreeScan(root.Left, height, level+1, seatIndex-(1<<uint(height-level-2)), matrix)
	PrintLevelBinaryTreeScan(root.Right, height, level+1, seatIndex+(1<<uint(height-level-2)), matrix)
}

func PrintBinaryTreeGraph(root *TreeNode) {
	if root == nil {
		return;
	}
	//队列
	var queue []*TreeNode

	cur := root
	queue = append(queue, cur)
	//此处特意放入一个nil，用于隔离每一层。
	queue = append(queue, nil)

	for len(queue) != 0 {

		//出队列
		cur := queue[0]
		queue = queue[1:]

		if cur == nil {

			if len(queue) == 0 {
				//整个遍历过程结束了。
				fmt.Println()
				break
			} else {
				//开始做准备，下一行了
				queue = append(queue, nil)
				//换行了。
				fmt.Println()
				continue
			}
		}

		fmt.Printf("%v ", cur.Val)

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

	}
}
