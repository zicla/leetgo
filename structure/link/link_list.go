package link

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//从一个数组中创建链表
func CreateLinkList(list []int) *ListNode {
	if list == nil {
		return nil
	}
	var node *ListNode
	var current *ListNode
	for _, v := range list {
		if node == nil {
			node = &ListNode{
				Val: v,
			}
			current = node
			continue
		} else {
			current.Next = &ListNode{
				Val: v,
			}
			current = current.Next
		}
	}
	return node
}

//反转一个链表。
func ReverseLinkList(root *ListNode) *ListNode {
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

//打印链表中的内容
func PrintLinkList(root *ListNode) {

	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
}
