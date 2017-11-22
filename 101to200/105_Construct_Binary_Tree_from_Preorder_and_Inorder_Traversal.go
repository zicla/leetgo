package main

import (
	. "leetgo/structure/btree"
)

func buildTree(preorder []int, inorder []int) *TreeNode {

	if preorder == nil {
		return nil
	}
	N := len(preorder)
	if N == 0 {
		return nil
	}
	rootIndex := -1
	for i := 0; i < N; i++ {
		if inorder[i] == preorder[0] {
			rootIndex = i
		}
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex]),
		Right: buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:]),
	}
	return root
}

func main() {
	tree := buildTree([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7})
	PrintLevelBinaryTree(tree)
	tree1 := buildTree([]int{}, []int{})
	PrintLevelBinaryTree(tree1)
	tree2 := buildTree([]int{1}, []int{1})
	PrintLevelBinaryTree(tree2)

}
