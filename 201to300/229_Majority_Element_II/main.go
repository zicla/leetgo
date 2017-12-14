package main

import (
	"fmt"
	"math"
)

func majorityElement(nums []int) []int {
	N := len(nums)
	var res []int
	if N == 0 {
		return res
	}

	//摩尔投票法

	m := math.MaxInt64
	n := math.MaxInt64
	cm := 0
	cn := 0
	for _, a := range nums {
		if a == m {
			cm++
		} else if a == n {
			cn++
		} else if cm == 0 {
			m = a
			cm = 1
		} else if cn == 0 {
			n = a
			cn = 1
		} else {
			cm--
			cn--
		}
	}

	cm = 0
	cn = 0
	for _, v := range nums {
		if v == m {
			cm++
		} else if v == n {
			cn++
		}
	}

	if cm > N/3 {
		res = append(res, m)
	}
	if cn > N/3 {
		res = append(res, n)
	}

	return res
}

func main() {

	fmt.Printf("%v\n", majorityElement([]int{1, 1, 1, 2, 3, 4, 5, 6}))
	fmt.Printf("%v\n", majorityElement([]int{1, 1, 2, 1, 2, 2}))
	fmt.Printf("%v\n", majorityElement([]int{1, 1}))
	fmt.Printf("%v\n", majorityElement([]int{0, 0, 0}))
}
