package main

import (
	"fmt"
)

type GasLink struct {
	Index int
	Val   int
	Cost  int
	Next  *GasLink
}

func canCompleteCircuit(gas []int, cost []int) int {
	N := len(gas)
	if N == 0 {
		return 1
	}
	if N == 1 {
		if gas[0] >= cost[0] {
			return 0
		} else {
			return -1
		}
	}

	root := &GasLink{
		Index: 0,
		Val:   gas[0],
		Cost:  cost[0],
	}
	cur := root
	for i := 1; i < N; i++ {
		cur.Next = &GasLink{
			Index: i,
			Val:   gas[i],
			Cost:  cost[i],
		}
		cur = cur.Next
	}
	cur.Next = root

	start := root
	for i := 0; i < N; i++ {

		cur = start
		var tank int
		ok := true
		times := 0
		for !(cur == start && times != 0) {
			times++

			//顺利到达一个加油站后补充油
			tank += cur.Val

			//判断能否顺利到达下一个站。
			tank = tank - cur.Cost
			if tank < 0 {
				ok = false
				break
			}
			cur = cur.Next

		}

		if ok {
			return start.Index
		}

		start = start.Next

	}
	return -1

}

func main() {

	fmt.Printf("%v\n", canCompleteCircuit([]int{1, 2}, []int{2, 1}))
	fmt.Printf("%v\n", canCompleteCircuit([]int{4}, []int{5}))
	fmt.Printf("%v\n", canCompleteCircuit([]int{5}, []int{4}))
}
