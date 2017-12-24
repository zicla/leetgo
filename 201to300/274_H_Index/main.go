package main

import (
	"math"
	"fmt"
)

func QuickSort(v []int, left int, right int) {

	if left < right {
		key := v[left]
		low := left
		high := right
		for low < high {
			for low < high && v[high] >= key {
				high--
			}
			v[low] = v[high]
			for low < high && v[low] <= key {
				low++
			}
			v[high] = v[low]
		}
		v[low] = key
		QuickSort(v, left, low-1)
		QuickSort(v, low+1, right)
	}

}

func hIndex(citations []int) int {

	N := len(citations)
	if N == 0 {
		return 0
	}
	QuickSort(citations, 0, len(citations)-1)

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
