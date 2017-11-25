package main

import "fmt"

func minimumTotal(triangle [][]int) int {

	N := len(triangle)

	if N == 0 {
		return 0
	}

	for i := 0; i < N; i++ {
		for j := 0; j <= i; j++ {
			if i == 0 {
			} else {
				if j == 0 {
					triangle[i][j] = triangle[i-1][j] + triangle[i][j]
				} else if j == i {
					triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
				} else {
					if triangle[i-1][j-1] < triangle[i-1][j] {
						triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
					} else {
						triangle[i][j] = triangle[i-1][j] + triangle[i][j]
					}
				}
			}
		}
	}

	min := triangle[N-1][0]
	for i := 1; i < N; i++ {
		if triangle[N-1][i] < min {
			min = triangle[N-1][i]
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
