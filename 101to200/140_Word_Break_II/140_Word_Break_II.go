package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
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

	var path []string
	var res []string

	if dp[N] {
		wordBreakRecursive(s, N, 0, dp, dictMap, &path, &res)
	}

	return res
}

//这里表示找 s[index:N]的word break.
func wordBreakRecursive(s string, N int, index int, dp []bool, dictMap map[string]bool, path *[]string, res *[]string) {

	//这条路可行，记录下来
	if index == N {

		tmp := strings.Join(*path, " ")
		*res = append(*res, tmp)

	} else {
		for i := index; i <= N; i++ {

			if dp[i] {
				if dictMap[s[index:i]] {
					*path = append(*path, s[index:i])
					wordBreakRecursive(s, N, i, dp, dictMap, path, res)
					*path = (*path)[:len(*path)-1]

				}

			}

		}
	}

}

func main() {
	ss := wordBreak("aaa", []string{"a", "aa"})
	fmt.Printf("%v\n", ss)
	fmt.Printf("%v\n", wordBreak("ab", []string{"a", "b"}))
	fmt.Printf("%v\n", wordBreak("ab", []string{"a", "c"}))
	fmt.Printf("%v\n", wordBreak("leetcode", []string{"leet", "code"}))

}
