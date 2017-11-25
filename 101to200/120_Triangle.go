package main

import "fmt"

func minimumTotal(triangle [][]int) int {

	N := len(triangle)

	if N == 0 {
		return 0
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, i+1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			if i == 0 {
				dp[0][0] = triangle[0][0]
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + triangle[i][j]
				} else if j == i {
					dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				} else {
					if dp[i-1][j-1] < dp[i-1][j] {
						dp[i][j] = dp[i-1][j-1] + triangle[i][j]
					} else {
						dp[i][j] = dp[i-1][j] + triangle[i][j]
					}
				}
			}
		}
	}

	min := dp[N-1][0]
	for i := 1; i < N; i++ {
		if dp[N-1][i] < min {
			min = dp[N-1][i]
		}
	}

	return min

}

func main() {
	arr := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(arr))
}
