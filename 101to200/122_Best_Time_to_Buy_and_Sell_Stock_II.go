package main

import "fmt"

func maxProfit122(prices []int) int {
	res := 0
	N := len(prices)
	for i := 1; i < N; i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func main() {

	fmt.Printf("%v\n", maxProfit122([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v\n", maxProfit122([]int{7, 6, 4, 3, 1}))
}
