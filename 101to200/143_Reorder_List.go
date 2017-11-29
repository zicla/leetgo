package main

import (
	. "leetgo/structure/link"
)

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	cur := head
	N := 0
	for cur != nil {
		N++
		cur = cur.Next
	}

	//只需要移动[(N-1)/2]
	var times int = (N - 1) / 2

	p1 := head

	for i := 1; i <= times; i++ {

		lastButOne := head
		last := head.Next
		for last.Next != nil {
			last = last.Next;
			lastButOne = lastButOne.Next;
		}
		lastButOne.Next = nil

		last.Next = p1.Next
		p1.Next = last

		last = last.Next
		p1 = last

	}

}
func main() {
	l := CreateLinkList([]int{1, 2, 3, 4, 5, 6})
	reorderList(l)
	PrintLinkList(l)

}
