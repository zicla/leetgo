package main

import "fmt"

func maxProfit(prices []int) int {
	N := len(prices)
	if N == 0 {
		return 0
	}
	max := 0
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if prices[j]-prices[i] > max {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

func main() {

	fmt.Printf("%v\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v\n", maxProfit([]int{7, 6, 4, 3, 1}))
}
