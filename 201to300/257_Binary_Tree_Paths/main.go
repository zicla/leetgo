package main

import (
	. "leetgo/structure/btree"
	"fmt"
	"math"
)

func binaryTreePaths(root *TreeNode) []string {

	var res []string
	if root == nil {
		return res
	} else if root.Left == nil && root.Right == nil {
		res = append(res, fmt.Sprintf("%d", root.Val))
		return res
	} else {

		left := binaryTreePaths(root.Left)
		for _, v := range left {
			res = append(res, fmt.Sprintf("%d->%s", root.Val, v))
		}

		right := binaryTreePaths(root.Right)
		for _, v := range right {
			res = append(res, fmt.Sprintf("%d->%s", root.Val, v))
		}

		return res
	}

}

func main() {

	fmt.Printf("%v\n", binaryTreePaths(CreateBinaryTree([]int{1, 2, 3, math.MaxInt64, 5}, 0)))
}
