package main

import (
	. "leetgo/structure/link"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	root := head
	p1 := head.Next
	p2 := p1.Next
	root.Next = nil

	for true {
		//找到p1的位置。
		cur := root

		if p1.Val <= cur.Val {
			p1.Next = cur
			root = p1
		} else {

			for true {
				if p1.Val >= cur.Val && (cur.Next == nil || p1.Val < cur.Next.Val) {
					break
				}
				cur = cur.Next
			}

			if cur.Next == nil {
				cur.Next = p1
				p1.Next = nil
			} else {
				p1.Next = cur.Next
				cur.Next = p1
			}
		}

		p1 = p2
		if p1 == nil {
			break
		}
		p2 = p1.Next

	}
	return root
}

func main() {
	a := CreateLinkList([]int{3, 4, 1, 2, 7, 9, 2})
	b := insertionSortList(a)
	PrintLinkList(b)
}
