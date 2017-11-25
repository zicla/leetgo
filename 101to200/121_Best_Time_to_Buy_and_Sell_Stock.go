package main

import (
	"fmt"
	"math"
)

func maxProfit121(prices []int) int {
	N := len(prices)
	if N == 0 {
		return 0
	}
	res := 0
	min := prices[0]
	for i := 0; i < N; i++ {

		//遇到更便宜的了，那就考虑在这里买进可能更好的
		if prices[i] < min {
			min = prices[i]
		} else {
			//遇到比买价更贵，那么就该考虑卖出了。
			tmp := prices[i] - min
			if tmp > res {
				res = tmp
			}

		}
	}

	return res
}

func maxProfitDP(prices []int) int {
	N := len(prices)
	if N == 0 {
		return 0
	}

	buy := math.MinInt32
	sell := 0
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < N; i++ {
		sell = max(sell, buy+prices[i])
		buy = max(buy, -prices[i])
	}

	return sell
}

func main() {

	fmt.Printf("%v\n", maxProfit121([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v\n", maxProfit121([]int{7, 6, 4, 3, 1}))

	fmt.Printf("%v\n", maxProfitDP([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v\n", maxProfitDP([]int{7, 6, 4, 3, 1}))

}
