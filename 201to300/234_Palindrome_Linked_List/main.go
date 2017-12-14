package main

import (
	"fmt"
	. "leetgo/structure/link"
)

//反转一个链表。
func isPalindromeReverse(root *ListNode) *ListNode {
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

func isPalindrome(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}

	//使用快慢指针，找到中点。
	fast := head
	slow := head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	//slow后面的链表反转
	head2 := isPalindromeReverse(slow.Next)
	slow.Next = nil

	p := head
	q := head2
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}

	return true
}

func main() {

	fmt.Printf("%v\n", isPalindrome(CreateLinkList([]int{1, 2})))

}
