package main

import "fmt"

func numSquares(n int) int {
	//dp[n]
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {

		min := dp[i-1] + 1
		for j := 2; j*j <= i; j++ {

			t := dp[i-j*j] + 1
			if t < min {
				min = t
			}
		}
		dp[i] = min

	}
	return dp[n]

}
func main() {

	fmt.Printf("%v\n", numSquares(12))
	fmt.Printf("%v\n", numSquares(13))
	fmt.Printf("%v\n", numSquares(14))
}
