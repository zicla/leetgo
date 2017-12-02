package btree

import "fmt"

type Stack struct {
	N   int
	Arr []*TreeNode
}

func (this *Stack) Push(node *TreeNode) {
	if this.Arr == nil {
		this.Arr = make([]*TreeNode, 1<<7-1)
	}
	this.Arr[this.N] = node
	this.N++
}

func (this *Stack) Pop() *TreeNode {
	if this.N == 0 {
		return nil
	} else {
		node := this.Arr[this.N-1]
		this.N--
		return node
	}
}

func (this *Stack) Empty() bool {
	return this.N == 0
}

//golang中的快速栈。
func StackSimple() {
	var stack []int
	// Push
	stack = append([]int{1}, stack...)

	// Top (just get next element, don't remove it)
	x := stack[0]
	fmt.Println(x)

	// Discard top element
	stack = stack[1:]
	// Is empty ?
	if len(stack) == 0 {
		fmt.Println("Stack is empty !")
	}
}
