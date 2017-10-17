package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	p1 := l1
	p2 := l2
	var res *ListNode = nil
	var p *ListNode = nil
	sum := 0
	for p1 != nil || p2 != nil {

		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}

		if res == nil {
			if sum > 9 {
				res = &ListNode{Val: sum - 10}
				sum = 1
			} else {
				res = &ListNode{Val: sum}
				sum = 0
			}
			p = res
		} else {
			if sum > 9 {
				p.Next = &ListNode{Val: sum - 10}
				sum = 1
			} else {
				p.Next = &ListNode{Val: sum}
				sum = 0
			}
			p = p.Next
		}

	}

	if sum != 0 {
		p.Next = &ListNode{Val: sum}
	}

	return res
}
func main() {
	//list1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	//list2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	list1 := &ListNode{Val: 5}
	list2 := &ListNode{Val: 5}

	list3 := addTwoNumbers(list1, list2)

	p := list3;
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
