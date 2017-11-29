package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dictMap := make(map[string]int)
	for _, v := range wordDict {
		dictMap[v] ++
	}
	N := len(s)

	//dp[i][j]表示 startIndex=i endIndex=j 这个字符串从dictMap中多少种组合形式。我们要求的便是 dp[0][N-1]
	var dp [][]int
	for i := 0; i < N; i++ {
		dp = append(dp, make([]int, N))
	}

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			dp[i][j] = dictMap[s[i:j+1]]
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {

			for k := 0; k < j-i; k++ {
				//x := dp[i][i+k]
				//y := dp[i+k+1][j]
				//fmt.Println(x)
				//fmt.Println(y)
				dp[i][j] += dp[i][i+k] * dp[i+k+1][j]
			}
		}
	}

	return dp[0][N-1] > 0

}

func main() {

	fmt.Printf("%v\n", wordBreak("ab", []string{"a", "c"}))
	fmt.Printf("%v\n", wordBreak("ab", []string{"a", "b"}))
	fmt.Printf("%v\n", wordBreak("leetcode", []string{"leet1", "code"}))

	fmt.Printf("%v\n", wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))

}
