package main

import (
	"math"
	"fmt"
)


func hIndex(citations []int) int {

	N := len(citations)
	if N == 0 {
		return 0
	}
	h := math.MinInt64
	for i := N - 1; i >= 0; i-- {
		num := N - i
		c := citations[i]

		tmpH := num
		if c < tmpH {
			tmpH = c
		}

		if tmpH > h {
			h = tmpH
		} else {
			break
		}

	}
	return h
}
func main() {

	fmt.Printf("%v \n ", hIndex([]int{3, 0, 6, 1, 5}))

}
