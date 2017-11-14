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

	p := head
	for i := 0; i < m-1; i++ {
		p = p.Next
	}

	z := (n - m) / 2
	if (n-m)%2 == 1 {
		z = z + 1
	}
	for i := 0; i < z; i++ {

		tmp := p
		k := n - m - 2*i
		for j := 0; j < k; j++ {
			tmp = tmp.Next
		}

		t := p.Val
		p.Val = tmp.Val
		tmp.Val = t

		p = p.Next

	}

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
	reverseBetween(list, 2, 4)
	PrintListNode(list)

	list = &ListNode092{Val: 1, Next: &ListNode092{Val: 2, Next: &ListNode092{Val: 3, Next: &ListNode092{Val: 4, Next: &ListNode092{Val: 5}}}}}
	reverseBetween(list, 1, 5)
	PrintListNode(list)

	list = &ListNode092{Val: 1, Next: &ListNode092{Val: 2, Next: &ListNode092{Val: 3, Next: &ListNode092{Val: 4, Next: &ListNode092{Val: 5, Next: &ListNode092{Val: 6}}}}}}
	reverseBetween(list, 1, 4)
	PrintListNode(list)

	list = &ListNode092{Val: 1}
	reverseBetween(list, 1, 1)
	PrintListNode(list)

}
