package main

import "fmt"

func summaryRanges(nums []int) []string {

	var res []string
	N := len(nums)
	if N == 0 {
		return res
	}

	p := 0
	for p < N {

		q := p + 1
		for q < N && nums[q]-nums[p] == q-p {
			q++
		}

		if q-p == 1 {
			res = append(res, fmt.Sprintf("%d", nums[p]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[p], nums[q-1]))
		}
		p = q

	}

	return res
}

func main() {

	fmt.Printf("%v\n", summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Printf("%v\n", summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
