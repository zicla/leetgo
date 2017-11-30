package main

import (
	. "leetgo/structure/link"
)

func sortList(head *ListNode) *ListNode {

}

func main() {
	a := CreateLinkList([]int{3, 4, 1, 2, 7, 9, 2})
	b := insertionSortList(a)
	PrintLinkList(b)
}
