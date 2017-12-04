package main

import (
	. "leetgo/structure/link"
)

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	//找到root.
	root := head
	for root != nil && root.Val == val {
		root = root.Next
	}

	if root == nil {
		return nil
	}

	pre := root
	cur := root.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}

	}
	return root
}

func main() {

	a := removeElements(CreateLinkList([]int{1, 2, 6, 3, 4, 5, 6}), 6)
	PrintLinkList(a)

}
