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

	this.stack = append(this.stack, x)

	M := len(this.min)
	if M == 0 || x <= this.min[M-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {

	N := len(this.stack)
	if N == 0 {
		return
	}
	top := this.stack[N-1]
	this.stack = this.stack[0:N-1]
	M := len(this.min)
	if top == this.min[M-1] {
		this.min = this.min[0:M-1]
	}
}

func (this *MinStack) Top() int {
	N := len(this.stack)
	if N == 0 {
		return math.MinInt64
	}
	return this.stack[N-1]
}

func (this *MinStack) GetMin() int {
	M := len(this.min)
	return this.min[M-1]
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
