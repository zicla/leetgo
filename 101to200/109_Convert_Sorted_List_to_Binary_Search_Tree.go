package main

import (
	. "leetgo/structure/btree"
	. "leetgo/structure/link"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//链表中的内容放到数组中
func linkToArray(root *ListNode) []int {
	var arr []int
	for root != nil {
		arr = append(arr, root.Val)
		root = root.Next
	}
	return arr
}
func sortedArrayToBST1(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	t := &TreeNode{}
	median := int(len(nums) / 2)

	t.Val = nums[median]
	t.Left = sortedArrayToBST1(nums[:median])
	t.Right = sortedArrayToBST1(nums[median+1:])

	return t
}
func sortedListToBST(head *ListNode) *TreeNode {

	array := linkToArray(head)
	return sortedArrayToBST1(array)

}

func main() {

	bst := sortedListToBST(CreateLinkList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	PrintLevelBinaryTree(bst)
}
