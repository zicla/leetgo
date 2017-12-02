package main

import (
	"math"
	"fmt"
)

type MinStack struct {
	stack []int
	//min也维护一个类似栈的东西
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{}
}

func (this *MinStack) Push(x int) {

	this.stack = append([]int{x}, this.stack...)

	M := len(this.min)
	if M == 0 || x <= this.min[0] {
		this.min = append([]int{x}, this.min...)
	}
}

func (this *MinStack) Pop() {

	if len(this.stack) == 0 {
		return
	}
	top := this.stack[0]
	this.stack = this.stack[1:]

	if top == this.min[0] {
		this.min = this.min[1:]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return math.MinInt64
	}
	return this.stack[0]
}

func (this *MinStack) GetMin() int {

	return this.min[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	minStack := Constructor()
	minStack.Push(512);
	minStack.Push(-1024);
	minStack.Push(-1024);
	minStack.Push(512);
	minStack.Pop();
	fmt.Println(minStack.GetMin())
	minStack.Pop();
	fmt.Println(minStack.GetMin())
	minStack.Pop();
	fmt.Println(minStack.GetMin())

}
