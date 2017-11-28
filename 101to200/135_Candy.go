package main

import (
	"fmt"
)

func candy(ratings []int) int {
	N := len(ratings)
	if N == 0 {
		return 0
	}
	var arr = make([]int, N)
	arr[0] = 1
	for i := 1; i < N; i++ {
		if ratings[i-1] < ratings[i] {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = 1
		}
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := N - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			arr[i-1] = max(arr[i-1], arr[i]+1)
		}
	}

	sum := 0
	for i := 0; i < N; i++ {
		sum += arr[i]
	}

	return sum;
}

func main() {

	fmt.Printf("%v\n", candy([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", candy([]int{2, 2, 3, 4}))
	fmt.Printf("%v\n", candy([]int{1, 2, 2}))
}
