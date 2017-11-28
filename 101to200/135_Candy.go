package main

import (
	"fmt"
)

func candy(ratings []int) int {
	var box = make(map[int]int)
	for _, v := range ratings {
		box[v]++
	}

	return 1;
}

func main() {

	fmt.Printf("%v\n", candy([]int{1, 2, 3, 4, 5}))
}
