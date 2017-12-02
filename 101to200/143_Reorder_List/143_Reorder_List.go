package main

//反转一个链表。
func reorderListReverseLinkList(root *ListNode) *ListNode {
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

func reorderList(head *ListNode) {

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	//使用快指针和慢指针找到中间点，方便我们打断。
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//后半部分的链表比前面部分要小。
	mid := slow.Next
	slow.Next = nil
	//把后半部分的链表反转了。
	mid = reorderListReverseLinkList(mid)

	//将后一部分表插入到前一部分中去。
	cur1 := head
	cur2 := mid

	for true {
		p1 := cur1.Next
		p2 := cur2.Next
		cur1.Next = cur2
		cur2.Next = p1
		if p2 == nil {
			break
		}
		cur2 = p2
		cur1 = p1
	}

}
func main() {

	l := CreateLinkList([]int{1, 2, 3, 4})

	reorderList(l)
	PrintLinkList(l)

	l = CreateLinkList([]int{1, 2, 3, 4, 5, 6})

	reorderList(l)
	PrintLinkList(l)

}
