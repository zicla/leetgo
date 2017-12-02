package main

import (
	"fmt"
)

func minCut(s string) int {
	N := len(s)
	//P记录了i到j的字串是否回文串。
	var P [][]bool
	for i := 0; i < N; i++ {
		P = append(P, make([]bool, N))
	}
	//dp[i]表示对于字符串 s[i:N] 最小需要的刀数。
	dp := make([]int, N+1)

	for i := 0; i <= N; i++ {
		dp[i] = N - i - 1;
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			P[i][j] = false;
		}
	}

	var min = func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := N - 1; i >= 0; i-- {

		//每增加一个字符串
		for j := i; j < N; j++ {
			si := s[i]
			sj := s[j]
			if si == sj && (j-i <= 1 || P[i+1][j-1]) {
				P[i][j] = true;
				dp[i] = min(dp[i], dp[j+1]+1);
			}
		}
	}
	return dp[0];
}

func main() {

	fmt.Printf("%v\n", minCut("ababa"))
	fmt.Printf("%v\n", minCut("ababababababababababababcbabababababababababababa"))
	fmt.Printf("%v\n", minCut("aab"))
	fmt.Printf("%v\n", minCut("abcd"))
}
