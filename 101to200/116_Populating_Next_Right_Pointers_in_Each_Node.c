package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {

	//第一个位置表示s的前i个字符，第二个位置表示t的前j个字符。
	sLen := len(s)
	tLen := len(t)

	var dp [][]int
	for i := 0; i <= sLen; i++ {
		row := make([]int, tLen+1)
		dp = append(dp, row)
	}

	//空串正好是空串的子字符串
	for i := 0; i < sLen+1; i++ {
		dp[i][0] = 1
	}
	for j := 1; j < tLen+1; j++ {
		dp[0][j] = 0
	}

	for i := 1; i < sLen+1; i++ {
		for j := 1; j < tLen+1; j++ {

			if s[i-1] == t[j-1] {
				//考虑是否用上母串中的第i个字符。
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[sLen][tLen]
}

func main() {

	fmt.Println(numDistinct("bbbi", "bi"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))

}
