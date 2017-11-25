package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {

	N := len(prices)
	//假设买入之前口袋中有0元
	//buy1[i]表示前i天做第一笔交易 买入 股票后能剩的最多的钱。
	//sell1[i]表示前i天做第一笔交易 卖出 股票后能剩的最多的钱。
	//buy2[i]表示前i天做第二笔交易 买入 股票后能剩的最多的钱。
	//sell2[i]表示前i天做第二笔交易 卖出 股票后能剩的最多的钱。
	buy1 := make([]int, N+1)
	sell1 := make([]int, N+1)
	buy2 := make([]int, N+1)
	sell2 := make([]int, N+1)
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	buy1[0] = math.MinInt32
	buy2[0] = math.MinInt32

	//i代表第i天
	for i := 1; i <= N; i++ {

		buy1[i] = max(buy1[i-1], -prices[i-1])
		sell1[i] = max(sell1[i-1], buy1[i-1]+prices[i-1])
		buy2[i] = max(buy2[i-1], sell1[i-1]-prices[i-1])
		sell2[i] = max(sell2[i-1], buy2[i-1]+prices[i-1])

	}

	return sell2[N]
}

func main() {

	fmt.Printf("%v\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("%v\n", maxProfit([]int{7, 6, 4, 3, 1}))
}
