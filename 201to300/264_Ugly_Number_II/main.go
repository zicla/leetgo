package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a;
	}
	return b
}
func nthUglyNumber(n int) int {
	var res = []int{1}

	i2 := 0
	i3 := 0
	i5 := 0

	for len(res) < n {
		m2 := res[i2] * 2
		m3 := res[i3] * 3
		m5 := res[i5] * 5
		mn := min(m2, min(m3, m5));
		if mn == m2 {
			i2++
		}
		if mn == m3 {
			i3++
		}
		if mn == m5 {
			i5++
		}

		res = append(res, mn)
	}
	return res[len(res)-1]
}

func main() {

	fmt.Printf("%v\n", nthUglyNumber(10))
	fmt.Printf("%v\n", nthUglyNumber(6))
	fmt.Printf("%v\n", nthUglyNumber(8))
	fmt.Printf("%v\n", nthUglyNumber(499))
}
