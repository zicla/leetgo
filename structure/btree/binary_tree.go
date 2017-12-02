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
	if index < N && data[index] != math.MinInt64 && data[index] != math.MaxInt64 {

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

//采用广度优先搜索的方式，按层打印出整个二叉树。
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

//二叉树前序遍历
func BinaryTreePreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%v ", root.Val)
	BinaryTreePreOrder(root.Left)
	BinaryTreePreOrder(root.Right)
}

//二叉树非递归方式前序遍历。
//递归的本质就是栈，如果非递归的话，那就相当于要自己来构建这个栈，顺着递归的思路写便可。
func BinaryTreePreOrderIterative(root *TreeNode) {
	if root == nil {
		return
	}

	var stack = []*TreeNode{root}

	var cur *TreeNode
	for len(stack) != 0 {

		//弹出一个结点。
		cur = stack[0]
		stack = stack[1:]

		//打印出当前结点。
		fmt.Printf("%v ", cur.Val)

		//由于栈是先进后出，因此这里右结点先入栈，左结点后入栈。

		//同时把右结点入栈。
		if cur.Right != nil {
			stack = append([]*TreeNode{cur.Right}, stack...)
		}

		//同时把左结点入栈。
		if cur.Left != nil {
			stack = append([]*TreeNode{cur.Left}, stack...)
		}

	}

}

//二叉树中序遍历
func BinaryTreeInOrder(root *TreeNode) {
	if root == nil {
		return
	}

	BinaryTreeInOrder(root.Left)
	fmt.Printf("%v ", root.Val)
	BinaryTreeInOrder(root.Right)
}

//二叉树非递归方式中序遍历
//递归的本质就是栈，如果非递归的话，那就相当于要自己来构建这个栈，顺着递归的思路写便可。
func BinaryTreeInOrderIterative(root *TreeNode) {
	if root == nil {
		return
	}
	var stack = []*TreeNode{}
	cur := root
	for len(stack) != 0 || cur != nil {

		//先一通左边疯狂入栈。
		for cur != nil {
			stack = append([]*TreeNode{cur}, stack...)
			cur = cur.Left
		}

		if len(stack) != 0 {
			//出栈一个元素。
			cur = stack[0]
			stack = stack[1:]

			//打印出当前结点。
			fmt.Printf("%v ", cur.Val)

			//把这个元素指到右边
			cur = cur.Right
		}

	}

}

//二叉树后序遍历
func BinaryTreePostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	BinaryTreePostOrder(root.Left)
	BinaryTreePostOrder(root.Right)
	fmt.Printf("%v ", root.Val)
}

//二叉树非递归方式后序遍历
//递归的本质就是栈，如果非递归的话，那就相当于要自己来构建这个栈，顺着递归的思路写便可。
func BinaryTreePostOrderUnify(root *TreeNode) {
	if root == nil {
		return
	}
	var stack = []*TreeNode{}
	visitedMap := make(map[*TreeNode]bool)

	//根节点入栈
	stack = append(stack, root)
	visitedMap[root] = false
	var visited bool
	for len(stack) != 0 {
		root := stack[0]
		stack = stack[1:]
		visited = visitedMap[root]
		if root == nil {
			continue;
		}
		if visited {
			fmt.Printf("%v ", root.Val)
		} else {
			//调整此处的顺序可以获取到前序中序和后续遍历
			stack = append([]*TreeNode{root}, stack...)
			visitedMap[root] = true
			stack = append([]*TreeNode{root.Right}, stack...)
			stack = append([]*TreeNode{root.Left}, stack...)
		}
	}

}

//二叉树非递归方式前序遍历
func BinaryTreePreOrderUnify(root *TreeNode) {
	if root == nil {
		return
	}
	var stack = []*TreeNode{}
	visitedMap := make(map[*TreeNode]bool)

	//根节点入栈
	stack = append(stack, root)
	visitedMap[root] = false
	var visited bool
	for len(stack) != 0 {
		root := stack[0]
		stack = stack[1:]
		visited = visitedMap[root]
		if root == nil {
			continue;
		}
		if visited {
			fmt.Printf("%v ", root.Val)
		} else {

			stack = append([]*TreeNode{root.Right}, stack...)
			stack = append([]*TreeNode{root.Left}, stack...)

			//调整此处的顺序可以获取到前序中序和后续遍历
			stack = append([]*TreeNode{root}, stack...)
			visitedMap[root] = true

		}
	}

}

//二叉树非递归方式中序遍历
func BinaryTreeInOrderUnify(root *TreeNode) {
	if root == nil {
		return
	}
	var stack = []*TreeNode{}
	visitedMap := make(map[*TreeNode]bool)

	//根节点入栈
	stack = append(stack, root)
	visitedMap[root] = false
	var visited bool
	for len(stack) != 0 {
		root := stack[0]
		stack = stack[1:]
		visited = visitedMap[root]
		if root == nil {
			continue;
		}
		if visited {
			fmt.Printf("%v ", root.Val)
		} else {

			stack = append([]*TreeNode{root.Right}, stack...)

			//调整此处的顺序可以获取到前序中序和后续遍历
			stack = append([]*TreeNode{root}, stack...)
			visitedMap[root] = true

			stack = append([]*TreeNode{root.Left}, stack...)



		}
	}

}
