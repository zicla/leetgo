package main

import (
	"fmt"
	. "leetgo/structure/btree"
)

//数一棵树的结点数量。
func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root.Left) + 1 + count(root.Right)
}

func kthSmallest(root *TreeNode, k int) int {
	countLeft := count(root.Left)
	if k == countLeft+1 {
		return root.Val
	} else if k > countLeft+1 {
		return kthSmallest(root.Right, k-(countLeft+1))
	} else {
		return kthSmallest(root.Left, k)
	}
}

func main() {

	fmt.Printf("%v \n", kthSmallest(CreateBinaryTree([]int{5, 3, 7, 2, 4, 6, 10}, 0), 7))

}
