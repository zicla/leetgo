package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) []string {
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

	var store [][][]string
	for i := 0; i < N; i++ {
		var tmp [][]string
		for j := 0; j < N; j++ {
			var st []string
			tmp = append(tmp, st)
		}
		store = append(store, tmp)
	}

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			dp[i][j] = dictMap[s[i:j+1]]

			if dp[i][j] == 1 {
				store[i][j] = append(store[i][j], s[i:j+1])
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {

			for k := 0; k < j-i; k++ {

				step := dp[i][i+k] * dp[i+k+1][j]
				dp[i][j] += step

				if step > 0 {
					for x := 0; x < len(store[i][i+k]); x++ {
						for y := 0; y < len(store[i+k+1][j]); y++ {

							str1 := store[i][i+k][x]
							str2 := store[i+k+1][j][y]
							var str string
							if str1 == "" && str2 == "" {
								str = ""
							} else if str1 == "" {
								str = str2
							} else if str2 == "" {
								str = str1
							} else {
								str = str1 + " " + str2
							}

							store[i][j] = append(store[i][j], str)
						}
					}
				}

			}
		}
	}

	return store[0][N-1]
}

func main() {
	ss := wordBreak("aaa", []string{"a", "aa"})
	fmt.Printf("%v\n", ss)
	fmt.Printf("%v\n", wordBreak("ab", []string{"a", "b"}))
	fmt.Printf("%v\n", wordBreak("ab", []string{"a", "c"}))
	fmt.Printf("%v\n", wordBreak("leetcode", []string{"leet", "code"}))

}
