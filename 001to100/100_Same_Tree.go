package main

import (
	"fmt"
	. "leetgo/structure/btree"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		b1 := isSameTree(p.Left, q.Left)

		if !b1 {
			return false
		}

		b2 := p.Val == q.Val
		if !b2 {
			return false
		}
		b3 := isSameTree(p.Right, q.Right)

		if !b3 {
			return false
		} else {
			return true
		}

	} else {
		return false
	}
}


func main() {

	tree1 := CreateBinaryTree([]int{8, 4, 10, 2, 6, 9, 11, 1, 3, 7, 5}, 0)
	tree2 := CreateBinaryTree([]int{8, 4, 10, 3, 6, 9, 11, 1, 3, 7, 5}, 0)

	fmt.Printf("%v\n", isSameTree(tree1, tree2))

}
