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

//打印链表中的内容
func PrintLinkList(root *ListNode) {

	for root != nil {
		fmt.Printf("%d ", root.Val)
		root = root.Next
	}
	fmt.Println()
}
