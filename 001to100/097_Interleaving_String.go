package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {

	N1 := len(s1)
	N2 := len(s2)
	N3 := len(s3)
	if N1+N2 != N3 {
		return false
	}

	dp := make([][]bool, N1+1)
	for k := 0; k < N1+1; k++ {
		dp[k] = make([]bool, N2+1)
	}

	for k1 := 1; k1 <= N1; k1++ {
		if s3[k1-1] == s1[k1-1] {
			dp[k1][0] = true
		} else {
			break
		}
	}

	for k2 := 1; k2 <= N2; k2++ {
		if s3[k2-1] == s2[k2-1] {
			dp[0][k2] = true
		} else {
			break
		}
	}

	dp[0][0] = true

	for i := 1; i <= N1; i++ {
		for j := 1; j <= N2; j++ {

			if dp[i][j-1] && s3[i+j-1] == s2[j-1] {
				dp[i][j] = true
			}

			if dp[i-1][j] && s3[i+j-1] == s1[i-1] {
				dp[i][j] = true
			}

		}
	}

	return dp[N1][N2]
}

func main() {
	fmt.Printf("%v\n", isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Printf("%v\n", isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Printf("%v\n", isInterleave("", "", ""))
}
