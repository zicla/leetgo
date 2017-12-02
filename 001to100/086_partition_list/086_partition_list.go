package main

import "fmt"

type ListNode086 struct {
	Val  int
	Next *ListNode086
}

func partition(head *ListNode086, x int) *ListNode086 {

	var head1 *ListNode086;
	var p1 *ListNode086;
	var head2 *ListNode086;
	var p2 *ListNode086;
	p := head
	for p != nil {
		if p.Val < x {
			if head1 == nil {
				head1 = &ListNode086{Val: p.Val}
				p1 = head1
			} else {
				p1.Next = &ListNode086{Val: p.Val}
				p1 = p1.Next
			}
		} else {
			if head2 == nil {
				head2 = &ListNode086{Val: p.Val}
				p2 = head2
			} else {
				p2.Next = &ListNode086{Val: p.Val}
				p2 = p2.Next
			}
		}
		p = p.Next
	}

	if head1 == nil {
		return head2
	} else {
		p1.Next = head2
		return head1
	}
}

func main() {

	list1 := &ListNode086{Val: 1, Next: &ListNode086{Val: 4, Next: &ListNode086{Val: 3, Next: &ListNode086{Val: 2, Next: &ListNode086{Val: 5, Next: &ListNode086{Val: 2}}}}}}

	p := partition(list1, 3)

	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}

}
