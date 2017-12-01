package main

import (
	. "leetgo/structure/link"
)

//使用归并排序。
func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	//使用快慢指针寻找中点。
	fast := head
	slow := head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	p1 := head
	p2 := slow.Next
	slow.Next = nil

	orderedP1 := sortList(p1)
	orderedP2 := sortList(p2)

	tmp := sortListMerge(orderedP1, orderedP2)
	return tmp
}

//合并两个排好序的的链表。两个链表要求不相连。
func sortListMerge(node1, node2 *ListNode) *ListNode {

	if node1 != nil && node2 != nil {

		p1 := node1
		p2 := node2

		var root *ListNode
		var tail *ListNode
		if p1.Val < p2.Val {
			root = p1
			p1 = p1.Next
			root.Next = nil
			tail = root
		} else {
			root = p2
			p2 = p2.Next
			root.Next = nil
			tail = root
		}

		for p1 != nil && p2 != nil {
			if p1.Val < p2.Val {
				tail.Next = p1
				p1 = p1.Next
				tail = tail.Next
				tail.Next = nil
			} else {
				tail.Next = p2
				p2 = p2.Next
				tail = tail.Next
				tail.Next = nil
			}
		}

		for p1 != nil {
			tail.Next = p1
			tail = tail.Next
			p1 = p1.Next
		}
		for p2 != nil {
			tail.Next = p2
			tail = tail.Next
			p2 = p2.Next
		}

		return root
	} else if node1 == nil {
		return node2
	} else {
		return node1
	}
}

func main() {
	a := CreateLinkList([]int{3, 4, 1, 2, 7, 9, 2})
	b := sortList(a)
	PrintLinkList(b)

	a = CreateLinkList([]int{})
	b = sortList(a)
	PrintLinkList(b)

	a = CreateLinkList([]int{1})
	b = sortList(a)
	PrintLinkList(b)

	a = CreateLinkList([]int{2, 1})
	b = sortList(a)
	PrintLinkList(b)

}
