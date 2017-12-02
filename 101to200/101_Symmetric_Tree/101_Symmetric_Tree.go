package main

import (
	"fmt"
)

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true;
	}
	return isSymmetricScan(root.Left, root.Right)
}

func isSymmetricScan(node1 *TreeNode, node2 *TreeNode) bool {

	if node1 == nil && node2 == nil {
		return true
	} else if node1 != nil && node2 != nil {

		return node1.Val == node2.Val && isSymmetricScan(node1.Left, node2.Right) && isSymmetricScan(node1.Right, node2.Left)

	} else {
		return false
	}

}

func main() {

	tree1 := CreateBinaryTree([]int{1, 2, 2, 3, 4, 4, 3}, 0)
	tree2 := CreateBinaryTree([]int{8, 4, 10, 3, 6, 9, 11, 1, 3, 7, 5}, 0)

	fmt.Printf("%v\n", isSymmetric(tree1))
	fmt.Printf("%v\n", isSymmetric(tree2))

}
