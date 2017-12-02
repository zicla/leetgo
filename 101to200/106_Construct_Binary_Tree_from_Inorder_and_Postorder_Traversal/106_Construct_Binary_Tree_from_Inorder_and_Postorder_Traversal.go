package main

func buildTree(inorder []int, postorder []int) *TreeNode {

	if inorder == nil {
		return nil
	}
	N := len(inorder)
	if N == 0 {
		return nil
	}
	rootIndex := -1
	for i := 0; i < N; i++ {
		if inorder[i] == postorder[N-1] {
			rootIndex = i
		}
	}

	root := &TreeNode{
		Val:   postorder[N-1],
		Left:  buildTree(inorder[0:rootIndex], postorder[0:rootIndex]),
		Right: buildTree(inorder[rootIndex+1:], postorder[rootIndex:N-1]),
	}
	return root
}

func main() {
	tree := buildTree([]int{4, 2, 5, 1, 6, 3, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	PrintLevelBinaryTree(tree)
	tree1 := buildTree([]int{}, []int{})
	PrintLevelBinaryTree(tree1)
	tree2 := buildTree([]int{1}, []int{1})
	PrintLevelBinaryTree(tree2)

}
