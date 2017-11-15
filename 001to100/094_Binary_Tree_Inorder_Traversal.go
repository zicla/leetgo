package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 1<<7-1)
	N := []int{0}
	inorderTraversalRecursive(root, result, N)
	return result[0:N[0]]
}

func inorderTraversalRecursive(root *TreeNode, result []int, N []int) {
	if root != nil {
		inorderTraversalRecursive(root.Left, result, N)
		result[N[0]] = root.Val
		N[0]++
		inorderTraversalRecursive(root.Right, result, N)
	}
}

func main() {

}
