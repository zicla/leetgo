package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dictMap := make(map[string]bool)
	for _, v := range wordDict {
		dictMap[v] = true
	}
	N := len(s)

	//dp[i]表示前i个字符串是否ok.
	dp := make([]bool, N+1)
	dp[0] = true
	for i := 0; i <= N; i++ {
		for j := 0; j < i; j++ {

			if dp[j] && dictMap[s[j:i]] {
				dp[i] = true;
				break
			}

		}
	}

	return dp[N]

}

func main() {

	fmt.Printf("%v\n", wordBreak("ab", []string{"a", "c"}))
	fmt.Printf("%v\n", wordBreak("ab", []string{"a", "b"}))
	fmt.Printf("%v\n", wordBreak("leetcode", []string{"leet", "code"}))

	fmt.Printf("%v\n", wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))

}
