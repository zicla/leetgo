package main

import "fmt"

type ListNode092 struct {
	Val  int
	Next *ListNode092
}

func reverseBetween(head *ListNode092, m int, n int) *ListNode092 {

	if m == n {
		return head
	}

	pleft := head
	p := head
	for i := 0; i < m-1; i++ {
		p = p.Next
		if i == m-3 {
			pleft = p
		}
	}

	newHead := p
	newTail := p
	p = p.Next
	for j := m - 1; j < n-1; j++ {

		q := p.Next
		newTail.Next = q
		p.Next = newHead
		newHead = p

		p = q
	}

	if m == 1 {
		return newHead
	}

	pleft.Next = newHead

	return head

}

func PrintListNode(head *ListNode092) {

	for head != nil {
		fmt.Printf("%v ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func main() {
	list := &ListNode092{Val: 1, Next: &ListNode092{Val: 2, Next: &ListNode092{Val: 3, Next: &ListNode092{Val: 4, Next: &ListNode092{Val: 5}}}}}
	newList := reverseBetween(list, 2, 4)
	PrintListNode(newList)

	list = &ListNode092{Val: 1, Next: &ListNode092{Val: 2, Next: &ListNode092{Val: 3, Next: &ListNode092{Val: 4, Next: &ListNode092{Val: 5}}}}}
	newList = reverseBetween(list, 1, 5)
	PrintListNode(newList)

	list = &ListNode092{Val: 1, Next: &ListNode092{Val: 2, Next: &ListNode092{Val: 3, Next: &ListNode092{Val: 4, Next: &ListNode092{Val: 5, Next: &ListNode092{Val: 6}}}}}}
	newList = reverseBetween(list, 1, 4)
	PrintListNode(newList)

	list = &ListNode092{Val: 1}
	newList = reverseBetween(list, 1, 1)
	PrintListNode(newList)

}
