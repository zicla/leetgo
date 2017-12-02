package main

import (
	"fmt"
)

func singleNumber136(nums []int) int {

	b := 0
	for _, v := range nums {
		b = b ^ v
	}
	return b

}

func main() {

	fmt.Printf("%v\n", singleNumber136([]int{1, 2, 2}))
	fmt.Printf("%v\n", singleNumber136([]int{2, 2, 3, 3, 5}))
	fmt.Printf("%v\n", singleNumber136([]int{1, 2, 2, 10, 1}))
}
