package main

import (
	"fmt"
)

func combinationSum3(k int, n int) [][]int {

	var res [][]int

	combinationSum3DFS(0, 0, 0, k, n, []int{}, &res)

	return res
}

func combinationSum3DFS(index int, k int, n int, K int, N int, path []int, res *[][]int) {

	if k > K {
		return
	}

	if k == K && n == N {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := index + 1; i <= 9; i++ {
		combinationSum3DFS(i, k+1, n+i, K, N, append(path, i), res)
	}
}

func main() {

	fmt.Printf("%v \n", combinationSum3(3, 7))
	fmt.Printf("%v \n", combinationSum3(3, 9))
	fmt.Printf("%v \n", combinationSum3(4, 24))
}
