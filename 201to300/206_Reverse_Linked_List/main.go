package main

import (
	. "leetgo/structure/link"
)
func reverseList(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}
	if root.Next.Next == nil {
		p := root.Next
		p.Next = root
		root.Next = nil
		return p
	}

	head := root
	p1 := root.Next
	p2 := root.Next.Next
	head.Next = nil

	for true {
		p1.Next = head
		head = p1
		p1 = p2
		p2 = p2.Next

		if p2 == nil {
			p1.Next = head
			head = p1
			break
		}
	}

	return head
}

func main() {


}
