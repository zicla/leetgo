package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	t := &TreeNode{}
	median := int(len(nums) / 2)

	t.Val = nums[median]
	t.Left = sortedArrayToBST(nums[:median])
	t.Right = sortedArrayToBST(nums[median+1:])

	return t
}
func main() {

	bst := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	PrintLevelBinaryTree(bst)
}
