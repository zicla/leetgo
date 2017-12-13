package main

import . "leetgo/structure/btree"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root

}

func main() {

}
